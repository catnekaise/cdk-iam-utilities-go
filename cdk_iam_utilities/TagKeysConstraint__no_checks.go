//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TagKeysConstraint) validateAssembleParameters(scope constructs.Construct, _arg *ConstraintAssembleContext) error {
	return nil
}

func (t *jsiiProxy_TagKeysConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validateTagKeysConstraint_CreateParameters(operator StringMultiValueConditionOperator, isNotNull *bool, value *string) error {
	return nil
}

func validateTagKeysConstraint_RequireTagsEqualsParameters(value *string) error {
	return nil
}

