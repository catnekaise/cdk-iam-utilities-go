//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StsTransitiveTagKeysConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (s *jsiiProxy_StsTransitiveTagKeysConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validateStsTransitiveTagKeysConstraint_TagsEqualsAndPresentParameters(value *string) error {
	return nil
}

