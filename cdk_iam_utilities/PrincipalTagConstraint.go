package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type PrincipalTagConstraint interface {
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

// The jsii proxy struct for PrincipalTagConstraint
type jsiiProxy_PrincipalTagConstraint struct {
	jsiiProxy_TagConstraint
}

func (j *jsiiProxy_PrincipalTagConstraint) TagName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalTagConstraint) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


// Experimental.
func PrincipalTagConstraint_StringEquals(tagName *string, value *string, values ...*string) PrincipalTagConstraint {
	_init_.Initialize()

	if err := validatePrincipalTagConstraint_StringEqualsParameters(tagName, value); err != nil {
		panic(err)
	}
	args := []interface{}{tagName, value}
	for _, a := range values {
		args = append(args, a)
	}

	var returns PrincipalTagConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.PrincipalTagConstraint",
		"stringEquals",
		args,
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalTagConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
	if err := p.validateAssembleParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *[]*ConstraintPolicyMutation

	_jsii_.Invoke(
		p,
		"assemble",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalTagConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
	if err := p.validateIsNotNullConditionParameters(key); err != nil {
		panic(err)
	}
	var returns *ConstraintPolicyMutation

	_jsii_.Invoke(
		p,
		"isNotNullCondition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

