// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Members.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package deployment_groups

import (
	vapiBindings_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/protocol"
	vmwarecloudVmc_core_inventoryModel "github.com/zhengxiexie/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/model"
	"reflect"
)

func membersGetCoreDeploymentGroupMembersAsOperatorInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["deployment_group_id"] = vapiBindings_.NewStringType()
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["org_id"] = "OrgId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MembersGetCoreDeploymentGroupMembersAsOperatorOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmwarecloudVmc_core_inventoryModel.GroupMemberBindingType), reflect.TypeOf([]vmwarecloudVmc_core_inventoryModel.GroupMember{}))
}

func membersGetCoreDeploymentGroupMembersAsOperatorRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["deployment_group_id"] = vapiBindings_.NewStringType()
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["org_id"] = "OrgId"
	paramsTypeMap["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["deployment_group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["deploymentGroupId"] = vapiBindings_.NewStringType()
	pathParams["deployment_group_id"] = "deploymentGroupId"
	queryParams["org_id"] = "org_id"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"/api/inventory/operator/core/deployment-groups/{deploymentGroupId}/members",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.not_found": 404})
}
