package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IClaimsContext interface {
	// Experimental.
	KnownClaims() *[]*string
	// Experimental.
	MappedClaims() IMappedClaims
}

// The jsii proxy for IClaimsContext
type jsiiProxy_IClaimsContext struct {
	_ byte // padding
}

func (j *jsiiProxy_IClaimsContext) KnownClaims() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"knownClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClaimsContext) MappedClaims() IMappedClaims {
	var returns IMappedClaims
	_jsii_.Get(
		j,
		"mappedClaims",
		&returns,
	)
	return returns
}

