// Experimental utilities intended for AWS CDK IAM
package cdk_iam_utilities

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@catnekaise/cdk-iam-utilities.ArnConditionOperator",
		reflect.TypeOf((*ArnConditionOperator)(nil)).Elem(),
		map[string]interface{}{
			"ARN_EQUALS": ArnConditionOperator_ARN_EQUALS,
			"ARN_LIKE": ArnConditionOperator_ARN_LIKE,
			"ARN_NOT_EQUALS": ArnConditionOperator_ARN_NOT_EQUALS,
			"ARN_NOT_LIKE": ArnConditionOperator_ARN_NOT_LIKE,
			"ARN_EQUALS_IFEXISTS": ArnConditionOperator_ARN_EQUALS_IFEXISTS,
			"ARN_LIKE_IFEXISTS": ArnConditionOperator_ARN_LIKE_IFEXISTS,
			"ARN_NOT_EQUALS_IFEXISTS": ArnConditionOperator_ARN_NOT_EQUALS_IFEXISTS,
			"ARN_NOT_LIKE_IFEXISTS": ArnConditionOperator_ARN_NOT_LIKE_IFEXISTS,
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.AwsFederatedProviderConditionKey",
		reflect.TypeOf((*AwsFederatedProviderConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toCognitoIdentityConstraint", GoMethod: "ToCognitoIdentityConstraint"},
			_jsii_.MemberMethod{JsiiMethod: "toConstraint", GoMethod: "ToConstraint"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsFederatedProviderConditionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConditionKey)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.AwsPrincipalTagConditionKey",
		reflect.TypeOf((*AwsPrincipalTagConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsPrincipalTagConditionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConditionKey)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.AwsRequestTagConditionKey",
		reflect.TypeOf((*AwsRequestTagConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsRequestTagConditionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConditionKey)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.AwsResourceTagConditionKey",
		reflect.TypeOf((*AwsResourceTagConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsResourceTagConditionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConditionKey)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.AwsSourceVpcConditionKey",
		reflect.TypeOf((*AwsSourceVpcConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toVpcConstraint", GoMethod: "ToVpcConstraint"},
		},
		func() interface{} {
			j := jsiiProxy_AwsSourceVpcConditionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConditionKey)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.BoolConditionKey",
		reflect.TypeOf((*BoolConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BoolConditionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConditionKey)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.BoolConstraint",
		reflect.TypeOf((*BoolConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
		},
		func() interface{} {
			j := jsiiProxy_BoolConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Constraint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.CalledViaConstraint",
		reflect.TypeOf((*CalledViaConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
		},
		func() interface{} {
			j := jsiiProxy_CalledViaConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Constraint)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@catnekaise/cdk-iam-utilities.CalledViaServicePrincipal",
		reflect.TypeOf((*CalledViaServicePrincipal)(nil)).Elem(),
		map[string]interface{}{
			"AOSS": CalledViaServicePrincipal_AOSS,
			"ATHENA": CalledViaServicePrincipal_ATHENA,
			"BACKUP": CalledViaServicePrincipal_BACKUP,
			"CLOUD9": CalledViaServicePrincipal_CLOUD9,
			"CLOUDFORMATION": CalledViaServicePrincipal_CLOUDFORMATION,
			"DATABREW": CalledViaServicePrincipal_DATABREW,
			"DATAEXCHANGE": CalledViaServicePrincipal_DATAEXCHANGE,
			"DYNAMODB": CalledViaServicePrincipal_DYNAMODB,
			"IMAGEBUILDER": CalledViaServicePrincipal_IMAGEBUILDER,
			"KMS": CalledViaServicePrincipal_KMS,
			"MGN": CalledViaServicePrincipal_MGN,
			"NIMBLE": CalledViaServicePrincipal_NIMBLE,
			"OMICS": CalledViaServicePrincipal_OMICS,
			"RAM": CalledViaServicePrincipal_RAM,
			"ROBOMAKER": CalledViaServicePrincipal_ROBOMAKER,
			"SERVICECATALOG_APPREGISTRY": CalledViaServicePrincipal_SERVICECATALOG_APPREGISTRY,
			"SQLWORKBENCH": CalledViaServicePrincipal_SQLWORKBENCH,
		},
	)
	_jsii_.RegisterStruct(
		"@catnekaise/cdk-iam-utilities.Claim",
		reflect.TypeOf((*Claim)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.ClaimConstraint",
		reflect.TypeOf((*ClaimConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberProperty{JsiiProperty: "claim", GoGetter: "Claim"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
		},
		func() interface{} {
			j := jsiiProxy_ClaimConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Constraint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.ClaimsIamResourcePathBuilder",
		reflect.TypeOf((*ClaimsIamResourcePathBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "appendClaim", GoMethod: "AppendClaim"},
			_jsii_.MemberMethod{JsiiMethod: "appendPolicyVariable", GoMethod: "AppendPolicyVariable"},
			_jsii_.MemberMethod{JsiiMethod: "appendText", GoMethod: "AppendText"},
			_jsii_.MemberMethod{JsiiMethod: "appendValue", GoMethod: "AppendValue"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClaimsIamResourcePathBuilder{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IamResourcePathBuilder)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@catnekaise/cdk-iam-utilities.ClaimsIamResourcePathBuilderSettings",
		reflect.TypeOf((*ClaimsIamResourcePathBuilderSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.ClaimsUtility",
		reflect.TypeOf((*ClaimsUtility)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "knownClaims", GoGetter: "KnownClaims"},
			_jsii_.MemberProperty{JsiiProperty: "mappedClaims", GoGetter: "MappedClaims"},
			_jsii_.MemberMethod{JsiiMethod: "principalTagCondition", GoMethod: "PrincipalTagCondition"},
			_jsii_.MemberMethod{JsiiMethod: "requestTagCondition", GoMethod: "RequestTagCondition"},
			_jsii_.MemberMethod{JsiiMethod: "tagName", GoMethod: "TagName"},
			_jsii_.MemberMethod{JsiiMethod: "tagNameForClaim", GoMethod: "TagNameForClaim"},
		},
		func() interface{} {
			return &jsiiProxy_ClaimsUtility{}
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.ConditionKey",
		reflect.TypeOf((*ConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_ConditionKey{}
		},
	)
	_jsii_.RegisterStruct(
		"@catnekaise/cdk-iam-utilities.ConditionKeySettings",
		reflect.TypeOf((*ConditionKeySettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@catnekaise/cdk-iam-utilities.ConditionOperator",
		reflect.TypeOf((*ConditionOperator)(nil)).Elem(),
		map[string]interface{}{
			"STRING_EQUALS": ConditionOperator_STRING_EQUALS,
			"STRING_NOT_EQUALS": ConditionOperator_STRING_NOT_EQUALS,
			"STRING_EQUALS_IGNORECASE": ConditionOperator_STRING_EQUALS_IGNORECASE,
			"STRING_NOT_EQUALS_IGNORECASE": ConditionOperator_STRING_NOT_EQUALS_IGNORECASE,
			"STRING_LIKE": ConditionOperator_STRING_LIKE,
			"STRING_NOT_LIKE": ConditionOperator_STRING_NOT_LIKE,
			"STRING_EQUALS_IFEXISTS": ConditionOperator_STRING_EQUALS_IFEXISTS,
			"STRING_NOT_EQUALS_IFEXISTS": ConditionOperator_STRING_NOT_EQUALS_IFEXISTS,
			"STRING_EQUALS_IGNORECASE_IFEXISTS": ConditionOperator_STRING_EQUALS_IGNORECASE_IFEXISTS,
			"STRING_NOT_EQUALS_IGNORECASE_IFEXISTS": ConditionOperator_STRING_NOT_EQUALS_IGNORECASE_IFEXISTS,
			"STRING_LIKE_IFEXISTS": ConditionOperator_STRING_LIKE_IFEXISTS,
			"STRING_NOT_LIKE_IFEXISTS": ConditionOperator_STRING_NOT_LIKE_IFEXISTS,
			"DATE_EQUALS": ConditionOperator_DATE_EQUALS,
			"DATE_NOT_EQUALS": ConditionOperator_DATE_NOT_EQUALS,
			"DATE_LESS_THAN": ConditionOperator_DATE_LESS_THAN,
			"DATE_LESS_THAN_EQUALS": ConditionOperator_DATE_LESS_THAN_EQUALS,
			"DATE_GREATER_THAN": ConditionOperator_DATE_GREATER_THAN,
			"DATE_GREATER_THAN_EQUALS": ConditionOperator_DATE_GREATER_THAN_EQUALS,
			"DATE_EQUALS_IFEXISTS": ConditionOperator_DATE_EQUALS_IFEXISTS,
			"DATE_NOT_EQUALS_IFEXISTS": ConditionOperator_DATE_NOT_EQUALS_IFEXISTS,
			"DATE_LESS_THAN_IFEXISTS": ConditionOperator_DATE_LESS_THAN_IFEXISTS,
			"DATE_LESS_THAN_EQUALS_IFEXISTS": ConditionOperator_DATE_LESS_THAN_EQUALS_IFEXISTS,
			"DATE_GREATER_THAN_IFEXISTS": ConditionOperator_DATE_GREATER_THAN_IFEXISTS,
			"DATE_GREATER_THAN_EQUALS_IFEXISTS": ConditionOperator_DATE_GREATER_THAN_EQUALS_IFEXISTS,
			"NUMERIC_EQUALS": ConditionOperator_NUMERIC_EQUALS,
			"NUMERIC_NOT_EQUALS": ConditionOperator_NUMERIC_NOT_EQUALS,
			"NUMERIC_LESS_THAN": ConditionOperator_NUMERIC_LESS_THAN,
			"NUMERIC_LESS_THAN_EQUALS": ConditionOperator_NUMERIC_LESS_THAN_EQUALS,
			"NUMERIC_GREATER_THAN": ConditionOperator_NUMERIC_GREATER_THAN,
			"NUMERIC_GREATER_THAN_EQUALS": ConditionOperator_NUMERIC_GREATER_THAN_EQUALS,
			"NUMERIC_EQUALS_IFEXISTS": ConditionOperator_NUMERIC_EQUALS_IFEXISTS,
			"NUMERIC_NOT_EQUALS_IFEXISTS": ConditionOperator_NUMERIC_NOT_EQUALS_IFEXISTS,
			"NUMERIC_LESS_THAN_IFEXISTS": ConditionOperator_NUMERIC_LESS_THAN_IFEXISTS,
			"NUMERIC_LESS_THAN_EQUALS_IFEXISTS": ConditionOperator_NUMERIC_LESS_THAN_EQUALS_IFEXISTS,
			"NUMERIC_GREATER_THAN_IFEXISTS": ConditionOperator_NUMERIC_GREATER_THAN_IFEXISTS,
			"NUMERIC_GREATER_THAN_EQUALS_IFEXISTS": ConditionOperator_NUMERIC_GREATER_THAN_EQUALS_IFEXISTS,
			"FOR_ANY_VALUE_STRING_LIKE": ConditionOperator_FOR_ANY_VALUE_STRING_LIKE,
			"FOR_ANY_VALUE_STRING_EQUALS": ConditionOperator_FOR_ANY_VALUE_STRING_EQUALS,
			"FOR_ALL_VALUES_STRING_LIKE": ConditionOperator_FOR_ALL_VALUES_STRING_LIKE,
			"FOR_ALL_VALUES_STRING_EQUALS": ConditionOperator_FOR_ALL_VALUES_STRING_EQUALS,
			"FOR_ALL_VALUES_STRING_EQUALS_IGNORECASE": ConditionOperator_FOR_ALL_VALUES_STRING_EQUALS_IGNORECASE,
			"BOOL": ConditionOperator_BOOL,
			"BOOL_IFEXISTS": ConditionOperator_BOOL_IFEXISTS,
			"BINARY_EQUALS": ConditionOperator_BINARY_EQUALS,
			"ARN_EQUALS": ConditionOperator_ARN_EQUALS,
			"ARN_LIKE": ConditionOperator_ARN_LIKE,
			"ARN_NOT_EQUALS": ConditionOperator_ARN_NOT_EQUALS,
			"ARN_NOT_LIKE": ConditionOperator_ARN_NOT_LIKE,
			"ARN_EQUALS_IFEXISTS": ConditionOperator_ARN_EQUALS_IFEXISTS,
			"ARN_LIKE_IFEXISTS": ConditionOperator_ARN_LIKE_IFEXISTS,
			"ARN_NOT_EQUALS_IFEXISTS": ConditionOperator_ARN_NOT_EQUALS_IFEXISTS,
			"ARN_NOT_LIKE_IFEXISTS": ConditionOperator_ARN_NOT_LIKE_IFEXISTS,
			"IP_ADDRESS": ConditionOperator_IP_ADDRESS,
			"IP_ADDRESS_IFEXISTS": ConditionOperator_IP_ADDRESS_IFEXISTS,
			"NOT_IP_ADDRESS": ConditionOperator_NOT_IP_ADDRESS,
			"NOT_IP_ADDRESS_IFEXISTS": ConditionOperator_NOT_IP_ADDRESS_IFEXISTS,
			"NULL": ConditionOperator_NULL,
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.Constraint",
		reflect.TypeOf((*Constraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
		},
		func() interface{} {
			return &jsiiProxy_Constraint{}
		},
	)
	_jsii_.RegisterStruct(
		"@catnekaise/cdk-iam-utilities.ConstraintAssembleContext",
		reflect.TypeOf((*ConstraintAssembleContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@catnekaise/cdk-iam-utilities.ConstraintPolicyMutation",
		reflect.TypeOf((*ConstraintPolicyMutation)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@catnekaise/cdk-iam-utilities.ConstraintPolicyMutationType",
		reflect.TypeOf((*ConstraintPolicyMutationType)(nil)).Elem(),
		map[string]interface{}{
			"CONDITION": ConstraintPolicyMutationType_CONDITION,
		},
	)
	_jsii_.RegisterStruct(
		"@catnekaise/cdk-iam-utilities.ConstraintUtilitySettings",
		reflect.TypeOf((*ConstraintUtilitySettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.ConstraintsBuilder",
		reflect.TypeOf((*ConstraintsBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
			_jsii_.MemberProperty{JsiiProperty: "constraints", GoGetter: "Constraints"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
		},
		func() interface{} {
			j := jsiiProxy_ConstraintsBuilder{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConstraintsBuilder)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@catnekaise/cdk-iam-utilities.ConstraintsBuilderSettings",
		reflect.TypeOf((*ConstraintsBuilderSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.ConstraintsUtility",
		reflect.TypeOf((*ConstraintsUtility)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "appendGrant", GoMethod: "AppendGrant"},
			_jsii_.MemberMethod{JsiiMethod: "appendPolicy", GoMethod: "AppendPolicy"},
		},
		func() interface{} {
			return &jsiiProxy_ConstraintsUtility{}
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.DateConstraint",
		reflect.TypeOf((*DateConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
		},
		func() interface{} {
			j := jsiiProxy_DateConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Constraint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.GenericClaimsIamResourcePathBuilder",
		reflect.TypeOf((*GenericClaimsIamResourcePathBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "appendClaim", GoMethod: "AppendClaim"},
			_jsii_.MemberMethod{JsiiMethod: "appendPolicyVariable", GoMethod: "AppendPolicyVariable"},
			_jsii_.MemberMethod{JsiiMethod: "appendText", GoMethod: "AppendText"},
			_jsii_.MemberMethod{JsiiMethod: "appendValue", GoMethod: "AppendValue"},
			_jsii_.MemberMethod{JsiiMethod: "claim", GoMethod: "Claim"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "policyVariable", GoMethod: "PolicyVariable"},
			_jsii_.MemberMethod{JsiiMethod: "text", GoMethod: "Text"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "value", GoMethod: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_GenericClaimsIamResourcePathBuilder{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ClaimsIamResourcePathBuilder)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.GenericConditionKey",
		reflect.TypeOf((*GenericConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GenericConditionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConditionKey)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.GenericConstraint",
		reflect.TypeOf((*GenericConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_GenericConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Constraint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.GlobalBoolConditionKey",
		reflect.TypeOf((*GlobalBoolConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toBoolFalseConstraint", GoMethod: "ToBoolFalseConstraint"},
			_jsii_.MemberMethod{JsiiMethod: "toBoolTrueConstraint", GoMethod: "ToBoolTrueConstraint"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GlobalBoolConditionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BoolConditionKey)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.GlobalConditionKey",
		reflect.TypeOf((*GlobalConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toConstraint", GoMethod: "ToConstraint"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GlobalConditionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConditionKey)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@catnekaise/cdk-iam-utilities.IClaimsContext",
		reflect.TypeOf((*IClaimsContext)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "knownClaims", GoGetter: "KnownClaims"},
			_jsii_.MemberProperty{JsiiProperty: "mappedClaims", GoGetter: "MappedClaims"},
		},
		func() interface{} {
			return &jsiiProxy_IClaimsContext{}
		},
	)
	_jsii_.RegisterInterface(
		"@catnekaise/cdk-iam-utilities.IConstraintsBuilder",
		reflect.TypeOf((*IConstraintsBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "constraints", GoGetter: "Constraints"},
		},
		func() interface{} {
			return &jsiiProxy_IConstraintsBuilder{}
		},
	)
	_jsii_.RegisterInterface(
		"@catnekaise/cdk-iam-utilities.IIamResourcePath",
		reflect.TypeOf((*IIamResourcePath)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_IIamResourcePath{}
		},
	)
	_jsii_.RegisterInterface(
		"@catnekaise/cdk-iam-utilities.IMappedClaims",
		reflect.TypeOf((*IMappedClaims)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "claims", GoGetter: "Claims"},
		},
		func() interface{} {
			return &jsiiProxy_IMappedClaims{}
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.IamResourcePathBuilder",
		reflect.TypeOf((*IamResourcePathBuilder)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "appendPolicyVariable", GoMethod: "AppendPolicyVariable"},
			_jsii_.MemberMethod{JsiiMethod: "appendText", GoMethod: "AppendText"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IamResourcePathBuilder{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIamResourcePath)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.IamResourceTagConditionKey",
		reflect.TypeOf((*IamResourceTagConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IamResourceTagConditionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConditionKey)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@catnekaise/cdk-iam-utilities.IpAddressConditionOperator",
		reflect.TypeOf((*IpAddressConditionOperator)(nil)).Elem(),
		map[string]interface{}{
			"IP_ADDRESS": IpAddressConditionOperator_IP_ADDRESS,
			"IP_ADDRESS_IFEXISTS": IpAddressConditionOperator_IP_ADDRESS_IFEXISTS,
			"NOT_IP_ADDRESS": IpAddressConditionOperator_NOT_IP_ADDRESS,
			"NOT_IP_ADDRESS_IFEXISTS": IpAddressConditionOperator_NOT_IP_ADDRESS_IFEXISTS,
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.MappedClaims",
		reflect.TypeOf((*MappedClaims)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "claims", GoGetter: "Claims"},
		},
		func() interface{} {
			j := jsiiProxy_MappedClaims{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMappedClaims)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.NullConstraint",
		reflect.TypeOf((*NullConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
		},
		func() interface{} {
			j := jsiiProxy_NullConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Constraint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.OperatorUtils",
		reflect.TypeOf((*OperatorUtils)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_OperatorUtils{}
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.PassClaimsConstraint",
		reflect.TypeOf((*PassClaimsConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
		},
		func() interface{} {
			j := jsiiProxy_PassClaimsConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Constraint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@catnekaise/cdk-iam-utilities.PassClaimsConstraintSettings",
		reflect.TypeOf((*PassClaimsConstraintSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.PolicyType",
		reflect.TypeOf((*PolicyType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "isIdentityPolicy", GoGetter: "IsIdentityPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "isResourcePolicy", GoGetter: "IsResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "isResourcePolicyForService", GoMethod: "IsResourcePolicyForService"},
			_jsii_.MemberProperty{JsiiProperty: "isTrustPolicy", GoGetter: "IsTrustPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_PolicyType{}
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.PolicyVariable",
		reflect.TypeOf((*PolicyVariable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "defaultValue", GoGetter: "DefaultValue"},
			_jsii_.MemberProperty{JsiiProperty: "isTag", GoGetter: "IsTag"},
			_jsii_.MemberProperty{JsiiProperty: "tagName", GoGetter: "TagName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_PolicyVariable{}
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.PrincipalTagConstraint",
		reflect.TypeOf((*PrincipalTagConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
			_jsii_.MemberProperty{JsiiProperty: "tagName", GoGetter: "TagName"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
		},
		func() interface{} {
			j := jsiiProxy_PrincipalTagConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TagConstraint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.PrincipalType",
		reflect.TypeOf((*PrincipalType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "isAws", GoGetter: "IsAws"},
			_jsii_.MemberProperty{JsiiProperty: "isFederated", GoGetter: "IsFederated"},
			_jsii_.MemberProperty{JsiiProperty: "isSaml", GoGetter: "IsSaml"},
			_jsii_.MemberProperty{JsiiProperty: "isService", GoGetter: "IsService"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_PrincipalType{}
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.RequestTagConstraint",
		reflect.TypeOf((*RequestTagConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
			_jsii_.MemberProperty{JsiiProperty: "tagName", GoGetter: "TagName"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
		},
		func() interface{} {
			j := jsiiProxy_RequestTagConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TagConstraint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.ResourcePolicyType",
		reflect.TypeOf((*ResourcePolicyType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_ResourcePolicyType{}
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.ResourceTagConstraint",
		reflect.TypeOf((*ResourceTagConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
			_jsii_.MemberProperty{JsiiProperty: "tagName", GoGetter: "TagName"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
		},
		func() interface{} {
			j := jsiiProxy_ResourceTagConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TagConstraint)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@catnekaise/cdk-iam-utilities.StringConditionOperator",
		reflect.TypeOf((*StringConditionOperator)(nil)).Elem(),
		map[string]interface{}{
			"STRING_EQUALS": StringConditionOperator_STRING_EQUALS,
			"STRING_NOT_EQUALS": StringConditionOperator_STRING_NOT_EQUALS,
			"STRING_EQUALS_IGNORECASE": StringConditionOperator_STRING_EQUALS_IGNORECASE,
			"STRING_NOT_EQUALS_IGNORECASE": StringConditionOperator_STRING_NOT_EQUALS_IGNORECASE,
			"STRING_LIKE": StringConditionOperator_STRING_LIKE,
			"STRING_NOT_LIKE": StringConditionOperator_STRING_NOT_LIKE,
			"STRING_EQUALS_IFEXISTS": StringConditionOperator_STRING_EQUALS_IFEXISTS,
			"STRING_NOT_EQUALS_IFEXISTS": StringConditionOperator_STRING_NOT_EQUALS_IFEXISTS,
			"STRING_EQUALS_IGNORECASE_IFEXISTS": StringConditionOperator_STRING_EQUALS_IGNORECASE_IFEXISTS,
			"STRING_NOT_EQUALS_IGNORECASE_IFEXISTS": StringConditionOperator_STRING_NOT_EQUALS_IGNORECASE_IFEXISTS,
			"STRING_LIKE_IFEXISTS": StringConditionOperator_STRING_LIKE_IFEXISTS,
			"STRING_NOT_LIKE_IFEXISTS": StringConditionOperator_STRING_NOT_LIKE_IFEXISTS,
		},
	)
	_jsii_.RegisterEnum(
		"@catnekaise/cdk-iam-utilities.StringMultiValueConditionOperator",
		reflect.TypeOf((*StringMultiValueConditionOperator)(nil)).Elem(),
		map[string]interface{}{
			"FOR_ANY_VALUE_STRING_LIKE": StringMultiValueConditionOperator_FOR_ANY_VALUE_STRING_LIKE,
			"FOR_ALL_VALUES_STRING_LIKE": StringMultiValueConditionOperator_FOR_ALL_VALUES_STRING_LIKE,
			"FOR_ALL_VALUES_STRING_EQUALS": StringMultiValueConditionOperator_FOR_ALL_VALUES_STRING_EQUALS,
			"FOR_ALL_VALUES_STRING_EQUALS_IGNORECASE": StringMultiValueConditionOperator_FOR_ALL_VALUES_STRING_EQUALS_IGNORECASE,
			"FOR_ANY_VALUE_STRING_EQUALS": StringMultiValueConditionOperator_FOR_ANY_VALUE_STRING_EQUALS,
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.StsCognitoIdentityConstraint",
		reflect.TypeOf((*StsCognitoIdentityConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amr", GoGetter: "Amr"},
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolId", GoGetter: "IdentityPoolId"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
		},
		func() interface{} {
			j := jsiiProxy_StsCognitoIdentityConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Constraint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.StsServiceConditionKey",
		reflect.TypeOf((*StsServiceConditionKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StsServiceConditionKey{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ConditionKey)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.StsTransitiveTagKeysConstraint",
		reflect.TypeOf((*StsTransitiveTagKeysConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
		},
		func() interface{} {
			j := jsiiProxy_StsTransitiveTagKeysConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Constraint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.TagConstraint",
		reflect.TypeOf((*TagConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
		},
		func() interface{} {
			j := jsiiProxy_TagConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Constraint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@catnekaise/cdk-iam-utilities.TagKeysConstraint",
		reflect.TypeOf((*TagKeysConstraint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assemble", GoMethod: "Assemble"},
			_jsii_.MemberMethod{JsiiMethod: "isNotNullCondition", GoMethod: "IsNotNullCondition"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
		},
		func() interface{} {
			j := jsiiProxy_TagKeysConstraint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Constraint)
			return &j
		},
	)
}
