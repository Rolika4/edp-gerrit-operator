//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gerrit) DeepCopyInto(out *Gerrit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gerrit.
func (in *Gerrit) DeepCopy() *Gerrit {
	if in == nil {
		return nil
	}
	out := new(Gerrit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gerrit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritGroup) DeepCopyInto(out *GerritGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritGroup.
func (in *GerritGroup) DeepCopy() *GerritGroup {
	if in == nil {
		return nil
	}
	out := new(GerritGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritGroupList) DeepCopyInto(out *GerritGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GerritGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritGroupList.
func (in *GerritGroupList) DeepCopy() *GerritGroupList {
	if in == nil {
		return nil
	}
	out := new(GerritGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritGroupMember) DeepCopyInto(out *GerritGroupMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritGroupMember.
func (in *GerritGroupMember) DeepCopy() *GerritGroupMember {
	if in == nil {
		return nil
	}
	out := new(GerritGroupMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritGroupMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritGroupMemberList) DeepCopyInto(out *GerritGroupMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GerritGroupMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritGroupMemberList.
func (in *GerritGroupMemberList) DeepCopy() *GerritGroupMemberList {
	if in == nil {
		return nil
	}
	out := new(GerritGroupMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritGroupMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritGroupMemberSpec) DeepCopyInto(out *GerritGroupMemberSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritGroupMemberSpec.
func (in *GerritGroupMemberSpec) DeepCopy() *GerritGroupMemberSpec {
	if in == nil {
		return nil
	}
	out := new(GerritGroupMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritGroupMemberStatus) DeepCopyInto(out *GerritGroupMemberStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritGroupMemberStatus.
func (in *GerritGroupMemberStatus) DeepCopy() *GerritGroupMemberStatus {
	if in == nil {
		return nil
	}
	out := new(GerritGroupMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritGroupSpec) DeepCopyInto(out *GerritGroupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritGroupSpec.
func (in *GerritGroupSpec) DeepCopy() *GerritGroupSpec {
	if in == nil {
		return nil
	}
	out := new(GerritGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritGroupStatus) DeepCopyInto(out *GerritGroupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritGroupStatus.
func (in *GerritGroupStatus) DeepCopy() *GerritGroupStatus {
	if in == nil {
		return nil
	}
	out := new(GerritGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritList) DeepCopyInto(out *GerritList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Gerrit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritList.
func (in *GerritList) DeepCopy() *GerritList {
	if in == nil {
		return nil
	}
	out := new(GerritList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritMergeRequest) DeepCopyInto(out *GerritMergeRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritMergeRequest.
func (in *GerritMergeRequest) DeepCopy() *GerritMergeRequest {
	if in == nil {
		return nil
	}
	out := new(GerritMergeRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritMergeRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritMergeRequestList) DeepCopyInto(out *GerritMergeRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GerritMergeRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritMergeRequestList.
func (in *GerritMergeRequestList) DeepCopy() *GerritMergeRequestList {
	if in == nil {
		return nil
	}
	out := new(GerritMergeRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritMergeRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritMergeRequestSpec) DeepCopyInto(out *GerritMergeRequestSpec) {
	*out = *in
	if in.AdditionalArguments != nil {
		in, out := &in.AdditionalArguments, &out.AdditionalArguments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritMergeRequestSpec.
func (in *GerritMergeRequestSpec) DeepCopy() *GerritMergeRequestSpec {
	if in == nil {
		return nil
	}
	out := new(GerritMergeRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritMergeRequestStatus) DeepCopyInto(out *GerritMergeRequestStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritMergeRequestStatus.
func (in *GerritMergeRequestStatus) DeepCopy() *GerritMergeRequestStatus {
	if in == nil {
		return nil
	}
	out := new(GerritMergeRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritProject) DeepCopyInto(out *GerritProject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritProject.
func (in *GerritProject) DeepCopy() *GerritProject {
	if in == nil {
		return nil
	}
	out := new(GerritProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritProject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritProjectAccess) DeepCopyInto(out *GerritProjectAccess) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritProjectAccess.
func (in *GerritProjectAccess) DeepCopy() *GerritProjectAccess {
	if in == nil {
		return nil
	}
	out := new(GerritProjectAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritProjectAccess) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritProjectAccessList) DeepCopyInto(out *GerritProjectAccessList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GerritProjectAccess, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritProjectAccessList.
func (in *GerritProjectAccessList) DeepCopy() *GerritProjectAccessList {
	if in == nil {
		return nil
	}
	out := new(GerritProjectAccessList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritProjectAccessList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritProjectAccessSpec) DeepCopyInto(out *GerritProjectAccessSpec) {
	*out = *in
	if in.References != nil {
		in, out := &in.References, &out.References
		*out = make([]Reference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritProjectAccessSpec.
func (in *GerritProjectAccessSpec) DeepCopy() *GerritProjectAccessSpec {
	if in == nil {
		return nil
	}
	out := new(GerritProjectAccessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritProjectAccessStatus) DeepCopyInto(out *GerritProjectAccessStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritProjectAccessStatus.
func (in *GerritProjectAccessStatus) DeepCopy() *GerritProjectAccessStatus {
	if in == nil {
		return nil
	}
	out := new(GerritProjectAccessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritProjectList) DeepCopyInto(out *GerritProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GerritProject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritProjectList.
func (in *GerritProjectList) DeepCopy() *GerritProjectList {
	if in == nil {
		return nil
	}
	out := new(GerritProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritProjectSpec) DeepCopyInto(out *GerritProjectSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritProjectSpec.
func (in *GerritProjectSpec) DeepCopy() *GerritProjectSpec {
	if in == nil {
		return nil
	}
	out := new(GerritProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritProjectStatus) DeepCopyInto(out *GerritProjectStatus) {
	*out = *in
	if in.Branches != nil {
		in, out := &in.Branches, &out.Branches
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritProjectStatus.
func (in *GerritProjectStatus) DeepCopy() *GerritProjectStatus {
	if in == nil {
		return nil
	}
	out := new(GerritProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritReplicationConfig) DeepCopyInto(out *GerritReplicationConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritReplicationConfig.
func (in *GerritReplicationConfig) DeepCopy() *GerritReplicationConfig {
	if in == nil {
		return nil
	}
	out := new(GerritReplicationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritReplicationConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritReplicationConfigList) DeepCopyInto(out *GerritReplicationConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GerritReplicationConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritReplicationConfigList.
func (in *GerritReplicationConfigList) DeepCopy() *GerritReplicationConfigList {
	if in == nil {
		return nil
	}
	out := new(GerritReplicationConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GerritReplicationConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritReplicationConfigSpec) DeepCopyInto(out *GerritReplicationConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritReplicationConfigSpec.
func (in *GerritReplicationConfigSpec) DeepCopy() *GerritReplicationConfigSpec {
	if in == nil {
		return nil
	}
	out := new(GerritReplicationConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritReplicationConfigStatus) DeepCopyInto(out *GerritReplicationConfigStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritReplicationConfigStatus.
func (in *GerritReplicationConfigStatus) DeepCopy() *GerritReplicationConfigStatus {
	if in == nil {
		return nil
	}
	out := new(GerritReplicationConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritSpec) DeepCopyInto(out *GerritSpec) {
	*out = *in
	out.KeycloakSpec = in.KeycloakSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritSpec.
func (in *GerritSpec) DeepCopy() *GerritSpec {
	if in == nil {
		return nil
	}
	out := new(GerritSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GerritStatus) DeepCopyInto(out *GerritStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GerritStatus.
func (in *GerritStatus) DeepCopy() *GerritStatus {
	if in == nil {
		return nil
	}
	out := new(GerritStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakSpec) DeepCopyInto(out *KeycloakSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakSpec.
func (in *KeycloakSpec) DeepCopy() *KeycloakSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reference) DeepCopyInto(out *Reference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reference.
func (in *Reference) DeepCopy() *Reference {
	if in == nil {
		return nil
	}
	out := new(Reference)
	in.DeepCopyInto(out)
	return out
}
