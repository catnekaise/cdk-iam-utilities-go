//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"
)

func (g *jsiiProxy_GlobalConditionKey) validateToConstraintParameters(operator ConditionOperator, value *string) error {
	if operator == "" {
		return fmt.Errorf("parameter operator is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateGlobalConditionKey_PrincipalTagParameters(tagName *string) error {
	if tagName == nil {
		return fmt.Errorf("parameter tagName is required, but nil was provided")
	}

	return nil
}

func validateGlobalConditionKey_RequestTagParameters(tagName *string) error {
	if tagName == nil {
		return fmt.Errorf("parameter tagName is required, but nil was provided")
	}

	return nil
}

func validateGlobalConditionKey_ResourceTagParameters(tagName *string) error {
	if tagName == nil {
		return fmt.Errorf("parameter tagName is required, but nil was provided")
	}

	return nil
}

