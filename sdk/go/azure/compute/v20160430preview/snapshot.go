// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160430preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Snapshot resource.
type Snapshot struct {
	pulumi.CustomResourceState

	// the storage account type of the disk.
	AccountType pulumi.StringPtrOutput `pulumi:"accountType"`
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataResponseOutput `pulumi:"creationData"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrOutput `pulumi:"diskSizeGB"`
	// Encryption settings for disk or snapshot
	EncryptionSettings EncryptionSettingsResponsePtrOutput `pulumi:"encryptionSettings"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The Operating System type.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// A relative URI containing the VM id that has the disk attached.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The disk provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The time when the disk was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CreationData == nil {
		return nil, errors.New("invalid value for required argument 'CreationData'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/v20160430preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20170330:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20191101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200630:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20201201:Snapshot"),
		},
	})
	opts = append(opts, aliases)
	var resource Snapshot
	err := ctx.RegisterResource("azure-native:compute/v20160430preview:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure-native:compute/v20160430preview:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// the storage account type of the disk.
	AccountType *string `pulumi:"accountType"`
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData *CreationDataResponse `pulumi:"creationData"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// Encryption settings for disk or snapshot
	EncryptionSettings *EncryptionSettingsResponse `pulumi:"encryptionSettings"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// The Operating System type.
	OsType *string `pulumi:"osType"`
	// A relative URI containing the VM id that has the disk attached.
	OwnerId *string `pulumi:"ownerId"`
	// The disk provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The time when the disk was created.
	TimeCreated *string `pulumi:"timeCreated"`
	// Resource type
	Type *string `pulumi:"type"`
}

type SnapshotState struct {
	// the storage account type of the disk.
	AccountType pulumi.StringPtrInput
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataResponsePtrInput
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrInput
	// Encryption settings for disk or snapshot
	EncryptionSettings EncryptionSettingsResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The Operating System type.
	OsType pulumi.StringPtrInput
	// A relative URI containing the VM id that has the disk attached.
	OwnerId pulumi.StringPtrInput
	// The disk provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The time when the disk was created.
	TimeCreated pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// the storage account type of the disk.
	AccountType *string `pulumi:"accountType"`
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationData `pulumi:"creationData"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// Encryption settings for disk or snapshot
	EncryptionSettings *EncryptionSettings `pulumi:"encryptionSettings"`
	// Resource location
	Location *string `pulumi:"location"`
	// The Operating System type.
	OsType *string `pulumi:"osType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the snapshot within the given subscription and resource group.
	SnapshotName *string `pulumi:"snapshotName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// the storage account type of the disk.
	AccountType *StorageAccountTypes
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataInput
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrInput
	// Encryption settings for disk or snapshot
	EncryptionSettings EncryptionSettingsPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The Operating System type.
	OsType *OperatingSystemTypes
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the snapshot within the given subscription and resource group.
	SnapshotName pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshot)(nil))
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

type SnapshotOutput struct {
	*pulumi.OutputState
}

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshot)(nil))
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SnapshotOutput{})
}
