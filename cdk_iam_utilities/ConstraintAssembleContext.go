package cdk_iam_utilities

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Experimental.
type ConstraintAssembleContext struct {
	// Experimental.
	Effect awsiam.Effect `field:"required" json:"effect" yaml:"effect"`
	// Experimental.
	PolicyType PolicyType `field:"required" json:"policyType" yaml:"policyType"`
	// Experimental.
	ClaimsContext IClaimsContext `field:"optional" json:"claimsContext" yaml:"claimsContext"`
}

