package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type ClaimsIamResourcePathBuilder interface {
	IamResourcePathBuilder
	// Experimental.
	Options() *ClaimsIamResourcePathBuilderSettings
	// Experimental.
	Path() *[]*string
	// Experimental.
	AppendClaim(claims ...*string) *[]*string
	// Experimental.
	AppendPolicyVariable(policyVariable PolicyVariable) *[]*string
	// Experimental.
	AppendText(values ...*string) *[]*string
	// Experimental.
	AppendValue(values ...*string) *[]*string
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClaimsIamResourcePathBuilder
type jsiiProxy_ClaimsIamResourcePathBuilder struct {
	jsiiProxy_IamResourcePathBuilder
}

func (j *jsiiProxy_ClaimsIamResourcePathBuilder) Options() *ClaimsIamResourcePathBuilderSettings {
	var returns *ClaimsIamResourcePathBuilderSettings
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClaimsIamResourcePathBuilder) Path() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


// Experimental.
func NewClaimsIamResourcePathBuilder_Override(c ClaimsIamResourcePathBuilder, options *ClaimsIamResourcePathBuilderSettings, path *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"@catnekaise/cdk-iam-utilities.ClaimsIamResourcePathBuilder",
		[]interface{}{options, path},
		c,
	)
}

func (c *jsiiProxy_ClaimsIamResourcePathBuilder) AppendClaim(claims ...*string) *[]*string {
	args := []interface{}{}
	for _, a := range claims {
		args = append(args, a)
	}

	var returns *[]*string

	_jsii_.Invoke(
		c,
		"appendClaim",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClaimsIamResourcePathBuilder) AppendPolicyVariable(policyVariable PolicyVariable) *[]*string {
	if err := c.validateAppendPolicyVariableParameters(policyVariable); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"appendPolicyVariable",
		[]interface{}{policyVariable},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClaimsIamResourcePathBuilder) AppendText(values ...*string) *[]*string {
	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns *[]*string

	_jsii_.Invoke(
		c,
		"appendText",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClaimsIamResourcePathBuilder) AppendValue(values ...*string) *[]*string {
	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns *[]*string

	_jsii_.Invoke(
		c,
		"appendValue",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClaimsIamResourcePathBuilder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

