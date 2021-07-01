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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1



import (
    "fmt"

    ackrtwh "github.com/aws-controllers-k8s/runtime/pkg/webhook"
    ctrlrtconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

    v2 "github.com/aws-controllers-k8s/ecr-controller/apis/v2"
)

// ConvertTo converts this Repository to the Hub version (v2).
func (src *Repository) ConvertTo(dstRaw ctrlrtconversion.Hub) error {
	dst := dstRaw.(*v2.Repository)
	if src.Spec.EncryptionConfiguration != nil {
		dst.Spec.EncryptionConfiguration.EncryptionType = src.Spec.EncryptionConfiguration.EncryptionType
		dst.Spec.EncryptionConfiguration.KMSKey = src.Spec.EncryptionConfiguration.KMSKey
	}
	dst.Spec.ImageTagMutability = src.Spec.ImageTagMutability
	dst.Spec.RepositoryName = src.Spec.RepositoryName
	if src.Spec.ScanConfig != nil {
		dst.Spec.ImageScanningConfiguration.ScanOnPushFlag = src.Spec.ImageScanningConfiguration.ScanOnPushFlag
	}
//unsupported: list
	return nil
}

// ConvertFrom converts the Hub version (v2) to this Repository.
func (dst *Repository) ConvertFrom(srcRaw ctrlrtconversion.Hub) error {
	src := srcRaw.(*v2.Repository)
	if src.Spec.EncryptionConfiguration != nil {
		dst.Spec.EncryptionConfiguration.EncryptionType = src.Spec.EncryptionConfiguration.EncryptionType
		dst.Spec.EncryptionConfiguration.KMSKey = src.Spec.EncryptionConfiguration.KMSKey
	}
	dst.Spec.ImageTagMutability = src.Spec.ImageTagMutability
	dst.Spec.RepositoryName = src.Spec.RepositoryName
	if src.Spec.ScanConfig != nil {
		dst.Spec.ImageScanningConfiguration.ScanOnPushFlag = src.Spec.ImageScanningConfiguration.ScanOnPushFlag
	}
//unsupported: list
	return nil
}

func init() {
    webhook := ackrtwh.NewWebhook(
        "conversion",
        "Repository",
        "v2",
    )
    if err := ackrtwh.RegisterWebhook(webhook); err != nil {
        msg := fmt.Sprintf("cannot register webhook: %v", err)
        panic(msg)
    }
}

// Assert convertible interface implementation Repository
var _ ctrlrtconversion.Convertible = &Repository{}