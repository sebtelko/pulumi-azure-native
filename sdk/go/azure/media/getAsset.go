// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Asset.
// API Version: 2020-05-01.
func LookupAsset(ctx *pulumi.Context, args *LookupAssetArgs, opts ...pulumi.InvokeOption) (*LookupAssetResult, error) {
	var rv LookupAssetResult
	err := ctx.Invoke("azure-native:media:getAsset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssetArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Asset name.
	AssetName string `pulumi:"assetName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Asset.
type LookupAssetResult struct {
	// The alternate ID of the Asset.
	AlternateId *string `pulumi:"alternateId"`
	// The Asset ID.
	AssetId string `pulumi:"assetId"`
	// The name of the asset blob container.
	Container *string `pulumi:"container"`
	// The creation date of the Asset.
	Created string `pulumi:"created"`
	// The Asset description.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The last modified date of the Asset.
	LastModified string `pulumi:"lastModified"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The name of the storage account.
	StorageAccountName *string `pulumi:"storageAccountName"`
	// The Asset encryption format. One of None or MediaStorageEncryption.
	StorageEncryptionFormat string `pulumi:"storageEncryptionFormat"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
