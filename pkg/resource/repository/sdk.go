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

package repository

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/ecr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/ecr-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ECR{}
	_ = &svcapitypes.Repository{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.DescribeRepositoriesWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeRepositories", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "RepositoryNotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.Repositories {
		if elem.CreatedAt != nil {
			ko.Status.CreatedAt = &metav1.Time{*elem.CreatedAt}
		} else {
			ko.Status.CreatedAt = nil
		}
		if elem.EncryptionConfiguration != nil {
			f1 := &svcapitypes.EncryptionConfiguration{}
			if elem.EncryptionConfiguration.EncryptionType != nil {
				f1.EncryptionType = elem.EncryptionConfiguration.EncryptionType
			}
			if elem.EncryptionConfiguration.KmsKey != nil {
				f1.KMSKey = elem.EncryptionConfiguration.KmsKey
			}
			ko.Spec.EncryptionConfiguration = f1
		} else {
			ko.Spec.EncryptionConfiguration = nil
		}
		if elem.ImageScanningConfiguration != nil {
			f2 := &svcapitypes.ImageScanningConfiguration{}
			if elem.ImageScanningConfiguration.ScanOnPush != nil {
				f2.ScanOnPush = elem.ImageScanningConfiguration.ScanOnPush
			}
			ko.Spec.ImageScanningConfiguration = f2
		} else {
			ko.Spec.ImageScanningConfiguration = nil
		}
		if elem.ImageTagMutability != nil {
			ko.Spec.ImageTagMutability = elem.ImageTagMutability
		} else {
			ko.Spec.ImageTagMutability = nil
		}
		if elem.RegistryId != nil {
			ko.Status.RegistryID = elem.RegistryId
		} else {
			ko.Status.RegistryID = nil
		}
		if elem.RepositoryArn != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.RepositoryArn)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.RepositoryUri != nil {
			ko.Status.RepositoryURI = elem.RepositoryUri
		} else {
			ko.Status.RepositoryURI = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeRepositoriesInput, error) {
	res := &svcsdk.DescribeRepositoriesInput{}

	if r.ko.Status.RegistryID != nil {
		res.SetRegistryId(*r.ko.Status.RegistryID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(ctx, r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateRepositoryWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateRepository", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Repository.CreatedAt != nil {
		ko.Status.CreatedAt = &metav1.Time{*resp.Repository.CreatedAt}
	} else {
		ko.Status.CreatedAt = nil
	}
	if resp.Repository.RegistryId != nil {
		ko.Status.RegistryID = resp.Repository.RegistryId
	} else {
		ko.Status.RegistryID = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Repository.RepositoryArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Repository.RepositoryArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Repository.RepositoryUri != nil {
		ko.Status.RepositoryURI = resp.Repository.RepositoryUri
	} else {
		ko.Status.RepositoryURI = nil
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateRepositoryInput, error) {
	res := &svcsdk.CreateRepositoryInput{}

	if r.ko.Spec.EncryptionConfiguration != nil {
		f0 := &svcsdk.EncryptionConfiguration{}
		if r.ko.Spec.EncryptionConfiguration.EncryptionType != nil {
			f0.SetEncryptionType(*r.ko.Spec.EncryptionConfiguration.EncryptionType)
		}
		if r.ko.Spec.EncryptionConfiguration.KMSKey != nil {
			f0.SetKmsKey(*r.ko.Spec.EncryptionConfiguration.KMSKey)
		}
		res.SetEncryptionConfiguration(f0)
	}
	if r.ko.Spec.ImageScanningConfiguration != nil {
		f1 := &svcsdk.ImageScanningConfiguration{}
		if r.ko.Spec.ImageScanningConfiguration.ScanOnPush != nil {
			f1.SetScanOnPush(*r.ko.Spec.ImageScanningConfiguration.ScanOnPush)
		}
		res.SetImageScanningConfiguration(f1)
	}
	if r.ko.Spec.ImageTagMutability != nil {
		res.SetImageTagMutability(*r.ko.Spec.ImageTagMutability)
	}
	if r.ko.Spec.Name != nil {
		res.SetRepositoryName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.Tags != nil {
		f4 := []*svcsdk.Tag{}
		for _, f4iter := range r.ko.Spec.Tags {
			f4elem := &svcsdk.Tag{}
			if f4iter.Key != nil {
				f4elem.SetKey(*f4iter.Key)
			}
			if f4iter.Value != nil {
				f4elem.SetValue(*f4iter.Value)
			}
			f4 = append(f4, f4elem)
		}
		res.SetTags(f4)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	return rm.customUpdateRepository(ctx, desired, latest, delta)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {

	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteRepositoryWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteRepository", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteRepositoryInput, error) {
	res := &svcsdk.DeleteRepositoryInput{}

	if r.ko.Status.RegistryID != nil {
		res.SetRegistryId(*r.ko.Status.RegistryID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Repository,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Message()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	if terminalCondition != nil || recoverableCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
