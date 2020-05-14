// +build !ignore_autogenerated

// Code generated by operator-sdk-2. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdpSpec) DeepCopyInto(out *EdpSpec) {
	*out = *in
	return
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
	return
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
func (in *JenkinsFolder) DeepCopyInto(out *JenkinsFolder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsFolder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
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
	if in.JobName != nil {
		in, out := &in.JobName, &out.JobName
		*out = new(string)
		**out = **in
	}
	return
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
	return
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
	return
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
func (in *JenkinsJobList) DeepCopyInto(out *JenkinsJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
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
	out.Job = in.Job
	return
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
	return
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Jenkins, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
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
	return
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsScript, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
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
	return
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
	return
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
	return
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsServiceAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
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
	return
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
	return
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
	return
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
		copy(*out, *in)
	}
	out.KeycloakSpec = in.KeycloakSpec
	out.EdpSpec = in.EdpSpec
	return
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
	if in.CiJobProvisions != nil {
		in, out := &in.CiJobProvisions, &out.CiJobProvisions
		*out = make([]JobProvision, len(*in))
		copy(*out, *in)
	}
	return
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
	return
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
	return
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
	return
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
	return
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
	return
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
