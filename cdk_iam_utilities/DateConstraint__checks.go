//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_DateConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
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

func (d *jsiiProxy_DateConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateDateConstraint_BetweenDatesParameters(key ConditionKey, from *time.Time, to *time.Time) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if from == nil {
		return fmt.Errorf("parameter from is required, but nil was provided")
	}

	if to == nil {
		return fmt.Errorf("parameter to is required, but nil was provided")
	}

	return nil
}

func validateDateConstraint_GreaterThanParameters(key ConditionKey, date *time.Time) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if date == nil {
		return fmt.Errorf("parameter date is required, but nil was provided")
	}

	return nil
}

func validateDateConstraint_LessThanParameters(key ConditionKey, date *time.Time) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if date == nil {
		return fmt.Errorf("parameter date is required, but nil was provided")
	}

	return nil
}

