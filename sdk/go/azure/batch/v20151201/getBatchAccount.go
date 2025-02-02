// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20151201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains information about an Azure Batch account.
func LookupBatchAccount(ctx *pulumi.Context, args *LookupBatchAccountArgs, opts ...pulumi.InvokeOption) (*LookupBatchAccountResult, error) {
	var rv LookupBatchAccountResult
	err := ctx.Invoke("azure-native:batch/v20151201:getBatchAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBatchAccountArgs struct {
	// The name of the account.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Contains information about an Azure Batch account.
type LookupBatchAccountResult struct {
	// The endpoint used by this account to interact with the Batch services.
	AccountEndpoint string `pulumi:"accountEndpoint"`
	// The active job and job schedule quota for this Batch account.
	ActiveJobAndJobScheduleQuota int `pulumi:"activeJobAndJobScheduleQuota"`
	// The properties and status of any auto storage account associated with the account.
	AutoStorage *AutoStoragePropertiesResponse `pulumi:"autoStorage"`
	// The core quota for this Batch account.
	CoreQuota int `pulumi:"coreQuota"`
	// The ID of the resource
	Id string `pulumi:"id"`
	// The location of the resource
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The pool quota for this Batch account.
	PoolQuota int `pulumi:"poolQuota"`
	// The provisioned state of the resource
	ProvisioningState *string `pulumi:"provisioningState"`
	// The tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource
	Type string `pulumi:"type"`
}
