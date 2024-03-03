//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClaimsIamResourcePathBuilder) validateAppendPolicyVariableParameters(policyVariable PolicyVariable) error {
	return nil
}

func validateNewClaimsIamResourcePathBuilderParameters(options *ClaimsIamResourcePathBuilderSettings, path *[]*string) error {
	return nil
}

