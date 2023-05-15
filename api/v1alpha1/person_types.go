/*
Copyright 2023.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// PersonSpec defines the desired state of Person
type PersonSpec struct {
	FirstName  string `json:"firstName,omitempty"`
	MiddleName string `json:"middleName,omitempty"`
	LastName   string `json:"lastName,omitempty"`

	// Groups this person belongs to (eg: teams, roles, streams...)
	Groups []string `json:"groups,omitempty"`

	// Custom values associated with person.
	Values runtime.RawExtension `json:"values,omitempty"`
}

// PersonStatus defines the observed state of Person
type PersonStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Person is the Schema for the people API
type Person struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PersonSpec   `json:"spec,omitempty"`
	Status PersonStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PersonList contains a list of Person
type PersonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Person `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Person{}, &PersonList{})
}
