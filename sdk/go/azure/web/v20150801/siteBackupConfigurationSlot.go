// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of a backup which will be performed
type SiteBackupConfigurationSlot struct {
	pulumi.CustomResourceState

	// Schedule for the backup if it is executed periodically
	BackupSchedule BackupScheduleResponsePtrOutput `pulumi:"backupSchedule"`
	// Databases included in the backup
	Databases DatabaseBackupSettingResponseArrayOutput `pulumi:"databases"`
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Kind of resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// SAS URL to the container
	StorageAccountUrl pulumi.StringPtrOutput `pulumi:"storageAccountUrl"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSiteBackupConfigurationSlot registers a new resource with the given unique name, arguments, and options.
func NewSiteBackupConfigurationSlot(ctx *pulumi.Context,
	name string, args *SiteBackupConfigurationSlotArgs, opts ...pulumi.ResourceOption) (*SiteBackupConfigurationSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:SiteBackupConfigurationSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteBackupConfigurationSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteBackupConfigurationSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteBackupConfigurationSlot gets an existing SiteBackupConfigurationSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteBackupConfigurationSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteBackupConfigurationSlotState, opts ...pulumi.ResourceOption) (*SiteBackupConfigurationSlot, error) {
	var resource SiteBackupConfigurationSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteBackupConfigurationSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteBackupConfigurationSlot resources.
type siteBackupConfigurationSlotState struct {
	// Schedule for the backup if it is executed periodically
	BackupSchedule *BackupScheduleResponse `pulumi:"backupSchedule"`
	// Databases included in the backup
	Databases []DatabaseBackupSettingResponse `pulumi:"databases"`
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled
	Enabled *bool `pulumi:"enabled"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// SAS URL to the container
	StorageAccountUrl *string `pulumi:"storageAccountUrl"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type SiteBackupConfigurationSlotState struct {
	// Schedule for the backup if it is executed periodically
	BackupSchedule BackupScheduleResponsePtrInput
	// Databases included in the backup
	Databases DatabaseBackupSettingResponseArrayInput
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled
	Enabled pulumi.BoolPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringPtrInput
	// SAS URL to the container
	StorageAccountUrl pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SiteBackupConfigurationSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteBackupConfigurationSlotState)(nil)).Elem()
}

type siteBackupConfigurationSlotArgs struct {
	// Schedule for the backup if it is executed periodically
	BackupSchedule *BackupSchedule `pulumi:"backupSchedule"`
	// Databases included in the backup
	Databases []DatabaseBackupSetting `pulumi:"databases"`
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled
	Enabled *bool `pulumi:"enabled"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
	// SAS URL to the container
	StorageAccountUrl *string `pulumi:"storageAccountUrl"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a SiteBackupConfigurationSlot resource.
type SiteBackupConfigurationSlotArgs struct {
	// Schedule for the backup if it is executed periodically
	BackupSchedule BackupSchedulePtrInput
	// Databases included in the backup
	Databases DatabaseBackupSettingArrayInput
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled
	Enabled pulumi.BoolPtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringInput
	// Name of resource group
	ResourceGroupName pulumi.StringInput
	// Name of web app slot. If not specified then will default to production slot.
	Slot pulumi.StringInput
	// SAS URL to the container
	StorageAccountUrl pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringInput
}

func (SiteBackupConfigurationSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteBackupConfigurationSlotArgs)(nil)).Elem()
}

type SiteBackupConfigurationSlotInput interface {
	pulumi.Input

	ToSiteBackupConfigurationSlotOutput() SiteBackupConfigurationSlotOutput
	ToSiteBackupConfigurationSlotOutputWithContext(ctx context.Context) SiteBackupConfigurationSlotOutput
}

func (*SiteBackupConfigurationSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteBackupConfigurationSlot)(nil))
}

func (i *SiteBackupConfigurationSlot) ToSiteBackupConfigurationSlotOutput() SiteBackupConfigurationSlotOutput {
	return i.ToSiteBackupConfigurationSlotOutputWithContext(context.Background())
}

func (i *SiteBackupConfigurationSlot) ToSiteBackupConfigurationSlotOutputWithContext(ctx context.Context) SiteBackupConfigurationSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteBackupConfigurationSlotOutput)
}

type SiteBackupConfigurationSlotOutput struct {
	*pulumi.OutputState
}

func (SiteBackupConfigurationSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteBackupConfigurationSlot)(nil))
}

func (o SiteBackupConfigurationSlotOutput) ToSiteBackupConfigurationSlotOutput() SiteBackupConfigurationSlotOutput {
	return o
}

func (o SiteBackupConfigurationSlotOutput) ToSiteBackupConfigurationSlotOutputWithContext(ctx context.Context) SiteBackupConfigurationSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteBackupConfigurationSlotOutput{})
}
