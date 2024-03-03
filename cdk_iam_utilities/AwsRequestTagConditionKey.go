package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type AwsRequestTagConditionKey interface {
	ConditionKey
	// Experimental.
	Name() *string
	// Experimental.
	Settings() *ConditionKeySettings
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsRequestTagConditionKey
type jsiiProxy_AwsRequestTagConditionKey struct {
	jsiiProxy_ConditionKey
}

func (j *jsiiProxy_AwsRequestTagConditionKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRequestTagConditionKey) Settings() *ConditionKeySettings {
	var returns *ConditionKeySettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func AwsRequestTagConditionKey_Tag(tagName *string) AwsRequestTagConditionKey {
	_init_.Initialize()

	if err := validateAwsRequestTagConditionKey_TagParameters(tagName); err != nil {
		panic(err)
	}
	var returns AwsRequestTagConditionKey

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.AwsRequestTagConditionKey",
		"tag",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRequestTagConditionKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

