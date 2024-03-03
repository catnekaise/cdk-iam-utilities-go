//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"
)

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) validateAppendPolicyVariableParameters(policyVariable PolicyVariable) error {
	if policyVariable == nil {
		return fmt.Errorf("parameter policyVariable is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) validateClaimParameters(claim *string) error {
	if claim == nil {
		return fmt.Errorf("parameter claim is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) validatePolicyVariableParameters(value PolicyVariable) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) validateTextParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GenericClaimsIamResourcePathBuilder) validateValueParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateGenericClaimsIamResourcePathBuilder_CreateParameters(claimsContext IClaimsContext) error {
	if claimsContext == nil {
		return fmt.Errorf("parameter claimsContext is required, but nil was provided")
	}

	return nil
}

