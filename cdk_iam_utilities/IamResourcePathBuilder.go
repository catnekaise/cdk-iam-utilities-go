package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type IamResourcePathBuilder interface {
	IIamResourcePath
	// Experimental.
	Path() *[]*string
	// Experimental.
	AppendPolicyVariable(policyVariable PolicyVariable) *[]*string
	// Experimental.
	AppendText(values ...*string) *[]*string
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamResourcePathBuilder
type jsiiProxy_IamResourcePathBuilder struct {
	jsiiProxy_IIamResourcePath
}

func (j *jsiiProxy_IamResourcePathBuilder) Path() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


// Experimental.
func NewIamResourcePathBuilder_Override(i IamResourcePathBuilder, path *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"@catnekaise/cdk-iam-utilities.IamResourcePathBuilder",
		[]interface{}{path},
		i,
	)
}

func (i *jsiiProxy_IamResourcePathBuilder) AppendPolicyVariable(policyVariable PolicyVariable) *[]*string {
	if err := i.validateAppendPolicyVariableParameters(policyVariable); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"appendPolicyVariable",
		[]interface{}{policyVariable},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamResourcePathBuilder) AppendText(values ...*string) *[]*string {
	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns *[]*string

	_jsii_.Invoke(
		i,
		"appendText",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamResourcePathBuilder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

