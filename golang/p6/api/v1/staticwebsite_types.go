package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type StaticWebsiteSpec struct {
	Image    string `json:"image"`
	Replicas *int32 `json:"replicas,omitempty"`
	Port     int32  `json:"port,omitempty"`
}

type StaticWebsiteStatus struct {
	Phase          string `json:"phase,omitempty"`
	Message        string `json:"message,omitempty"`
	DeploymentName string `json:"deploymentName,omitempty"`
	ServiceName    string `json:"serviceName,omitempty"`
	ReadyReplicas  int32  `json:"readyReplicas,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// StaticWebsite is the Schema for the staticwebsites API
type StaticWebsite struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`
	// +required
	Spec StaticWebsiteSpec `json:"spec"`
	// +optional
	Status StaticWebsiteStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// StaticWebsiteList contains a list of StaticWebsite
type StaticWebsiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []StaticWebsite `json:"items"`
}

func init() {
	SchemeBuilder.Register(func(s *runtime.Scheme) error {
		s.AddKnownTypes(SchemeGroupVersion, &StaticWebsite{}, &StaticWebsiteList{})
		return nil
	})
}
