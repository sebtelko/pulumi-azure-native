// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Type of the Storage Target.
func LookupStorageTarget(ctx *pulumi.Context, args *LookupStorageTargetArgs, opts ...pulumi.InvokeOption) (*LookupStorageTargetResult, error) {
	var rv LookupStorageTargetResult
	err := ctx.Invoke("azure-native:storagecache/v20201001:getStorageTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageTargetArgs struct {
	// Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
	CacheName string `pulumi:"cacheName"`
	// Target resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Storage Target. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
	StorageTargetName string `pulumi:"storageTargetName"`
}

// Type of the Storage Target.
type LookupStorageTargetResult struct {
	// Properties when targetType is clfs.
	Clfs *ClfsTargetResponse `pulumi:"clfs"`
	// Resource ID of the Storage Target.
	Id string `pulumi:"id"`
	// List of Cache namespace junctions to target for namespace associations.
	Junctions []NamespaceJunctionResponse `pulumi:"junctions"`
	// Region name string.
	Location string `pulumi:"location"`
	// Name of the Storage Target.
	Name string `pulumi:"name"`
	// Properties when targetType is nfs3.
	Nfs3 *Nfs3TargetResponse `pulumi:"nfs3"`
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState *string `pulumi:"provisioningState"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the Storage Target.
	TargetType string `pulumi:"targetType"`
	// Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
	Type string `pulumi:"type"`
	// Properties when targetType is unknown.
	Unknown *UnknownTargetResponse `pulumi:"unknown"`
}
