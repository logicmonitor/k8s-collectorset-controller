package client

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	crv1alpha2 "github.com/logicmonitor/k8s-collectorset-controller/pkg/apis/v1alpha2"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const crdName = crv1alpha2.CollectorSetResourcePlural + "." + crv1alpha2.GroupName

// Client represents the CollectorSet client.
type Client struct {
	Clientset              *clientset.Clientset
	RESTClient             *rest.RESTClient
	APIExtensionsClientset *apiextensionsclientset.Clientset
}

// NewForConfig instantiates and returns the client and scheme.
func NewForConfig(cfg *rest.Config) (*Client, *runtime.Scheme, error) {
	s := runtime.NewScheme()
	err := crv1alpha2.AddToScheme(s)
	if err != nil {
		return nil, nil, err
	}

	client, err := clientset.NewForConfig(cfg)
	if err != nil {
		return nil, nil, err
	}

	config := *cfg
	config.GroupVersion = &crv1alpha2.SchemaGroupVersion
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.NewCodecFactory(s)
	restclient, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, nil, err
	}

	// Instantiate the Kubernetes API extensions client.
	apiextensionsclient, err := apiextensionsclientset.NewForConfig(&config)
	if err != nil {
		return nil, nil, err
	}

	c := &Client{
		Clientset:              client,
		RESTClient:             restclient,
		APIExtensionsClientset: apiextensionsclient,
	}

	return c, s, nil
}

func getCustomResourceDefinationSchema() *apiextensionsv1.JSONSchemaProps {
	minValue := 0.0
	minReplicas := 1.0
	minGroupID := -1.0
	defaultReplicas, _ := json.Marshal(1) //nolint:gosec

	return &apiextensionsv1.JSONSchemaProps{
		Description: "The collectorset specefication schema",
		Type:        "object",
		Required:    []string{"spec"},
		Properties: map[string]apiextensionsv1.JSONSchemaProps{
			"spec": {
				Type:     "object",
				Required: []string{"imageRepository", "imageTag", "imagePullPolicy", "replicas", "size", "clusterName"},
				Properties: map[string]apiextensionsv1.JSONSchemaProps{
					"imageRepository": {
						Description: "The image repository of the collector container",
						Type:        "string",
						Default: &apiextensionsv1.JSON{
							Raw: []byte("\"logicmonitor/collector\""),
						},
					},
					"imageTag": {
						Description: "The image tag of the collector container",
						Type:        "string",
						Default: &apiextensionsv1.JSON{
							Raw: []byte("\"latest\""),
						},
					},
					"imagePullPolicy": {
						Description: "The image pull policy of the collector container",
						Type:        "string",
						Default: &apiextensionsv1.JSON{
							Raw: []byte("\"Always\""),
						},
						Enum: []apiextensionsv1.JSON{
							{
								Raw: []byte("\"Always\""),
							},
							{
								Raw: []byte("\"IfNotPresent\""),
							},
							{
								Raw: []byte("\"Never\""),
							},
						},
					},
					"replicas": {
						Description: "The number of collector replicas",
						Type:        "integer",
						Minimum:     &minReplicas,
						Default: &apiextensionsv1.JSON{
							Raw: defaultReplicas,
						},
					},
					"size": {
						Description: "The collector size. Available collector sizes: nano, small, medium, large, extra_large, double_extra_large",
						Type:        "string",
						Default: &apiextensionsv1.JSON{
							Raw: []byte("\"nano\""),
						},
						Enum: []apiextensionsv1.JSON{
							{
								Raw: []byte("\"nano\""),
							},
							{
								Raw: []byte("\"small\""),
							},
							{
								Raw: []byte("\"medium\""),
							},
							{
								Raw: []byte("\"large\""),
							},
							{
								Raw: []byte("\"extra_large\""),
							},
							{
								Raw: []byte("\"double_extra_large\""),
							},
						},
					},
					"clusterName": {
						Description: "The clustername of the collector",
						Type:        "string",
					},
					"groupID": {
						Description: "The groupId of the collector",
						Type:        "integer",
						Minimum:     &minGroupID,
					},
					"escalationChainID": {
						Description: "The escalation chain Id of the collectors",
						Type:        "integer",
						Minimum:     &minValue,
					},
					"collectorVersion": {
						Description: "The collector version (Fractional numbered version is invalid. For ex: 29.101 is invalid, correct input is 29101)",
						Type:        "integer",
						Minimum:     &minValue,
					},
					"useEA": {
						Description: "Flag to opt for EA collector versions",
						Type:        "boolean",
					},
					"proxyURL": {
						Description: "The Http/Https proxy url of the collector",
						Type:        "string",
					},
					"secretName": {
						Description: "The Secret resource name of the collector",
						Type:        "string",
					},
					"statefulsetspec": {
						Description: "The collector StatefulSet specification for customizations",
						Type:        "object",
					},
					"policy": {
						Type: "object",
						Properties: map[string]apiextensionsv1.JSONSchemaProps{
							"distributionStrategy": {
								Description: "Distribution strategy to provide collector ID to the client requests from available running collectors",
								Type:        "string",
								Default: &apiextensionsv1.JSON{
									Raw: []byte("\"RoundRobin\""),
								},
							},
							"orchestrator": {
								Description: "The container orchestration platform designed to automate the deployment, scaling, and management of containerized applications",
								Type:        "string",
								Default: &apiextensionsv1.JSON{
									Raw: []byte("\"Kubernetes\""),
								},
							},
						},
					},
				},
			},
		},
	}
}

// CreateCustomResourceDefinition creates the CRD for collectors.
// nolint: gocyclo
func (c *Client) CreateCustomResourceDefinition() (*apiextensionsv1.CustomResourceDefinition, error) {
	schema := &apiextensionsv1.CustomResourceValidation{}
	preserveUnknownFields := true
	schema.OpenAPIV3Schema = getCustomResourceDefinationSchema()
	crd := &apiextensionsv1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: crdName,
		},
		Spec: apiextensionsv1.CustomResourceDefinitionSpec{
			Group: crv1alpha2.GroupName,
			Names: apiextensionsv1.CustomResourceDefinitionNames{
				Plural: crv1alpha2.CollectorSetResourcePlural,
				Kind:   reflect.TypeOf(crv1alpha2.CollectorSet{}).Name(),
			},
			Scope: apiextensionsv1.NamespaceScoped,
			Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
				{
					Name:    "v1alpha1",
					Served:  true,
					Storage: false,
					Schema: &apiextensionsv1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
							Description:            "The collectorset specification schema",
							Type:                   "object",
							XPreserveUnknownFields: &preserveUnknownFields,
						},
					},
				},
				{
					Name:    "v1alpha2",
					Served:  true,
					Storage: true,
					Schema:  schema,
				}},
		},
	}

	_, err := c.APIExtensionsClientset.ApiextensionsV1().CustomResourceDefinitions().Create(crd)
	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			if err1 := c.updateCRD(crd); err1 != nil {
				return nil, err1
			}
			return nil, nil
		}
		return nil, fmt.Errorf("error while creating crd: %w", err)
	}

	// wait for CRD being established
	err = wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
		crd, err = c.APIExtensionsClientset.ApiextensionsV1().CustomResourceDefinitions().Get(crdName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, cond := range crd.Status.Conditions {
			switch cond.Type {
			case apiextensionsv1.Established:
				if cond.Status == apiextensionsv1.ConditionTrue {
					return true, err
				}
			case apiextensionsv1.NamesAccepted:
				if cond.Status == apiextensionsv1.ConditionFalse {
					fmt.Printf("Name conflict: %v\n", cond.Reason)
				}
			}
		}
		return false, err
	})
	if err != nil {
		deleteErr := c.APIExtensionsClientset.ApiextensionsV1().CustomResourceDefinitions().Delete(crdName, nil)
		if deleteErr != nil {
			return nil, errors.NewAggregate([]error{err, deleteErr})
		}
		return nil, err
	}

	return crd, nil
}

func (c *Client) updateCRD(crd *apiextensionsv1.CustomResourceDefinition) error {
	// Get current CRD object for retrieving ResourceVersion
	// ResourceVersion is required for updating newer CRD object
	existingCrd, err := c.APIExtensionsClientset.ApiextensionsV1().CustomResourceDefinitions().Get(crdName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("error while retrieving existing crd: %w", err)
	}

	crd.SetResourceVersion(existingCrd.GetResourceVersion())
	_, err1 := c.APIExtensionsClientset.ApiextensionsV1().CustomResourceDefinitions().Update(crd)
	if err1 != nil {
		return fmt.Errorf("error while updating crd: %w", err1)
	}
	return nil
}
