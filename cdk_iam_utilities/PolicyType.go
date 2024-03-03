package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type PolicyType interface {
	// Experimental.
	IsIdentityPolicy() *bool
	// Experimental.
	IsResourcePolicy() *bool
	// Experimental.
	IsTrustPolicy() *bool
	// Experimental.
	PrincipalType() PrincipalType
	// Experimental.
	Service() ResourcePolicyType
	// Experimental.
	Type() *string
	// Experimental.
	IsResourcePolicyForService(service ResourcePolicyType) *bool
}

// The jsii proxy struct for PolicyType
type jsiiProxy_PolicyType struct {
	_ byte // padding
}

func (j *jsiiProxy_PolicyType) IsIdentityPolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isIdentityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyType) IsResourcePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isResourcePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyType) IsTrustPolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isTrustPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyType) PrincipalType() PrincipalType {
	var returns PrincipalType
	_jsii_.Get(
		j,
		"principalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyType) Service() ResourcePolicyType {
	var returns ResourcePolicyType
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyType) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func PolicyType_ResourcePolicy(type_ ResourcePolicyType) PolicyType {
	_init_.Initialize()

	if err := validatePolicyType_ResourcePolicyParameters(type_); err != nil {
		panic(err)
	}
	var returns PolicyType

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.PolicyType",
		"resourcePolicy",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyType_TrustPolicy(principalType PrincipalType) PolicyType {
	_init_.Initialize()

	if err := validatePolicyType_TrustPolicyParameters(principalType); err != nil {
		panic(err)
	}
	var returns PolicyType

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.PolicyType",
		"trustPolicy",
		[]interface{}{principalType},
		&returns,
	)

	return returns
}

func PolicyType_IdentityPolicy() PolicyType {
	_init_.Initialize()
	var returns PolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.PolicyType",
		"IdentityPolicy",
		&returns,
	)
	return returns
}

func PolicyType_NonSpecific() PolicyType {
	_init_.Initialize()
	var returns PolicyType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.PolicyType",
		"NonSpecific",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PolicyType) IsResourcePolicyForService(service ResourcePolicyType) *bool {
	if err := p.validateIsResourcePolicyForServiceParameters(service); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		p,
		"isResourcePolicyForService",
		[]interface{}{service},
		&returns,
	)

	return returns
}

