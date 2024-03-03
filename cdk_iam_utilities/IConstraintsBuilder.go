package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IConstraintsBuilder interface {
	// Experimental.
	Constraints() *[]Constraint
}

// The jsii proxy for IConstraintsBuilder
type jsiiProxy_IConstraintsBuilder struct {
	_ byte // padding
}

func (j *jsiiProxy_IConstraintsBuilder) Constraints() *[]Constraint {
	var returns *[]Constraint
	_jsii_.Get(
		j,
		"constraints",
		&returns,
	)
	return returns
}

