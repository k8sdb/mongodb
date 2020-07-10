// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigNode) DeepCopyInto(out *ConfigNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigNode.
func (in *ConfigNode) DeepCopy() *ConfigNode {
	if in == nil {
		return nil
	}
	out := new(ConfigNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchHorizontalScalingSpec) DeepCopyInto(out *ElasticsearchHorizontalScalingSpec) {
	*out = *in
	if in.Master != nil {
		in, out := &in.Master, &out.Master
		*out = new(int32)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(int32)
		**out = **in
	}
	if in.Client != nil {
		in, out := &in.Client, &out.Client
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchHorizontalScalingSpec.
func (in *ElasticsearchHorizontalScalingSpec) DeepCopy() *ElasticsearchHorizontalScalingSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchHorizontalScalingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchOpsRequest) DeepCopyInto(out *ElasticsearchOpsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchOpsRequest.
func (in *ElasticsearchOpsRequest) DeepCopy() *ElasticsearchOpsRequest {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchOpsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchOpsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchOpsRequestList) DeepCopyInto(out *ElasticsearchOpsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticsearchOpsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchOpsRequestList.
func (in *ElasticsearchOpsRequestList) DeepCopy() *ElasticsearchOpsRequestList {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchOpsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchOpsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchOpsRequestSpec) DeepCopyInto(out *ElasticsearchOpsRequestSpec) {
	*out = *in
	out.DatabaseRef = in.DatabaseRef
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeSpec)
		**out = **in
	}
	if in.HorizontalScaling != nil {
		in, out := &in.HorizontalScaling, &out.HorizontalScaling
		*out = new(ElasticsearchHorizontalScalingSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchOpsRequestSpec.
func (in *ElasticsearchOpsRequestSpec) DeepCopy() *ElasticsearchOpsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchOpsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchOpsRequestStatus) DeepCopyInto(out *ElasticsearchOpsRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchOpsRequestStatus.
func (in *ElasticsearchOpsRequestStatus) DeepCopy() *ElasticsearchOpsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchOpsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdOpsRequest) DeepCopyInto(out *EtcdOpsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdOpsRequest.
func (in *EtcdOpsRequest) DeepCopy() *EtcdOpsRequest {
	if in == nil {
		return nil
	}
	out := new(EtcdOpsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EtcdOpsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdOpsRequestList) DeepCopyInto(out *EtcdOpsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EtcdOpsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdOpsRequestList.
func (in *EtcdOpsRequestList) DeepCopy() *EtcdOpsRequestList {
	if in == nil {
		return nil
	}
	out := new(EtcdOpsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EtcdOpsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdOpsRequestSpec) DeepCopyInto(out *EtcdOpsRequestSpec) {
	*out = *in
	out.DatabaseRef = in.DatabaseRef
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdOpsRequestSpec.
func (in *EtcdOpsRequestSpec) DeepCopy() *EtcdOpsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(EtcdOpsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdOpsRequestStatus) DeepCopyInto(out *EtcdOpsRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdOpsRequestStatus.
func (in *EtcdOpsRequestStatus) DeepCopy() *EtcdOpsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(EtcdOpsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcachedOpsRequest) DeepCopyInto(out *MemcachedOpsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcachedOpsRequest.
func (in *MemcachedOpsRequest) DeepCopy() *MemcachedOpsRequest {
	if in == nil {
		return nil
	}
	out := new(MemcachedOpsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemcachedOpsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcachedOpsRequestList) DeepCopyInto(out *MemcachedOpsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MemcachedOpsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcachedOpsRequestList.
func (in *MemcachedOpsRequestList) DeepCopy() *MemcachedOpsRequestList {
	if in == nil {
		return nil
	}
	out := new(MemcachedOpsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemcachedOpsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcachedOpsRequestSpec) DeepCopyInto(out *MemcachedOpsRequestSpec) {
	*out = *in
	out.DatabaseRef = in.DatabaseRef
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcachedOpsRequestSpec.
func (in *MemcachedOpsRequestSpec) DeepCopy() *MemcachedOpsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(MemcachedOpsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcachedOpsRequestStatus) DeepCopyInto(out *MemcachedOpsRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcachedOpsRequestStatus.
func (in *MemcachedOpsRequestStatus) DeepCopy() *MemcachedOpsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(MemcachedOpsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBHorizontalScalingSpec) DeepCopyInto(out *MongoDBHorizontalScalingSpec) {
	*out = *in
	if in.Shard != nil {
		in, out := &in.Shard, &out.Shard
		*out = new(MongoDBShardNode)
		**out = **in
	}
	if in.ConfigServer != nil {
		in, out := &in.ConfigServer, &out.ConfigServer
		*out = new(ConfigNode)
		**out = **in
	}
	if in.Mongos != nil {
		in, out := &in.Mongos, &out.Mongos
		*out = new(MongosNode)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBHorizontalScalingSpec.
func (in *MongoDBHorizontalScalingSpec) DeepCopy() *MongoDBHorizontalScalingSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBHorizontalScalingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsRequest) DeepCopyInto(out *MongoDBOpsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsRequest.
func (in *MongoDBOpsRequest) DeepCopy() *MongoDBOpsRequest {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBOpsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsRequestList) DeepCopyInto(out *MongoDBOpsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDBOpsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsRequestList.
func (in *MongoDBOpsRequestList) DeepCopy() *MongoDBOpsRequestList {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBOpsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsRequestSpec) DeepCopyInto(out *MongoDBOpsRequestSpec) {
	*out = *in
	out.DatabaseRef = in.DatabaseRef
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeSpec)
		**out = **in
	}
	if in.HorizontalScaling != nil {
		in, out := &in.HorizontalScaling, &out.HorizontalScaling
		*out = new(MongoDBHorizontalScalingSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VerticalScaling != nil {
		in, out := &in.VerticalScaling, &out.VerticalScaling
		*out = new(MongoDBVerticalScalingSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsRequestSpec.
func (in *MongoDBOpsRequestSpec) DeepCopy() *MongoDBOpsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsRequestStatus) DeepCopyInto(out *MongoDBOpsRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsRequestStatus.
func (in *MongoDBOpsRequestStatus) DeepCopy() *MongoDBOpsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBShardNode) DeepCopyInto(out *MongoDBShardNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBShardNode.
func (in *MongoDBShardNode) DeepCopy() *MongoDBShardNode {
	if in == nil {
		return nil
	}
	out := new(MongoDBShardNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBVerticalScalingSpec) DeepCopyInto(out *MongoDBVerticalScalingSpec) {
	*out = *in
	if in.Standalone != nil {
		in, out := &in.Standalone, &out.Standalone
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Mongos != nil {
		in, out := &in.Mongos, &out.Mongos
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigServer != nil {
		in, out := &in.ConfigServer, &out.ConfigServer
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Shard != nil {
		in, out := &in.Shard, &out.Shard
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Exporter != nil {
		in, out := &in.Exporter, &out.Exporter
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBVerticalScalingSpec.
func (in *MongoDBVerticalScalingSpec) DeepCopy() *MongoDBVerticalScalingSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBVerticalScalingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongosNode) DeepCopyInto(out *MongosNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongosNode.
func (in *MongosNode) DeepCopy() *MongosNode {
	if in == nil {
		return nil
	}
	out := new(MongosNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLHorizontalScalingSpec) DeepCopyInto(out *MySQLHorizontalScalingSpec) {
	*out = *in
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLHorizontalScalingSpec.
func (in *MySQLHorizontalScalingSpec) DeepCopy() *MySQLHorizontalScalingSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLHorizontalScalingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLOpsRequest) DeepCopyInto(out *MySQLOpsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLOpsRequest.
func (in *MySQLOpsRequest) DeepCopy() *MySQLOpsRequest {
	if in == nil {
		return nil
	}
	out := new(MySQLOpsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLOpsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLOpsRequestList) DeepCopyInto(out *MySQLOpsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLOpsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLOpsRequestList.
func (in *MySQLOpsRequestList) DeepCopy() *MySQLOpsRequestList {
	if in == nil {
		return nil
	}
	out := new(MySQLOpsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLOpsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLOpsRequestSpec) DeepCopyInto(out *MySQLOpsRequestSpec) {
	*out = *in
	out.DatabaseRef = in.DatabaseRef
	if in.StatefulSetOrdinal != nil {
		in, out := &in.StatefulSetOrdinal, &out.StatefulSetOrdinal
		*out = new(int32)
		**out = **in
	}
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeSpec)
		**out = **in
	}
	if in.HorizontalScaling != nil {
		in, out := &in.HorizontalScaling, &out.HorizontalScaling
		*out = new(MySQLHorizontalScalingSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VerticalScaling != nil {
		in, out := &in.VerticalScaling, &out.VerticalScaling
		*out = new(MySQLVerticalScalingSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLOpsRequestSpec.
func (in *MySQLOpsRequestSpec) DeepCopy() *MySQLOpsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLOpsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLOpsRequestStatus) DeepCopyInto(out *MySQLOpsRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLOpsRequestStatus.
func (in *MySQLOpsRequestStatus) DeepCopy() *MySQLOpsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(MySQLOpsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLVerticalScalingSpec) DeepCopyInto(out *MySQLVerticalScalingSpec) {
	*out = *in
	if in.MySQL != nil {
		in, out := &in.MySQL, &out.MySQL
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Exporter != nil {
		in, out := &in.Exporter, &out.Exporter
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLVerticalScalingSpec.
func (in *MySQLVerticalScalingSpec) DeepCopy() *MySQLVerticalScalingSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLVerticalScalingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBOpsRequest) DeepCopyInto(out *PerconaXtraDBOpsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBOpsRequest.
func (in *PerconaXtraDBOpsRequest) DeepCopy() *PerconaXtraDBOpsRequest {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBOpsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBOpsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBOpsRequestList) DeepCopyInto(out *PerconaXtraDBOpsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaXtraDBOpsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBOpsRequestList.
func (in *PerconaXtraDBOpsRequestList) DeepCopy() *PerconaXtraDBOpsRequestList {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBOpsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBOpsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBOpsRequestSpec) DeepCopyInto(out *PerconaXtraDBOpsRequestSpec) {
	*out = *in
	out.DatabaseRef = in.DatabaseRef
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBOpsRequestSpec.
func (in *PerconaXtraDBOpsRequestSpec) DeepCopy() *PerconaXtraDBOpsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBOpsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBOpsRequestStatus) DeepCopyInto(out *PerconaXtraDBOpsRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBOpsRequestStatus.
func (in *PerconaXtraDBOpsRequestStatus) DeepCopy() *PerconaXtraDBOpsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBOpsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgBouncerOpsRequest) DeepCopyInto(out *PgBouncerOpsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgBouncerOpsRequest.
func (in *PgBouncerOpsRequest) DeepCopy() *PgBouncerOpsRequest {
	if in == nil {
		return nil
	}
	out := new(PgBouncerOpsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgBouncerOpsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgBouncerOpsRequestList) DeepCopyInto(out *PgBouncerOpsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PgBouncerOpsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgBouncerOpsRequestList.
func (in *PgBouncerOpsRequestList) DeepCopy() *PgBouncerOpsRequestList {
	if in == nil {
		return nil
	}
	out := new(PgBouncerOpsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PgBouncerOpsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgBouncerOpsRequestSpec) DeepCopyInto(out *PgBouncerOpsRequestSpec) {
	*out = *in
	out.DatabaseRef = in.DatabaseRef
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgBouncerOpsRequestSpec.
func (in *PgBouncerOpsRequestSpec) DeepCopy() *PgBouncerOpsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(PgBouncerOpsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgBouncerOpsRequestStatus) DeepCopyInto(out *PgBouncerOpsRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgBouncerOpsRequestStatus.
func (in *PgBouncerOpsRequestStatus) DeepCopy() *PgBouncerOpsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(PgBouncerOpsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresOpsRequest) DeepCopyInto(out *PostgresOpsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresOpsRequest.
func (in *PostgresOpsRequest) DeepCopy() *PostgresOpsRequest {
	if in == nil {
		return nil
	}
	out := new(PostgresOpsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresOpsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresOpsRequestList) DeepCopyInto(out *PostgresOpsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresOpsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresOpsRequestList.
func (in *PostgresOpsRequestList) DeepCopy() *PostgresOpsRequestList {
	if in == nil {
		return nil
	}
	out := new(PostgresOpsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresOpsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresOpsRequestSpec) DeepCopyInto(out *PostgresOpsRequestSpec) {
	*out = *in
	out.DatabaseRef = in.DatabaseRef
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresOpsRequestSpec.
func (in *PostgresOpsRequestSpec) DeepCopy() *PostgresOpsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresOpsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresOpsRequestStatus) DeepCopyInto(out *PostgresOpsRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresOpsRequestStatus.
func (in *PostgresOpsRequestStatus) DeepCopy() *PostgresOpsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresOpsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySQLOpsRequest) DeepCopyInto(out *ProxySQLOpsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySQLOpsRequest.
func (in *ProxySQLOpsRequest) DeepCopy() *ProxySQLOpsRequest {
	if in == nil {
		return nil
	}
	out := new(ProxySQLOpsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxySQLOpsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySQLOpsRequestList) DeepCopyInto(out *ProxySQLOpsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProxySQLOpsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySQLOpsRequestList.
func (in *ProxySQLOpsRequestList) DeepCopy() *ProxySQLOpsRequestList {
	if in == nil {
		return nil
	}
	out := new(ProxySQLOpsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxySQLOpsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySQLOpsRequestSpec) DeepCopyInto(out *ProxySQLOpsRequestSpec) {
	*out = *in
	out.DatabaseRef = in.DatabaseRef
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySQLOpsRequestSpec.
func (in *ProxySQLOpsRequestSpec) DeepCopy() *ProxySQLOpsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ProxySQLOpsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySQLOpsRequestStatus) DeepCopyInto(out *ProxySQLOpsRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySQLOpsRequestStatus.
func (in *ProxySQLOpsRequestStatus) DeepCopy() *ProxySQLOpsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ProxySQLOpsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisOpsRequest) DeepCopyInto(out *RedisOpsRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisOpsRequest.
func (in *RedisOpsRequest) DeepCopy() *RedisOpsRequest {
	if in == nil {
		return nil
	}
	out := new(RedisOpsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisOpsRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisOpsRequestList) DeepCopyInto(out *RedisOpsRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisOpsRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisOpsRequestList.
func (in *RedisOpsRequestList) DeepCopy() *RedisOpsRequestList {
	if in == nil {
		return nil
	}
	out := new(RedisOpsRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisOpsRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisOpsRequestSpec) DeepCopyInto(out *RedisOpsRequestSpec) {
	*out = *in
	out.DatabaseRef = in.DatabaseRef
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisOpsRequestSpec.
func (in *RedisOpsRequestSpec) DeepCopy() *RedisOpsRequestSpec {
	if in == nil {
		return nil
	}
	out := new(RedisOpsRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisOpsRequestStatus) DeepCopyInto(out *RedisOpsRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisOpsRequestStatus.
func (in *RedisOpsRequestStatus) DeepCopy() *RedisOpsRequestStatus {
	if in == nil {
		return nil
	}
	out := new(RedisOpsRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeSpec) DeepCopyInto(out *UpgradeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeSpec.
func (in *UpgradeSpec) DeepCopy() *UpgradeSpec {
	if in == nil {
		return nil
	}
	out := new(UpgradeSpec)
	in.DeepCopyInto(out)
	return out
}
