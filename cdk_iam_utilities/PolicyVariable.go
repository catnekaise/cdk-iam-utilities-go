package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type PolicyVariable interface {
	// Experimental.
	DefaultValue() *string
	// Experimental.
	IsTag() *bool
	// Experimental.
	TagName() *string
	// Experimental.
	Type() *string
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PolicyVariable
type jsiiProxy_PolicyVariable struct {
	_ byte // padding
}

func (j *jsiiProxy_PolicyVariable) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVariable) IsTag() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVariable) TagName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVariable) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func PolicyVariable_PrincipalOrgId(defaultValue *string) PolicyVariable {
	_init_.Initialize()

	var returns PolicyVariable

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.PolicyVariable",
		"principalOrgId",
		[]interface{}{defaultValue},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyVariable_PrincipalTag(tagName *string, defaultValue *string) PolicyVariable {
	_init_.Initialize()

	if err := validatePolicyVariable_PrincipalTagParameters(tagName); err != nil {
		panic(err)
	}
	var returns PolicyVariable

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.PolicyVariable",
		"principalTag",
		[]interface{}{tagName, defaultValue},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyVariable_PrincipalType(defaultValue *string) PolicyVariable {
	_init_.Initialize()

	var returns PolicyVariable

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.PolicyVariable",
		"principalType",
		[]interface{}{defaultValue},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyVariable_RequestTag(tagName *string, defaultValue *string) PolicyVariable {
	_init_.Initialize()

	if err := validatePolicyVariable_RequestTagParameters(tagName); err != nil {
		panic(err)
	}
	var returns PolicyVariable

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.PolicyVariable",
		"requestTag",
		[]interface{}{tagName, defaultValue},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyVariable_ResourceTag(tagName *string, defaultValue *string) PolicyVariable {
	_init_.Initialize()

	if err := validatePolicyVariable_ResourceTagParameters(tagName); err != nil {
		panic(err)
	}
	var returns PolicyVariable

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.PolicyVariable",
		"resourceTag",
		[]interface{}{tagName, defaultValue},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyVariable_UserId(defaultValue *string) PolicyVariable {
	_init_.Initialize()

	var returns PolicyVariable

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.PolicyVariable",
		"userId",
		[]interface{}{defaultValue},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyVariable_Username(defaultValue *string) PolicyVariable {
	_init_.Initialize()

	var returns PolicyVariable

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.PolicyVariable",
		"username",
		[]interface{}{defaultValue},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVariable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

