/*
Copyright 2024.

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

// UnifiNetworkSpec defines the desired state of UnifiNetwork
type UnifiNetworkSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of UnifiNetwork. Edit unifinetwork_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// UnifiNetworkStatus defines the observed state of UnifiNetwork
type UnifiNetworkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UnifiNetwork is the Schema for the unifinetworks API
type UnifiNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UnifiNetworkSpec   `json:"spec,omitempty"`
	Status UnifiNetworkStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UnifiNetworkList contains a list of UnifiNetwork
type UnifiNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UnifiNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UnifiNetwork{}, &UnifiNetworkList{})
}
