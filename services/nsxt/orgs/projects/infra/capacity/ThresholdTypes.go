// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Threshold.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package capacity

import (
	vapiBindings_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/zhengxiexie/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func thresholdGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ThresholdGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.CapacityThresholdListBindingType)
}

func thresholdGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/capacity/threshold",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func thresholdPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["capacity_threshold"] = vapiBindings_.NewReferenceType(nsx_policyModel.CapacityThresholdBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["capacity_threshold"] = "CapacityThreshold"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ThresholdPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func thresholdPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["capacity_threshold"] = vapiBindings_.NewReferenceType(nsx_policyModel.CapacityThresholdBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["capacity_threshold"] = "CapacityThreshold"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["capacity_threshold"] = vapiBindings_.NewReferenceType(nsx_policyModel.CapacityThresholdBindingType)
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"capacity_threshold",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/capacity/threshold",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func thresholdUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["capacity_threshold_list"] = vapiBindings_.NewReferenceType(nsx_policyModel.CapacityThresholdListBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["capacity_threshold_list"] = "CapacityThresholdList"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ThresholdUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.CapacityThresholdListBindingType)
}

func thresholdUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["capacity_threshold_list"] = vapiBindings_.NewReferenceType(nsx_policyModel.CapacityThresholdListBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["capacity_threshold_list"] = "CapacityThresholdList"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["capacity_threshold_list"] = vapiBindings_.NewReferenceType(nsx_policyModel.CapacityThresholdListBindingType)
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"capacity_threshold_list",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/capacity/threshold",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
