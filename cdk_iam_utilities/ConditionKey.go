package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type ConditionKey interface {
	// Experimental.
	Name() *string
	// Experimental.
	Settings() *ConditionKeySettings
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConditionKey
type jsiiProxy_ConditionKey struct {
	_ byte // padding
}

func (j *jsiiProxy_ConditionKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionKey) Settings() *ConditionKeySettings {
	var returns *ConditionKeySettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func NewConditionKey_Override(c ConditionKey, name *string, settings *ConditionKeySettings) {
	_init_.Initialize()

	_jsii_.Create(
		"@catnekaise/cdk-iam-utilities.ConditionKey",
		[]interface{}{name, settings},
		c,
	)
}

func (c *jsiiProxy_ConditionKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

