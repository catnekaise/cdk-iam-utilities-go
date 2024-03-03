//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"
)

func validateMappedClaims_CreateParameters(claim *string) error {
	if claim == nil {
		return fmt.Errorf("parameter claim is required, but nil was provided")
	}

	return nil
}

func validateMappedClaims_CreateMappedParameters(claims *map[string]*string) error {
	if claims == nil {
		return fmt.Errorf("parameter claims is required, but nil was provided")
	}

	return nil
}

