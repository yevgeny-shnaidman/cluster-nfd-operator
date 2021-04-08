// +build !ignore_autogenerated

/*
Copyright 2021. The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	conditionsv1 "github.com/openshift/custom-resource-status/conditions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMap) DeepCopyInto(out *ConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMap.
func (in *ConfigMap) DeepCopy() *ConfigMap {
	if in == nil {
		return nil
	}
	out := new(ConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureDiscovery) DeepCopyInto(out *NodeFeatureDiscovery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureDiscovery.
func (in *NodeFeatureDiscovery) DeepCopy() *NodeFeatureDiscovery {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFeatureDiscovery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureDiscoveryList) DeepCopyInto(out *NodeFeatureDiscoveryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeFeatureDiscovery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureDiscoveryList.
func (in *NodeFeatureDiscoveryList) DeepCopy() *NodeFeatureDiscoveryList {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureDiscoveryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFeatureDiscoveryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureDiscoverySpec) DeepCopyInto(out *NodeFeatureDiscoverySpec) {
	*out = *in
	out.Operand = in.Operand
	out.WorkerConfig = in.WorkerConfig
	out.CustomConfig = in.CustomConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureDiscoverySpec.
func (in *NodeFeatureDiscoverySpec) DeepCopy() *NodeFeatureDiscoverySpec {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureDiscoverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureDiscoveryStatus) DeepCopyInto(out *NodeFeatureDiscoveryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditionsv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureDiscoveryStatus.
func (in *NodeFeatureDiscoveryStatus) DeepCopy() *NodeFeatureDiscoveryStatus {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureDiscoveryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperandSpec) DeepCopyInto(out *OperandSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperandSpec.
func (in *OperandSpec) DeepCopy() *OperandSpec {
	if in == nil {
		return nil
	}
	out := new(OperandSpec)
	in.DeepCopyInto(out)
	return out
}
