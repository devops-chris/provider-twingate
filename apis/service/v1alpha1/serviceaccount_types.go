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

// ServiceAccountParameters are the configurable fields of a ServiceAccount.
type ServiceAccountParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// ServiceAccountObservation are the observable fields of a ServiceAccount.
type ServiceAccountObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A ServiceAccountSpec defines the desired state of a ServiceAccount.
type ServiceAccountSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServiceAccountParameters `json:"forProvider"`
}

// A ServiceAccountStatus represents the observed state of a ServiceAccount.
type ServiceAccountStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServiceAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A ServiceAccount is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,twingate}
type ServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceAccountSpec   `json:"spec"`
	Status ServiceAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountList contains a list of ServiceAccount
type ServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccount `json:"items"`
}

// ServiceAccount type metadata.
var (
	ServiceAccountKind             = reflect.TypeOf(ServiceAccount{}).Name()
	ServiceAccountGroupKind        = schema.GroupKind{Group: Group, Kind: ServiceAccountKind}.String()
	ServiceAccountKindAPIVersion   = ServiceAccountKind + "." + SchemeGroupVersion.String()
	ServiceAccountGroupVersionKind = SchemeGroupVersion.WithKind(ServiceAccountKind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccount{}, &ServiceAccountList{})
}
