// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Media Services account.
func LookupMediaService(ctx *pulumi.Context, args *LookupMediaServiceArgs, opts ...pulumi.InvokeOption) (*LookupMediaServiceResult, error) {
	var rv LookupMediaServiceResult
	err := ctx.Invoke("azure-native:media/v20180701:getMediaService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMediaServiceArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Media Services account.
type LookupMediaServiceResult struct {
	// Fully qualified resource ID for the resource.
	Id string `pulumi:"id"`
	// The Azure Region of the resource.
	Location *string `pulumi:"location"`
	// The Media Services account ID.
	MediaServiceId string `pulumi:"mediaServiceId"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The storage accounts for this resource.
	StorageAccounts []StorageAccountResponse `pulumi:"storageAccounts"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
