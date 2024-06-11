// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: VpcPeeringscreate.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package deployments

import (
	vapiBindings_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/protocol"
	vmwarecloudVmc_core_networkModel "github.com/zhengxiexie/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
	"reflect"
)

func vpcPeeringscreateCreateVpcPeeringInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["deployment_id"] = vapiBindings_.NewStringType()
	fields["vpc_peering_request"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.VpcPeeringRequestBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_id"] = "DeploymentId"
	fieldNameMap["vpc_peering_request"] = "VpcPeeringRequest"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func VpcPeeringscreateCreateVpcPeeringOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.ActivityBindingType)
}

func vpcPeeringscreateCreateVpcPeeringRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["deployment_id"] = vapiBindings_.NewStringType()
	fields["vpc_peering_request"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.VpcPeeringRequestBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_id"] = "DeploymentId"
	fieldNameMap["vpc_peering_request"] = "VpcPeeringRequest"
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_peering_request"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.VpcPeeringRequestBindingType)
	paramsTypeMap["deployment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["deploymentId"] = vapiBindings_.NewStringType()
	pathParams["deployment_id"] = "deploymentId"
	pathParams["org_id"] = "orgId"
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
		"vpc_peering_request",
		"POST",
		"/api/network/{orgId}/core/deployments/{deploymentId}/vpc-peerings:create",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}
