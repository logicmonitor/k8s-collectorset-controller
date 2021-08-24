package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeBuilder is a SchemeBuilder.
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// AddToScheme is the SchemeBuilder AddToScheme function.
	AddToScheme = SchemeBuilder.AddToScheme
)

// GroupName is the group name used in this package.
const GroupName = "logicmonitor.com"

// SchemeGroupVersion is the group version used to register these objects.
var SchemaGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha2"}
var SchemeGroupVersionInternal = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}

// Resource takes an unqualified resource and returns a Group-qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemaGroupVersion.WithResource(resource).GroupResource()
}

// addKnownTypes adds the set of types defined in this package to the supplied scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemaGroupVersion,
		&CollectorSet{},
		&CollectorSetList{},
	)
	metav1.AddToGroupVersion(scheme, SchemaGroupVersion)

	scheme.AddKnownTypes(SchemeGroupVersionInternal,
		&CollectorSet{},
		&CollectorSetList{},
	)
	//metav1.AddToGroupVersion(scheme, SchemeGroupVersionInternal)
	return nil
}
