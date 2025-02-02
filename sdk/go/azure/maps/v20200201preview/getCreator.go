// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure resource which represents Maps Creator product and provides ability to manage private location data.
func LookupCreator(ctx *pulumi.Context, args *LookupCreatorArgs, opts ...pulumi.InvokeOption) (*LookupCreatorResult, error) {
	var rv LookupCreatorResult
	err := ctx.Invoke("azure-native:maps/v20200201preview:getCreator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCreatorArgs struct {
	// The name of the Maps Account.
	AccountName string `pulumi:"accountName"`
	// The name of the Maps Creator instance.
	CreatorName string `pulumi:"creatorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure resource which represents Maps Creator product and provides ability to manage private location data.
type LookupCreatorResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The Creator resource properties.
	Properties CreatorPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
