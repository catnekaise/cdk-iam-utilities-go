package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IIamResourcePath interface {
	// Experimental.
	ToString() *string
}

// The jsii proxy for IIamResourcePath
type jsiiProxy_IIamResourcePath struct {
	_ byte // padding
}

func (i *jsiiProxy_IIamResourcePath) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

