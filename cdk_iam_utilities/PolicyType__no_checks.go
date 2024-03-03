//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyType) validateIsResourcePolicyForServiceParameters(service ResourcePolicyType) error {
	return nil
}

func validatePolicyType_ResourcePolicyParameters(type_ ResourcePolicyType) error {
	return nil
}

func validatePolicyType_TrustPolicyParameters(principalType PrincipalType) error {
	return nil
}

