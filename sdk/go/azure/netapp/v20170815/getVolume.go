// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170815

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Volume resource
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:netapp/v20170815:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
}

// Volume resource
type LookupVolumeResult struct {
	// A unique file path for the volume. Used when creating mount targets
	CreationToken string `pulumi:"creationToken"`
	// Export policy rule
	ExportPolicy *VolumePropertiesResponseExportPolicy `pulumi:"exportPolicy"`
	// Unique FileSystem Identifier.
	FileSystemId string `pulumi:"fileSystemId"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// The service level of the file system
	ServiceLevel string `pulumi:"serviceLevel"`
	// The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
	SubnetId *string `pulumi:"subnetId"`
	// Resource tags
	Tags interface{} `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB.
	UsageThreshold *float64 `pulumi:"usageThreshold"`
}
