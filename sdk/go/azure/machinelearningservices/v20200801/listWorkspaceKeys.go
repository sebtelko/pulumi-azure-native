// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceKeys(ctx *pulumi.Context, args *ListWorkspaceKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceKeysResult, error) {
	var rv ListWorkspaceKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200801:listWorkspaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceKeysArgs struct {
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type ListWorkspaceKeysResult struct {
	AppInsightsInstrumentationKey string                                 `pulumi:"appInsightsInstrumentationKey"`
	ContainerRegistryCredentials  RegistryListCredentialsResultResponse  `pulumi:"containerRegistryCredentials"`
	NotebookAccessKeys            *NotebookListCredentialsResultResponse `pulumi:"notebookAccessKeys"`
	UserStorageKey                string                                 `pulumi:"userStorageKey"`
	UserStorageResourceId         string                                 `pulumi:"userStorageResourceId"`
}
