//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IamResourcePathBuilder) validateAppendPolicyVariableParameters(policyVariable PolicyVariable) error {
	return nil
}

func validateNewIamResourcePathBuilderParameters(path *[]*string) error {
	return nil
}

