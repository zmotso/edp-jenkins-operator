//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDStageJenkinsDeployment) DeepCopyInto(out *CDStageJenkinsDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDStageJenkinsDeployment.
func (in *CDStageJenkinsDeployment) DeepCopy() *CDStageJenkinsDeployment {
	if in == nil {
		return nil
	}
	out := new(CDStageJenkinsDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDStageJenkinsDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDStageJenkinsDeploymentList) DeepCopyInto(out *CDStageJenkinsDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CDStageJenkinsDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDStageJenkinsDeploymentList.
func (in *CDStageJenkinsDeploymentList) DeepCopy() *CDStageJenkinsDeploymentList {
	if in == nil {
		return nil
	}
	out := new(CDStageJenkinsDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDStageJenkinsDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDStageJenkinsDeploymentSpec) DeepCopyInto(out *CDStageJenkinsDeploymentSpec) {
	*out = *in
	out.Tag = in.Tag
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDStageJenkinsDeploymentSpec.
func (in *CDStageJenkinsDeploymentSpec) DeepCopy() *CDStageJenkinsDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(CDStageJenkinsDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDStageJenkinsDeploymentStatus) DeepCopyInto(out *CDStageJenkinsDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDStageJenkinsDeploymentStatus.
func (in *CDStageJenkinsDeploymentStatus) DeepCopy() *CDStageJenkinsDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(CDStageJenkinsDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdpSpec) DeepCopyInto(out *EdpSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdpSpec.
func (in *EdpSpec) DeepCopy() *EdpSpec {
	if in == nil {
		return nil
	}
	out := new(EdpSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jenkins) DeepCopyInto(out *Jenkins) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jenkins.
func (in *Jenkins) DeepCopy() *Jenkins {
	if in == nil {
		return nil
	}
	out := new(Jenkins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Jenkins) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAgent) DeepCopyInto(out *JenkinsAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAgent.
func (in *JenkinsAgent) DeepCopy() *JenkinsAgent {
	if in == nil {
		return nil
	}
	out := new(JenkinsAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAgentList) DeepCopyInto(out *JenkinsAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAgentList.
func (in *JenkinsAgentList) DeepCopy() *JenkinsAgentList {
	if in == nil {
		return nil
	}
	out := new(JenkinsAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAgentSpec) DeepCopyInto(out *JenkinsAgentSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAgentSpec.
func (in *JenkinsAgentSpec) DeepCopy() *JenkinsAgentSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAgentStatus) DeepCopyInto(out *JenkinsAgentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAgentStatus.
func (in *JenkinsAgentStatus) DeepCopy() *JenkinsAgentStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAuthorizationRole) DeepCopyInto(out *JenkinsAuthorizationRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAuthorizationRole.
func (in *JenkinsAuthorizationRole) DeepCopy() *JenkinsAuthorizationRole {
	if in == nil {
		return nil
	}
	out := new(JenkinsAuthorizationRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsAuthorizationRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAuthorizationRoleList) DeepCopyInto(out *JenkinsAuthorizationRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsAuthorizationRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAuthorizationRoleList.
func (in *JenkinsAuthorizationRoleList) DeepCopy() *JenkinsAuthorizationRoleList {
	if in == nil {
		return nil
	}
	out := new(JenkinsAuthorizationRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsAuthorizationRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAuthorizationRoleMapping) DeepCopyInto(out *JenkinsAuthorizationRoleMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAuthorizationRoleMapping.
func (in *JenkinsAuthorizationRoleMapping) DeepCopy() *JenkinsAuthorizationRoleMapping {
	if in == nil {
		return nil
	}
	out := new(JenkinsAuthorizationRoleMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsAuthorizationRoleMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAuthorizationRoleMappingList) DeepCopyInto(out *JenkinsAuthorizationRoleMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsAuthorizationRoleMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAuthorizationRoleMappingList.
func (in *JenkinsAuthorizationRoleMappingList) DeepCopy() *JenkinsAuthorizationRoleMappingList {
	if in == nil {
		return nil
	}
	out := new(JenkinsAuthorizationRoleMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsAuthorizationRoleMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAuthorizationRoleMappingSpec) DeepCopyInto(out *JenkinsAuthorizationRoleMappingSpec) {
	*out = *in
	if in.OwnerName != nil {
		in, out := &in.OwnerName, &out.OwnerName
		*out = new(string)
		**out = **in
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAuthorizationRoleMappingSpec.
func (in *JenkinsAuthorizationRoleMappingSpec) DeepCopy() *JenkinsAuthorizationRoleMappingSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsAuthorizationRoleMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAuthorizationRoleMappingStatus) DeepCopyInto(out *JenkinsAuthorizationRoleMappingStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAuthorizationRoleMappingStatus.
func (in *JenkinsAuthorizationRoleMappingStatus) DeepCopy() *JenkinsAuthorizationRoleMappingStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsAuthorizationRoleMappingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAuthorizationRoleSpec) DeepCopyInto(out *JenkinsAuthorizationRoleSpec) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OwnerName != nil {
		in, out := &in.OwnerName, &out.OwnerName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAuthorizationRoleSpec.
func (in *JenkinsAuthorizationRoleSpec) DeepCopy() *JenkinsAuthorizationRoleSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsAuthorizationRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAuthorizationRoleStatus) DeepCopyInto(out *JenkinsAuthorizationRoleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAuthorizationRoleStatus.
func (in *JenkinsAuthorizationRoleStatus) DeepCopy() *JenkinsAuthorizationRoleStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsAuthorizationRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsFolder) DeepCopyInto(out *JenkinsFolder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsFolder.
func (in *JenkinsFolder) DeepCopy() *JenkinsFolder {
	if in == nil {
		return nil
	}
	out := new(JenkinsFolder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsFolder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsFolderList) DeepCopyInto(out *JenkinsFolderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsFolder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsFolderList.
func (in *JenkinsFolderList) DeepCopy() *JenkinsFolderList {
	if in == nil {
		return nil
	}
	out := new(JenkinsFolderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsFolderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsFolderSpec) DeepCopyInto(out *JenkinsFolderSpec) {
	*out = *in
	if in.CodebaseName != nil {
		in, out := &in.CodebaseName, &out.CodebaseName
		*out = new(string)
		**out = **in
	}
	if in.OwnerName != nil {
		in, out := &in.OwnerName, &out.OwnerName
		*out = new(string)
		**out = **in
	}
	if in.Job != nil {
		in, out := &in.Job, &out.Job
		*out = new(Job)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsFolderSpec.
func (in *JenkinsFolderSpec) DeepCopy() *JenkinsFolderSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsFolderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsFolderStatus) DeepCopyInto(out *JenkinsFolderStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsFolderStatus.
func (in *JenkinsFolderStatus) DeepCopy() *JenkinsFolderStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsFolderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsJob) DeepCopyInto(out *JenkinsJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsJob.
func (in *JenkinsJob) DeepCopy() *JenkinsJob {
	if in == nil {
		return nil
	}
	out := new(JenkinsJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsJobBuildRun) DeepCopyInto(out *JenkinsJobBuildRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsJobBuildRun.
func (in *JenkinsJobBuildRun) DeepCopy() *JenkinsJobBuildRun {
	if in == nil {
		return nil
	}
	out := new(JenkinsJobBuildRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsJobBuildRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsJobBuildRunList) DeepCopyInto(out *JenkinsJobBuildRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsJobBuildRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsJobBuildRunList.
func (in *JenkinsJobBuildRunList) DeepCopy() *JenkinsJobBuildRunList {
	if in == nil {
		return nil
	}
	out := new(JenkinsJobBuildRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsJobBuildRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsJobBuildRunSpec) DeepCopyInto(out *JenkinsJobBuildRunSpec) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.OwnerName != nil {
		in, out := &in.OwnerName, &out.OwnerName
		*out = new(string)
		**out = **in
	}
	if in.DeleteAfterCompletionInterval != nil {
		in, out := &in.DeleteAfterCompletionInterval, &out.DeleteAfterCompletionInterval
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsJobBuildRunSpec.
func (in *JenkinsJobBuildRunSpec) DeepCopy() *JenkinsJobBuildRunSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsJobBuildRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsJobBuildRunStatus) DeepCopyInto(out *JenkinsJobBuildRunStatus) {
	*out = *in
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsJobBuildRunStatus.
func (in *JenkinsJobBuildRunStatus) DeepCopy() *JenkinsJobBuildRunStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsJobBuildRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsJobList) DeepCopyInto(out *JenkinsJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsJobList.
func (in *JenkinsJobList) DeepCopy() *JenkinsJobList {
	if in == nil {
		return nil
	}
	out := new(JenkinsJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsJobSpec) DeepCopyInto(out *JenkinsJobSpec) {
	*out = *in
	if in.OwnerName != nil {
		in, out := &in.OwnerName, &out.OwnerName
		*out = new(string)
		**out = **in
	}
	if in.StageName != nil {
		in, out := &in.StageName, &out.StageName
		*out = new(string)
		**out = **in
	}
	if in.JenkinsFolder != nil {
		in, out := &in.JenkinsFolder, &out.JenkinsFolder
		*out = new(string)
		**out = **in
	}
	in.Job.DeepCopyInto(&out.Job)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsJobSpec.
func (in *JenkinsJobSpec) DeepCopy() *JenkinsJobSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsJobStatus) DeepCopyInto(out *JenkinsJobStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsJobStatus.
func (in *JenkinsJobStatus) DeepCopy() *JenkinsJobStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsList) DeepCopyInto(out *JenkinsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Jenkins, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsList.
func (in *JenkinsList) DeepCopy() *JenkinsList {
	if in == nil {
		return nil
	}
	out := new(JenkinsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsScript) DeepCopyInto(out *JenkinsScript) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsScript.
func (in *JenkinsScript) DeepCopy() *JenkinsScript {
	if in == nil {
		return nil
	}
	out := new(JenkinsScript)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsScript) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsScriptList) DeepCopyInto(out *JenkinsScriptList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsScript, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsScriptList.
func (in *JenkinsScriptList) DeepCopy() *JenkinsScriptList {
	if in == nil {
		return nil
	}
	out := new(JenkinsScriptList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsScriptList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsScriptSpec) DeepCopyInto(out *JenkinsScriptSpec) {
	*out = *in
	if in.OwnerName != nil {
		in, out := &in.OwnerName, &out.OwnerName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsScriptSpec.
func (in *JenkinsScriptSpec) DeepCopy() *JenkinsScriptSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsScriptSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsScriptStatus) DeepCopyInto(out *JenkinsScriptStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsScriptStatus.
func (in *JenkinsScriptStatus) DeepCopy() *JenkinsScriptStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsScriptStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsServiceAccount) DeepCopyInto(out *JenkinsServiceAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsServiceAccount.
func (in *JenkinsServiceAccount) DeepCopy() *JenkinsServiceAccount {
	if in == nil {
		return nil
	}
	out := new(JenkinsServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsServiceAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsServiceAccountList) DeepCopyInto(out *JenkinsServiceAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsServiceAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsServiceAccountList.
func (in *JenkinsServiceAccountList) DeepCopy() *JenkinsServiceAccountList {
	if in == nil {
		return nil
	}
	out := new(JenkinsServiceAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsServiceAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsServiceAccountSpec) DeepCopyInto(out *JenkinsServiceAccountSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsServiceAccountSpec.
func (in *JenkinsServiceAccountSpec) DeepCopy() *JenkinsServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsServiceAccountStatus) DeepCopyInto(out *JenkinsServiceAccountStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsServiceAccountStatus.
func (in *JenkinsServiceAccountStatus) DeepCopy() *JenkinsServiceAccountStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsServiceAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsSharedLibraries) DeepCopyInto(out *JenkinsSharedLibraries) {
	*out = *in
	if in.CredentialID != nil {
		in, out := &in.CredentialID, &out.CredentialID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsSharedLibraries.
func (in *JenkinsSharedLibraries) DeepCopy() *JenkinsSharedLibraries {
	if in == nil {
		return nil
	}
	out := new(JenkinsSharedLibraries)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsSharedLibrary) DeepCopyInto(out *JenkinsSharedLibrary) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsSharedLibrary.
func (in *JenkinsSharedLibrary) DeepCopy() *JenkinsSharedLibrary {
	if in == nil {
		return nil
	}
	out := new(JenkinsSharedLibrary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsSharedLibrary) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsSharedLibraryList) DeepCopyInto(out *JenkinsSharedLibraryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsSharedLibrary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsSharedLibraryList.
func (in *JenkinsSharedLibraryList) DeepCopy() *JenkinsSharedLibraryList {
	if in == nil {
		return nil
	}
	out := new(JenkinsSharedLibraryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsSharedLibraryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsSharedLibrarySpec) DeepCopyInto(out *JenkinsSharedLibrarySpec) {
	*out = *in
	if in.OwnerName != nil {
		in, out := &in.OwnerName, &out.OwnerName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsSharedLibrarySpec.
func (in *JenkinsSharedLibrarySpec) DeepCopy() *JenkinsSharedLibrarySpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsSharedLibrarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsSharedLibraryStatus) DeepCopyInto(out *JenkinsSharedLibraryStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsSharedLibraryStatus.
func (in *JenkinsSharedLibraryStatus) DeepCopy() *JenkinsSharedLibraryStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsSharedLibraryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsSpec) DeepCopyInto(out *JenkinsSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]JenkinsVolumes, len(*in))
		copy(*out, *in)
	}
	if in.SharedLibraries != nil {
		in, out := &in.SharedLibraries, &out.SharedLibraries
		*out = make([]JenkinsSharedLibraries, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.KeycloakSpec = in.KeycloakSpec
	out.EdpSpec = in.EdpSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsSpec.
func (in *JenkinsSpec) DeepCopy() *JenkinsSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsStatus) DeepCopyInto(out *JenkinsStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
	if in.Slaves != nil {
		in, out := &in.Slaves, &out.Slaves
		*out = make([]Slave, len(*in))
		copy(*out, *in)
	}
	if in.JobProvisions != nil {
		in, out := &in.JobProvisions, &out.JobProvisions
		*out = make([]JobProvision, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsStatus.
func (in *JenkinsStatus) DeepCopy() *JenkinsStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsVolumes) DeepCopyInto(out *JenkinsVolumes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsVolumes.
func (in *JenkinsVolumes) DeepCopy() *JenkinsVolumes {
	if in == nil {
		return nil
	}
	out := new(JenkinsVolumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Job) DeepCopyInto(out *Job) {
	*out = *in
	if in.AutoTriggerPeriod != nil {
		in, out := &in.AutoTriggerPeriod, &out.AutoTriggerPeriod
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Job.
func (in *Job) DeepCopy() *Job {
	if in == nil {
		return nil
	}
	out := new(Job)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobProvision) DeepCopyInto(out *JobProvision) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobProvision.
func (in *JobProvision) DeepCopy() *JobProvision {
	if in == nil {
		return nil
	}
	out := new(JobProvision)
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
func (in *Slave) DeepCopyInto(out *Slave) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Slave.
func (in *Slave) DeepCopy() *Slave {
	if in == nil {
		return nil
	}
	out := new(Slave)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}
