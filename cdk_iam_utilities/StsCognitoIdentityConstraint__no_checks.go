//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StsCognitoIdentityConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (s *jsiiProxy_StsCognitoIdentityConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validateStsCognitoIdentityConstraint_IdentityPoolParameters(identityPoolId *string, amr *string) error {
	return nil
}

