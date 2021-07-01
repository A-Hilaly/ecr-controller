// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package v2

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Z struct {
	S    *string
	Tags []*Tag
}

type Y struct {
	S *string
	I *int
	Z *Z
}

type X struct {
	Y *Y
	S *string
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	// The encryption configuration for the repository. This determines how the
	// contents of your repository are encrypted at rest.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`
	// The tag mutability setting for the repository. If this parameter is omitted,
	// the default setting of MUTABLE will be used which will allow image tags to
	// be overwritten. If IMMUTABLE is specified, all image tags within the repository
	// will be immutable which will prevent them from being overwritten.
	ImageTagMutability *string `json:"imageTagMutability,omitempty"`
	// The name to use for the repository. The repository name may be specified
	// on its own (such as nginx-web-app) or it can be prepended with a namespace
	// to group the repository into a category (such as project-a/nginx-web-app).
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The image scanning configuration for the repository. This determines whether
	// images are scanned for known vulnerabilities after being pushed to the repository.
	ScanConfig *ImageScanningConfiguration `json:"scanConfig,omitempty"`

	X *X
	// The metadata that you apply to the repository to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of
	// which you define. Tag keys can have a maximum character length of 128 characters,
	// and tag values can have a maximum length of 256 characters.
	Tags []*Tag `json:"tags,omitempty"`
}

// RepositoryStatus defines the observed state of Repository
type RepositoryStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The date and time, in JavaScript date format, when the repository was created.
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// The AWS account ID associated with the registry that contains the repository.
	RegistryID *string `json:"registryID,omitempty"`
	// The URI for the repository. You can use this URI for container image push
	// and pull operations.
	RepositoryURI *string `json:"repositoryURI,omitempty"`
}

// Repository is the Schema for the Repositories API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositorySpec   `json:"spec,omitempty"`
	Status            RepositoryStatus `json:"status,omitempty"`
}

// RepositoryList contains a list of Repository
// +kubebuilder:object:root=true
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}
