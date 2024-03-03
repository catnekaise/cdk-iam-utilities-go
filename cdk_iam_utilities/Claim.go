package cdk_iam_utilities


// Experimental.
type Claim struct {
	// Name represents the original value of the claim/attribute.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tag Name is name of the tag corresponding to name.
	//
	// It can either match name or be a different value.
	// Experimental.
	TagName *string `field:"required" json:"tagName" yaml:"tagName"`
}

