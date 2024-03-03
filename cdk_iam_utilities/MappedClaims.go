package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type MappedClaims interface {
	IMappedClaims
	// Experimental.
	Claims() *[]*Claim
}

// The jsii proxy struct for MappedClaims
type jsiiProxy_MappedClaims struct {
	jsiiProxy_IMappedClaims
}

func (j *jsiiProxy_MappedClaims) Claims() *[]*Claim {
	var returns *[]*Claim
	_jsii_.Get(
		j,
		"claims",
		&returns,
	)
	return returns
}


// Experimental.
func MappedClaims_Create(claim *string, additionalClaims ...*string) MappedClaims {
	_init_.Initialize()

	if err := validateMappedClaims_CreateParameters(claim); err != nil {
		panic(err)
	}
	args := []interface{}{claim}
	for _, a := range additionalClaims {
		args = append(args, a)
	}

	var returns MappedClaims

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.MappedClaims",
		"create",
		args,
		&returns,
	)

	return returns
}

// Experimental.
func MappedClaims_CreateMapped(claims *map[string]*string) MappedClaims {
	_init_.Initialize()

	if err := validateMappedClaims_CreateMappedParameters(claims); err != nil {
		panic(err)
	}
	var returns MappedClaims

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.MappedClaims",
		"createMapped",
		[]interface{}{claims},
		&returns,
	)

	return returns
}

