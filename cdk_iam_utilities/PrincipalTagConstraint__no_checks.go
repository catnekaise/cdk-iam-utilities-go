//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrincipalTagConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (p *jsiiProxy_PrincipalTagConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validatePrincipalTagConstraint_StringEqualsParameters(tagName *string, value *string) error {
	return nil
}

