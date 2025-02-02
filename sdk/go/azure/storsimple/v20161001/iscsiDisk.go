// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The iSCSI disk.
type IscsiDisk struct {
	pulumi.CustomResourceState

	// The access control records.
	AccessControlRecords pulumi.StringArrayOutput `pulumi:"accessControlRecords"`
	// The data policy.
	DataPolicy pulumi.StringOutput `pulumi:"dataPolicy"`
	// The description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The disk status.
	DiskStatus pulumi.StringOutput `pulumi:"diskStatus"`
	// The local used capacity in bytes.
	LocalUsedCapacityInBytes pulumi.Float64Output `pulumi:"localUsedCapacityInBytes"`
	// The monitoring.
	MonitoringStatus pulumi.StringOutput `pulumi:"monitoringStatus"`
	// The name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioned capacity in bytes.
	ProvisionedCapacityInBytes pulumi.Float64Output `pulumi:"provisionedCapacityInBytes"`
	// The type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The used capacity in bytes.
	UsedCapacityInBytes pulumi.Float64Output `pulumi:"usedCapacityInBytes"`
}

// NewIscsiDisk registers a new resource with the given unique name, arguments, and options.
func NewIscsiDisk(ctx *pulumi.Context,
	name string, args *IscsiDiskArgs, opts ...pulumi.ResourceOption) (*IscsiDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessControlRecords == nil {
		return nil, errors.New("invalid value for required argument 'AccessControlRecords'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.IscsiServerName == nil {
		return nil, errors.New("invalid value for required argument 'IscsiServerName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ProvisionedCapacityInBytes == nil {
		return nil, errors.New("invalid value for required argument 'ProvisionedCapacityInBytes'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20161001:IscsiDisk"),
		},
	})
	opts = append(opts, aliases)
	var resource IscsiDisk
	err := ctx.RegisterResource("azure-native:storsimple/v20161001:IscsiDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIscsiDisk gets an existing IscsiDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIscsiDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IscsiDiskState, opts ...pulumi.ResourceOption) (*IscsiDisk, error) {
	var resource IscsiDisk
	err := ctx.ReadResource("azure-native:storsimple/v20161001:IscsiDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IscsiDisk resources.
type iscsiDiskState struct {
	// The access control records.
	AccessControlRecords []string `pulumi:"accessControlRecords"`
	// The data policy.
	DataPolicy *string `pulumi:"dataPolicy"`
	// The description.
	Description *string `pulumi:"description"`
	// The disk status.
	DiskStatus *string `pulumi:"diskStatus"`
	// The local used capacity in bytes.
	LocalUsedCapacityInBytes *float64 `pulumi:"localUsedCapacityInBytes"`
	// The monitoring.
	MonitoringStatus *string `pulumi:"monitoringStatus"`
	// The name.
	Name *string `pulumi:"name"`
	// The provisioned capacity in bytes.
	ProvisionedCapacityInBytes *float64 `pulumi:"provisionedCapacityInBytes"`
	// The type.
	Type *string `pulumi:"type"`
	// The used capacity in bytes.
	UsedCapacityInBytes *float64 `pulumi:"usedCapacityInBytes"`
}

type IscsiDiskState struct {
	// The access control records.
	AccessControlRecords pulumi.StringArrayInput
	// The data policy.
	DataPolicy pulumi.StringPtrInput
	// The description.
	Description pulumi.StringPtrInput
	// The disk status.
	DiskStatus pulumi.StringPtrInput
	// The local used capacity in bytes.
	LocalUsedCapacityInBytes pulumi.Float64PtrInput
	// The monitoring.
	MonitoringStatus pulumi.StringPtrInput
	// The name.
	Name pulumi.StringPtrInput
	// The provisioned capacity in bytes.
	ProvisionedCapacityInBytes pulumi.Float64PtrInput
	// The type.
	Type pulumi.StringPtrInput
	// The used capacity in bytes.
	UsedCapacityInBytes pulumi.Float64PtrInput
}

func (IscsiDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiDiskState)(nil)).Elem()
}

type iscsiDiskArgs struct {
	// The access control records.
	AccessControlRecords []string `pulumi:"accessControlRecords"`
	// The data policy.
	DataPolicy string `pulumi:"dataPolicy"`
	// The description.
	Description *string `pulumi:"description"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The disk name.
	DiskName *string `pulumi:"diskName"`
	// The disk status.
	DiskStatus string `pulumi:"diskStatus"`
	// The iSCSI server name.
	IscsiServerName string `pulumi:"iscsiServerName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The monitoring.
	MonitoringStatus string `pulumi:"monitoringStatus"`
	// The provisioned capacity in bytes.
	ProvisionedCapacityInBytes float64 `pulumi:"provisionedCapacityInBytes"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IscsiDisk resource.
type IscsiDiskArgs struct {
	// The access control records.
	AccessControlRecords pulumi.StringArrayInput
	// The data policy.
	DataPolicy DataPolicy
	// The description.
	Description pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// The disk name.
	DiskName pulumi.StringPtrInput
	// The disk status.
	DiskStatus DiskStatus
	// The iSCSI server name.
	IscsiServerName pulumi.StringInput
	// The manager name
	ManagerName pulumi.StringInput
	// The monitoring.
	MonitoringStatus MonitoringStatus
	// The provisioned capacity in bytes.
	ProvisionedCapacityInBytes pulumi.Float64Input
	// The resource group name
	ResourceGroupName pulumi.StringInput
}

func (IscsiDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiDiskArgs)(nil)).Elem()
}

type IscsiDiskInput interface {
	pulumi.Input

	ToIscsiDiskOutput() IscsiDiskOutput
	ToIscsiDiskOutputWithContext(ctx context.Context) IscsiDiskOutput
}

func (*IscsiDisk) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiDisk)(nil))
}

func (i *IscsiDisk) ToIscsiDiskOutput() IscsiDiskOutput {
	return i.ToIscsiDiskOutputWithContext(context.Background())
}

func (i *IscsiDisk) ToIscsiDiskOutputWithContext(ctx context.Context) IscsiDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiDiskOutput)
}

type IscsiDiskOutput struct {
	*pulumi.OutputState
}

func (IscsiDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiDisk)(nil))
}

func (o IscsiDiskOutput) ToIscsiDiskOutput() IscsiDiskOutput {
	return o
}

func (o IscsiDiskOutput) ToIscsiDiskOutputWithContext(ctx context.Context) IscsiDiskOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IscsiDiskOutput{})
}
