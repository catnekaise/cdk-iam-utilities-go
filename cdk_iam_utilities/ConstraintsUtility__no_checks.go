//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConstraintsUtility) validateAppendGrantParameters(scope constructs.Construct, settings *ConstraintUtilitySettings, grant awsiam.Grant) error {
	return nil
}

func (c *jsiiProxy_ConstraintsUtility) validateAppendPolicyParameters(scope constructs.Construct, settings *ConstraintUtilitySettings, policyStatement awsiam.PolicyStatement) error {
	return nil
}

func validateConstraintsUtility_ForConstraintsParameters(constraints *[]Constraint) error {
	return nil
}

