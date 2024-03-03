package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type CalledViaConstraint interface {
	Constraint
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for CalledViaConstraint
type jsiiProxy_CalledViaConstraint struct {
	jsiiProxy_Constraint
}

// Experimental.
func CalledViaConstraint_CalledVia(service CalledViaServicePrincipal) CalledViaConstraint {
	_init_.Initialize()

	if err := validateCalledViaConstraint_CalledViaParameters(service); err != nil {
		panic(err)
	}
	var returns CalledViaConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.CalledViaConstraint",
		"calledVia",
		[]interface{}{service},
		&returns,
	)

	return returns
}

// Experimental.
func CalledViaConstraint_CalledViaFirst(service CalledViaServicePrincipal) CalledViaConstraint {
	_init_.Initialize()

	if err := validateCalledViaConstraint_CalledViaFirstParameters(service); err != nil {
		panic(err)
	}
	var returns CalledViaConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.CalledViaConstraint",
		"calledViaFirst",
		[]interface{}{service},
		&returns,
	)

	return returns
}

// Experimental.
func CalledViaConstraint_CalledViaFirstAndLast(firstService CalledViaServicePrincipal, lastService CalledViaServicePrincipal) CalledViaConstraint {
	_init_.Initialize()

	if err := validateCalledViaConstraint_CalledViaFirstAndLastParameters(firstService, lastService); err != nil {
		panic(err)
	}
	var returns CalledViaConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.CalledViaConstraint",
		"calledViaFirstAndLast",
		[]interface{}{firstService, lastService},
		&returns,
	)

	return returns
}

// Experimental.
func CalledViaConstraint_CalledViaLast(service CalledViaServicePrincipal) CalledViaConstraint {
	_init_.Initialize()

	if err := validateCalledViaConstraint_CalledViaLastParameters(service); err != nil {
		panic(err)
	}
	var returns CalledViaConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.CalledViaConstraint",
		"calledViaLast",
		[]interface{}{service},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CalledViaConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
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

func (c *jsiiProxy_CalledViaConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
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

