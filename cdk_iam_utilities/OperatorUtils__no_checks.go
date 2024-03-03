//go:build no_runtime_type_checking

package cdk_iam_utilities

// Building without runtime type checking enabled, so all the below just return nil

func validateOperatorUtils_ArraySupportParameters(value ConditionOperator) error {
	return nil
}

func validateOperatorUtils_ConvertParameters(value interface{}) error {
	return nil
}

func validateOperatorUtils_OperatorIsSupportedParameters(supportedOperators *[]*string, operator ConditionOperator) error {
	return nil
}

func validateOperatorUtils_OperatorShortNameParameters(operator ConditionOperator) error {
	return nil
}

