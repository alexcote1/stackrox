// +build !ignore_autogenerated

/*
Copyright 2021 Red Hat.

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
	commonv1alpha1 "github.com/stackrox/rox/operator/api/common/v1alpha1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerComponentSpec) DeepCopyInto(out *AnalyzerComponentSpec) {
	*out = *in
	if in.ScannerComponent != nil {
		in, out := &in.ScannerComponent, &out.ScannerComponent
		*out = new(ScannerComponentPolicy)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(ScannerReplicas)
		(*in).DeepCopyInto(*out)
	}
	if in.Scanner != nil {
		in, out := &in.Scanner, &out.Scanner
		*out = new(commonv1alpha1.DeploymentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ScannerDB != nil {
		in, out := &in.ScannerDB, &out.ScannerDB
		*out = new(commonv1alpha1.DeploymentSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerComponentSpec.
func (in *AnalyzerComponentSpec) DeepCopy() *AnalyzerComponentSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyzerComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Central) DeepCopyInto(out *Central) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Central.
func (in *Central) DeepCopy() *Central {
	if in == nil {
		return nil
	}
	out := new(Central)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Central) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CentralComponentSpec) DeepCopyInto(out *CentralComponentSpec) {
	*out = *in
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
	if in.TelemetryPolicy != nil {
		in, out := &in.TelemetryPolicy, &out.TelemetryPolicy
		*out = new(TelemetryPolicy)
		**out = **in
	}
	if in.AdminPasswordSecret != nil {
		in, out := &in.AdminPasswordSecret, &out.AdminPasswordSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(Persistence)
		(*in).DeepCopyInto(*out)
	}
	if in.Exposure != nil {
		in, out := &in.Exposure, &out.Exposure
		*out = new(Exposure)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CentralComponentSpec.
func (in *CentralComponentSpec) DeepCopy() *CentralComponentSpec {
	if in == nil {
		return nil
	}
	out := new(CentralComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CentralComponentStatus) DeepCopyInto(out *CentralComponentStatus) {
	*out = *in
	if in.GeneratedAdminPasswordSecret != nil {
		in, out := &in.GeneratedAdminPasswordSecret, &out.GeneratedAdminPasswordSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CentralComponentStatus.
func (in *CentralComponentStatus) DeepCopy() *CentralComponentStatus {
	if in == nil {
		return nil
	}
	out := new(CentralComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CentralList) DeepCopyInto(out *CentralList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Central, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CentralList.
func (in *CentralList) DeepCopy() *CentralList {
	if in == nil {
		return nil
	}
	out := new(CentralList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CentralList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CentralSpec) DeepCopyInto(out *CentralSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = new(Egress)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(commonv1alpha1.TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Central != nil {
		in, out := &in.Central, &out.Central
		*out = new(CentralComponentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Analyzer != nil {
		in, out := &in.Analyzer, &out.Analyzer
		*out = new(AnalyzerComponentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Customize != nil {
		in, out := &in.Customize, &out.Customize
		*out = new(commonv1alpha1.CustomizeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CentralSpec.
func (in *CentralSpec) DeepCopy() *CentralSpec {
	if in == nil {
		return nil
	}
	out := new(CentralSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CentralStatus) DeepCopyInto(out *CentralStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]commonv1alpha1.StackRoxCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeployedRelease != nil {
		in, out := &in.DeployedRelease, &out.DeployedRelease
		*out = new(commonv1alpha1.StackRoxRelease)
		**out = **in
	}
	if in.CentralStatus != nil {
		in, out := &in.CentralStatus, &out.CentralStatus
		*out = new(CentralComponentStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CentralStatus.
func (in *CentralStatus) DeepCopy() *CentralStatus {
	if in == nil {
		return nil
	}
	out := new(CentralStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Egress) DeepCopyInto(out *Egress) {
	*out = *in
	if in.ConnectivityPolicy != nil {
		in, out := &in.ConnectivityPolicy, &out.ConnectivityPolicy
		*out = new(ConnectivityPolicy)
		**out = **in
	}
	if in.ProxyConfigSecret != nil {
		in, out := &in.ProxyConfigSecret, &out.ProxyConfigSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Egress.
func (in *Egress) DeepCopy() *Egress {
	if in == nil {
		return nil
	}
	out := new(Egress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Exposure) DeepCopyInto(out *Exposure) {
	*out = *in
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(ExposureLoadBalancer)
		(*in).DeepCopyInto(*out)
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(ExposureNodePort)
		(*in).DeepCopyInto(*out)
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(ExposureRoute)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Exposure.
func (in *Exposure) DeepCopy() *Exposure {
	if in == nil {
		return nil
	}
	out := new(Exposure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposureLoadBalancer) DeepCopyInto(out *ExposureLoadBalancer) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposureLoadBalancer.
func (in *ExposureLoadBalancer) DeepCopy() *ExposureLoadBalancer {
	if in == nil {
		return nil
	}
	out := new(ExposureLoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposureNodePort) DeepCopyInto(out *ExposureNodePort) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposureNodePort.
func (in *ExposureNodePort) DeepCopy() *ExposureNodePort {
	if in == nil {
		return nil
	}
	out := new(ExposureNodePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposureRoute) DeepCopyInto(out *ExposureRoute) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposureRoute.
func (in *ExposureRoute) DeepCopy() *ExposureRoute {
	if in == nil {
		return nil
	}
	out := new(ExposureRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence) DeepCopyInto(out *Persistence) {
	*out = *in
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(string)
		**out = **in
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence.
func (in *Persistence) DeepCopy() *Persistence {
	if in == nil {
		return nil
	}
	out := new(Persistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaim) DeepCopyInto(out *PersistentVolumeClaim) {
	*out = *in
	if in.ClaimName != nil {
		in, out := &in.ClaimName, &out.ClaimName
		*out = new(string)
		**out = **in
	}
	if in.CreateClaim != nil {
		in, out := &in.CreateClaim, &out.CreateClaim
		*out = new(ClaimCreatePolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaim.
func (in *PersistentVolumeClaim) DeepCopy() *PersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScannerReplicas) DeepCopyInto(out *ScannerReplicas) {
	*out = *in
	if in.AutoScaling != nil {
		in, out := &in.AutoScaling, &out.AutoScaling
		*out = new(AutoScalingPolicy)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScannerReplicas.
func (in *ScannerReplicas) DeepCopy() *ScannerReplicas {
	if in == nil {
		return nil
	}
	out := new(ScannerReplicas)
	in.DeepCopyInto(out)
	return out
}
