//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CalledViaConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	return nil
}

func (c *jsiiProxy_CalledViaConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	return nil
}

func validateCalledViaConstraint_CalledViaParameters(service CalledViaServicePrincipal) error {
	return nil
}

func validateCalledViaConstraint_CalledViaFirstParameters(service CalledViaServicePrincipal) error {
	return nil
}

func validateCalledViaConstraint_CalledViaFirstAndLastParameters(firstService CalledViaServicePrincipal, lastService CalledViaServicePrincipal) error {
	return nil
}

func validateCalledViaConstraint_CalledViaLastParameters(service CalledViaServicePrincipal) error {
	return nil
}

