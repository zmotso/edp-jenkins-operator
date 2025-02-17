package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// JenkinsJobSpec defines the desired state of Jenkins job
// +k8s:openapi-gen=true

type ActionType string
type Result string

const (
	RoleBinding           ActionType = "role_binding"
	CreateJenkinsPipeline ActionType = "create_jenkins_pipeline"
	TriggerJobProvision   ActionType = "trigger_job_provision"

	Success Result = "success"
	Error   Result = "error"
)

type JenkinsJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	// +nullable
	// +optional
	OwnerName *string `json:"ownerName,omitempty"`
	// +nullable
	// +optional
	StageName *string `json:"stageName,omitempty"`
	// +nullable
	// +optional
	JenkinsFolder *string `json:"jenkinsFolder,omitempty"`
	Job           Job     `json:"job"`
}

type Job struct {
	Name   string `json:"name"`
	Config string `json:"config"`
	// +nullable
	// +optional
	AutoTriggerPeriod *int32 `json:"autoTriggerPeriod,omitempty"`
}

// JenkinsFolderStatus defines the observed state of Jenkins
// +k8s:openapi-gen=true
type JenkinsJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	// +optional
	Available bool `json:"available,omitempty"`
	// +optional
	LastTimeUpdated metav1.Time `json:"lastTimeUpdated,omitempty"`
	// +optional
	Status          string     `json:"status,omitempty"`
	Username        string     `json:"username"`
	Action          ActionType `json:"action"`
	Result          Result     `json:"result"`
	DetailedMessage string     `json:"detailedMessage"`
	Value           string     `json:"value"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JenkinsFolder is the Schema for the jenkins API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:deprecatedversion
type JenkinsJob struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec JenkinsJobSpec `json:"spec,omitempty"`
	// +optional
	Status JenkinsJobStatus `json:"status,omitempty"`
}

func (jj JenkinsJob) IsAutoTriggerEnabled() bool {
	period := jj.Spec.Job.AutoTriggerPeriod
	if period == nil || *period == 0 {
		return false
	}
	if *period < 5 || *period > 7200 {
		ctrl.Log.WithName("jenkins-job-api").Info("autoTriggerPeriod value is incorrect. disable auto trigger",
			"value", *period)
		return false
	}
	return true
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JenkinsFolderList contains a list of Jenkins
type JenkinsJobList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JenkinsJob `json:"items"`
}
