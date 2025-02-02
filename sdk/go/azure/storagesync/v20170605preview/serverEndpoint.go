// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170605preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Server Endpoint object.
type ServerEndpoint struct {
	pulumi.CustomResourceState

	// Bytes in progress
	ByteProgress pulumi.IntPtrOutput `pulumi:"byteProgress"`
	// Bytes total
	ByteTotal pulumi.IntPtrOutput `pulumi:"byteTotal"`
	// Cloud Tiering.
	CloudTiering pulumi.StringPtrOutput `pulumi:"cloudTiering"`
	// current progress type.
	CurrentProgressType pulumi.StringPtrOutput `pulumi:"currentProgressType"`
	// Friendly Name
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// Item download error count.
	ItemDownloadErrorCount pulumi.IntPtrOutput `pulumi:"itemDownloadErrorCount"`
	// Item Progress Count
	ItemProgressCount pulumi.IntPtrOutput `pulumi:"itemProgressCount"`
	// Item Total Count
	ItemTotalCount pulumi.IntPtrOutput `pulumi:"itemTotalCount"`
	// Item Upload Error Count.
	ItemUploadErrorCount pulumi.IntPtrOutput `pulumi:"itemUploadErrorCount"`
	// Last Sync Success
	LastSyncSuccess pulumi.StringPtrOutput `pulumi:"lastSyncSuccess"`
	// ServerEndpoint lastWorkflowId
	LastWorkflowId pulumi.StringPtrOutput `pulumi:"lastWorkflowId"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// ServerEndpoint Provisioning State
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Server Local path.
	ServerLocalPath pulumi.StringPtrOutput `pulumi:"serverLocalPath"`
	// Server Resource Id.
	ServerResourceId pulumi.StringPtrOutput `pulumi:"serverResourceId"`
	// sync error context.
	SyncErrorContext pulumi.StringPtrOutput `pulumi:"syncErrorContext"`
	// Sync Error Direction.
	SyncErrorDirection pulumi.StringPtrOutput `pulumi:"syncErrorDirection"`
	// Sync Error State
	SyncErrorState pulumi.StringPtrOutput `pulumi:"syncErrorState"`
	// Sync Error State Timestamp
	SyncErrorStateTimestamp pulumi.StringPtrOutput `pulumi:"syncErrorStateTimestamp"`
	// Total progress
	TotalProgress pulumi.IntPtrOutput `pulumi:"totalProgress"`
	// The type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent pulumi.IntPtrOutput `pulumi:"volumeFreeSpacePercent"`
}

// NewServerEndpoint registers a new resource with the given unique name, arguments, and options.
func NewServerEndpoint(ctx *pulumi.Context,
	name string, args *ServerEndpointArgs, opts ...pulumi.ResourceOption) (*ServerEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageSyncServiceName == nil {
		return nil, errors.New("invalid value for required argument 'StorageSyncServiceName'")
	}
	if args.SyncGroupName == nil {
		return nil, errors.New("invalid value for required argument 'SyncGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storagesync/v20170605preview:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagesync:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180402:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagesync/v20180402:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180701:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagesync/v20180701:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20181001:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagesync/v20181001:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190201:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagesync/v20190201:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190301:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagesync/v20190301:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190601:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagesync/v20190601:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20191001:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagesync/v20191001:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200301:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagesync/v20200301:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200901:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:storagesync/v20200901:ServerEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerEndpoint
	err := ctx.RegisterResource("azure-native:storagesync/v20170605preview:ServerEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerEndpoint gets an existing ServerEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerEndpointState, opts ...pulumi.ResourceOption) (*ServerEndpoint, error) {
	var resource ServerEndpoint
	err := ctx.ReadResource("azure-native:storagesync/v20170605preview:ServerEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerEndpoint resources.
type serverEndpointState struct {
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
	Name *string `pulumi:"name"`
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
	Type *string `pulumi:"type"`
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent *int `pulumi:"volumeFreeSpacePercent"`
}

type ServerEndpointState struct {
	// Bytes in progress
	ByteProgress pulumi.IntPtrInput
	// Bytes total
	ByteTotal pulumi.IntPtrInput
	// Cloud Tiering.
	CloudTiering pulumi.StringPtrInput
	// current progress type.
	CurrentProgressType pulumi.StringPtrInput
	// Friendly Name
	FriendlyName pulumi.StringPtrInput
	// Item download error count.
	ItemDownloadErrorCount pulumi.IntPtrInput
	// Item Progress Count
	ItemProgressCount pulumi.IntPtrInput
	// Item Total Count
	ItemTotalCount pulumi.IntPtrInput
	// Item Upload Error Count.
	ItemUploadErrorCount pulumi.IntPtrInput
	// Last Sync Success
	LastSyncSuccess pulumi.StringPtrInput
	// ServerEndpoint lastWorkflowId
	LastWorkflowId pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// ServerEndpoint Provisioning State
	ProvisioningState pulumi.StringPtrInput
	// Server Local path.
	ServerLocalPath pulumi.StringPtrInput
	// Server Resource Id.
	ServerResourceId pulumi.StringPtrInput
	// sync error context.
	SyncErrorContext pulumi.StringPtrInput
	// Sync Error Direction.
	SyncErrorDirection pulumi.StringPtrInput
	// Sync Error State
	SyncErrorState pulumi.StringPtrInput
	// Sync Error State Timestamp
	SyncErrorStateTimestamp pulumi.StringPtrInput
	// Total progress
	TotalProgress pulumi.IntPtrInput
	// The type of the resource
	Type pulumi.StringPtrInput
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent pulumi.IntPtrInput
}

func (ServerEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverEndpointState)(nil)).Elem()
}

type serverEndpointArgs struct {
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
	// ServerEndpoint Provisioning State
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Server Endpoint object.
	ServerEndpointName *string `pulumi:"serverEndpointName"`
	// Server Local path.
	ServerLocalPath *string `pulumi:"serverLocalPath"`
	// Server Resource Id.
	ServerResourceId *string `pulumi:"serverResourceId"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	// sync error context.
	SyncErrorContext *string `pulumi:"syncErrorContext"`
	// Sync Error Direction.
	SyncErrorDirection *string `pulumi:"syncErrorDirection"`
	// Sync Error State
	SyncErrorState *string `pulumi:"syncErrorState"`
	// Sync Error State Timestamp
	SyncErrorStateTimestamp *string `pulumi:"syncErrorStateTimestamp"`
	// Name of Sync Group resource.
	SyncGroupName string `pulumi:"syncGroupName"`
	// Total progress
	TotalProgress *int `pulumi:"totalProgress"`
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent *int `pulumi:"volumeFreeSpacePercent"`
}

// The set of arguments for constructing a ServerEndpoint resource.
type ServerEndpointArgs struct {
	// Bytes in progress
	ByteProgress pulumi.IntPtrInput
	// Bytes total
	ByteTotal pulumi.IntPtrInput
	// Cloud Tiering.
	CloudTiering pulumi.StringPtrInput
	// current progress type.
	CurrentProgressType pulumi.StringPtrInput
	// Friendly Name
	FriendlyName pulumi.StringPtrInput
	// Item download error count.
	ItemDownloadErrorCount pulumi.IntPtrInput
	// Item Progress Count
	ItemProgressCount pulumi.IntPtrInput
	// Item Total Count
	ItemTotalCount pulumi.IntPtrInput
	// Item Upload Error Count.
	ItemUploadErrorCount pulumi.IntPtrInput
	// Last Sync Success
	LastSyncSuccess pulumi.StringPtrInput
	// ServerEndpoint lastWorkflowId
	LastWorkflowId pulumi.StringPtrInput
	// ServerEndpoint Provisioning State
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Server Endpoint object.
	ServerEndpointName pulumi.StringPtrInput
	// Server Local path.
	ServerLocalPath pulumi.StringPtrInput
	// Server Resource Id.
	ServerResourceId pulumi.StringPtrInput
	// Name of Storage Sync Service resource.
	StorageSyncServiceName pulumi.StringInput
	// sync error context.
	SyncErrorContext pulumi.StringPtrInput
	// Sync Error Direction.
	SyncErrorDirection pulumi.StringPtrInput
	// Sync Error State
	SyncErrorState pulumi.StringPtrInput
	// Sync Error State Timestamp
	SyncErrorStateTimestamp pulumi.StringPtrInput
	// Name of Sync Group resource.
	SyncGroupName pulumi.StringInput
	// Total progress
	TotalProgress pulumi.IntPtrInput
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent pulumi.IntPtrInput
}

func (ServerEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverEndpointArgs)(nil)).Elem()
}

type ServerEndpointInput interface {
	pulumi.Input

	ToServerEndpointOutput() ServerEndpointOutput
	ToServerEndpointOutputWithContext(ctx context.Context) ServerEndpointOutput
}

func (*ServerEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpoint)(nil))
}

func (i *ServerEndpoint) ToServerEndpointOutput() ServerEndpointOutput {
	return i.ToServerEndpointOutputWithContext(context.Background())
}

func (i *ServerEndpoint) ToServerEndpointOutputWithContext(ctx context.Context) ServerEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointOutput)
}

type ServerEndpointOutput struct {
	*pulumi.OutputState
}

func (ServerEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpoint)(nil))
}

func (o ServerEndpointOutput) ToServerEndpointOutput() ServerEndpointOutput {
	return o
}

func (o ServerEndpointOutput) ToServerEndpointOutputWithContext(ctx context.Context) ServerEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServerEndpointOutput{})
}
