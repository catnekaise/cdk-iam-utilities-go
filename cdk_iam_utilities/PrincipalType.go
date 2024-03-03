package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type PrincipalType interface {
	// Experimental.
	IsAws() *bool
	// Experimental.
	IsFederated() *bool
	// Experimental.
	IsSaml() *bool
	// Experimental.
	IsService() *bool
	// Experimental.
	Type() *string
}

// The jsii proxy struct for PrincipalType
type jsiiProxy_PrincipalType struct {
	_ byte // padding
}

func (j *jsiiProxy_PrincipalType) IsAws() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isAws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalType) IsFederated() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFederated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalType) IsSaml() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSaml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalType) IsService() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalType) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func PrincipalType_Aws() PrincipalType {
	_init_.Initialize()
	var returns PrincipalType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.PrincipalType",
		"Aws",
		&returns,
	)
	return returns
}

func PrincipalType_Federated() PrincipalType {
	_init_.Initialize()
	var returns PrincipalType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.PrincipalType",
		"Federated",
		&returns,
	)
	return returns
}

func PrincipalType_Saml() PrincipalType {
	_init_.Initialize()
	var returns PrincipalType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.PrincipalType",
		"Saml",
		&returns,
	)
	return returns
}

func PrincipalType_Service() PrincipalType {
	_init_.Initialize()
	var returns PrincipalType
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.PrincipalType",
		"Service",
		&returns,
	)
	return returns
}

