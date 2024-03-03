package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type ResourcePolicyType interface {
	// Experimental.
	Name() *string
}

// The jsii proxy struct for ResourcePolicyType
type jsiiProxy_ResourcePolicyType struct {
	_ byte // padding
}

func (j *jsiiProxy_ResourcePolicyType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func ResourcePolicyType_API_GATEWAY() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"API_GATEWAY",
		&returns,
	)
	return returns
}

func ResourcePolicyType_BACKUP() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"BACKUP",
		&returns,
	)
	return returns
}

func ResourcePolicyType_CODE_BUILD() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"CODE_BUILD",
		&returns,
	)
	return returns
}

func ResourcePolicyType_ECR() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"ECR",
		&returns,
	)
	return returns
}

func ResourcePolicyType_EVENTBRIDGE() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"EVENTBRIDGE",
		&returns,
	)
	return returns
}

func ResourcePolicyType_GLUE() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"GLUE",
		&returns,
	)
	return returns
}

func ResourcePolicyType_KMS() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"KMS",
		&returns,
	)
	return returns
}

func ResourcePolicyType_LAMBDA() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"LAMBDA",
		&returns,
	)
	return returns
}

func ResourcePolicyType_S3() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"S3",
		&returns,
	)
	return returns
}

func ResourcePolicyType_SECRETS_MANAGER() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"SECRETS_MANAGER",
		&returns,
	)
	return returns
}

func ResourcePolicyType_SNS() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"SNS",
		&returns,
	)
	return returns
}

func ResourcePolicyType_SQS() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"SQS",
		&returns,
	)
	return returns
}

func ResourcePolicyType_STS() ResourcePolicyType {
	_init_.Initialize()
	var returns ResourcePolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		"STS",
		&returns,
	)
	return returns
}

