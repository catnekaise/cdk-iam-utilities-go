//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Constraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (c *jsiiProxy_Constraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

