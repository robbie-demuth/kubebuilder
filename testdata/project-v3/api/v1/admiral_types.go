/*
Copyright 2022 The Kubernetes authors.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AdmiralSpec defines the desired state of Admiral
type AdmiralSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Admiral. Edit admiral_types.go to remove/update
	// +optional
	Foo string `json:"foo,omitempty"`

	// Bar is an example field of Admiral. Edit admiral_types.go to remove/update
	// +optional
	Bar map[string]string `json:"bar,omitempty"`
}

// AdmiralStatus defines the observed state of Admiral
type AdmiralStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=admirales,scope=Cluster

// Admiral is the Schema for the admirales API
type Admiral struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AdmiralSpec   `json:"spec,omitempty"`
	Status AdmiralStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AdmiralList contains a list of Admiral
type AdmiralList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Admiral `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Admiral{}, &AdmiralList{})
}
