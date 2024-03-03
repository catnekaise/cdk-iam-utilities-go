package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type GenericClaimsIamResourcePathBuilder interface {
	ClaimsIamResourcePathBuilder
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
	Claim(claim *string, additionalClaims ...*string) GenericClaimsIamResourcePathBuilder
	// Experimental.
	PolicyVariable(value PolicyVariable) GenericClaimsIamResourcePathBuilder
	// Experimental.
	Text(value *string, additionalValues ...*string) GenericClaimsIamResourcePathBuilder
	// Experimental.
	ToString() *string
	// Experimental.
	Value(value *string, additionalValues ...*string) GenericClaimsIamResourcePathBuilder
}

// The jsii proxy struct for GenericClaimsIamResourcePathBuilder
type jsiiProxy_GenericClaimsIamResourcePathBuilder struct {
	jsiiProxy_ClaimsIamResourcePathBuilder
}

func (j *jsiiProxy_GenericClaimsIamResourcePathBuilder) Options() *ClaimsIamResourcePathBuilderSettings {
	var returns *ClaimsIamResourcePathBuilderSettings
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GenericClaimsIamResourcePathBuilder) Path() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


// Experimental.
func GenericClaimsIamResourcePathBuilder_Create(claimsContext IClaimsContext) GenericClaimsIamResourcePathBuilder {
	_init_.Initialize()

	if err := validateGenericClaimsIamResourcePathBuilder_CreateParameters(claimsContext); err != nil {
		panic(err)
	}
	var returns GenericClaimsIamResourcePathBuilder

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.GenericClaimsIamResourcePathBuilder",
		"create",
		[]interface{}{claimsContext},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) AppendClaim(claims ...*string) *[]*string {
	args := []interface{}{}
	for _, a := range claims {
		args = append(args, a)
	}

	var returns *[]*string

	_jsii_.Invoke(
		g,
		"appendClaim",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) AppendPolicyVariable(policyVariable PolicyVariable) *[]*string {
	if err := g.validateAppendPolicyVariableParameters(policyVariable); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"appendPolicyVariable",
		[]interface{}{policyVariable},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) AppendText(values ...*string) *[]*string {
	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns *[]*string

	_jsii_.Invoke(
		g,
		"appendText",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) AppendValue(values ...*string) *[]*string {
	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns *[]*string

	_jsii_.Invoke(
		g,
		"appendValue",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) Claim(claim *string, additionalClaims ...*string) GenericClaimsIamResourcePathBuilder {
	if err := g.validateClaimParameters(claim); err != nil {
		panic(err)
	}
	args := []interface{}{claim}
	for _, a := range additionalClaims {
		args = append(args, a)
	}

	var returns GenericClaimsIamResourcePathBuilder

	_jsii_.Invoke(
		g,
		"claim",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) PolicyVariable(value PolicyVariable) GenericClaimsIamResourcePathBuilder {
	if err := g.validatePolicyVariableParameters(value); err != nil {
		panic(err)
	}
	var returns GenericClaimsIamResourcePathBuilder

	_jsii_.Invoke(
		g,
		"policyVariable",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) Text(value *string, additionalValues ...*string) GenericClaimsIamResourcePathBuilder {
	if err := g.validateTextParameters(value); err != nil {
		panic(err)
	}
	args := []interface{}{value}
	for _, a := range additionalValues {
		args = append(args, a)
	}

	var returns GenericClaimsIamResourcePathBuilder

	_jsii_.Invoke(
		g,
		"text",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) Value(value *string, additionalValues ...*string) GenericClaimsIamResourcePathBuilder {
	if err := g.validateValueParameters(value); err != nil {
		panic(err)
	}
	args := []interface{}{value}
	for _, a := range additionalValues {
		args = append(args, a)
	}

	var returns GenericClaimsIamResourcePathBuilder

	_jsii_.Invoke(
		g,
		"value",
		args,
		&returns,
	)

	return returns
}

