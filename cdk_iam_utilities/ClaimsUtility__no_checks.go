//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClaimsUtility) validatePrincipalTagConditionParameters(claim *string) error {
	return nil
}

func (c *jsiiProxy_ClaimsUtility) validateRequestTagConditionParameters(claim *string) error {
	return nil
}

func (c *jsiiProxy_ClaimsUtility) validateTagNameParameters(scope constructs.Construct, claim *string) error {
	return nil
}

func (c *jsiiProxy_ClaimsUtility) validateTagNameForClaimParameters(claim *string) error {
	return nil
}

func validateClaimsUtility_ForContextParameters(context IClaimsContext) error {
	return nil
}

