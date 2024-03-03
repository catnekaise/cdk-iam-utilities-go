package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type TagConstraint interface {
	Constraint
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for TagConstraint
type jsiiProxy_TagConstraint struct {
	jsiiProxy_Constraint
}

// Experimental.
func NewTagConstraint_Override(t TagConstraint) {
	_init_.Initialize()

	_jsii_.Create(
		"@catnekaise/cdk-iam-utilities.TagConstraint",
		nil, // no parameters
		t,
	)
}

func (t *jsiiProxy_TagConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
	if err := t.validateAssembleParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *[]*ConstraintPolicyMutation

	_jsii_.Invoke(
		t,
		"assemble",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
	if err := t.validateIsNotNullConditionParameters(key); err != nil {
		panic(err)
	}
	var returns *ConstraintPolicyMutation

	_jsii_.Invoke(
		t,
		"isNotNullCondition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

