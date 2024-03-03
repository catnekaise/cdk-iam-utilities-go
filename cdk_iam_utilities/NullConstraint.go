package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type NullConstraint interface {
	Constraint
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for NullConstraint
type jsiiProxy_NullConstraint struct {
	jsiiProxy_Constraint
}

// Experimental.
func NewNullConstraint(key ConditionKey, isNull *bool) NullConstraint {
	_init_.Initialize()

	if err := validateNewNullConstraintParameters(key, isNull); err != nil {
		panic(err)
	}
	j := jsiiProxy_NullConstraint{}

	_jsii_.Create(
		"@catnekaise/cdk-iam-utilities.NullConstraint",
		[]interface{}{key, isNull},
		&j,
	)

	return &j
}

// Experimental.
func NewNullConstraint_Override(n NullConstraint, key ConditionKey, isNull *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@catnekaise/cdk-iam-utilities.NullConstraint",
		[]interface{}{key, isNull},
		n,
	)
}

// Experimental.
func NullConstraint_IsNotNull(key ConditionKey) NullConstraint {
	_init_.Initialize()

	if err := validateNullConstraint_IsNotNullParameters(key); err != nil {
		panic(err)
	}
	var returns NullConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.NullConstraint",
		"isNotNull",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Experimental.
func NullConstraint_IsNull(key ConditionKey) NullConstraint {
	_init_.Initialize()

	if err := validateNullConstraint_IsNullParameters(key); err != nil {
		panic(err)
	}
	var returns NullConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.NullConstraint",
		"isNull",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NullConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
	if err := n.validateAssembleParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *[]*ConstraintPolicyMutation

	_jsii_.Invoke(
		n,
		"assemble",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NullConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
	if err := n.validateIsNotNullConditionParameters(key); err != nil {
		panic(err)
	}
	var returns *ConstraintPolicyMutation

	_jsii_.Invoke(
		n,
		"isNotNullCondition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

