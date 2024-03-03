package cdk_iam_utilities

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type DateConstraint interface {
	Constraint
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for DateConstraint
type jsiiProxy_DateConstraint struct {
	jsiiProxy_Constraint
}

// Experimental.
func DateConstraint_BetweenDates(key ConditionKey, from *time.Time, to *time.Time) DateConstraint {
	_init_.Initialize()

	if err := validateDateConstraint_BetweenDatesParameters(key, from, to); err != nil {
		panic(err)
	}
	var returns DateConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.DateConstraint",
		"betweenDates",
		[]interface{}{key, from, to},
		&returns,
	)

	return returns
}

// Experimental.
func DateConstraint_GreaterThan(key ConditionKey, date *time.Time) DateConstraint {
	_init_.Initialize()

	if err := validateDateConstraint_GreaterThanParameters(key, date); err != nil {
		panic(err)
	}
	var returns DateConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.DateConstraint",
		"greaterThan",
		[]interface{}{key, date},
		&returns,
	)

	return returns
}

// Experimental.
func DateConstraint_LessThan(key ConditionKey, date *time.Time) DateConstraint {
	_init_.Initialize()

	if err := validateDateConstraint_LessThanParameters(key, date); err != nil {
		panic(err)
	}
	var returns DateConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.DateConstraint",
		"lessThan",
		[]interface{}{key, date},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DateConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
	if err := d.validateAssembleParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *[]*ConstraintPolicyMutation

	_jsii_.Invoke(
		d,
		"assemble",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DateConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
	if err := d.validateIsNotNullConditionParameters(key); err != nil {
		panic(err)
	}
	var returns *ConstraintPolicyMutation

	_jsii_.Invoke(
		d,
		"isNotNullCondition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

