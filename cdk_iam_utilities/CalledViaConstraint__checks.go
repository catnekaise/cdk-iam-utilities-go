//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CalledViaConstraint) validateAssembleParameters(scope constructs.Construct, context *ConstraintAssembleContext) error {
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

func (c *jsiiProxy_CalledViaConstraint) validateIsNotNullConditionParameters(key ConditionKey) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateCalledViaConstraint_CalledViaParameters(service CalledViaServicePrincipal) error {
	if service == "" {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

func validateCalledViaConstraint_CalledViaFirstParameters(service CalledViaServicePrincipal) error {
	if service == "" {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

func validateCalledViaConstraint_CalledViaFirstAndLastParameters(firstService CalledViaServicePrincipal, lastService CalledViaServicePrincipal) error {
	if firstService == "" {
		return fmt.Errorf("parameter firstService is required, but nil was provided")
	}

	if lastService == "" {
		return fmt.Errorf("parameter lastService is required, but nil was provided")
	}

	return nil
}

func validateCalledViaConstraint_CalledViaLastParameters(service CalledViaServicePrincipal) error {
	if service == "" {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

