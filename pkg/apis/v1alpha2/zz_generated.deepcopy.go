// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	distributor "github.com/logicmonitor/k8s-collectorset-controller/pkg/distributor"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorSet) DeepCopyInto(out *CollectorSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorSet.
func (in *CollectorSet) DeepCopy() *CollectorSet {
	if in == nil {
		return nil
	}
	out := new(CollectorSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectorSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorSetList) DeepCopyInto(out *CollectorSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CollectorSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorSetList.
func (in *CollectorSetList) DeepCopy() *CollectorSetList {
	if in == nil {
		return nil
	}
	out := new(CollectorSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectorSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorSetPolicy) DeepCopyInto(out *CollectorSetPolicy) {
	*out = *in
	if in.DistibutionStrategy != nil {
		in, out := &in.DistibutionStrategy, &out.DistibutionStrategy
		*out = new(distributor.Type)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorSetPolicy.
func (in *CollectorSetPolicy) DeepCopy() *CollectorSetPolicy {
	if in == nil {
		return nil
	}
	out := new(CollectorSetPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorSetSpec) DeepCopyInto(out *CollectorSetSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.CollectorStatefulSetSpec.DeepCopyInto(&out.CollectorStatefulSetSpec)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(CollectorSetPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorSetSpec.
func (in *CollectorSetSpec) DeepCopy() *CollectorSetSpec {
	if in == nil {
		return nil
	}
	out := new(CollectorSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorSetStatus) DeepCopyInto(out *CollectorSetStatus) {
	*out = *in
	if in.IDs != nil {
		in, out := &in.IDs, &out.IDs
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorSetStatus.
func (in *CollectorSetStatus) DeepCopy() *CollectorSetStatus {
	if in == nil {
		return nil
	}
	out := new(CollectorSetStatus)
	in.DeepCopyInto(out)
	return out
}
