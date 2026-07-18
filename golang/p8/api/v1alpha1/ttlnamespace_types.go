package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TTLNamespaceSpec struct {
	TargetNamespace string `json:"targetNamespace,omitempty"`
	TTL             string `json:"ttl,omitempty"`
}

// +kubebuilder:object:generate=true
type TTLNamespaceStatus struct {
	ExpiresAt *metav1.Time `json:"expiresAt,omitempty"`
	Phase     string       `json:"phase,omitempty"`
	Message   string       `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type TTLNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TTLNamespaceSpec   `json:"spec,omitempty"`
	Status TTLNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type TTLNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []TTLNamespace `json:"items"`
}
