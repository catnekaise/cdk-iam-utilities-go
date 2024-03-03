//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"
)

func (a *jsiiProxy_AwsFederatedProviderConditionKey) validateToConstraintParameters(operator ConditionOperator, value *string) error {
	if operator == "" {
		return fmt.Errorf("parameter operator is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

