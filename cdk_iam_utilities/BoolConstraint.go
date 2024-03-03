package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type BoolConstraint interface {
	Constraint
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for BoolConstraint
type jsiiProxy_BoolConstraint struct {
	jsiiProxy_Constraint
}

// Experimental.
func BoolConstraint_WhenFalse(key ConditionKey, ifExists *bool) BoolConstraint {
	_init_.Initialize()

	if err := validateBoolConstraint_WhenFalseParameters(key); err != nil {
		panic(err)
	}
	var returns BoolConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.BoolConstraint",
		"whenFalse",
		[]interface{}{key, ifExists},
		&returns,
	)

	return returns
}

// Experimental.
func BoolConstraint_WhenTrue(key ConditionKey, ifExists *bool) BoolConstraint {
	_init_.Initialize()

	if err := validateBoolConstraint_WhenTrueParameters(key); err != nil {
		panic(err)
	}
	var returns BoolConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.BoolConstraint",
		"whenTrue",
		[]interface{}{key, ifExists},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BoolConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
	if err := b.validateAssembleParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *[]*ConstraintPolicyMutation

	_jsii_.Invoke(
		b,
		"assemble",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BoolConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
	if err := b.validateIsNotNullConditionParameters(key); err != nil {
		panic(err)
	}
	var returns *ConstraintPolicyMutation

	_jsii_.Invoke(
		b,
		"isNotNullCondition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

