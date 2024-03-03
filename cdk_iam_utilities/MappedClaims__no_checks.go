//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func validateMappedClaims_CreateParameters(claim *string) error {
	return nil
}

func validateMappedClaims_CreateMappedParameters(claims *map[string]*string) error {
	return nil
}

