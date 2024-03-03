package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type IamResourceTagConditionKey interface {
	ConditionKey
	// Experimental.
	Name() *string
	// Experimental.
	Settings() *ConditionKeySettings
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamResourceTagConditionKey
type jsiiProxy_IamResourceTagConditionKey struct {
	jsiiProxy_ConditionKey
}

func (j *jsiiProxy_IamResourceTagConditionKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamResourceTagConditionKey) Settings() *ConditionKeySettings {
	var returns *ConditionKeySettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func IamResourceTagConditionKey_Tag(tagName *string) IamResourceTagConditionKey {
	_init_.Initialize()

	if err := validateIamResourceTagConditionKey_TagParameters(tagName); err != nil {
		panic(err)
	}
	var returns IamResourceTagConditionKey

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.IamResourceTagConditionKey",
		"tag",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamResourceTagConditionKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

