// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: StatusSummary.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package upgrade

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``selectionStatus`` of method StatusSummary#get.
const StatusSummary_GET_SELECTION_STATUS_SELECTED = "SELECTED"

// Possible value for ``selectionStatus`` of method StatusSummary#get.
const StatusSummary_GET_SELECTION_STATUS_DESELECTED = "DESELECTED"

// Possible value for ``selectionStatus`` of method StatusSummary#get.
const StatusSummary_GET_SELECTION_STATUS_ALL = "ALL"

func statusSummaryGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["selection_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["show_history"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["selection_status"] = "SelectionStatus"
	fieldNameMap["show_history"] = "ShowHistory"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statusSummaryGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.UpgradeStatusBindingType)
}

func statusSummaryGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["selection_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["show_history"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["selection_status"] = "SelectionStatus"
	fieldNameMap["show_history"] = "ShowHistory"
	paramsTypeMap["selection_status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["show_history"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["component_type"] = "component_type"
	queryParams["selection_status"] = "selection_status"
	queryParams["show_history"] = "show_history"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
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
		"/api/v1/upgrade/status-summary",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}