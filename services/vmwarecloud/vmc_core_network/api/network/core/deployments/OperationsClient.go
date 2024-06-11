// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Operations
// Used by client-side stubs.

package deployments

import (
	vapiStdErrors_ "github.com/zhengxiexie/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/zhengxiexie/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type OperationsClient interface {

	// create deployment operation
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @param operationParam Operation request payload
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Bad Request
	// @throws ConcurrentChange Conflict
	CreateDeploymentOperation(orgIdParam string, deploymentIdParam string, operationParam vmwarecloudVmc_core_networkModel.Operation) (vmwarecloudVmc_core_networkModel.Operation, error)
}

type operationsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewOperationsClient(connector vapiProtocolClient_.Connector) *operationsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.operations")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create_deployment_operation": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create_deployment_operation"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	oIface := operationsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &oIface
}

func (oIface *operationsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := oIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (oIface *operationsClient) CreateDeploymentOperation(orgIdParam string, deploymentIdParam string, operationParam vmwarecloudVmc_core_networkModel.Operation) (vmwarecloudVmc_core_networkModel.Operation, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := operationsCreateDeploymentOperationRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(operationsCreateDeploymentOperationInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.Operation
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.operations", "create_deployment_operation", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.Operation
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OperationsCreateDeploymentOperationOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.Operation), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
