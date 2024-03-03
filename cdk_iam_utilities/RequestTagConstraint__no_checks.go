//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RequestTagConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (r *jsiiProxy_RequestTagConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validateRequestTagConstraint_StringEqualsParameters(tagName *string, value *string) error {
	return nil
}

func validateRequestTagConstraint_StringLikeParameters(tagName *string, value *string) error {
	return nil
}

