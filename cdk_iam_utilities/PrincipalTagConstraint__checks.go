//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_PrincipalTagConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(context, func() string { return "parameter context" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PrincipalTagConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validatePrincipalTagConstraint_StringEqualsParameters(tagName *string, value *string) error {
	if tagName == nil {
		return fmt.Errorf("parameter tagName is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

