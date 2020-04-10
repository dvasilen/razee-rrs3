package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +k8s:openapi-gen=true
type RemoteResourceS3 struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec Spec `json:"spec,omitempty"`
}

// +k8s:openapi-gen=true
type SecretKeyRef struct {
	Name      string `json:"name"`
	Key       string `json:"key"`
	Namespace string `json:"namespace"`
}

// +k8s:openapi-gen=true
type ValueFrom struct {
	SecretKeyRef SecretKeyRef `json:"secretKeyRef"`
}

// +k8s:openapi-gen=true
type APIKeyRef struct {
	ValueFrom ValueFrom `json:"valueFrom"`
}

// +k8s:openapi-gen=true
type Iam struct {
	ResponseType string    `json:"responseType"`
	GrantType    string    `json:"grantType"`
	URL          string    `json:"url"`
	APIKeyRef    APIKeyRef `json:"apiKeyRef"`
}

// +k8s:openapi-gen=true
type Auth struct {
	Iam Iam `json:"iam"`
}

// +k8s:openapi-gen=true
type Options struct {
	URL string `json:"url"`
}

// +k8s:openapi-gen=true
type Requests struct {
	Options Options `json:"options"`
}

// +k8s:openapi-gen=true
type Spec struct {
	Auth     Auth       `json:"auth"`
	Requests []Requests `json:"requests"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RemoteResourceS3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RemoteResourceS3 `json:"items"`
}
