// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ExternalAddress
// Used by client-side stubs.

package ports

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ExternalAddressClient interface {

	// Delete a external address binding
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param vpcIdParam (required)
	// @param subnetIdParam (required)
	// @param portIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, projectIdParam string, vpcIdParam string, subnetIdParam string, portIdParam string) error

	// Read a external address binding.
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param vpcIdParam (required)
	// @param subnetIdParam (required)
	// @param portIdParam (required)
	// @return com.vmware.nsx_policy.model.ExternalAddressBinding
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, vpcIdParam string, subnetIdParam string, portIdParam string) (nsx_policyModel.ExternalAddressBinding, error)

	// If a external address binding is not already present, create a new external address binding . If it already exists, update the external address binding.
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param vpcIdParam (required)
	// @param subnetIdParam (required)
	// @param portIdParam (required)
	// @param externalAddressBindingParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, vpcIdParam string, subnetIdParam string, portIdParam string, externalAddressBindingParam nsx_policyModel.ExternalAddressBinding) error

	// If a external address binding is not already present, create a new external address binding. If it already exists, update the external address binding. This operation will fully replace the object.
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param vpcIdParam (required)
	// @param subnetIdParam (required)
	// @param portIdParam (required)
	// @param externalAddressBindingParam (required)
	// @return com.vmware.nsx_policy.model.ExternalAddressBinding
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, vpcIdParam string, subnetIdParam string, portIdParam string, externalAddressBindingParam nsx_policyModel.ExternalAddressBinding) (nsx_policyModel.ExternalAddressBinding, error)
}

type externalAddressClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewExternalAddressClient(connector vapiProtocolClient_.Connector) *externalAddressClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.vpcs.subnets.ports.external_address")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	eIface := externalAddressClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *externalAddressClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *externalAddressClient) Delete(orgIdParam string, projectIdParam string, vpcIdParam string, subnetIdParam string, portIdParam string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := externalAddressDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(externalAddressDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("SubnetId", subnetIdParam)
	sv.AddStructField("PortId", portIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.subnets.ports.external_address", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *externalAddressClient) Get(orgIdParam string, projectIdParam string, vpcIdParam string, subnetIdParam string, portIdParam string) (nsx_policyModel.ExternalAddressBinding, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := externalAddressGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(externalAddressGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("SubnetId", subnetIdParam)
	sv.AddStructField("PortId", portIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ExternalAddressBinding
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.subnets.ports.external_address", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ExternalAddressBinding
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ExternalAddressGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ExternalAddressBinding), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *externalAddressClient) Patch(orgIdParam string, projectIdParam string, vpcIdParam string, subnetIdParam string, portIdParam string, externalAddressBindingParam nsx_policyModel.ExternalAddressBinding) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := externalAddressPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(externalAddressPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("SubnetId", subnetIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("ExternalAddressBinding", externalAddressBindingParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.subnets.ports.external_address", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *externalAddressClient) Update(orgIdParam string, projectIdParam string, vpcIdParam string, subnetIdParam string, portIdParam string, externalAddressBindingParam nsx_policyModel.ExternalAddressBinding) (nsx_policyModel.ExternalAddressBinding, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := externalAddressUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(externalAddressUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("SubnetId", subnetIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("ExternalAddressBinding", externalAddressBindingParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ExternalAddressBinding
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.subnets.ports.external_address", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ExternalAddressBinding
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ExternalAddressUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ExternalAddressBinding), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
