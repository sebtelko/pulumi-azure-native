// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150114preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// iOS Policy entity for Intune MAM.
type IoMAMPolicyByName struct {
	pulumi.CustomResourceState

	AccessRecheckOfflineTimeout pulumi.StringPtrOutput `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  pulumi.StringPtrOutput `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         pulumi.StringPtrOutput `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           pulumi.StringPtrOutput `pulumi:"appSharingToLevel"`
	Authentication              pulumi.StringPtrOutput `pulumi:"authentication"`
	ClipboardSharingLevel       pulumi.StringPtrOutput `pulumi:"clipboardSharingLevel"`
	DataBackup                  pulumi.StringPtrOutput `pulumi:"dataBackup"`
	Description                 pulumi.StringPtrOutput `pulumi:"description"`
	DeviceCompliance            pulumi.StringPtrOutput `pulumi:"deviceCompliance"`
	FileEncryptionLevel         pulumi.StringPtrOutput `pulumi:"fileEncryptionLevel"`
	FileSharingSaveAs           pulumi.StringPtrOutput `pulumi:"fileSharingSaveAs"`
	FriendlyName                pulumi.StringOutput    `pulumi:"friendlyName"`
	GroupStatus                 pulumi.StringOutput    `pulumi:"groupStatus"`
	LastModifiedTime            pulumi.StringOutput    `pulumi:"lastModifiedTime"`
	// Resource Location
	Location       pulumi.StringPtrOutput `pulumi:"location"`
	ManagedBrowser pulumi.StringPtrOutput `pulumi:"managedBrowser"`
	// Resource name
	Name               pulumi.StringOutput    `pulumi:"name"`
	NumOfApps          pulumi.IntOutput       `pulumi:"numOfApps"`
	OfflineWipeTimeout pulumi.StringPtrOutput `pulumi:"offlineWipeTimeout"`
	Pin                pulumi.StringPtrOutput `pulumi:"pin"`
	PinNumRetry        pulumi.IntPtrOutput    `pulumi:"pinNumRetry"`
	// Resource Tags
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TouchId pulumi.StringPtrOutput `pulumi:"touchId"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIoMAMPolicyByName registers a new resource with the given unique name, arguments, and options.
func NewIoMAMPolicyByName(ctx *pulumi.Context,
	name string, args *IoMAMPolicyByNameArgs, opts ...pulumi.ResourceOption) (*IoMAMPolicyByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FriendlyName == nil {
		return nil, errors.New("invalid value for required argument 'FriendlyName'")
	}
	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.AppSharingFromLevel == nil {
		args.AppSharingFromLevel = pulumi.StringPtr("none")
	}
	if args.AppSharingToLevel == nil {
		args.AppSharingToLevel = pulumi.StringPtr("none")
	}
	if args.Authentication == nil {
		args.Authentication = pulumi.StringPtr("required")
	}
	if args.ClipboardSharingLevel == nil {
		args.ClipboardSharingLevel = pulumi.StringPtr("blocked")
	}
	if args.DataBackup == nil {
		args.DataBackup = pulumi.StringPtr("allow")
	}
	if args.DeviceCompliance == nil {
		args.DeviceCompliance = pulumi.StringPtr("enable")
	}
	if args.FileEncryptionLevel == nil {
		args.FileEncryptionLevel = pulumi.StringPtr("deviceLocked")
	}
	if args.FileSharingSaveAs == nil {
		args.FileSharingSaveAs = pulumi.StringPtr("allow")
	}
	if args.ManagedBrowser == nil {
		args.ManagedBrowser = pulumi.StringPtr("required")
	}
	if args.Pin == nil {
		args.Pin = pulumi.StringPtr("required")
	}
	if args.TouchId == nil {
		args.TouchId = pulumi.StringPtr("enable")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:intune/v20150114preview:IoMAMPolicyByName"),
		},
		{
			Type: pulumi.String("azure-native:intune:IoMAMPolicyByName"),
		},
		{
			Type: pulumi.String("azure-nextgen:intune:IoMAMPolicyByName"),
		},
		{
			Type: pulumi.String("azure-native:intune/v20150114privatepreview:IoMAMPolicyByName"),
		},
		{
			Type: pulumi.String("azure-nextgen:intune/v20150114privatepreview:IoMAMPolicyByName"),
		},
	})
	opts = append(opts, aliases)
	var resource IoMAMPolicyByName
	err := ctx.RegisterResource("azure-native:intune/v20150114preview:IoMAMPolicyByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIoMAMPolicyByName gets an existing IoMAMPolicyByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIoMAMPolicyByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IoMAMPolicyByNameState, opts ...pulumi.ResourceOption) (*IoMAMPolicyByName, error) {
	var resource IoMAMPolicyByName
	err := ctx.ReadResource("azure-native:intune/v20150114preview:IoMAMPolicyByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IoMAMPolicyByName resources.
type ioMAMPolicyByNameState struct {
	AccessRecheckOfflineTimeout *string `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  *string `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         *string `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           *string `pulumi:"appSharingToLevel"`
	Authentication              *string `pulumi:"authentication"`
	ClipboardSharingLevel       *string `pulumi:"clipboardSharingLevel"`
	DataBackup                  *string `pulumi:"dataBackup"`
	Description                 *string `pulumi:"description"`
	DeviceCompliance            *string `pulumi:"deviceCompliance"`
	FileEncryptionLevel         *string `pulumi:"fileEncryptionLevel"`
	FileSharingSaveAs           *string `pulumi:"fileSharingSaveAs"`
	FriendlyName                *string `pulumi:"friendlyName"`
	GroupStatus                 *string `pulumi:"groupStatus"`
	LastModifiedTime            *string `pulumi:"lastModifiedTime"`
	// Resource Location
	Location       *string `pulumi:"location"`
	ManagedBrowser *string `pulumi:"managedBrowser"`
	// Resource name
	Name               *string `pulumi:"name"`
	NumOfApps          *int    `pulumi:"numOfApps"`
	OfflineWipeTimeout *string `pulumi:"offlineWipeTimeout"`
	Pin                *string `pulumi:"pin"`
	PinNumRetry        *int    `pulumi:"pinNumRetry"`
	// Resource Tags
	Tags    map[string]string `pulumi:"tags"`
	TouchId *string           `pulumi:"touchId"`
	// Resource type
	Type *string `pulumi:"type"`
}

type IoMAMPolicyByNameState struct {
	AccessRecheckOfflineTimeout pulumi.StringPtrInput
	AccessRecheckOnlineTimeout  pulumi.StringPtrInput
	AppSharingFromLevel         pulumi.StringPtrInput
	AppSharingToLevel           pulumi.StringPtrInput
	Authentication              pulumi.StringPtrInput
	ClipboardSharingLevel       pulumi.StringPtrInput
	DataBackup                  pulumi.StringPtrInput
	Description                 pulumi.StringPtrInput
	DeviceCompliance            pulumi.StringPtrInput
	FileEncryptionLevel         pulumi.StringPtrInput
	FileSharingSaveAs           pulumi.StringPtrInput
	FriendlyName                pulumi.StringPtrInput
	GroupStatus                 pulumi.StringPtrInput
	LastModifiedTime            pulumi.StringPtrInput
	// Resource Location
	Location       pulumi.StringPtrInput
	ManagedBrowser pulumi.StringPtrInput
	// Resource name
	Name               pulumi.StringPtrInput
	NumOfApps          pulumi.IntPtrInput
	OfflineWipeTimeout pulumi.StringPtrInput
	Pin                pulumi.StringPtrInput
	PinNumRetry        pulumi.IntPtrInput
	// Resource Tags
	Tags    pulumi.StringMapInput
	TouchId pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (IoMAMPolicyByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*ioMAMPolicyByNameState)(nil)).Elem()
}

type ioMAMPolicyByNameArgs struct {
	AccessRecheckOfflineTimeout *string `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  *string `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         *string `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           *string `pulumi:"appSharingToLevel"`
	Authentication              *string `pulumi:"authentication"`
	ClipboardSharingLevel       *string `pulumi:"clipboardSharingLevel"`
	DataBackup                  *string `pulumi:"dataBackup"`
	Description                 *string `pulumi:"description"`
	DeviceCompliance            *string `pulumi:"deviceCompliance"`
	FileEncryptionLevel         *string `pulumi:"fileEncryptionLevel"`
	FileSharingSaveAs           *string `pulumi:"fileSharingSaveAs"`
	FriendlyName                string  `pulumi:"friendlyName"`
	// Location hostName for the tenant
	HostName string `pulumi:"hostName"`
	// Resource Location
	Location           *string `pulumi:"location"`
	ManagedBrowser     *string `pulumi:"managedBrowser"`
	OfflineWipeTimeout *string `pulumi:"offlineWipeTimeout"`
	Pin                *string `pulumi:"pin"`
	PinNumRetry        *int    `pulumi:"pinNumRetry"`
	// Unique name for the policy
	PolicyName *string `pulumi:"policyName"`
	// Resource Tags
	Tags    map[string]string `pulumi:"tags"`
	TouchId *string           `pulumi:"touchId"`
}

// The set of arguments for constructing a IoMAMPolicyByName resource.
type IoMAMPolicyByNameArgs struct {
	AccessRecheckOfflineTimeout pulumi.StringPtrInput
	AccessRecheckOnlineTimeout  pulumi.StringPtrInput
	AppSharingFromLevel         pulumi.StringPtrInput
	AppSharingToLevel           pulumi.StringPtrInput
	Authentication              pulumi.StringPtrInput
	ClipboardSharingLevel       pulumi.StringPtrInput
	DataBackup                  pulumi.StringPtrInput
	Description                 pulumi.StringPtrInput
	DeviceCompliance            pulumi.StringPtrInput
	FileEncryptionLevel         pulumi.StringPtrInput
	FileSharingSaveAs           pulumi.StringPtrInput
	FriendlyName                pulumi.StringInput
	// Location hostName for the tenant
	HostName pulumi.StringInput
	// Resource Location
	Location           pulumi.StringPtrInput
	ManagedBrowser     pulumi.StringPtrInput
	OfflineWipeTimeout pulumi.StringPtrInput
	Pin                pulumi.StringPtrInput
	PinNumRetry        pulumi.IntPtrInput
	// Unique name for the policy
	PolicyName pulumi.StringPtrInput
	// Resource Tags
	Tags    pulumi.StringMapInput
	TouchId pulumi.StringPtrInput
}

func (IoMAMPolicyByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ioMAMPolicyByNameArgs)(nil)).Elem()
}

type IoMAMPolicyByNameInput interface {
	pulumi.Input

	ToIoMAMPolicyByNameOutput() IoMAMPolicyByNameOutput
	ToIoMAMPolicyByNameOutputWithContext(ctx context.Context) IoMAMPolicyByNameOutput
}

func (*IoMAMPolicyByName) ElementType() reflect.Type {
	return reflect.TypeOf((*IoMAMPolicyByName)(nil))
}

func (i *IoMAMPolicyByName) ToIoMAMPolicyByNameOutput() IoMAMPolicyByNameOutput {
	return i.ToIoMAMPolicyByNameOutputWithContext(context.Background())
}

func (i *IoMAMPolicyByName) ToIoMAMPolicyByNameOutputWithContext(ctx context.Context) IoMAMPolicyByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoMAMPolicyByNameOutput)
}

type IoMAMPolicyByNameOutput struct {
	*pulumi.OutputState
}

func (IoMAMPolicyByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IoMAMPolicyByName)(nil))
}

func (o IoMAMPolicyByNameOutput) ToIoMAMPolicyByNameOutput() IoMAMPolicyByNameOutput {
	return o
}

func (o IoMAMPolicyByNameOutput) ToIoMAMPolicyByNameOutputWithContext(ctx context.Context) IoMAMPolicyByNameOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IoMAMPolicyByNameOutput{})
}
