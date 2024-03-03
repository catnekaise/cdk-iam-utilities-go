package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type ClaimsUtility interface {
	// Experimental.
	KnownClaims() *[]*string
	// Experimental.
	MappedClaims() IMappedClaims
	// Experimental.
	PrincipalTagCondition(claim *string) AwsPrincipalTagConditionKey
	// Experimental.
	RequestTagCondition(claim *string) AwsRequestTagConditionKey
	// Experimental.
	TagName(scope constructs.Construct, claim *string) *string
	// Experimental.
	TagNameForClaim(claim *string) *string
}

// The jsii proxy struct for ClaimsUtility
type jsiiProxy_ClaimsUtility struct {
	_ byte // padding
}

func (j *jsiiProxy_ClaimsUtility) KnownClaims() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"knownClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClaimsUtility) MappedClaims() IMappedClaims {
	var returns IMappedClaims
	_jsii_.Get(
		j,
		"mappedClaims",
		&returns,
	)
	return returns
}


// Experimental.
func ClaimsUtility_ForContext(context IClaimsContext) ClaimsUtility {
	_init_.Initialize()

	if err := validateClaimsUtility_ForContextParameters(context); err != nil {
		panic(err)
	}
	var returns ClaimsUtility

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.ClaimsUtility",
		"forContext",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClaimsUtility) PrincipalTagCondition(claim *string) AwsPrincipalTagConditionKey {
	if err := c.validatePrincipalTagConditionParameters(claim); err != nil {
		panic(err)
	}
	var returns AwsPrincipalTagConditionKey

	_jsii_.Invoke(
		c,
		"principalTagCondition",
		[]interface{}{claim},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClaimsUtility) RequestTagCondition(claim *string) AwsRequestTagConditionKey {
	if err := c.validateRequestTagConditionParameters(claim); err != nil {
		panic(err)
	}
	var returns AwsRequestTagConditionKey

	_jsii_.Invoke(
		c,
		"requestTagCondition",
		[]interface{}{claim},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClaimsUtility) TagName(scope constructs.Construct, claim *string) *string {
	if err := c.validateTagNameParameters(scope, claim); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"tagName",
		[]interface{}{scope, claim},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClaimsUtility) TagNameForClaim(claim *string) *string {
	if err := c.validateTagNameForClaimParameters(claim); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"tagNameForClaim",
		[]interface{}{claim},
		&returns,
	)

	return returns
}

