// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package intune

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Android Policy entity for Intune MAM.
// API Version: 2015-01-14-preview.
type AndroidMAMPolicyByName struct {
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
	FileEncryption              pulumi.StringPtrOutput `pulumi:"fileEncryption"`
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
	ScreenCapture      pulumi.StringPtrOutput `pulumi:"screenCapture"`
	// Resource Tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAndroidMAMPolicyByName registers a new resource with the given unique name, arguments, and options.
func NewAndroidMAMPolicyByName(ctx *pulumi.Context,
	name string, args *AndroidMAMPolicyByNameArgs, opts ...pulumi.ResourceOption) (*AndroidMAMPolicyByName, error) {
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
	if args.FileEncryption == nil {
		args.FileEncryption = pulumi.StringPtr("required")
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
	if args.ScreenCapture == nil {
		args.ScreenCapture = pulumi.StringPtr("allow")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:intune:AndroidMAMPolicyByName"),
		},
		{
			Type: pulumi.String("azure-native:intune/v20150114preview:AndroidMAMPolicyByName"),
		},
		{
			Type: pulumi.String("azure-nextgen:intune/v20150114preview:AndroidMAMPolicyByName"),
		},
		{
			Type: pulumi.String("azure-native:intune/v20150114privatepreview:AndroidMAMPolicyByName"),
		},
		{
			Type: pulumi.String("azure-nextgen:intune/v20150114privatepreview:AndroidMAMPolicyByName"),
		},
	})
	opts = append(opts, aliases)
	var resource AndroidMAMPolicyByName
	err := ctx.RegisterResource("azure-native:intune:AndroidMAMPolicyByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAndroidMAMPolicyByName gets an existing AndroidMAMPolicyByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAndroidMAMPolicyByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AndroidMAMPolicyByNameState, opts ...pulumi.ResourceOption) (*AndroidMAMPolicyByName, error) {
	var resource AndroidMAMPolicyByName
	err := ctx.ReadResource("azure-native:intune:AndroidMAMPolicyByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AndroidMAMPolicyByName resources.
type androidMAMPolicyByNameState struct {
	AccessRecheckOfflineTimeout *string `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  *string `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         *string `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           *string `pulumi:"appSharingToLevel"`
	Authentication              *string `pulumi:"authentication"`
	ClipboardSharingLevel       *string `pulumi:"clipboardSharingLevel"`
	DataBackup                  *string `pulumi:"dataBackup"`
	Description                 *string `pulumi:"description"`
	DeviceCompliance            *string `pulumi:"deviceCompliance"`
	FileEncryption              *string `pulumi:"fileEncryption"`
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
	ScreenCapture      *string `pulumi:"screenCapture"`
	// Resource Tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type AndroidMAMPolicyByNameState struct {
	AccessRecheckOfflineTimeout pulumi.StringPtrInput
	AccessRecheckOnlineTimeout  pulumi.StringPtrInput
	AppSharingFromLevel         pulumi.StringPtrInput
	AppSharingToLevel           pulumi.StringPtrInput
	Authentication              pulumi.StringPtrInput
	ClipboardSharingLevel       pulumi.StringPtrInput
	DataBackup                  pulumi.StringPtrInput
	Description                 pulumi.StringPtrInput
	DeviceCompliance            pulumi.StringPtrInput
	FileEncryption              pulumi.StringPtrInput
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
	ScreenCapture      pulumi.StringPtrInput
	// Resource Tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (AndroidMAMPolicyByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*androidMAMPolicyByNameState)(nil)).Elem()
}

type androidMAMPolicyByNameArgs struct {
	AccessRecheckOfflineTimeout *string `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  *string `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         *string `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           *string `pulumi:"appSharingToLevel"`
	Authentication              *string `pulumi:"authentication"`
	ClipboardSharingLevel       *string `pulumi:"clipboardSharingLevel"`
	DataBackup                  *string `pulumi:"dataBackup"`
	Description                 *string `pulumi:"description"`
	DeviceCompliance            *string `pulumi:"deviceCompliance"`
	FileEncryption              *string `pulumi:"fileEncryption"`
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
	PolicyName    *string `pulumi:"policyName"`
	ScreenCapture *string `pulumi:"screenCapture"`
	// Resource Tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AndroidMAMPolicyByName resource.
type AndroidMAMPolicyByNameArgs struct {
	AccessRecheckOfflineTimeout pulumi.StringPtrInput
	AccessRecheckOnlineTimeout  pulumi.StringPtrInput
	AppSharingFromLevel         pulumi.StringPtrInput
	AppSharingToLevel           pulumi.StringPtrInput
	Authentication              pulumi.StringPtrInput
	ClipboardSharingLevel       pulumi.StringPtrInput
	DataBackup                  pulumi.StringPtrInput
	Description                 pulumi.StringPtrInput
	DeviceCompliance            pulumi.StringPtrInput
	FileEncryption              pulumi.StringPtrInput
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
	PolicyName    pulumi.StringPtrInput
	ScreenCapture pulumi.StringPtrInput
	// Resource Tags
	Tags pulumi.StringMapInput
}

func (AndroidMAMPolicyByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*androidMAMPolicyByNameArgs)(nil)).Elem()
}

type AndroidMAMPolicyByNameInput interface {
	pulumi.Input

	ToAndroidMAMPolicyByNameOutput() AndroidMAMPolicyByNameOutput
	ToAndroidMAMPolicyByNameOutputWithContext(ctx context.Context) AndroidMAMPolicyByNameOutput
}

func (*AndroidMAMPolicyByName) ElementType() reflect.Type {
	return reflect.TypeOf((*AndroidMAMPolicyByName)(nil))
}

func (i *AndroidMAMPolicyByName) ToAndroidMAMPolicyByNameOutput() AndroidMAMPolicyByNameOutput {
	return i.ToAndroidMAMPolicyByNameOutputWithContext(context.Background())
}

func (i *AndroidMAMPolicyByName) ToAndroidMAMPolicyByNameOutputWithContext(ctx context.Context) AndroidMAMPolicyByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AndroidMAMPolicyByNameOutput)
}

type AndroidMAMPolicyByNameOutput struct {
	*pulumi.OutputState
}

func (AndroidMAMPolicyByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AndroidMAMPolicyByName)(nil))
}

func (o AndroidMAMPolicyByNameOutput) ToAndroidMAMPolicyByNameOutput() AndroidMAMPolicyByNameOutput {
	return o
}

func (o AndroidMAMPolicyByNameOutput) ToAndroidMAMPolicyByNameOutputWithContext(ctx context.Context) AndroidMAMPolicyByNameOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AndroidMAMPolicyByNameOutput{})
}
