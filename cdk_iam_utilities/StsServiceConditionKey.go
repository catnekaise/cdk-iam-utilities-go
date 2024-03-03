package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type StsServiceConditionKey interface {
	ConditionKey
	// Experimental.
	Name() *string
	// Experimental.
	Settings() *ConditionKeySettings
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StsServiceConditionKey
type jsiiProxy_StsServiceConditionKey struct {
	jsiiProxy_ConditionKey
}

func (j *jsiiProxy_StsServiceConditionKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StsServiceConditionKey) Settings() *ConditionKeySettings {
	var returns *ConditionKeySettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Filters access by the tags that are attached to the role that is being assumed.
// Experimental.
func StsServiceConditionKey_IamResourceTag(tagName *string) IamResourceTagConditionKey {
	_init_.Initialize()

	if err := validateStsServiceConditionKey_IamResourceTagParameters(tagName); err != nil {
		panic(err)
	}
	var returns IamResourceTagConditionKey

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.StsServiceConditionKey",
		"iamResourceTag",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

func StsServiceConditionKey_AWSServiceName() StsServiceConditionKey {
	_init_.Initialize()
	var returns StsServiceConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.StsServiceConditionKey",
		"AWSServiceName",
		&returns,
	)
	return returns
}

func StsServiceConditionKey_CognitoIdentityAmr() StsServiceConditionKey {
	_init_.Initialize()
	var returns StsServiceConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.StsServiceConditionKey",
		"CognitoIdentityAmr",
		&returns,
	)
	return returns
}

func StsServiceConditionKey_CognitoIdentityAud() StsServiceConditionKey {
	_init_.Initialize()
	var returns StsServiceConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.StsServiceConditionKey",
		"CognitoIdentityAud",
		&returns,
	)
	return returns
}

func StsServiceConditionKey_CognitoIdentitySub() StsServiceConditionKey {
	_init_.Initialize()
	var returns StsServiceConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.StsServiceConditionKey",
		"CognitoIdentitySub",
		&returns,
	)
	return returns
}

func StsServiceConditionKey_DurationSeconds() StsServiceConditionKey {
	_init_.Initialize()
	var returns StsServiceConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.StsServiceConditionKey",
		"DurationSeconds",
		&returns,
	)
	return returns
}

func StsServiceConditionKey_ExternalId() StsServiceConditionKey {
	_init_.Initialize()
	var returns StsServiceConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.StsServiceConditionKey",
		"ExternalId",
		&returns,
	)
	return returns
}

func StsServiceConditionKey_RoleSessionName() StsServiceConditionKey {
	_init_.Initialize()
	var returns StsServiceConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.StsServiceConditionKey",
		"RoleSessionName",
		&returns,
	)
	return returns
}

func StsServiceConditionKey_SourceIdentity() StsServiceConditionKey {
	_init_.Initialize()
	var returns StsServiceConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.StsServiceConditionKey",
		"SourceIdentity",
		&returns,
	)
	return returns
}

func StsServiceConditionKey_TransitiveTagKeys() StsServiceConditionKey {
	_init_.Initialize()
	var returns StsServiceConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.StsServiceConditionKey",
		"TransitiveTagKeys",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StsServiceConditionKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

