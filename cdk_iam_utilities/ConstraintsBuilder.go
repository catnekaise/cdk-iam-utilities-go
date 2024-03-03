package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type ConstraintsBuilder interface {
	IConstraintsBuilder
	// Experimental.
	Constraints() *[]Constraint
	// Experimental.
	Settings() *ConstraintsBuilderSettings
	// Experimental.
	Add(constraint Constraint, additionalConstraints ...Constraint) ConstraintsBuilder
}

// The jsii proxy struct for ConstraintsBuilder
type jsiiProxy_ConstraintsBuilder struct {
	jsiiProxy_IConstraintsBuilder
}

func (j *jsiiProxy_ConstraintsBuilder) Constraints() *[]Constraint {
	var returns *[]Constraint
	_jsii_.Get(
		j,
		"constraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstraintsBuilder) Settings() *ConstraintsBuilderSettings {
	var returns *ConstraintsBuilderSettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func NewConstraintsBuilder_Override(c ConstraintsBuilder, settings *ConstraintsBuilderSettings) {
	_init_.Initialize()

	_jsii_.Create(
		"@catnekaise/cdk-iam-utilities.ConstraintsBuilder",
		[]interface{}{settings},
		c,
	)
}

func (c *jsiiProxy_ConstraintsBuilder) Add(constraint Constraint, additionalConstraints ...Constraint) ConstraintsBuilder {
	if err := c.validateAddParameters(constraint); err != nil {
		panic(err)
	}
	args := []interface{}{constraint}
	for _, a := range additionalConstraints {
		args = append(args, a)
	}

	var returns ConstraintsBuilder

	_jsii_.Invoke(
		c,
		"add",
		args,
		&returns,
	)

	return returns
}

