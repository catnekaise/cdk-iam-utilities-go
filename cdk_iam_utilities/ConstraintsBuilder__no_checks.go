//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConstraintsBuilder) validateAddParameters(constraint Constraint) error {
	return nil
}

func validateNewConstraintsBuilderParameters(settings *ConstraintsBuilderSettings) error {
	return nil
}

