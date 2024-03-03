//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"
)

func (i *jsiiProxy_IamResourcePathBuilder) validateAppendPolicyVariableParameters(policyVariable PolicyVariable) error {
	if policyVariable == nil {
		return fmt.Errorf("parameter policyVariable is required, but nil was provided")
	}

	return nil
}

func validateNewIamResourcePathBuilderParameters(path *[]*string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

