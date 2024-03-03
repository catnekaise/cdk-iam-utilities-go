//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlobalConditionKey) validateToConstraintParameters(operator ConditionOperator, value *string) error {
	return nil
}

func validateGlobalConditionKey_PrincipalTagParameters(tagName *string) error {
	return nil
}

func validateGlobalConditionKey_RequestTagParameters(tagName *string) error {
	return nil
}

func validateGlobalConditionKey_ResourceTagParameters(tagName *string) error {
	return nil
}

