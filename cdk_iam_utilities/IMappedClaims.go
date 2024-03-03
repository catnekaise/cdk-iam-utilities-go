package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IMappedClaims interface {
	// Experimental.
	Claims() *[]*Claim
}

// The jsii proxy for IMappedClaims
type jsiiProxy_IMappedClaims struct {
	_ byte // padding
}

func (j *jsiiProxy_IMappedClaims) Claims() *[]*Claim {
	var returns *[]*Claim
	_jsii_.Get(
		j,
		"claims",
		&returns,
	)
	return returns
}

