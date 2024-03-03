package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type GlobalBoolConditionKey interface {
	BoolConditionKey
	// Experimental.
	Name() *string
	// Experimental.
	Settings() *ConditionKeySettings
	// Experimental.
	ToBoolFalseConstraint(ifExists *bool) BoolConstraint
	// Experimental.
	ToBoolTrueConstraint(ifExists *bool) BoolConstraint
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlobalBoolConditionKey
type jsiiProxy_GlobalBoolConditionKey struct {
	jsiiProxy_BoolConditionKey
}

func (j *jsiiProxy_GlobalBoolConditionKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalBoolConditionKey) Settings() *ConditionKeySettings {
	var returns *ConditionKeySettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


func GlobalBoolConditionKey_MultiFactorAuthPresent() GlobalBoolConditionKey {
	_init_.Initialize()
	var returns GlobalBoolConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalBoolConditionKey",
		"MultiFactorAuthPresent",
		&returns,
	)
	return returns
}

func GlobalBoolConditionKey_PrincipalIsAWSService() GlobalBoolConditionKey {
	_init_.Initialize()
	var returns GlobalBoolConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalBoolConditionKey",
		"PrincipalIsAWSService",
		&returns,
	)
	return returns
}

func GlobalBoolConditionKey_SecureTransport() GlobalBoolConditionKey {
	_init_.Initialize()
	var returns GlobalBoolConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalBoolConditionKey",
		"SecureTransport",
		&returns,
	)
	return returns
}

func GlobalBoolConditionKey_ViaAWSService() GlobalBoolConditionKey {
	_init_.Initialize()
	var returns GlobalBoolConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalBoolConditionKey",
		"ViaAWSService",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GlobalBoolConditionKey) ToBoolFalseConstraint(ifExists *bool) BoolConstraint {
	var returns BoolConstraint

	_jsii_.Invoke(
		g,
		"toBoolFalseConstraint",
		[]interface{}{ifExists},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalBoolConditionKey) ToBoolTrueConstraint(ifExists *bool) BoolConstraint {
	var returns BoolConstraint

	_jsii_.Invoke(
		g,
		"toBoolTrueConstraint",
		[]interface{}{ifExists},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalBoolConditionKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

