//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BoolConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (b *jsiiProxy_BoolConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validateBoolConstraint_WhenFalseParameters(key ConditionKey) error {
	return nil
}

func validateBoolConstraint_WhenTrueParameters(key ConditionKey) error {
	return nil
}

