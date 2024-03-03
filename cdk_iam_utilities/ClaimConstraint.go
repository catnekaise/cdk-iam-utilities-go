package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type ClaimConstraint interface {
	Constraint
	// Experimental.
	Claim() *string
	// Experimental.
	Operator() ConditionOperator
	// Experimental.
	Values() *[]*string
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for ClaimConstraint
type jsiiProxy_ClaimConstraint struct {
	jsiiProxy_Constraint
}

func (j *jsiiProxy_ClaimConstraint) Claim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"claim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClaimConstraint) Operator() ConditionOperator {
	var returns ConditionOperator
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClaimConstraint) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


// Experimental.
func NewClaimConstraint_Override(c ClaimConstraint, operator ConditionOperator, claim *string, values *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"@catnekaise/cdk-iam-utilities.ClaimConstraint",
		[]interface{}{operator, claim, values},
		c,
	)
}

func (c *jsiiProxy_ClaimConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
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

func (c *jsiiProxy_ClaimConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
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

