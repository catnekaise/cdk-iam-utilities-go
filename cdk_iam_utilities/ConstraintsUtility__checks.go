//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_ConstraintsUtility) validateAppendGrantParameters(scope constructs.Construct, settings *ConstraintUtilitySettings, grant awsiam.Grant) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if settings == nil {
		return fmt.Errorf("parameter settings is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(settings, func() string { return "parameter settings" }); err != nil {
		return err
	}

	if grant == nil {
		return fmt.Errorf("parameter grant is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConstraintsUtility) validateAppendPolicyParameters(scope constructs.Construct, settings *ConstraintUtilitySettings, policyStatement awsiam.PolicyStatement) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if settings == nil {
		return fmt.Errorf("parameter settings is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(settings, func() string { return "parameter settings" }); err != nil {
		return err
	}

	if policyStatement == nil {
		return fmt.Errorf("parameter policyStatement is required, but nil was provided")
	}

	return nil
}

func validateConstraintsUtility_ForConstraintsParameters(constraints *[]Constraint) error {
	if constraints == nil {
		return fmt.Errorf("parameter constraints is required, but nil was provided")
	}

	return nil
}

