// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200515preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Linked workspace.
func LookupLinkedWorkspace(ctx *pulumi.Context, args *LookupLinkedWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedWorkspaceResult, error) {
	var rv LookupLinkedWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200515preview:getLinkedWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedWorkspaceArgs struct {
	// Friendly name of the linked workspace
	LinkName string `pulumi:"linkName"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Linked workspace.
type LookupLinkedWorkspaceResult struct {
	// ResourceId of the link of the linked workspace.
	Id string `pulumi:"id"`
	// Friendly name of the linked workspace.
	Name string `pulumi:"name"`
	// LinkedWorkspace specific properties.
	Properties LinkedWorkspacePropsResponse `pulumi:"properties"`
	// Resource type of linked workspace.
	Type string `pulumi:"type"`
}
