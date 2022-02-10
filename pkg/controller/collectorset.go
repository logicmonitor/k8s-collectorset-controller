package controller

import (
	"fmt"
	"strings"

	crv1alpha2 "github.com/logicmonitor/k8s-collectorset-controller/pkg/apis/v1alpha2"
	"github.com/logicmonitor/k8s-collectorset-controller/pkg/constants"
	"github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
	"github.com/logicmonitor/lm-sdk-go/models"
	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	clientset "k8s.io/client-go/kubernetes"
)

// CreateOrUpdateCollectorSet creates a replicaset for each collector in
// a CollectorSet
func CreateOrUpdateCollectorSet(collectorset *crv1alpha2.CollectorSet, controller *Controller) ([]int32, error) {
	groupID := collectorset.Spec.GroupID
	if groupID == 0 || !checkCollectorGroupExistsByID(controller.LogicmonitorClient, groupID) {
		groupName := constants.ClusterCollectorGroupPrefix + collectorset.Spec.ClusterName
		log.Infof("Group name is %s", groupName)

		newGroupID, err := getCollectorGroupID(controller.LogicmonitorClient, groupName, collectorset)
		if err != nil {
			return nil, err
		}
		log.Infof("Adding collector group %q with ID %d", strings.Title(groupName), newGroupID)
		groupID = newGroupID
	}

	ids, err := getCollectorIDs(controller.LogicmonitorClient, groupID, collectorset)
	if err != nil {
		return nil, err
	}

	statefulset, err := createStsObject(collectorset, ids, controller.CollectorsetConfig.IgnoreSSL)
	if err != nil {
		return nil, err
	}

	setProxyConfiguration(collectorset, statefulset)

	if _, _err := controller.Clientset.AppsV1().StatefulSets(statefulset.ObjectMeta.Namespace).Create(statefulset); _err != nil {
		if !apierrors.IsAlreadyExists(_err) {
			return nil, _err
		}
		if _, _err := controller.Clientset.AppsV1().StatefulSets(statefulset.ObjectMeta.Namespace).Update(statefulset); _err != nil {
			return nil, _err
		}
	}

	collectorset.Status.IDs = ids

	err = updateCollectors(controller.LogicmonitorClient, ids)
	if err != nil {
		log.Warnf("Failed to set collector backup agents: %v", err)
	}
	return collectorset.Status.IDs, nil
}

func createStsObject(collectorset *crv1alpha2.CollectorSet, ids []int32, ignoreSSL bool) (*appsv1.StatefulSet, error) {
	secretIsOptional := false
	collectorSize := strings.ToLower(collectorset.Spec.Size)
	log.Infof("Collector size is %s", collectorSize)

	imagePullPolicy, err := getCollectorImagePullPolicy(collectorset)
	if err != nil {
		return nil, err
	}

	statefulset := appsv1.StatefulSet{}

	// Default statefulset params - not allowed to edit by user
	statefulset.TypeMeta = metav1.TypeMeta{
		APIVersion: "apps/v1",
		Kind:       "StatefulSet",
	}

	if collectorset.Spec.Labels != nil {
		statefulset.ObjectMeta.Labels = collectorset.Spec.Labels
	} else {
		statefulset.ObjectMeta.Labels = make(map[string]string)
	}
	// Add label in user defined labels
	statefulset.ObjectMeta.Labels["logicmonitor.com/collectorset"] = collectorset.Name
	log.Debugf("statefulset.ObjectMeta.Labels = %v", statefulset.ObjectMeta.Labels)

	statefulset.ObjectMeta.Name = collectorset.Name
	statefulset.ObjectMeta.Namespace = collectorset.Namespace

	statefulset.Annotations = collectorset.Spec.Annotations

	statefulset.Spec = collectorset.Spec.CollectorStatefulSetSpec

	// validate tolerations and log error message accordingly
	validateTolerations(statefulset.Spec.Template.Spec.Tolerations)

	statefulset.Spec.Replicas = collectorset.Spec.Replicas
	statefulset.Spec.Selector = &metav1.LabelSelector{
		MatchLabels: map[string]string{
			"logicmonitor.com/collectorset": collectorset.Name,
		},
	}

	// configuring pod template
	// load podLables
	podLabels := make(map[string]string)
	if statefulset.Spec.Template.ObjectMeta.Labels != nil {
		podLabels = statefulset.Spec.Template.ObjectMeta.Labels
	}
	if statefulset.ObjectMeta.Labels != nil {
		for key, value := range statefulset.ObjectMeta.Labels {
			podLabels[key] = value
		}
	}
	podLabels["logicmonitor.com/collectorset"] = collectorset.Name
	log.Debugf("podLabels = %v", podLabels)

	// load annotations
	annotations := make(map[string]string)
	if statefulset.Spec.Template.ObjectMeta.Annotations != nil {
		annotations = statefulset.Spec.Template.ObjectMeta.Annotations
	}
	if statefulset.ObjectMeta.Annotations != nil {
		for key, value := range statefulset.ObjectMeta.Annotations {
			annotations[key] = value
		}
	}
	log.Debugf("annotations = %v", annotations)

	statefulset.Spec.Template.ObjectMeta = metav1.ObjectMeta{
		Namespace:   collectorset.Namespace,
		Labels:      podLabels,
		Annotations: annotations,
	}

	statefulset.Spec.Template.Spec.ServiceAccountName = constants.CollectorServiceAccountName
	statefulset.Spec.Template.Spec.Affinity = &apiv1.Affinity{
		PodAntiAffinity: &apiv1.PodAntiAffinity{
			RequiredDuringSchedulingIgnoredDuringExecution: []apiv1.PodAffinityTerm{
				{
					LabelSelector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"logicmonitor.com/collectorset": collectorset.Name,
						},
					},
					TopologyKey: "kubernetes.io/hostname",
				},
			},
		},
	}

	statefulset.Spec.Template.Spec.Containers = []apiv1.Container{
		{
			Name:            "collector",
			Image:           getCollectorImage(collectorset),
			ImagePullPolicy: imagePullPolicy,
			Env: []apiv1.EnvVar{
				{
					Name: "account",
					ValueFrom: &apiv1.EnvVarSource{
						SecretKeyRef: &apiv1.SecretKeySelector{
							LocalObjectReference: apiv1.LocalObjectReference{
								Name: constants.CollectorsetControllerSecretName,
							},
							Key:      "account",
							Optional: &secretIsOptional,
						},
					},
				},
				{
					Name: "access_id",
					ValueFrom: &apiv1.EnvVarSource{
						SecretKeyRef: &apiv1.SecretKeySelector{
							LocalObjectReference: apiv1.LocalObjectReference{
								Name: constants.CollectorsetControllerSecretName,
							},
							Key:      "accessID",
							Optional: &secretIsOptional,
						},
					},
				},
				{
					Name: "access_key",
					ValueFrom: &apiv1.EnvVarSource{
						SecretKeyRef: &apiv1.SecretKeySelector{
							LocalObjectReference: apiv1.LocalObjectReference{
								Name: constants.CollectorsetControllerSecretName,
							},
							Key:      "accessKey",
							Optional: &secretIsOptional,
						},
					},
				},
				{
					Name:  "kubernetes",
					Value: "true",
				},
				{
					Name:  "collector_size",
					Value: collectorSize,
				},
				{
					Name:  "collector_version",
					Value: fmt.Sprint(collectorset.Spec.CollectorVersion), // the default value is 0, santaba will assign the latest version
				},
				{
					Name:  "use_ea",
					Value: fmt.Sprint(collectorset.Spec.UseEA), // the default value is false, santaba will assign the latest GD version
				},
				{
					Name:  "COLLECTOR_IDS",
					Value: strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]"),
				},
				{
					Name:  "ignore_ssl",
					Value: fmt.Sprint(ignoreSSL), // the default value is false
				},
			},
			Resources: getResourceRequirements(collectorSize, collectorset.Spec.CollectorStatefulSetSpec),
		},
	}

	statefulset.Spec.UpdateStrategy = appsv1.StatefulSetUpdateStrategy{
		Type: appsv1.RollingUpdateStatefulSetStrategyType,
	}
	statefulset.Spec.PodManagementPolicy = appsv1.ParallelPodManagement

	log.Infof("statefulset object= %v", statefulset.Spec.Template.Spec)
	return &statefulset, nil
}

func getCollectorImage(collectorset *crv1alpha2.CollectorSet) string {
	if collectorset.Spec.ImageRepository == "" {
		return constants.DefaultCollectorImage
	}
	imageTag := collectorset.Spec.ImageTag
	if imageTag == "" {
		imageTag = constants.DefaultCollectorImageTag
	}
	return collectorset.Spec.ImageRepository + ":" + imageTag
}

func getCollectorImagePullPolicy(collectorset *crv1alpha2.CollectorSet) (apiv1.PullPolicy, error) {
	if collectorset.Spec.ImagePullPolicy == "" {
		return constants.DefaultCollectorImagePullPolicy, nil
	}
	switch collectorset.Spec.ImagePullPolicy {
	case apiv1.PullAlways, apiv1.PullNever, apiv1.PullIfNotPresent:
		return collectorset.Spec.ImagePullPolicy, nil
	}
	return "", fmt.Errorf("unsupported imagePullPolicy value: %v, supported values: [%v, %v, %v]", collectorset.Spec.ImagePullPolicy, apiv1.PullAlways, apiv1.PullNever, apiv1.PullIfNotPresent)
}

func setProxyConfiguration(collectorset *crv1alpha2.CollectorSet, statefulset *appsv1.StatefulSet) {
	if collectorset.Spec.ProxyURL == "" {
		return
	}
	container := &statefulset.Spec.Template.Spec.Containers[0]
	container.Env = append(container.Env,
		apiv1.EnvVar{
			Name:  "proxy_url",
			Value: collectorset.Spec.ProxyURL,
		},
	)
	if collectorset.Spec.SecretName != "" {
		secretIsOptionalTrue := true
		container.Env = append(container.Env,
			apiv1.EnvVar{
				Name: "proxy_user",
				ValueFrom: &apiv1.EnvVarSource{
					SecretKeyRef: &apiv1.SecretKeySelector{
						LocalObjectReference: apiv1.LocalObjectReference{
							Name: collectorset.Spec.SecretName,
						},
						Key:      "proxyUser",
						Optional: &secretIsOptionalTrue,
					},
				},
			},
			apiv1.EnvVar{
				Name: "proxy_pass",
				ValueFrom: &apiv1.EnvVarSource{
					SecretKeyRef: &apiv1.SecretKeySelector{
						LocalObjectReference: apiv1.LocalObjectReference{
							Name: collectorset.Spec.SecretName,
						},
						Key:      "proxyPass",
						Optional: &secretIsOptionalTrue,
					},
				},
			},
		)
	}
}

func validateTolerations(tolerations []apiv1.Toleration) {
	valid := true
	if tolerations != nil {
		for _, toleration := range tolerations {
			if toleration.Operator == apiv1.TolerationOpExists && toleration.Value != "" {
				log.Errorf("Value must be empty when 'operator' is 'Exists'. Toleration: %v", toleration)
				valid = false
			} else if toleration.Operator != apiv1.TolerationOpExists && toleration.Key == "" {
				log.Errorf("Operator must be 'Exists' when 'key' is empty. Toleration: %v", toleration)
				valid = false
			} else if toleration.Effect != apiv1.TaintEffectNoExecute && toleration.TolerationSeconds != nil {
				log.Errorf("Effect must be 'NoExecute' when 'tolerationSeconds' is set. Toleration: %v", toleration)
				valid = false
			}
		}
	}
	if valid {
		log.Debug("Valid configuration for tolerations")
	}
}

func updateCollectors(client *client.LMSdkGo, ids []int32) error {
	// if there is only one collector, there will be no backup for it
	if len(ids) < 2 {
		return nil
	}

	for i := 0; i < len(ids); i++ {
		var backupAgentID int32
		if i == 0 {
			backupAgentID = ids[len(ids)-1]
		} else {
			backupAgentID = ids[i-1]
		}
		err := updateCollectorBackupAgent(client, ids[i], backupAgentID)
		if err != nil {
			log.Warnf("Failed to update the backup collector id: %v", err)
		}
	}

	return nil
}

// DeleteCollectorSet deletes the collectorset.
func DeleteCollectorSet(collectorset *crv1alpha2.CollectorSet, client clientset.Interface) error {
	data := []byte(`[{"op":"add","path":"/spec/replicas","value": 0}]`)
	if _, err := client.AppsV1().StatefulSets(collectorset.Namespace).Patch(collectorset.Name, types.JSONPatchType, data); err != nil {
		return err
	}

	deleteOpts := metav1.DeleteOptions{}
	return client.AppsV1().StatefulSets(collectorset.Namespace).Delete(collectorset.Name, &deleteOpts)
}

func checkCollectorGroupExistsByID(client *client.LMSdkGo, id int32) bool {
	params := lm.NewGetCollectorGroupByIDParams()
	params.SetID(id)
	fields := "id"
	params.SetFields(&fields)
	restResponse, err := client.LM.GetCollectorGroupByID(params)
	if err != nil || restResponse.Payload == nil {
		log.Warnf("Failed to get collector group with id %d", id)
		return false
	}
	return true
}

func getCollectorGroupID(client *client.LMSdkGo, name string, collectorset *crv1alpha2.CollectorSet) (int32, error) {
	params := lm.NewGetCollectorGroupListParams()
	filter := fmt.Sprintf("name:\"%s\"", name)
	params.SetFilter(&filter)
	restResponse, err := client.LM.GetCollectorGroupList(params)
	if err != nil {
		return -1, err
	}

	if restResponse.Payload == nil || restResponse.Payload.Total == 0 {
		log.Infof("Adding collector group with name %q", name)
		return addCollectorGroup(client, name, collectorset)
	}
	if restResponse.Payload.Total == 1 {
		return restResponse.Payload.Items[0].ID, err
	}
	return -1, fmt.Errorf("failed to get collector group ID")
}

func addCollectorGroup(client *client.LMSdkGo, name string, collectorset *crv1alpha2.CollectorSet) (int32, error) {
	kubernetesLabelApp := constants.CustomPropertyKubernetesLabelApp
	kubernetesLabelAppValue := constants.CustomPropertyKubernetesLabelAppValue
	autoClusterName := constants.CustomPropertyAutoClusterName
	AutoClusterNameValue := collectorset.Spec.ClusterName
	customProperties := []*models.NameAndValue{
		{Name: &kubernetesLabelApp, Value: &kubernetesLabelAppValue},
		{Name: &autoClusterName, Value: &AutoClusterNameValue},
	}

	body := &models.CollectorGroup{
		Name:             &name,
		CustomProperties: customProperties,
	}
	params := lm.NewAddCollectorGroupParams()
	params.SetBody(body)
	restResponse, err := client.LM.AddCollectorGroup(params)
	if err != nil {
		return -1, err
	}
	return restResponse.Payload.ID, nil
}

// $(statefulset name)-$(ordinal)
func getCollectorIDs(client *client.LMSdkGo, groupID int32, collectorset *crv1alpha2.CollectorSet) ([]int32, error) {
	var ids []int32
	for ordinal := int32(0); ordinal < *collectorset.Spec.Replicas; ordinal++ {
		name := fmt.Sprintf("%s%s-%d", constants.ClusterCollectorGroupPrefix, collectorset.Spec.ClusterName, ordinal)
		filter := fmt.Sprintf("collectorGroupId:%v,description:\"%v\"", groupID, name)
		params := lm.NewGetCollectorListParams()
		params.SetFilter(&filter)
		restResponse, err := client.LM.GetCollectorList(params)
		if err != nil {
			return nil, err
		}
		var id int32
		if restResponse.Payload == nil || restResponse.Payload.Total == 0 {
			log.Infof("Adding collector with description %q", name)
			kubernetesLabelApp := constants.CustomPropertyKubernetesLabelApp
			kubernetesLabelAppValue := constants.CustomPropertyKubernetesLabelAppValue
			autoClusterName := constants.CustomPropertyAutoClusterName
			AutoClusterNameValue := collectorset.Spec.ClusterName
			customProperties := []*models.NameAndValue{
				{Name: &kubernetesLabelApp, Value: &kubernetesLabelAppValue},
				{Name: &autoClusterName, Value: &AutoClusterNameValue},
			}

			body := &models.Collector{
				Description:                   name,
				CollectorGroupID:              groupID,
				NeedAutoCreateCollectorDevice: false,
				CustomProperties:              customProperties,
			}
			id, err = addCollector(client, body)
			if err != nil {
				return nil, err
			}

			// update the escalating chain id, if failed the value will be the default value
			// the default value of this option param is 0, which means disable notification
			collector, err := getCollectorByID(client, id)
			if err != nil || collector == nil {
				log.Warnf("Failed to get the collector, err: %v", err)
				collector = body
				collector.ID = id
			}

			collector.EscalatingChainID = collectorset.Spec.EscalationChainID
			err = updateCollector(client, collector)
			if err != nil {
				log.Warnf("Failed to update the escalation chain id. The default value will be used. %v", err)
			}
		} else {
			id = restResponse.Payload.Items[0].ID
		}
		ids = append(ids, id)
	}

	return ids, nil
}

// nolint: gocyclo
func getResourceRequirements(size string, spec appsv1.StatefulSetSpec) apiv1.ResourceRequirements {
	resourceList := apiv1.ResourceList{}
	var quantity *resource.Quantity
	switch size {
	case "nano":
		quantity = resource.NewQuantity(2*1024*1024*1024, resource.BinarySI)
	case "small":
		quantity = resource.NewQuantity(2*1024*1024*1024, resource.BinarySI)
	case "medium":
		quantity = resource.NewQuantity(4*1024*1024*1024, resource.BinarySI)
	case "large":
		quantity = resource.NewQuantity(8*1024*1024*1024, resource.BinarySI)
	case "extra_large":
		quantity = resource.NewQuantity(16*1024*1024*1024, resource.BinarySI)
	case "double_extra_large":
		quantity = resource.NewQuantity(32*1024*1024*1024, resource.BinarySI)
	default:
		break
	}
	var userRequests apiv1.ResourceList
	resourceList[apiv1.ResourceMemory] = *quantity
	if len(spec.Template.Spec.Containers) >= 0 {
		for _, container := range spec.Template.Spec.Containers {
			if container.Name == constants.CollectorServiceAccountName {
				userLimits := container.Resources.Limits
				userRequests = container.Resources.Requests
				if es, ok := userLimits[apiv1.ResourceEphemeralStorage]; ok {
					resourceList[apiv1.ResourceEphemeralStorage] = es
				}
			}
		}
	}

	return apiv1.ResourceRequirements{
		Limits:   resourceList,
		Requests: userRequests,
	}
}

func addCollector(client *client.LMSdkGo, body *models.Collector) (int32, error) {
	params := lm.NewAddCollectorParams()
	params.SetBody(body)
	restResponse, err := client.LM.AddCollector(params)
	if err != nil {
		return -1, err
	}
	return restResponse.Payload.ID, nil
}

func getCollectorByID(client *client.LMSdkGo, id int32) (*models.Collector, error) {
	params := lm.NewGetCollectorByIDParams()
	params.SetID(id)
	restResponse, err := client.LM.GetCollectorByID(params)
	if err != nil {
		return nil, err
	}
	return restResponse.Payload, nil
}

func updateCollector(client *client.LMSdkGo, body *models.Collector) error {
	params := lm.NewUpdateCollectorByIDParams()
	params.SetBody(body)
	params.SetID(body.ID)
	_, err := client.LM.UpdateCollectorByID(params)
	if err != nil {
		return err
	}

	return nil
}

func updateCollectorBackupAgent(client *client.LMSdkGo, id, backupID int32) error {
	// Get all the fields before updating to prevent setting default values to the other fields
	restResponse, err := getCollectorByID(client, id)
	if err != nil || restResponse == nil {
		return fmt.Errorf("failed to get the collector: %v", err)
	}

	collector := restResponse
	collector.EnableFailBack = true
	collector.BackupAgentID = backupID
	updateErr := updateCollector(client, collector)
	if updateErr != nil {
		return fmt.Errorf("failed to update the collector: %v", updateErr)
	}
	return nil
}
