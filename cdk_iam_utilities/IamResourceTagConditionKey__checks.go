//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"
)

func validateIamResourceTagConditionKey_TagParameters(tagName *string) error {
	if tagName == nil {
		return fmt.Errorf("parameter tagName is required, but nil was provided")
	}

	return nil
}

