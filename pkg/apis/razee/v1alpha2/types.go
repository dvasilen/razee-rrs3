/**
 * Copyright 2019 IBM Corp. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
	Iam  Iam  `json:"iam,omitempty"`
	Hmac Hmac `json:"hmac,omitempty"`
}

// +k8s:openapi-gen=true
type Hmac struct {
	AccessKeyID     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
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
