// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170605preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Server Endpoint object.
func LookupServerEndpoint(ctx *pulumi.Context, args *LookupServerEndpointArgs, opts ...pulumi.InvokeOption) (*LookupServerEndpointResult, error) {
	var rv LookupServerEndpointResult
	err := ctx.Invoke("azure-native:storagesync/v20170605preview:getServerEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerEndpointArgs struct {
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Server Endpoint object.
	ServerEndpointName string `pulumi:"serverEndpointName"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	// Name of Sync Group resource.
	SyncGroupName string `pulumi:"syncGroupName"`
}

// Server Endpoint object.
type LookupServerEndpointResult struct {
	// Bytes in progress
	ByteProgress *int `pulumi:"byteProgress"`
	// Bytes total
	ByteTotal *int `pulumi:"byteTotal"`
	// Cloud Tiering.
	CloudTiering *string `pulumi:"cloudTiering"`
	// current progress type.
	CurrentProgressType *string `pulumi:"currentProgressType"`
	// Friendly Name
	FriendlyName *string `pulumi:"friendlyName"`
	// The id of the resource.
	Id string `pulumi:"id"`
	// Item download error count.
	ItemDownloadErrorCount *int `pulumi:"itemDownloadErrorCount"`
	// Item Progress Count
	ItemProgressCount *int `pulumi:"itemProgressCount"`
	// Item Total Count
	ItemTotalCount *int `pulumi:"itemTotalCount"`
	// Item Upload Error Count.
	ItemUploadErrorCount *int `pulumi:"itemUploadErrorCount"`
	// Last Sync Success
	LastSyncSuccess *string `pulumi:"lastSyncSuccess"`
	// ServerEndpoint lastWorkflowId
	LastWorkflowId *string `pulumi:"lastWorkflowId"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// ServerEndpoint Provisioning State
	ProvisioningState *string `pulumi:"provisioningState"`
	// Server Local path.
	ServerLocalPath *string `pulumi:"serverLocalPath"`
	// Server Resource Id.
	ServerResourceId *string `pulumi:"serverResourceId"`
	// sync error context.
	SyncErrorContext *string `pulumi:"syncErrorContext"`
	// Sync Error Direction.
	SyncErrorDirection *string `pulumi:"syncErrorDirection"`
	// Sync Error State
	SyncErrorState *string `pulumi:"syncErrorState"`
	// Sync Error State Timestamp
	SyncErrorStateTimestamp *string `pulumi:"syncErrorStateTimestamp"`
	// Total progress
	TotalProgress *int `pulumi:"totalProgress"`
	// The type of the resource
	Type string `pulumi:"type"`
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent *int `pulumi:"volumeFreeSpacePercent"`
}
