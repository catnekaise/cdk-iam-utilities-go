package cdk_iam_utilities


// Experimental.
type PassClaimsConstraintSettings struct {
	// Experimental.
	AllowAnyTags *bool `field:"required" json:"allowAnyTags" yaml:"allowAnyTags"`
	// Experimental.
	Claims *map[string]*string `field:"required" json:"claims" yaml:"claims"`
	// Experimental.
	SpecificallyAllowedTags *[]*string `field:"optional" json:"specificallyAllowedTags" yaml:"specificallyAllowedTags"`
}

