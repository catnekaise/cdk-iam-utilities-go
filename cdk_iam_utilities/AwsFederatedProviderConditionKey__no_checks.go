//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsFederatedProviderConditionKey) validateToConstraintParameters(operator ConditionOperator, value *string) error {
	return nil
}

