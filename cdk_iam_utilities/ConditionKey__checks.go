//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNewConditionKeyParameters(name *string, settings *ConditionKeySettings) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if settings == nil {
		return fmt.Errorf("parameter settings is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(settings, func() string { return "parameter settings" }); err != nil {
		return err
	}

	return nil
}

