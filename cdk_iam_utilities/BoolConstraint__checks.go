//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (b *jsiiProxy_BoolConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(context, func() string { return "parameter context" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BoolConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateBoolConstraint_WhenFalseParameters(key ConditionKey) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateBoolConstraint_WhenTrueParameters(key ConditionKey) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

