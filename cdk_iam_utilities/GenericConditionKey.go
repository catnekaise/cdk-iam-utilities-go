package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type GenericConditionKey interface {
	ConditionKey
	// Experimental.
	Name() *string
	// Experimental.
	Settings() *ConditionKeySettings
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GenericConditionKey
type jsiiProxy_GenericConditionKey struct {
	jsiiProxy_ConditionKey
}

func (j *jsiiProxy_GenericConditionKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenericConditionKey) Settings() *ConditionKeySettings {
	var returns *ConditionKeySettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func GenericConditionKey_Custom(name *string, settings *ConditionKeySettings) GenericConditionKey {
	_init_.Initialize()

	if err := validateGenericConditionKey_CustomParameters(name, settings); err != nil {
		panic(err)
	}
	var returns GenericConditionKey

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.GenericConditionKey",
		"custom",
		[]interface{}{name, settings},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericConditionKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

