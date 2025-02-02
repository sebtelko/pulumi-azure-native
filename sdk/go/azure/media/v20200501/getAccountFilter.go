// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Account Filter.
func LookupAccountFilter(ctx *pulumi.Context, args *LookupAccountFilterArgs, opts ...pulumi.InvokeOption) (*LookupAccountFilterResult, error) {
	var rv LookupAccountFilterResult
	err := ctx.Invoke("azure-native:media/v20200501:getAccountFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountFilterArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Account Filter name
	FilterName string `pulumi:"filterName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Account Filter.
type LookupAccountFilterResult struct {
	// The first quality.
	FirstQuality *FirstQualityResponse `pulumi:"firstQuality"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The presentation time range.
	PresentationTimeRange *PresentationTimeRangeResponse `pulumi:"presentationTimeRange"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tracks selection conditions.
	Tracks []FilterTrackSelectionResponse `pulumi:"tracks"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
