//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GenericConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (g *jsiiProxy_GenericConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validateNewGenericConstraintParameters(operator ConditionOperator, key ConditionKey, value *string) error {
	return nil
}

