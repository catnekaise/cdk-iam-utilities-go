package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type TagKeysConstraint interface {
	Constraint
	// Experimental.
	Operator() StringMultiValueConditionOperator
	// Experimental.
	Values() *[]*string
	// Experimental.
	Assemble(scope constructs.Construct, _arg *ConstraintAssembleContext) *[]*ConstraintPolicyMutation
	// Experimental.
	IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation
}

// The jsii proxy struct for TagKeysConstraint
type jsiiProxy_TagKeysConstraint struct {
	jsiiProxy_Constraint
}

func (j *jsiiProxy_TagKeysConstraint) Operator() StringMultiValueConditionOperator {
	var returns StringMultiValueConditionOperator
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagKeysConstraint) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


// Experimental.
func TagKeysConstraint_Create(operator StringMultiValueConditionOperator, isNotNull *bool, value *string, values ...*string) TagKeysConstraint {
	_init_.Initialize()

	if err := validateTagKeysConstraint_CreateParameters(operator, isNotNull, value); err != nil {
		panic(err)
	}
	args := []interface{}{operator, isNotNull, value}
	for _, a := range values {
		args = append(args, a)
	}

	var returns TagKeysConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.TagKeysConstraint",
		"create",
		args,
		&returns,
	)

	return returns
}

// Limit request tags to those specified and check for null.
// Experimental.
func TagKeysConstraint_RequireTagsEquals(value *string, values ...*string) TagKeysConstraint {
	_init_.Initialize()

	if err := validateTagKeysConstraint_RequireTagsEqualsParameters(value); err != nil {
		panic(err)
	}
	args := []interface{}{value}
	for _, a := range values {
		args = append(args, a)
	}

	var returns TagKeysConstraint

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.TagKeysConstraint",
		"requireTagsEquals",
		args,
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagKeysConstraint) Assemble(scope constructs.Construct, _arg *ConstraintAssembleContext) *[]*ConstraintPolicyMutation {
	if err := t.validateAssembleParameters(scope, _arg); err != nil {
		panic(err)
	}
	var returns *[]*ConstraintPolicyMutation

	_jsii_.Invoke(
		t,
		"assemble",
		[]interface{}{scope, _arg},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagKeysConstraint) IsNotNullCondition(key ConditionKey) *ConstraintPolicyMutation {
	if err := t.validateIsNotNullConditionParameters(key); err != nil {
		panic(err)
	}
	var returns *ConstraintPolicyMutation

	_jsii_.Invoke(
		t,
		"isNotNullCondition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

