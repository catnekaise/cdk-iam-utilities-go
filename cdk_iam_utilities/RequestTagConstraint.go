package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type RequestTagConstraint interface {
	TagConstraint
	// Experimental.
	TagName() *string
	// Experimental.
	Values() *[]*string
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for RequestTagConstraint
type jsiiProxy_RequestTagConstraint struct {
	jsiiProxy_TagConstraint
}

func (j *jsiiProxy_RequestTagConstraint) TagName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestTagConstraint) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


// Experimental.
func RequestTagConstraint_StringEquals(tagName *string, value *string, values ...*string) RequestTagConstraint {
	_init_.Initialize()

	if err := validateRequestTagConstraint_StringEqualsParameters(tagName, value); err != nil {
		panic(err)
	}
	args := []interface{}{tagName, value}
	for _, a := range values {
		args = append(args, a)
	}

	var returns RequestTagConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.RequestTagConstraint",
		"stringEquals",
		args,
		&returns,
	)

	return returns
}

// Experimental.
func RequestTagConstraint_StringLike(tagName *string, value *string, values ...*string) RequestTagConstraint {
	_init_.Initialize()

	if err := validateRequestTagConstraint_StringLikeParameters(tagName, value); err != nil {
		panic(err)
	}
	args := []interface{}{tagName, value}
	for _, a := range values {
		args = append(args, a)
	}

	var returns RequestTagConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.RequestTagConstraint",
		"stringLike",
		args,
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestTagConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
	if err := r.validateAssembleParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *[]*ConstraintPolicyMutation

	_jsii_.Invoke(
		r,
		"assemble",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestTagConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
	if err := r.validateIsNotNullConditionParameters(key); err != nil {
		panic(err)
	}
	var returns *ConstraintPolicyMutation

	_jsii_.Invoke(
		r,
		"isNotNullCondition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

