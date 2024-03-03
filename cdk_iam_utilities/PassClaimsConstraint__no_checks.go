//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PassClaimsConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (p *jsiiProxy_PassClaimsConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validatePassClaimsConstraint_CreateParameters(claims *PassClaimsConstraintSettings) error {
	return nil
}

