/*
Copyright 2023 Eirik Opheim.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SecretmapperSpec defines the desired state of Secretmapper
type SecretmapperSpec struct {
	source      NamespacedName `json:"source"`
	destination NamespacedName `json:"destination"`
}

type NamespacedName struct {
	namespace string `json:"namespace"`
	name      string `json:"name"`
}

// SecretmapperStatus defines the observed state of Secretmapper
type SecretmapperStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Secretmapper is the Schema for the secretmappers API
type Secretmapper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretmapperSpec   `json:"spec,omitempty"`
	Status SecretmapperStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SecretmapperList contains a list of Secretmapper
type SecretmapperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secretmapper `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Secretmapper{}, &SecretmapperList{})
}
