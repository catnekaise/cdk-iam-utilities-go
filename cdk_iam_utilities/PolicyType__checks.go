//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"
)

func (p *jsiiProxy_PolicyType) validateIsResourcePolicyForServiceParameters(service ResourcePolicyType) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

func validatePolicyType_ResourcePolicyParameters(type_ ResourcePolicyType) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

func validatePolicyType_TrustPolicyParameters(principalType PrincipalType) error {
	if principalType == nil {
		return fmt.Errorf("parameter principalType is required, but nil was provided")
	}

	return nil
}

