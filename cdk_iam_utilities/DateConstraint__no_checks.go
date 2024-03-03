//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DateConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (d *jsiiProxy_DateConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validateDateConstraint_BetweenDatesParameters(key ConditionKey, from *time.Time, to *time.Time) error {
	return nil
}

func validateDateConstraint_GreaterThanParameters(key ConditionKey, date *time.Time) error {
	return nil
}

func validateDateConstraint_LessThanParameters(key ConditionKey, date *time.Time) error {
	return nil
}

