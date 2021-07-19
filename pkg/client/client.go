package client

import (
	"fmt"
	"reflect"
	"time"

	crv1alpha1 "github.com/logicmonitor/k8s-collectorset-controller/pkg/apis/v1alpha1"
	log "github.com/sirupsen/logrus"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const crdName = crv1alpha1.CollectorSetResourcePlural + "." + crv1alpha1.GroupName

// Client represents the CollectorSet client.
type Client struct {
	Clientset              *clientset.Clientset
	RESTClient             *rest.RESTClient
	APIExtensionsClientset *apiextensionsclientset.Clientset
}

// NewForConfig instantiates and returns the client and scheme.
func NewForConfig(cfg *rest.Config) (*Client, *runtime.Scheme, error) {
	s := runtime.NewScheme()
	err := crv1alpha1.AddToScheme(s)
	if err != nil {
		return nil, nil, err
	}

	client, err := clientset.NewForConfig(cfg)
	if err != nil {
		return nil, nil, err
	}

	config := *cfg
	config.GroupVersion = &crv1alpha1.SchemeGroupVersion
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
	return &apiextensionsv1.JSONSchemaProps{
		Description: "collectorset controller's spec schema",
		Type:        "object",
		Properties: map[string]apiextensionsv1.JSONSchemaProps{
			"spec": {
				Type: "object",
				Properties: map[string]apiextensionsv1.JSONSchemaProps{
					"imageRepository": {
						Description: "The image repository of the Collector container",
						Type:        "string",
					},
					"imageTag": {
						Description: "The image tag of the Collector container",
						Type:        "string",
					},
					"imagePullPolicy": {
						Description: "The image pull policy of the Collector container",
						Type:        "string",
					},
					"replicas": {
						Description: "The number of collectors to create",
						Type:        "integer",
					},
					"size": {
						Description: "The collector size to install. Can be nano, small, medium, or large",
						Type:        "string",
					},
					"clusterName": {
						Description: "The collector clustername",
						Type:        "string",
					},
					"groupID": {
						Description: "The ID of the group of the collectors",
						Type:        "integer",
					},
					"escalationChainID": {
						Description: "The ID of the escalation chain of the collectors",
						Type:        "integer",
					},
					"collectorVersion": {
						Description: "The version of the collectors",
						Type:        "integer",
					},
					"useEA": {
						Description: "On a collector downloading event, either download the latest EA version or the latest GD version",
						Type:        "boolean",
					},
					"proxyURL": {
						Description: "The Http/s proxy url of the collectors",
						Type:        "string",
					},
					"secretName": {
						Description: "The Secret resource name of the collectors",
						Type:        "string",
					},
					"statefulsetspec": {
						Description: "Collector statefulset template",
						Type:        "object",
					},
					"policy": {
						Type: "object",
						Properties: map[string]apiextensionsv1.JSONSchemaProps{
							"distributionStrategy": {
								Description: "distribution stratergy policy",
								Type:        "string",
							},
							"orchestrator": {
								Description: "orchestrator type",
								Type:        "string",
							},
						},
					},
				},
			},
		},
	}
}

// CreateCustomResourceDefinition creates the CRD for collectors.
func (c *Client) CreateCustomResourceDefinition() (*apiextensionsv1.CustomResourceDefinition, error) {
	schema := &apiextensionsv1.CustomResourceValidation{}

	schema.OpenAPIV3Schema = getCustomResourceDefinationSchema()
	crd := &apiextensionsv1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: crdName,
		},
		Spec: apiextensionsv1.CustomResourceDefinitionSpec{
			Group: crv1alpha1.GroupName,
			Names: apiextensionsv1.CustomResourceDefinitionNames{
				Plural: crv1alpha1.CollectorSetResourcePlural,
				Kind:   reflect.TypeOf(crv1alpha1.CollectorSet{}).Name(),
			},
			Scope: apiextensionsv1.ClusterScoped,
			Versions: []apiextensionsv1.CustomResourceDefinitionVersion{{
				Name:    crv1alpha1.SchemeGroupVersion.Version,
				Served:  true,
				Storage: true,
				Schema:  schema,
			}},
		},
	}

	crd1, err := c.APIExtensionsClientset.ApiextensionsV1().CustomResourceDefinitions().Create(crd)
	if err != nil {
		log.Debugf("error while creating crd- %v", err)
		return nil, err
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

	return crd1, nil
}

// // WaitForCollectorMonitoring creates a collector and waits for it to be ready.
// func WaitForCollectorMonitoring(clientset *rest.RESTClient, name string) error {
// 	return wait.Poll(100*time.Millisecond, 10*time.Second, func() (bool, error) {
// 		var collector crv1alpha1.CollectorSet
// 		err := clientset.Get().
// 			Resource(crv1alpha1.CollectorSetResourcePlural).
// 			Namespace(apiv1.NamespaceDefault).
// 			Name(name).
// 			Do().Into(&collector)

// 		if err == nil && collector.Status.State == crv1alpha1.CollectorSetStateMonitoring {
// 			return true, nil
// 		}

// 		return false, err
// 	})
// }
