package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type AwsPrincipalTagConditionKey interface {
	ConditionKey
	// Experimental.
	Name() *string
	// Experimental.
	Settings() *ConditionKeySettings
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsPrincipalTagConditionKey
type jsiiProxy_AwsPrincipalTagConditionKey struct {
	jsiiProxy_ConditionKey
}

func (j *jsiiProxy_AwsPrincipalTagConditionKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPrincipalTagConditionKey) Settings() *ConditionKeySettings {
	var returns *ConditionKeySettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func AwsPrincipalTagConditionKey_Tag(tagName *string) AwsPrincipalTagConditionKey {
	_init_.Initialize()

	if err := validateAwsPrincipalTagConditionKey_TagParameters(tagName); err != nil {
		panic(err)
	}
	var returns AwsPrincipalTagConditionKey

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.AwsPrincipalTagConditionKey",
		"tag",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPrincipalTagConditionKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

