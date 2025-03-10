//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Katalyst Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlConfig) DeepCopyInto(out *AccessControlConfig) {
	*out = *in
	if in.AccessControlPolicies != nil {
		in, out := &in.AccessControlPolicies, &out.AccessControlPolicies
		*out = make([]AccessControlPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlConfig.
func (in *AccessControlConfig) DeepCopy() *AccessControlConfig {
	if in == nil {
		return nil
	}
	out := new(AccessControlConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicy) DeepCopyInto(out *AccessControlPolicy) {
	*out = *in
	in.PolicyRule.DeepCopyInto(&out.PolicyRule)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicy.
func (in *AccessControlPolicy) DeepCopy() *AccessControlPolicy {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminQoSConfig) DeepCopyInto(out *AdminQoSConfig) {
	*out = *in
	if in.ReclaimedResourceConfig != nil {
		in, out := &in.ReclaimedResourceConfig, &out.ReclaimedResourceConfig
		*out = new(ReclaimedResourceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EvictionConfig != nil {
		in, out := &in.EvictionConfig, &out.EvictionConfig
		*out = new(EvictionConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminQoSConfig.
func (in *AdminQoSConfig) DeepCopy() *AdminQoSConfig {
	if in == nil {
		return nil
	}
	out := new(AdminQoSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminQoSConfiguration) DeepCopyInto(out *AdminQoSConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminQoSConfiguration.
func (in *AdminQoSConfiguration) DeepCopy() *AdminQoSConfiguration {
	if in == nil {
		return nil
	}
	out := new(AdminQoSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminQoSConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminQoSConfigurationList) DeepCopyInto(out *AdminQoSConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdminQoSConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminQoSConfigurationList.
func (in *AdminQoSConfigurationList) DeepCopy() *AdminQoSConfigurationList {
	if in == nil {
		return nil
	}
	out := new(AdminQoSConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminQoSConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminQoSConfigurationSpec) DeepCopyInto(out *AdminQoSConfigurationSpec) {
	*out = *in
	in.GenericConfigSpec.DeepCopyInto(&out.GenericConfigSpec)
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminQoSConfigurationSpec.
func (in *AdminQoSConfigurationSpec) DeepCopy() *AdminQoSConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(AdminQoSConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfig) DeepCopyInto(out *AuthConfig) {
	*out = *in
	if in.BasicAuthConfig != nil {
		in, out := &in.BasicAuthConfig, &out.BasicAuthConfig
		*out = new(BasicAuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AccessControlConfig != nil {
		in, out := &in.AccessControlConfig, &out.AccessControlConfig
		*out = new(AccessControlConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfig.
func (in *AuthConfig) DeepCopy() *AuthConfig {
	if in == nil {
		return nil
	}
	out := new(AuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfiguration) DeepCopyInto(out *AuthConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfiguration.
func (in *AuthConfiguration) DeepCopy() *AuthConfiguration {
	if in == nil {
		return nil
	}
	out := new(AuthConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfigurationList) DeepCopyInto(out *AuthConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfigurationList.
func (in *AuthConfigurationList) DeepCopy() *AuthConfigurationList {
	if in == nil {
		return nil
	}
	out := new(AuthConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfigurationSpec) DeepCopyInto(out *AuthConfigurationSpec) {
	*out = *in
	in.GenericConfigSpec.DeepCopyInto(&out.GenericConfigSpec)
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfigurationSpec.
func (in *AuthConfigurationSpec) DeepCopy() *AuthConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuthConfig) DeepCopyInto(out *BasicAuthConfig) {
	*out = *in
	if in.UserPasswordPairs != nil {
		in, out := &in.UserPasswordPairs, &out.UserPasswordPairs
		*out = make([]UserPasswordPair, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuthConfig.
func (in *BasicAuthConfig) DeepCopy() *BasicAuthConfig {
	if in == nil {
		return nil
	}
	out := new(BasicAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUHeadroomConfig) DeepCopyInto(out *CPUHeadroomConfig) {
	*out = *in
	if in.UtilBasedConfig != nil {
		in, out := &in.UtilBasedConfig, &out.UtilBasedConfig
		*out = new(CPUHeadroomUtilBasedConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUHeadroomConfig.
func (in *CPUHeadroomConfig) DeepCopy() *CPUHeadroomConfig {
	if in == nil {
		return nil
	}
	out := new(CPUHeadroomConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUHeadroomUtilBasedConfig) DeepCopyInto(out *CPUHeadroomUtilBasedConfig) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.TargetReclaimedCoreUtilization != nil {
		in, out := &in.TargetReclaimedCoreUtilization, &out.TargetReclaimedCoreUtilization
		*out = new(float64)
		**out = **in
	}
	if in.MaxReclaimedCoreUtilization != nil {
		in, out := &in.MaxReclaimedCoreUtilization, &out.MaxReclaimedCoreUtilization
		*out = new(float64)
		**out = **in
	}
	if in.MaxOversoldRate != nil {
		in, out := &in.MaxOversoldRate, &out.MaxOversoldRate
		*out = new(float64)
		**out = **in
	}
	if in.MaxHeadroomCapacityRate != nil {
		in, out := &in.MaxHeadroomCapacityRate, &out.MaxHeadroomCapacityRate
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUHeadroomUtilBasedConfig.
func (in *CPUHeadroomUtilBasedConfig) DeepCopy() *CPUHeadroomUtilBasedConfig {
	if in == nil {
		return nil
	}
	out := new(CPUHeadroomUtilBasedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUPressureEvictionConfig) DeepCopyInto(out *CPUPressureEvictionConfig) {
	*out = *in
	if in.EnableLoadEviction != nil {
		in, out := &in.EnableLoadEviction, &out.EnableLoadEviction
		*out = new(bool)
		**out = **in
	}
	if in.LoadUpperBoundRatio != nil {
		in, out := &in.LoadUpperBoundRatio, &out.LoadUpperBoundRatio
		*out = new(float64)
		**out = **in
	}
	if in.LoadThresholdMetPercentage != nil {
		in, out := &in.LoadThresholdMetPercentage, &out.LoadThresholdMetPercentage
		*out = new(float64)
		**out = **in
	}
	if in.LoadMetricRingSize != nil {
		in, out := &in.LoadMetricRingSize, &out.LoadMetricRingSize
		*out = new(int)
		**out = **in
	}
	if in.LoadEvictionCoolDownTime != nil {
		in, out := &in.LoadEvictionCoolDownTime, &out.LoadEvictionCoolDownTime
		*out = new(v1.Duration)
		**out = **in
	}
	if in.EnableSuppressionEviction != nil {
		in, out := &in.EnableSuppressionEviction, &out.EnableSuppressionEviction
		*out = new(bool)
		**out = **in
	}
	if in.MaxSuppressionToleranceRate != nil {
		in, out := &in.MaxSuppressionToleranceRate, &out.MaxSuppressionToleranceRate
		*out = new(float64)
		**out = **in
	}
	if in.MinSuppressionToleranceDuration != nil {
		in, out := &in.MinSuppressionToleranceDuration, &out.MinSuppressionToleranceDuration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.GracePeriod != nil {
		in, out := &in.GracePeriod, &out.GracePeriod
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUPressureEvictionConfig.
func (in *CPUPressureEvictionConfig) DeepCopy() *CPUPressureEvictionConfig {
	if in == nil {
		return nil
	}
	out := new(CPUPressureEvictionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeConfig) DeepCopyInto(out *CustomNodeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeConfig.
func (in *CustomNodeConfig) DeepCopy() *CustomNodeConfig {
	if in == nil {
		return nil
	}
	out := new(CustomNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomNodeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeConfigList) DeepCopyInto(out *CustomNodeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomNodeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeConfigList.
func (in *CustomNodeConfigList) DeepCopy() *CustomNodeConfigList {
	if in == nil {
		return nil
	}
	out := new(CustomNodeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomNodeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeConfigSpec) DeepCopyInto(out *CustomNodeConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeConfigSpec.
func (in *CustomNodeConfigSpec) DeepCopy() *CustomNodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CustomNodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNodeConfigStatus) DeepCopyInto(out *CustomNodeConfigStatus) {
	*out = *in
	if in.KatalystCustomConfigList != nil {
		in, out := &in.KatalystCustomConfigList, &out.KatalystCustomConfigList
		*out = make([]TargetConfig, len(*in))
		copy(*out, *in)
	}
	if in.ServiceProfileConfigList != nil {
		in, out := &in.ServiceProfileConfigList, &out.ServiceProfileConfigList
		*out = make([]TargetConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNodeConfigStatus.
func (in *CustomNodeConfigStatus) DeepCopy() *CustomNodeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(CustomNodeConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralSelector) DeepCopyInto(out *EphemeralSelector) {
	*out = *in
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LastDuration != nil {
		in, out := &in.LastDuration, &out.LastDuration
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralSelector.
func (in *EphemeralSelector) DeepCopy() *EphemeralSelector {
	if in == nil {
		return nil
	}
	out := new(EphemeralSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvictionConfig) DeepCopyInto(out *EvictionConfig) {
	*out = *in
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CPUPressureEvictionConfig != nil {
		in, out := &in.CPUPressureEvictionConfig, &out.CPUPressureEvictionConfig
		*out = new(CPUPressureEvictionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemLoadPressureEvictionConfig != nil {
		in, out := &in.SystemLoadPressureEvictionConfig, &out.SystemLoadPressureEvictionConfig
		*out = new(SystemLoadPressureEvictionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MemoryPressureEvictionConfig != nil {
		in, out := &in.MemoryPressureEvictionConfig, &out.MemoryPressureEvictionConfig
		*out = new(MemoryPressureEvictionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RootfsPressureEvictionConfig != nil {
		in, out := &in.RootfsPressureEvictionConfig, &out.RootfsPressureEvictionConfig
		*out = new(RootfsPressureEvictionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ReclaimedResourcesEvictionConfig != nil {
		in, out := &in.ReclaimedResourcesEvictionConfig, &out.ReclaimedResourcesEvictionConfig
		*out = new(ReclaimedResourcesEvictionConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvictionConfig.
func (in *EvictionConfig) DeepCopy() *EvictionConfig {
	if in == nil {
		return nil
	}
	out := new(EvictionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericConfigCondition) DeepCopyInto(out *GenericConfigCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericConfigCondition.
func (in *GenericConfigCondition) DeepCopy() *GenericConfigCondition {
	if in == nil {
		return nil
	}
	out := new(GenericConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericConfigSpec) DeepCopyInto(out *GenericConfigSpec) {
	*out = *in
	in.EphemeralSelector.DeepCopyInto(&out.EphemeralSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericConfigSpec.
func (in *GenericConfigSpec) DeepCopy() *GenericConfigSpec {
	if in == nil {
		return nil
	}
	out := new(GenericConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericConfigStatus) DeepCopyInto(out *GenericConfigStatus) {
	*out = *in
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]GenericConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericConfigStatus.
func (in *GenericConfigStatus) DeepCopy() *GenericConfigStatus {
	if in == nil {
		return nil
	}
	out := new(GenericConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KatalystCustomConfig) DeepCopyInto(out *KatalystCustomConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KatalystCustomConfig.
func (in *KatalystCustomConfig) DeepCopy() *KatalystCustomConfig {
	if in == nil {
		return nil
	}
	out := new(KatalystCustomConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KatalystCustomConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KatalystCustomConfigCondition) DeepCopyInto(out *KatalystCustomConfigCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KatalystCustomConfigCondition.
func (in *KatalystCustomConfigCondition) DeepCopy() *KatalystCustomConfigCondition {
	if in == nil {
		return nil
	}
	out := new(KatalystCustomConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KatalystCustomConfigList) DeepCopyInto(out *KatalystCustomConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KatalystCustomConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KatalystCustomConfigList.
func (in *KatalystCustomConfigList) DeepCopy() *KatalystCustomConfigList {
	if in == nil {
		return nil
	}
	out := new(KatalystCustomConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KatalystCustomConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KatalystCustomConfigSpec) DeepCopyInto(out *KatalystCustomConfigSpec) {
	*out = *in
	out.TargetType = in.TargetType
	if in.NodeLabelSelectorAllowedKeyList != nil {
		in, out := &in.NodeLabelSelectorAllowedKeyList, &out.NodeLabelSelectorAllowedKeyList
		*out = make([]PriorityNodeLabelSelectorAllowedKeyList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KatalystCustomConfigSpec.
func (in *KatalystCustomConfigSpec) DeepCopy() *KatalystCustomConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KatalystCustomConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KatalystCustomConfigStatus) DeepCopyInto(out *KatalystCustomConfigStatus) {
	*out = *in
	if in.InvalidTargetConfigList != nil {
		in, out := &in.InvalidTargetConfigList, &out.InvalidTargetConfigList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]KatalystCustomConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KatalystCustomConfigStatus.
func (in *KatalystCustomConfigStatus) DeepCopy() *KatalystCustomConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KatalystCustomConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryHeadroomConfig) DeepCopyInto(out *MemoryHeadroomConfig) {
	*out = *in
	if in.UtilBasedConfig != nil {
		in, out := &in.UtilBasedConfig, &out.UtilBasedConfig
		*out = new(MemoryHeadroomUtilBasedConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryHeadroomConfig.
func (in *MemoryHeadroomConfig) DeepCopy() *MemoryHeadroomConfig {
	if in == nil {
		return nil
	}
	out := new(MemoryHeadroomConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryHeadroomUtilBasedConfig) DeepCopyInto(out *MemoryHeadroomUtilBasedConfig) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.FreeBasedRatio != nil {
		in, out := &in.FreeBasedRatio, &out.FreeBasedRatio
		*out = new(float64)
		**out = **in
	}
	if in.StaticBasedCapacity != nil {
		in, out := &in.StaticBasedCapacity, &out.StaticBasedCapacity
		*out = new(float64)
		**out = **in
	}
	if in.CacheBasedRatio != nil {
		in, out := &in.CacheBasedRatio, &out.CacheBasedRatio
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryHeadroomUtilBasedConfig.
func (in *MemoryHeadroomUtilBasedConfig) DeepCopy() *MemoryHeadroomUtilBasedConfig {
	if in == nil {
		return nil
	}
	out := new(MemoryHeadroomUtilBasedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryPressureEvictionConfig) DeepCopyInto(out *MemoryPressureEvictionConfig) {
	*out = *in
	if in.EnableNumaLevelEviction != nil {
		in, out := &in.EnableNumaLevelEviction, &out.EnableNumaLevelEviction
		*out = new(bool)
		**out = **in
	}
	if in.EnableSystemLevelEviction != nil {
		in, out := &in.EnableSystemLevelEviction, &out.EnableSystemLevelEviction
		*out = new(bool)
		**out = **in
	}
	if in.NumaVictimMinimumUtilizationThreshold != nil {
		in, out := &in.NumaVictimMinimumUtilizationThreshold, &out.NumaVictimMinimumUtilizationThreshold
		*out = new(float64)
		**out = **in
	}
	if in.NumaFreeBelowWatermarkTimesThreshold != nil {
		in, out := &in.NumaFreeBelowWatermarkTimesThreshold, &out.NumaFreeBelowWatermarkTimesThreshold
		*out = new(int)
		**out = **in
	}
	if in.SystemKswapdRateThreshold != nil {
		in, out := &in.SystemKswapdRateThreshold, &out.SystemKswapdRateThreshold
		*out = new(int)
		**out = **in
	}
	if in.SystemKswapdRateExceedDurationThreshold != nil {
		in, out := &in.SystemKswapdRateExceedDurationThreshold, &out.SystemKswapdRateExceedDurationThreshold
		*out = new(int)
		**out = **in
	}
	if in.SystemFreeMemoryThresholdMinimum != nil {
		in, out := &in.SystemFreeMemoryThresholdMinimum, &out.SystemFreeMemoryThresholdMinimum
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.NumaEvictionRankingMetrics != nil {
		in, out := &in.NumaEvictionRankingMetrics, &out.NumaEvictionRankingMetrics
		*out = make([]NumaEvictionRankingMetric, len(*in))
		copy(*out, *in)
	}
	if in.SystemEvictionRankingMetrics != nil {
		in, out := &in.SystemEvictionRankingMetrics, &out.SystemEvictionRankingMetrics
		*out = make([]SystemEvictionRankingMetric, len(*in))
		copy(*out, *in)
	}
	if in.EnableRSSOveruseEviction != nil {
		in, out := &in.EnableRSSOveruseEviction, &out.EnableRSSOveruseEviction
		*out = new(bool)
		**out = **in
	}
	if in.RSSOveruseRateThreshold != nil {
		in, out := &in.RSSOveruseRateThreshold, &out.RSSOveruseRateThreshold
		*out = new(float64)
		**out = **in
	}
	if in.GracePeriod != nil {
		in, out := &in.GracePeriod, &out.GracePeriod
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryPressureEvictionConfig.
func (in *MemoryPressureEvictionConfig) DeepCopy() *MemoryPressureEvictionConfig {
	if in == nil {
		return nil
	}
	out := new(MemoryPressureEvictionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRule) DeepCopyInto(out *PolicyRule) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRule.
func (in *PolicyRule) DeepCopy() *PolicyRule {
	if in == nil {
		return nil
	}
	out := new(PolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityNodeLabelSelectorAllowedKeyList) DeepCopyInto(out *PriorityNodeLabelSelectorAllowedKeyList) {
	*out = *in
	if in.KeyList != nil {
		in, out := &in.KeyList, &out.KeyList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityNodeLabelSelectorAllowedKeyList.
func (in *PriorityNodeLabelSelectorAllowedKeyList) DeepCopy() *PriorityNodeLabelSelectorAllowedKeyList {
	if in == nil {
		return nil
	}
	out := new(PriorityNodeLabelSelectorAllowedKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimedResourceConfig) DeepCopyInto(out *ReclaimedResourceConfig) {
	*out = *in
	if in.EnableReclaim != nil {
		in, out := &in.EnableReclaim, &out.EnableReclaim
		*out = new(bool)
		**out = **in
	}
	if in.ReservedResourceForReport != nil {
		in, out := &in.ReservedResourceForReport, &out.ReservedResourceForReport
		*out = new(corev1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[corev1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.MinReclaimedResourceForReport != nil {
		in, out := &in.MinReclaimedResourceForReport, &out.MinReclaimedResourceForReport
		*out = new(corev1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[corev1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.ReservedResourceForAllocate != nil {
		in, out := &in.ReservedResourceForAllocate, &out.ReservedResourceForAllocate
		*out = new(corev1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[corev1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.MinReclaimedResourceForAllocate != nil {
		in, out := &in.MinReclaimedResourceForAllocate, &out.MinReclaimedResourceForAllocate
		*out = new(corev1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[corev1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.CPUHeadroomConfig != nil {
		in, out := &in.CPUHeadroomConfig, &out.CPUHeadroomConfig
		*out = new(CPUHeadroomConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MemoryHeadroomConfig != nil {
		in, out := &in.MemoryHeadroomConfig, &out.MemoryHeadroomConfig
		*out = new(MemoryHeadroomConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimedResourceConfig.
func (in *ReclaimedResourceConfig) DeepCopy() *ReclaimedResourceConfig {
	if in == nil {
		return nil
	}
	out := new(ReclaimedResourceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimedResourcesEvictionConfig) DeepCopyInto(out *ReclaimedResourcesEvictionConfig) {
	*out = *in
	if in.EvictionThreshold != nil {
		in, out := &in.EvictionThreshold, &out.EvictionThreshold
		*out = make(map[corev1.ResourceName]float64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.GracePeriod != nil {
		in, out := &in.GracePeriod, &out.GracePeriod
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimedResourcesEvictionConfig.
func (in *ReclaimedResourcesEvictionConfig) DeepCopy() *ReclaimedResourcesEvictionConfig {
	if in == nil {
		return nil
	}
	out := new(ReclaimedResourcesEvictionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootfsPressureEvictionConfig) DeepCopyInto(out *RootfsPressureEvictionConfig) {
	*out = *in
	if in.EnableRootfsPressureEviction != nil {
		in, out := &in.EnableRootfsPressureEviction, &out.EnableRootfsPressureEviction
		*out = new(bool)
		**out = **in
	}
	if in.GracePeriod != nil {
		in, out := &in.GracePeriod, &out.GracePeriod
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootfsPressureEvictionConfig.
func (in *RootfsPressureEvictionConfig) DeepCopy() *RootfsPressureEvictionConfig {
	if in == nil {
		return nil
	}
	out := new(RootfsPressureEvictionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemLoadPressureEvictionConfig) DeepCopyInto(out *SystemLoadPressureEvictionConfig) {
	*out = *in
	if in.SoftThreshold != nil {
		in, out := &in.SoftThreshold, &out.SoftThreshold
		*out = new(int64)
		**out = **in
	}
	if in.HardThreshold != nil {
		in, out := &in.HardThreshold, &out.HardThreshold
		*out = new(int64)
		**out = **in
	}
	if in.HistorySize != nil {
		in, out := &in.HistorySize, &out.HistorySize
		*out = new(int64)
		**out = **in
	}
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(int64)
		**out = **in
	}
	if in.CoolDownTime != nil {
		in, out := &in.CoolDownTime, &out.CoolDownTime
		*out = new(int64)
		**out = **in
	}
	if in.GracePeriod != nil {
		in, out := &in.GracePeriod, &out.GracePeriod
		*out = new(int64)
		**out = **in
	}
	if in.ThresholdMetPercentage != nil {
		in, out := &in.ThresholdMetPercentage, &out.ThresholdMetPercentage
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemLoadPressureEvictionConfig.
func (in *SystemLoadPressureEvictionConfig) DeepCopy() *SystemLoadPressureEvictionConfig {
	if in == nil {
		return nil
	}
	out := new(SystemLoadPressureEvictionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetConfig) DeepCopyInto(out *TargetConfig) {
	*out = *in
	out.ConfigType = in.ConfigType
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetConfig.
func (in *TargetConfig) DeepCopy() *TargetConfig {
	if in == nil {
		return nil
	}
	out := new(TargetConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPasswordPair) DeepCopyInto(out *UserPasswordPair) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPasswordPair.
func (in *UserPasswordPair) DeepCopy() *UserPasswordPair {
	if in == nil {
		return nil
	}
	out := new(UserPasswordPair)
	in.DeepCopyInto(out)
	return out
}
