package client

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	crv1alpha2 "github.com/logicmonitor/k8s-collectorset-controller/pkg/apis/v1alpha2"
	log "github.com/sirupsen/logrus"
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

func getCustomResourceDefinationSchema() (*apiextensionsv1.JSONSchemaProps, error) {
	schema := &apiextensionsv1.JSONSchemaProps{}
	err := json.Unmarshal([]byte(schemaStr), schema)
	if err != nil {
		log.Errorf("Failed to parse schema definition: %v: %v", err, schemaStr)
		return nil, fmt.Errorf("failed to parse schema definition: %w: %v", err, schemaStr)
	}
	log.Debugf("Unmarshaled schema: %v", schema)
	return schema, nil
}

// CreateCustomResourceDefinition creates the CRD for collectors.
// nolint: gocyclo
func (c *Client) CreateCustomResourceDefinition() (*apiextensionsv1.CustomResourceDefinition, error) {
	schema := &apiextensionsv1.CustomResourceValidation{}
	preserveUnknownFields := true
	var err error
	schema.OpenAPIV3Schema, err = getCustomResourceDefinationSchema()
	if err != nil {
		return nil, err
	}
	crd := &apiextensionsv1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: crdName,
		},
		Spec: apiextensionsv1.CustomResourceDefinitionSpec{
			Group: crv1alpha2.GroupName,
			Names: apiextensionsv1.CustomResourceDefinitionNames{
				Plural: crv1alpha2.CollectorSetResourcePlural,
				Kind:   reflect.TypeOf(crv1alpha2.CollectorSet{}).Name(),
				ShortNames: []string{
					// cs is default shortname in kubernetes for ComponentStatuses, it is deprecated from 1.19+,
					// when they remove it, we can use cs shortname here
					"lmcs",
				},
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
				},
			},
		},
	}

	_, err = c.APIExtensionsClientset.ApiextensionsV1().CustomResourceDefinitions().Create(crd)
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
