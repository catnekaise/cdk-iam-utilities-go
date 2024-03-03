//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func validateNewConditionKeyParameters(name *string, settings *ConditionKeySettings) error {
	return nil
}

