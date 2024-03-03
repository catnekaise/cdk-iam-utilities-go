//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourceTagConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (r *jsiiProxy_ResourceTagConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validateResourceTagConstraint_CreateParameters(operator StringConditionOperator, tagName *string, value *string) error {
	return nil
}

func validateResourceTagConstraint_StringEqualsParameters(tagName *string, value *string) error {
	return nil
}

