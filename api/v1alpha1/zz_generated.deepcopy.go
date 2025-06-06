//go:build !ignore_autogenerated

/*
Copyright 2025.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelMetadata) DeepCopyInto(out *KernelMetadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelMetadata.
func (in *KernelMetadata) DeepCopy() *KernelMetadata {
	if in == nil {
		return nil
	}
	out := new(KernelMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelProperties) DeepCopyInto(out *KernelProperties) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]KernelMetadata, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelProperties.
func (in *KernelProperties) DeepCopy() *KernelProperties {
	if in == nil {
		return nil
	}
	out := new(KernelProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TritonKernelCache) DeepCopyInto(out *TritonKernelCache) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TritonKernelCache.
func (in *TritonKernelCache) DeepCopy() *TritonKernelCache {
	if in == nil {
		return nil
	}
	out := new(TritonKernelCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TritonKernelCache) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TritonKernelCacheList) DeepCopyInto(out *TritonKernelCacheList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TritonKernelCache, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TritonKernelCacheList.
func (in *TritonKernelCacheList) DeepCopy() *TritonKernelCacheList {
	if in == nil {
		return nil
	}
	out := new(TritonKernelCacheList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TritonKernelCacheList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TritonKernelCacheSpec) DeepCopyInto(out *TritonKernelCacheSpec) {
	*out = *in
	in.KernelProperties.DeepCopyInto(&out.KernelProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TritonKernelCacheSpec.
func (in *TritonKernelCacheSpec) DeepCopy() *TritonKernelCacheSpec {
	if in == nil {
		return nil
	}
	out := new(TritonKernelCacheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TritonKernelCacheStatus) DeepCopyInto(out *TritonKernelCacheStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastSynced.DeepCopyInto(&out.LastSynced)
	if in.ObservedMetadata != nil {
		in, out := &in.ObservedMetadata, &out.ObservedMetadata
		*out = make([]KernelMetadata, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TritonKernelCacheStatus.
func (in *TritonKernelCacheStatus) DeepCopy() *TritonKernelCacheStatus {
	if in == nil {
		return nil
	}
	out := new(TritonKernelCacheStatus)
	in.DeepCopyInto(out)
	return out
}
