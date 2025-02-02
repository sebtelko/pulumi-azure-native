// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package engagementfabric

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List of the EngagementFabric channel descriptions
// API Version: 2018-09-01-preview.
func ListAccountChannelTypes(ctx *pulumi.Context, args *ListAccountChannelTypesArgs, opts ...pulumi.InvokeOption) (*ListAccountChannelTypesResult, error) {
	var rv ListAccountChannelTypesResult
	err := ctx.Invoke("azure-native:engagementfabric:listAccountChannelTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountChannelTypesArgs struct {
	// Account Name
	AccountName string `pulumi:"accountName"`
	// Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of the EngagementFabric channel descriptions
type ListAccountChannelTypesResult struct {
	// Channel descriptions
	Value []ChannelTypeDescriptionResponse `pulumi:"value"`
}
