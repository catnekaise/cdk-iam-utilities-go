package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type AwsFederatedProviderConditionKey interface {
	ConditionKey
	// Experimental.
	Name() *string
	// Experimental.
	Settings() *ConditionKeySettings
	// Experimental.
	ToCognitoIdentityConstraint() Constraint
	// Experimental.
	ToConstraint(operator ConditionOperator, value *string, additionalValues ...*string) GenericConstraint
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFederatedProviderConditionKey
type jsiiProxy_AwsFederatedProviderConditionKey struct {
	jsiiProxy_ConditionKey
}

func (j *jsiiProxy_AwsFederatedProviderConditionKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFederatedProviderConditionKey) Settings() *ConditionKeySettings {
	var returns *ConditionKeySettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


func AwsFederatedProviderConditionKey_Create() AwsFederatedProviderConditionKey {
	_init_.Initialize()
	var returns AwsFederatedProviderConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.AwsFederatedProviderConditionKey",
		"Create",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsFederatedProviderConditionKey) ToCognitoIdentityConstraint() Constraint {
	var returns Constraint

	_jsii_.Invoke(
		a,
		"toCognitoIdentityConstraint",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFederatedProviderConditionKey) ToConstraint(operator ConditionOperator, value *string, additionalValues ...*string) GenericConstraint {
	if err := a.validateToConstraintParameters(operator, value); err != nil {
		panic(err)
	}
	args := []interface{}{operator, value}
	for _, a := range additionalValues {
		args = append(args, a)
	}

	var returns GenericConstraint

	_jsii_.Invoke(
		a,
		"toConstraint",
		args,
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFederatedProviderConditionKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

