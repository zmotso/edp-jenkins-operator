package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type JenkinsAuthorizationRoleSpec struct {
	Name        string   `json:"name"`
	RoleType    string   `json:"roleType"`
	Permissions []string `json:"permissions"`
	Pattern     string   `json:"pattern"`
	// +nullable
	// +optional
	OwnerName *string `json:"ownerName,omitempty"`
}

type JenkinsAuthorizationRoleStatus struct {
	Value string `json:"value"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion
type JenkinsAuthorizationRole struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec JenkinsAuthorizationRoleSpec `json:"spec,omitempty"`
	// +optional
	Status JenkinsAuthorizationRoleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
type JenkinsAuthorizationRoleList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JenkinsAuthorizationRole `json:"items"`
}
