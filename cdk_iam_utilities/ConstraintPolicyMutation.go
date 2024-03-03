package cdk_iam_utilities


// Experimental.
type ConstraintPolicyMutation struct {
	// Experimental.
	Key ConditionKey `field:"required" json:"key" yaml:"key"`
	// Experimental.
	Operator ConditionOperator `field:"required" json:"operator" yaml:"operator"`
	// Experimental.
	Type ConstraintPolicyMutationType `field:"required" json:"type" yaml:"type"`
	// Experimental.
	Value *[]interface{} `field:"required" json:"value" yaml:"value"`
	// Experimental.
	ActionsMatchService *string `field:"optional" json:"actionsMatchService" yaml:"actionsMatchService"`
	// Experimental.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Experimental.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

