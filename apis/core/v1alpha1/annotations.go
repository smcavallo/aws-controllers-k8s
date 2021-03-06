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

package v1alpha1

const (
	// AnnotationPrefix is the prefix for all ACK annotations
	AnnotationPrefix = "services.k8s.aws/"
	// AnnotationARN is an annotation whose value is an Amazon Resource Name,
	// which is a globally-unique identifier, for the backend AWS service API
	// resource. If this annotation is SET on a CR, that means the user is
	// indicating to the ACK service controller that it should expect a backend
	// AWS service API resource to already exist (and that ACK should "adopt"
	// the resource into its management). If this annotation is NOT SET on a
	// CR, that means the user expects the ACK service controller to create the
	// backend AWS service API resource.
	AnnotationARN = AnnotationPrefix + "arn"
	// AnnotationOwnerAccountID is an annotation whoe value is the identifier
	// for the AWS account to which the resource belongs.  If this annotation
	// is set on a CR, the Kubernetes user is indicating that the ACK service
	// controller should create/patch/delete the resource in the specified AWS
	// Account. In order for this cross-account resource management to succeed,
	// the AWS IAM Role that the ACK service controller runs as needs to have
	// the ability to call the AWS STS::AssumeRole API call and assume an IAM
	// Role in the target AWS Account.
	// TODO(jaypipes): Link to documentation on cross-account resource
	// management
	AnnotationOwnerAccountID = AnnotationPrefix + "owner-account-id"
)
