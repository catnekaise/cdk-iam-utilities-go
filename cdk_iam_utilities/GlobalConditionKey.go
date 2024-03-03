package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type GlobalConditionKey interface {
	ConditionKey
	// Experimental.
	Name() *string
	// Experimental.
	Settings() *ConditionKeySettings
	// Experimental.
	ToConstraint(operator ConditionOperator, value *string, additionalValues ...*string) GenericConstraint
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlobalConditionKey
type jsiiProxy_GlobalConditionKey struct {
	jsiiProxy_ConditionKey
}

func (j *jsiiProxy_GlobalConditionKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalConditionKey) Settings() *ConditionKeySettings {
	var returns *ConditionKeySettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Use this key to compare the tag attached to the principal making the request with the tag that you specify in the policy.
// Experimental.
func GlobalConditionKey_PrincipalTag(tagName *string) AwsPrincipalTagConditionKey {
	_init_.Initialize()

	if err := validateGlobalConditionKey_PrincipalTagParameters(tagName); err != nil {
		panic(err)
	}
	var returns AwsPrincipalTagConditionKey

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"principalTag",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

// Use this key to compare the tag key-value pair that was passed in the request with the tag pair that you specify in the policy.
// Experimental.
func GlobalConditionKey_RequestTag(tagName *string) AwsRequestTagConditionKey {
	_init_.Initialize()

	if err := validateGlobalConditionKey_RequestTagParameters(tagName); err != nil {
		panic(err)
	}
	var returns AwsRequestTagConditionKey

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"requestTag",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

// Use this key to compare the tag key-value pair that you specify in the policy with the key-value pair attached to the resource.
// Experimental.
func GlobalConditionKey_ResourceTag(tagName *string) AwsResourceTagConditionKey {
	_init_.Initialize()

	if err := validateGlobalConditionKey_ResourceTagParameters(tagName); err != nil {
		panic(err)
	}
	var returns AwsResourceTagConditionKey

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"resourceTag",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

func GlobalConditionKey_CalledVia() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"CalledVia",
		&returns,
	)
	return returns
}

func GlobalConditionKey_CalledViaFirst() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"CalledViaFirst",
		&returns,
	)
	return returns
}

func GlobalConditionKey_CalledViaLast() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"CalledViaLast",
		&returns,
	)
	return returns
}

func GlobalConditionKey_CurrentTime() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"CurrentTime",
		&returns,
	)
	return returns
}

func GlobalConditionKey_Ec2InstanceSourcePrivateIPv4() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"Ec2InstanceSourcePrivateIPv4",
		&returns,
	)
	return returns
}

func GlobalConditionKey_Ec2InstanceSourceVpc() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"Ec2InstanceSourceVpc",
		&returns,
	)
	return returns
}

func GlobalConditionKey_EpochTime() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"EpochTime",
		&returns,
	)
	return returns
}

func GlobalConditionKey_FederatedProvider() AwsFederatedProviderConditionKey {
	_init_.Initialize()
	var returns AwsFederatedProviderConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"FederatedProvider",
		&returns,
	)
	return returns
}

func GlobalConditionKey_MultiFactorAuthAge() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"MultiFactorAuthAge",
		&returns,
	)
	return returns
}

func GlobalConditionKey_MultiFactorAuthPresent() GlobalBoolConditionKey {
	_init_.Initialize()
	var returns GlobalBoolConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"MultiFactorAuthPresent",
		&returns,
	)
	return returns
}

func GlobalConditionKey_PrincipalAccount() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"PrincipalAccount",
		&returns,
	)
	return returns
}

func GlobalConditionKey_PrincipalArn() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"PrincipalArn",
		&returns,
	)
	return returns
}

func GlobalConditionKey_PrincipalIsAWSService() GlobalBoolConditionKey {
	_init_.Initialize()
	var returns GlobalBoolConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"PrincipalIsAWSService",
		&returns,
	)
	return returns
}

func GlobalConditionKey_PrincipalOrgID() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"PrincipalOrgID",
		&returns,
	)
	return returns
}

func GlobalConditionKey_PrincipalOrgPaths() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"PrincipalOrgPaths",
		&returns,
	)
	return returns
}

func GlobalConditionKey_PrincipalServiceName() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"PrincipalServiceName",
		&returns,
	)
	return returns
}

func GlobalConditionKey_PrincipalServiceNamesList() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"PrincipalServiceNamesList",
		&returns,
	)
	return returns
}

func GlobalConditionKey_PrincipalType() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"PrincipalType",
		&returns,
	)
	return returns
}

func GlobalConditionKey_Referer() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"Referer",
		&returns,
	)
	return returns
}

func GlobalConditionKey_RequestedRegion() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"RequestedRegion",
		&returns,
	)
	return returns
}

func GlobalConditionKey_ResourceAccount() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"ResourceAccount",
		&returns,
	)
	return returns
}

func GlobalConditionKey_ResourceOrgID() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"ResourceOrgID",
		&returns,
	)
	return returns
}

func GlobalConditionKey_ResourceOrgPaths() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"ResourceOrgPaths",
		&returns,
	)
	return returns
}

func GlobalConditionKey_SecureTransport() GlobalBoolConditionKey {
	_init_.Initialize()
	var returns GlobalBoolConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"SecureTransport",
		&returns,
	)
	return returns
}

func GlobalConditionKey_SourceAccount() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"SourceAccount",
		&returns,
	)
	return returns
}

func GlobalConditionKey_SourceArn() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"SourceArn",
		&returns,
	)
	return returns
}

func GlobalConditionKey_SourceIdentity() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"SourceIdentity",
		&returns,
	)
	return returns
}

func GlobalConditionKey_SourceIp() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"SourceIp",
		&returns,
	)
	return returns
}

func GlobalConditionKey_SourceOrgID() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"SourceOrgID",
		&returns,
	)
	return returns
}

func GlobalConditionKey_SourceOrgPaths() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"SourceOrgPaths",
		&returns,
	)
	return returns
}

func GlobalConditionKey_SourceVpc() AwsSourceVpcConditionKey {
	_init_.Initialize()
	var returns AwsSourceVpcConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"SourceVpc",
		&returns,
	)
	return returns
}

func GlobalConditionKey_SourceVpce() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"SourceVpce",
		&returns,
	)
	return returns
}

func GlobalConditionKey_TagKeys() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"TagKeys",
		&returns,
	)
	return returns
}

func GlobalConditionKey_TokenIssueTime() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"TokenIssueTime",
		&returns,
	)
	return returns
}

func GlobalConditionKey_UserAgent() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"UserAgent",
		&returns,
	)
	return returns
}

func GlobalConditionKey_Userid() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"Userid",
		&returns,
	)
	return returns
}

func GlobalConditionKey_Username() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"Username",
		&returns,
	)
	return returns
}

func GlobalConditionKey_ViaAWSService() GlobalBoolConditionKey {
	_init_.Initialize()
	var returns GlobalBoolConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"ViaAWSService",
		&returns,
	)
	return returns
}

func GlobalConditionKey_VpcSourceIp() GlobalConditionKey {
	_init_.Initialize()
	var returns GlobalConditionKey
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		"VpcSourceIp",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GlobalConditionKey) ToConstraint(operator ConditionOperator, value *string, additionalValues ...*string) GenericConstraint {
	if err := g.validateToConstraintParameters(operator, value); err != nil {
		panic(err)
	}
	args := []interface{}{operator, value}
	for _, a := range additionalValues {
		args = append(args, a)
	}

	var returns GenericConstraint

	_jsii_.Invoke(
		g,
		"toConstraint",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalConditionKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

