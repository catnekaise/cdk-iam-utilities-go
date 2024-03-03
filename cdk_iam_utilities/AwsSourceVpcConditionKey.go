package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Experimental.
type AwsSourceVpcConditionKey interface {
	ConditionKey
	// Experimental.
	Name() *string
	// Experimental.
	Settings() *ConditionKeySettings
	// Experimental.
	ToString() *string
	// Experimental.
	ToVpcConstraint(vpc awsec2.IVpc) GenericConstraint
}

// The jsii proxy struct for AwsSourceVpcConditionKey
type jsiiProxy_AwsSourceVpcConditionKey struct {
	jsiiProxy_ConditionKey
}

func (j *jsiiProxy_AwsSourceVpcConditionKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSourceVpcConditionKey) Settings() *ConditionKeySettings {
	var returns *ConditionKeySettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


func AwsSourceVpcConditionKey_Create() AwsSourceVpcConditionKey {
	_init_.Initialize()
	var returns AwsSourceVpcConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.AwsSourceVpcConditionKey",
		"Create",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsSourceVpcConditionKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSourceVpcConditionKey) ToVpcConstraint(vpc awsec2.IVpc) GenericConstraint {
	if err := a.validateToVpcConstraintParameters(vpc); err != nil {
		panic(err)
	}
	var returns GenericConstraint

	_jsii_.Invoke(
		a,
		"toVpcConstraint",
		[]interface{}{vpc},
		&returns,
	)

	return returns
}

