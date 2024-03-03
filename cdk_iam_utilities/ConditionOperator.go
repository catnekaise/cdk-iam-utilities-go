package cdk_iam_utilities


// Experimental.
type ConditionOperator string

const (
	// Experimental.
	ConditionOperator_STRING_EQUALS ConditionOperator = "STRING_EQUALS"
	// Experimental.
	ConditionOperator_STRING_NOT_EQUALS ConditionOperator = "STRING_NOT_EQUALS"
	// Experimental.
	ConditionOperator_STRING_EQUALS_IGNORECASE ConditionOperator = "STRING_EQUALS_IGNORECASE"
	// Experimental.
	ConditionOperator_STRING_NOT_EQUALS_IGNORECASE ConditionOperator = "STRING_NOT_EQUALS_IGNORECASE"
	// Experimental.
	ConditionOperator_STRING_LIKE ConditionOperator = "STRING_LIKE"
	// Experimental.
	ConditionOperator_STRING_NOT_LIKE ConditionOperator = "STRING_NOT_LIKE"
	// Experimental.
	ConditionOperator_STRING_EQUALS_IFEXISTS ConditionOperator = "STRING_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_STRING_NOT_EQUALS_IFEXISTS ConditionOperator = "STRING_NOT_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_STRING_EQUALS_IGNORECASE_IFEXISTS ConditionOperator = "STRING_EQUALS_IGNORECASE_IFEXISTS"
	// Experimental.
	ConditionOperator_STRING_NOT_EQUALS_IGNORECASE_IFEXISTS ConditionOperator = "STRING_NOT_EQUALS_IGNORECASE_IFEXISTS"
	// Experimental.
	ConditionOperator_STRING_LIKE_IFEXISTS ConditionOperator = "STRING_LIKE_IFEXISTS"
	// Experimental.
	ConditionOperator_STRING_NOT_LIKE_IFEXISTS ConditionOperator = "STRING_NOT_LIKE_IFEXISTS"
	// Experimental.
	ConditionOperator_DATE_EQUALS ConditionOperator = "DATE_EQUALS"
	// Experimental.
	ConditionOperator_DATE_NOT_EQUALS ConditionOperator = "DATE_NOT_EQUALS"
	// Experimental.
	ConditionOperator_DATE_LESS_THAN ConditionOperator = "DATE_LESS_THAN"
	// Experimental.
	ConditionOperator_DATE_LESS_THAN_EQUALS ConditionOperator = "DATE_LESS_THAN_EQUALS"
	// Experimental.
	ConditionOperator_DATE_GREATER_THAN ConditionOperator = "DATE_GREATER_THAN"
	// Experimental.
	ConditionOperator_DATE_GREATER_THAN_EQUALS ConditionOperator = "DATE_GREATER_THAN_EQUALS"
	// Experimental.
	ConditionOperator_DATE_EQUALS_IFEXISTS ConditionOperator = "DATE_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_DATE_NOT_EQUALS_IFEXISTS ConditionOperator = "DATE_NOT_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_DATE_LESS_THAN_IFEXISTS ConditionOperator = "DATE_LESS_THAN_IFEXISTS"
	// Experimental.
	ConditionOperator_DATE_LESS_THAN_EQUALS_IFEXISTS ConditionOperator = "DATE_LESS_THAN_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_DATE_GREATER_THAN_IFEXISTS ConditionOperator = "DATE_GREATER_THAN_IFEXISTS"
	// Experimental.
	ConditionOperator_DATE_GREATER_THAN_EQUALS_IFEXISTS ConditionOperator = "DATE_GREATER_THAN_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_NUMERIC_EQUALS ConditionOperator = "NUMERIC_EQUALS"
	// Experimental.
	ConditionOperator_NUMERIC_NOT_EQUALS ConditionOperator = "NUMERIC_NOT_EQUALS"
	// Experimental.
	ConditionOperator_NUMERIC_LESS_THAN ConditionOperator = "NUMERIC_LESS_THAN"
	// Experimental.
	ConditionOperator_NUMERIC_LESS_THAN_EQUALS ConditionOperator = "NUMERIC_LESS_THAN_EQUALS"
	// Experimental.
	ConditionOperator_NUMERIC_GREATER_THAN ConditionOperator = "NUMERIC_GREATER_THAN"
	// Experimental.
	ConditionOperator_NUMERIC_GREATER_THAN_EQUALS ConditionOperator = "NUMERIC_GREATER_THAN_EQUALS"
	// Experimental.
	ConditionOperator_NUMERIC_EQUALS_IFEXISTS ConditionOperator = "NUMERIC_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_NUMERIC_NOT_EQUALS_IFEXISTS ConditionOperator = "NUMERIC_NOT_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_NUMERIC_LESS_THAN_IFEXISTS ConditionOperator = "NUMERIC_LESS_THAN_IFEXISTS"
	// Experimental.
	ConditionOperator_NUMERIC_LESS_THAN_EQUALS_IFEXISTS ConditionOperator = "NUMERIC_LESS_THAN_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_NUMERIC_GREATER_THAN_IFEXISTS ConditionOperator = "NUMERIC_GREATER_THAN_IFEXISTS"
	// Experimental.
	ConditionOperator_NUMERIC_GREATER_THAN_EQUALS_IFEXISTS ConditionOperator = "NUMERIC_GREATER_THAN_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_FOR_ANY_VALUE_STRING_LIKE ConditionOperator = "FOR_ANY_VALUE_STRING_LIKE"
	// Experimental.
	ConditionOperator_FOR_ANY_VALUE_STRING_EQUALS ConditionOperator = "FOR_ANY_VALUE_STRING_EQUALS"
	// Experimental.
	ConditionOperator_FOR_ALL_VALUES_STRING_LIKE ConditionOperator = "FOR_ALL_VALUES_STRING_LIKE"
	// Experimental.
	ConditionOperator_FOR_ALL_VALUES_STRING_EQUALS ConditionOperator = "FOR_ALL_VALUES_STRING_EQUALS"
	// Experimental.
	ConditionOperator_FOR_ALL_VALUES_STRING_EQUALS_IGNORECASE ConditionOperator = "FOR_ALL_VALUES_STRING_EQUALS_IGNORECASE"
	// Experimental.
	ConditionOperator_BOOL ConditionOperator = "BOOL"
	// Experimental.
	ConditionOperator_BOOL_IFEXISTS ConditionOperator = "BOOL_IFEXISTS"
	// Experimental.
	ConditionOperator_BINARY_EQUALS ConditionOperator = "BINARY_EQUALS"
	// Experimental.
	ConditionOperator_ARN_EQUALS ConditionOperator = "ARN_EQUALS"
	// Experimental.
	ConditionOperator_ARN_LIKE ConditionOperator = "ARN_LIKE"
	// Experimental.
	ConditionOperator_ARN_NOT_EQUALS ConditionOperator = "ARN_NOT_EQUALS"
	// Experimental.
	ConditionOperator_ARN_NOT_LIKE ConditionOperator = "ARN_NOT_LIKE"
	// Experimental.
	ConditionOperator_ARN_EQUALS_IFEXISTS ConditionOperator = "ARN_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_ARN_LIKE_IFEXISTS ConditionOperator = "ARN_LIKE_IFEXISTS"
	// Experimental.
	ConditionOperator_ARN_NOT_EQUALS_IFEXISTS ConditionOperator = "ARN_NOT_EQUALS_IFEXISTS"
	// Experimental.
	ConditionOperator_ARN_NOT_LIKE_IFEXISTS ConditionOperator = "ARN_NOT_LIKE_IFEXISTS"
	// Experimental.
	ConditionOperator_IP_ADDRESS ConditionOperator = "IP_ADDRESS"
	// Experimental.
	ConditionOperator_IP_ADDRESS_IFEXISTS ConditionOperator = "IP_ADDRESS_IFEXISTS"
	// Experimental.
	ConditionOperator_NOT_IP_ADDRESS ConditionOperator = "NOT_IP_ADDRESS"
	// Experimental.
	ConditionOperator_NOT_IP_ADDRESS_IFEXISTS ConditionOperator = "NOT_IP_ADDRESS_IFEXISTS"
	// Experimental.
	ConditionOperator_NULL ConditionOperator = "NULL"
)

