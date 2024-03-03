package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type PassClaimsConstraint interface {
	Constraint
	// Experimental.
	Settings() *PassClaimsConstraintSettings
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for PassClaimsConstraint
type jsiiProxy_PassClaimsConstraint struct {
	jsiiProxy_Constraint
}

func (j *jsiiProxy_PassClaimsConstraint) Settings() *PassClaimsConstraintSettings {
	var returns *PassClaimsConstraintSettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func PassClaimsConstraint_Create(claims *PassClaimsConstraintSettings) PassClaimsConstraint {
	_init_.Initialize()

	if err := validatePassClaimsConstraint_CreateParameters(claims); err != nil {
		panic(err)
	}
	var returns PassClaimsConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.PassClaimsConstraint",
		"create",
		[]interface{}{claims},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PassClaimsConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
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

func (p *jsiiProxy_PassClaimsConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
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

