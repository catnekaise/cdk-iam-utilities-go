package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type StsTransitiveTagKeysConstraint interface {
	Constraint
	// Experimental.
	Values() *[]*string
	// Experimental.
	Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for StsTransitiveTagKeysConstraint
type jsiiProxy_StsTransitiveTagKeysConstraint struct {
	jsiiProxy_Constraint
}

func (j *jsiiProxy_StsTransitiveTagKeysConstraint) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


// Limit transitive tags to those specified and check for null.
// Experimental.
func StsTransitiveTagKeysConstraint_TagsEqualsAndPresent(value *string, values ...*string) StsTransitiveTagKeysConstraint {
	_init_.Initialize()

	if err := validateStsTransitiveTagKeysConstraint_TagsEqualsAndPresentParameters(value); err != nil {
		panic(err)
	}
	args := []interface{}{value}
	for _, a := range values {
		args = append(args, a)
	}

	var returns StsTransitiveTagKeysConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.StsTransitiveTagKeysConstraint",
		"tagsEqualsAndPresent",
		args,
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StsTransitiveTagKeysConstraint) Assemble(scope constructs.Construct, context *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
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

func (s *jsiiProxy_StsTransitiveTagKeysConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
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

