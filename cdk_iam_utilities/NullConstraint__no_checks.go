//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NullConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (n *jsiiProxy_NullConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validateNullConstraint_IsNotNullParameters(key ConditionKey) error {
	return nil
}

func validateNullConstraint_IsNullParameters(key ConditionKey) error {
	return nil
}

func validateNewNullConstraintParameters(key ConditionKey, isNull *bool) error {
	return nil
}

