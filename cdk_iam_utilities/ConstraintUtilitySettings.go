package cdk_iam_utilities


// Experimental.
type ConstraintUtilitySettings struct {
	// Experimental.
	PolicyType PolicyType `field:"required" json:"policyType" yaml:"policyType"`
	// Experimental.
	AppendConditionValues *bool `field:"optional" json:"appendConditionValues" yaml:"appendConditionValues"`
	// Experimental.
	ClaimsContext IClaimsContext `field:"optional" json:"claimsContext" yaml:"claimsContext"`
}

