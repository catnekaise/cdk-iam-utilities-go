//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func validatePolicyVariable_PrincipalTagParameters(tagName *string) error {
	return nil
}

func validatePolicyVariable_RequestTagParameters(tagName *string) error {
	return nil
}

func validatePolicyVariable_ResourceTagParameters(tagName *string) error {
	return nil
}

