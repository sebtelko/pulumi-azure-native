// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AzureStorageInfo dictionary resource.
// API Version: 2020-12-01.
func ListWebAppAzureStorageAccounts(ctx *pulumi.Context, args *ListWebAppAzureStorageAccountsArgs, opts ...pulumi.InvokeOption) (*ListWebAppAzureStorageAccountsResult, error) {
	var rv ListWebAppAzureStorageAccountsResult
	err := ctx.Invoke("azure-native:web:listWebAppAzureStorageAccounts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppAzureStorageAccountsArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// AzureStorageInfo dictionary resource.
type ListWebAppAzureStorageAccountsResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Azure storage accounts.
	Properties map[string]AzureStorageInfoValueResponse `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}
