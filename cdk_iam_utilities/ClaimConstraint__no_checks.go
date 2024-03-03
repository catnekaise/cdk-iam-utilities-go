//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClaimConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (c *jsiiProxy_ClaimConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validateNewClaimConstraintParameters(operator ConditionOperator, claim *string, values *[]*string) error {
	return nil
}

