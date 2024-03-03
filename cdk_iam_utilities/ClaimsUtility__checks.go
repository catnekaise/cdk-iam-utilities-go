//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_ClaimsUtility) validatePrincipalTagConditionParameters(claim *string) error {
	if claim == nil {
		return fmt.Errorf("parameter claim is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ClaimsUtility) validateRequestTagConditionParameters(claim *string) error {
	if claim == nil {
		return fmt.Errorf("parameter claim is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ClaimsUtility) validateTagNameParameters(scope constructs.Construct, claim *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if claim == nil {
		return fmt.Errorf("parameter claim is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ClaimsUtility) validateTagNameForClaimParameters(claim *string) error {
	if claim == nil {
		return fmt.Errorf("parameter claim is required, but nil was provided")
	}

	return nil
}

func validateClaimsUtility_ForContextParameters(context IClaimsContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

