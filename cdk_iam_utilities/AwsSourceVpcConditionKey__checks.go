//go:build !no_runtime_type_checking

package cdk_iam_utilities

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func (a *jsiiProxy_AwsSourceVpcConditionKey) validateToVpcConstraintParameters(vpc awsec2.IVpc) error {
	if vpc == nil {
		return fmt.Errorf("parameter vpc is required, but nil was provided")
	}

	return nil
}

