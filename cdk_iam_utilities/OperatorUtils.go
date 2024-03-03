package cdk_iam_utilities

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/catnekaise/cdk-iam-utilities-go/cdk_iam_utilities/jsii"
)

// Experimental.
type OperatorUtils interface {
}

// The jsii proxy struct for OperatorUtils
type jsiiProxy_OperatorUtils struct {
	_ byte // padding
}

// Experimental.
func OperatorUtils_ArraySupport(value ConditionOperator) *bool {
	_init_.Initialize()

	if err := validateOperatorUtils_ArraySupportParameters(value); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"arraySupport",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func OperatorUtils_Convert(value interface{}) ConditionOperator {
	_init_.Initialize()

	if err := validateOperatorUtils_ConvertParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionOperator

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"convert",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func OperatorUtils_OperatorIsSupported(supportedOperators *[]*string, operator ConditionOperator) *bool {
	_init_.Initialize()

	if err := validateOperatorUtils_OperatorIsSupportedParameters(supportedOperators, operator); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"operatorIsSupported",
		[]interface{}{supportedOperators, operator},
		&returns,
	)

	return returns
}

// Experimental.
func OperatorUtils_OperatorShortName(operator ConditionOperator) *string {
	_init_.Initialize()

	if err := validateOperatorUtils_OperatorShortNameParameters(operator); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"operatorShortName",
		[]interface{}{operator},
		&returns,
	)

	return returns
}

func OperatorUtils_Arn() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"Arn",
		&returns,
	)
	return returns
}

func OperatorUtils_Binary() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"Binary",
		&returns,
	)
	return returns
}

func OperatorUtils_Bool() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"Bool",
		&returns,
	)
	return returns
}

func OperatorUtils_Date() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"Date",
		&returns,
	)
	return returns
}

func OperatorUtils_IpAddress() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"IpAddress",
		&returns,
	)
	return returns
}

func OperatorUtils_Many() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"Many",
		&returns,
	)
	return returns
}

func OperatorUtils_Numeric() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"Numeric",
		&returns,
	)
	return returns
}

func OperatorUtils_String() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		"String",
		&returns,
	)
	return returns
}

