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

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	// Full display name of the group.
	FullName string `json:"fullName,omitempty"`

	// Optional email address of the group.
	Email string `json:"email,omitempty"`

	// Optional emoji to picture group in a more visual way.
	Emoji string `json:"emoji,omitempty"`

	// Type of group (eg: team, role, stream...)
	Type string `json:"type,omitempty"`

	// Parent group that will be inherited by all persons belonging to this group and its subgroups.
	Parent string `json:"parent,omitempty"`

	// Arbitrary custom values associated with group.
	Values runtime.RawExtension `json:"values,omitempty"`
}

//+kubebuilder:object:root=true

// Group is the Schema for the groups API
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec GroupSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// GroupList contains a list of Group
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
