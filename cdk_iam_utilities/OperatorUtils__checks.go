//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"
)

func validateOperatorUtils_ArraySupportParameters(value ConditionOperator) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateOperatorUtils_ConvertParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateOperatorUtils_OperatorIsSupportedParameters(supportedOperators *[]*string, operator ConditionOperator) error {
	if supportedOperators == nil {
		return fmt.Errorf("parameter supportedOperators is required, but nil was provided")
	}

	if operator == "" {
		return fmt.Errorf("parameter operator is required, but nil was provided")
	}

	return nil
}

func validateOperatorUtils_OperatorShortNameParameters(operator ConditionOperator) error {
	if operator == "" {
		return fmt.Errorf("parameter operator is required, but nil was provided")
	}

	return nil
}

