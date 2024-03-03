//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (t *jsiiProxy_TagKeysConstraint) validateAssembleParameters(scope constructs.Construct, _arg *ConstraintAssembleContext) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if _arg == nil {
		return fmt.Errorf("parameter _arg is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_arg, func() string { return "parameter _arg" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TagKeysConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateTagKeysConstraint_CreateParameters(operator StringMultiValueConditionOperator, isNotNull *bool, value *string) error {
	if operator == "" {
		return fmt.Errorf("parameter operator is required, but nil was provided")
	}

	if isNotNull == nil {
		return fmt.Errorf("parameter isNotNull is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateTagKeysConstraint_RequireTagsEqualsParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

