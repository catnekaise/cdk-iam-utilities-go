package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type ConstraintsUtility interface {
	// Experimental.
	AppendGrant(scope constructs.Construct, settings *ConstraintUtilitySettings, grant awsiam.Grant)
	// Experimental.
	AppendPolicy(scope constructs.Construct, settings *ConstraintUtilitySettings, policyStatement awsiam.PolicyStatement)
}

// The jsii proxy struct for ConstraintsUtility
type jsiiProxy_ConstraintsUtility struct {
	_ byte // padding
}

// Experimental.
func ConstraintsUtility_ForConstraints(constraints *[]Constraint) ConstraintsUtility {
	_init_.Initialize()

	if err := validateConstraintsUtility_ForConstraintsParameters(constraints); err != nil {
		panic(err)
	}
	var returns ConstraintsUtility

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.ConstraintsUtility",
		"forConstraints",
		[]interface{}{constraints},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstraintsUtility) AppendGrant(scope constructs.Construct, settings *ConstraintUtilitySettings, grant awsiam.Grant) {
	if err := c.validateAppendGrantParameters(scope, settings, grant); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"appendGrant",
		[]interface{}{scope, settings, grant},
	)
}

func (c *jsiiProxy_ConstraintsUtility) AppendPolicy(scope constructs.Construct, settings *ConstraintUtilitySettings, policyStatement awsiam.PolicyStatement) {
	if err := c.validateAppendPolicyParameters(scope, settings, policyStatement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"appendPolicy",
		[]interface{}{scope, settings, policyStatement},
	)
}

