package v1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	JobBuildRunStatusCreated   = "created"
	JobBuildRunStatusCompleted = "completed"
	JobBuildRunStatusFailed    = "failed"
	JobBuildRunStatusRetrying  = "retrying"
	JobBuildRunStatusNotFound  = "jobNotFound"
)

type JenkinsJobBuildRunSpec struct {
	JobPath string `json:"jobpath"`
	// +nullable
	// +optional
	Params map[string]string `json:"params,omitempty"`
	Retry  int               `json:"retry"`
	// +nullable
	// +optional
	OwnerName *string `json:"ownerName,omitempty"`
	// +nullable
	// +optional
	DeleteAfterCompletionInterval *string `json:"deleteAfterCompletionInterval,omitempty"`
}

func (in *JenkinsJobBuildRun) GetDeleteAfterCompletionInterval() time.Duration {
	if in.Spec.DeleteAfterCompletionInterval == nil {
		return time.Hour
	}

	dur, err := time.ParseDuration(*in.Spec.DeleteAfterCompletionInterval)
	if err != nil {
		return time.Hour
	}

	return dur
}

type JenkinsJobBuildRunStatus struct {
	Status      string      `json:"status"`
	Launches    int         `json:"launches"`
	BuildNumber int64       `json:"buildNumber"`
	LastUpdated metav1.Time `json:"lastUpdated"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion
type JenkinsJobBuildRun struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec JenkinsJobBuildRunSpec `json:"spec,omitempty"`
	// +optional
	Status JenkinsJobBuildRunStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
type JenkinsJobBuildRunList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JenkinsJobBuildRun `json:"items"`
}
