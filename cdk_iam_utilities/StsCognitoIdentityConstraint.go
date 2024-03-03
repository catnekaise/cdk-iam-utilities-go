package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type StsCognitoIdentityConstraint interface {
	Constraint
	// Experimental.
	Amr() *string
	// Experimental.
	IdentityPoolId() *string
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for StsCognitoIdentityConstraint
type jsiiProxy_StsCognitoIdentityConstraint struct {
	jsiiProxy_Constraint
}

func (j *jsiiProxy_StsCognitoIdentityConstraint) Amr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StsCognitoIdentityConstraint) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}


// Experimental.
func StsCognitoIdentityConstraint_IdentityPool(identityPoolId *string, amr *string) StsCognitoIdentityConstraint {
	_init_.Initialize()

	if err := validateStsCognitoIdentityConstraint_IdentityPoolParameters(identityPoolId, amr); err != nil {
		panic(err)
	}
	var returns StsCognitoIdentityConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.StsCognitoIdentityConstraint",
		"identityPool",
		[]interface{}{identityPoolId, amr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StsCognitoIdentityConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
	if err := s.validateAssembleParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *[]*ConstraintPolicyMutation

	_jsii_.Invoke(
		s,
		"assemble",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StsCognitoIdentityConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
	if err := s.validateIsNotNullConditionParameters(key); err != nil {
		panic(err)
	}
	var returns *ConstraintPolicyMutation

	_jsii_.Invoke(
		s,
		"isNotNullCondition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

