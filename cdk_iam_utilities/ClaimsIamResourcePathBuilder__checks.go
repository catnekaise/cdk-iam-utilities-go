//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (c *jsiiProxy_ClaimsIamResourcePathBuilder) validateAppendPolicyVariableParameters(policyVariable PolicyVariable) error {
	if policyVariable == nil {
		return fmt.Errorf("parameter policyVariable is required, but nil was provided")
	}

	return nil
}

func validateNewClaimsIamResourcePathBuilderParameters(options *ClaimsIamResourcePathBuilderSettings, path *[]*string) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

