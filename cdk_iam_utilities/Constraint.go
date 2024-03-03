package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type Constraint interface {
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for Constraint
type jsiiProxy_Constraint struct {
	_ byte // padding
}

// Experimental.
func NewConstraint_Override(c Constraint) {
	_init_.Initialize()

	_jsii_.Create(
		"@catnekaise/cdk-iam-utilities.Constraint",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_Constraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
	if err := c.validateAssembleParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *[]*ConstraintPolicyMutation

	_jsii_.Invoke(
		c,
		"assemble",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Constraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
	if err := c.validateIsNotNullConditionParameters(key); err != nil {
		panic(err)
	}
	var returns *ConstraintPolicyMutation

	_jsii_.Invoke(
		c,
		"isNotNullCondition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

