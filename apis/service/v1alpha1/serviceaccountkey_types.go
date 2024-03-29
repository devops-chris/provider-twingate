/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ServiceAccountKeyParameters are the configurable fields of a ServiceAccountKey.
type ServiceAccountKeyParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// ServiceAccountKeyObservation are the observable fields of a ServiceAccountKey.
type ServiceAccountKeyObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A ServiceAccountKeySpec defines the desired state of a ServiceAccountKey.
type ServiceAccountKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServiceAccountKeyParameters `json:"forProvider"`
}

// A ServiceAccountKeyStatus represents the observed state of a ServiceAccountKey.
type ServiceAccountKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServiceAccountKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A ServiceAccountKey is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,twingate}
type ServiceAccountKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceAccountKeySpec   `json:"spec"`
	Status ServiceAccountKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountKeyList contains a list of ServiceAccountKey
type ServiceAccountKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccountKey `json:"items"`
}

// ServiceAccountKey type metadata.
var (
	ServiceAccountKeyKind             = reflect.TypeOf(ServiceAccountKey{}).Name()
	ServiceAccountKeyGroupKind        = schema.GroupKind{Group: Group, Kind: ServiceAccountKeyKind}.String()
	ServiceAccountKeyKindAPIVersion   = ServiceAccountKeyKind + "." + SchemeGroupVersion.String()
	ServiceAccountKeyGroupVersionKind = SchemeGroupVersion.WithKind(ServiceAccountKeyKind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccountKey{}, &ServiceAccountKeyList{})
}
