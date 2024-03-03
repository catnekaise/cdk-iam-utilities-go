//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (c *jsiiProxy_ConstraintsBuilder) validateAddParameters(constraint Constraint) error {
	if constraint == nil {
		return fmt.Errorf("parameter constraint is required, but nil was provided")
	}

	return nil
}

func validateNewConstraintsBuilderParameters(settings *ConstraintsBuilderSettings) error {
	if err := _jsii_.ValidateStruct(settings, func() string { return "parameter settings" }); err != nil {
		return err
	}

	return nil
}

