package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type AwsResourceTagConditionKey interface {
	ConditionKey
	// Experimental.
	Name() *string
	// Experimental.
	Settings() *ConditionKeySettings
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsResourceTagConditionKey
type jsiiProxy_AwsResourceTagConditionKey struct {
	jsiiProxy_ConditionKey
}

func (j *jsiiProxy_AwsResourceTagConditionKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResourceTagConditionKey) Settings() *ConditionKeySettings {
	var returns *ConditionKeySettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func AwsResourceTagConditionKey_Tag(tagName *string) AwsResourceTagConditionKey {
	_init_.Initialize()

	if err := validateAwsResourceTagConditionKey_TagParameters(tagName); err != nil {
		panic(err)
	}
	var returns AwsResourceTagConditionKey

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.AwsResourceTagConditionKey",
		"tag",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResourceTagConditionKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

