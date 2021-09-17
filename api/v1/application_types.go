/*
Copyright 2021.

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
	"github.com/operator-framework/operator-lib/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApplicationKind is kind string
const (
	ApplicationKind = "Applications"
)

// Condition keys for IApplication
const (
	ApplicationConditionReady             = status.ConditionType("ready")
	ApplicationConditionWebhookRegistered = status.ConditionType("webhook-registered")
)

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	// Git config for target repository
	Git GitConfig `json:"git"`
}

// ApplicationStatus defines the observed state of Application
type ApplicationStatus struct {
	// Conditions of IntegrationConfig
	Conditions status.Conditions `json:"conditions"`
	Secrets    string            `json:"secrets,omitempty"` // TODO 왜 필요해?
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Application is the Schema for the applications API
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
