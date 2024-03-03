package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type GenericConstraint interface {
	Constraint
	// Experimental.
	Key() ConditionKey
	// Experimental.
	Operator() ConditionOperator
	// Experimental.
	Value() *[]*string
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for GenericConstraint
type jsiiProxy_GenericConstraint struct {
	jsiiProxy_Constraint
}

func (j *jsiiProxy_GenericConstraint) Key() ConditionKey {
	var returns ConditionKey
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenericConstraint) Operator() ConditionOperator {
	var returns ConditionOperator
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenericConstraint) Value() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewGenericConstraint(operator ConditionOperator, key ConditionKey, value *string, additionalValues ...*string) GenericConstraint {
	_init_.Initialize()

	if err := validateNewGenericConstraintParameters(operator, key, value); err != nil {
		panic(err)
	}
	args := []interface{}{operator, key, value}
	for _, a := range additionalValues {
		args = append(args, a)
	}

	j := jsiiProxy_GenericConstraint{}

	_jsii_.Create(
		"@catnekaise/cdk-iam-utilities.GenericConstraint",
		args,
		&j,
	)

	return &j
}

// Experimental.
func NewGenericConstraint_Override(g GenericConstraint, operator ConditionOperator, key ConditionKey, value *string, additionalValues ...*string) {
	_init_.Initialize()

	args := []interface{}{operator, key, value}
	for _, a := range additionalValues {
		args = append(args, a)
	}

	_jsii_.Create(
		"@catnekaise/cdk-iam-utilities.GenericConstraint",
		args,
		g,
	)
}

func (g *jsiiProxy_GenericConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
	if err := g.validateAssembleParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *[]*ConstraintPolicyMutation

	_jsii_.Invoke(
		g,
		"assemble",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
	if err := g.validateIsNotNullConditionParameters(key); err != nil {
		panic(err)
	}
	var returns *ConstraintPolicyMutation

	_jsii_.Invoke(
		g,
		"isNotNullCondition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

