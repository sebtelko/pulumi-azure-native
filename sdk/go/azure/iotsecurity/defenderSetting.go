// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iotsecurity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IoT Defender settings
// API Version: 2021-02-01-preview.
type DefenderSetting struct {
	pulumi.CustomResourceState

	// Size of the device quota (as a opposed to a Pay as You Go billing model). Value is required to be in multiples of 1000.
	DeviceQuota pulumi.IntOutput `pulumi:"deviceQuota"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The kind of onboarding for the subscription
	OnboardingKind pulumi.StringOutput `pulumi:"onboardingKind"`
	// Sentinel Workspace Resource Ids
	SentinelWorkspaceResourceIds pulumi.StringArrayOutput `pulumi:"sentinelWorkspaceResourceIds"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDefenderSetting registers a new resource with the given unique name, arguments, and options.
func NewDefenderSetting(ctx *pulumi.Context,
	name string, args *DefenderSettingArgs, opts ...pulumi.ResourceOption) (*DefenderSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceQuota == nil {
		return nil, errors.New("invalid value for required argument 'DeviceQuota'")
	}
	if args.OnboardingKind == nil {
		return nil, errors.New("invalid value for required argument 'OnboardingKind'")
	}
	if args.SentinelWorkspaceResourceIds == nil {
		return nil, errors.New("invalid value for required argument 'SentinelWorkspaceResourceIds'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:iotsecurity:DefenderSetting"),
		},
		{
			Type: pulumi.String("azure-native:iotsecurity/v20210201preview:DefenderSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:iotsecurity/v20210201preview:DefenderSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource DefenderSetting
	err := ctx.RegisterResource("azure-native:iotsecurity:DefenderSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefenderSetting gets an existing DefenderSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefenderSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefenderSettingState, opts ...pulumi.ResourceOption) (*DefenderSetting, error) {
	var resource DefenderSetting
	err := ctx.ReadResource("azure-native:iotsecurity:DefenderSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefenderSetting resources.
type defenderSettingState struct {
	// Size of the device quota (as a opposed to a Pay as You Go billing model). Value is required to be in multiples of 1000.
	DeviceQuota *int `pulumi:"deviceQuota"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The kind of onboarding for the subscription
	OnboardingKind *string `pulumi:"onboardingKind"`
	// Sentinel Workspace Resource Ids
	SentinelWorkspaceResourceIds []string `pulumi:"sentinelWorkspaceResourceIds"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type DefenderSettingState struct {
	// Size of the device quota (as a opposed to a Pay as You Go billing model). Value is required to be in multiples of 1000.
	DeviceQuota pulumi.IntPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The kind of onboarding for the subscription
	OnboardingKind pulumi.StringPtrInput
	// Sentinel Workspace Resource Ids
	SentinelWorkspaceResourceIds pulumi.StringArrayInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (DefenderSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*defenderSettingState)(nil)).Elem()
}

type defenderSettingArgs struct {
	// Size of the device quota (as a opposed to a Pay as You Go billing model). Value is required to be in multiples of 1000.
	DeviceQuota int `pulumi:"deviceQuota"`
	// The kind of onboarding for the subscription
	OnboardingKind string `pulumi:"onboardingKind"`
	// Sentinel Workspace Resource Ids
	SentinelWorkspaceResourceIds []string `pulumi:"sentinelWorkspaceResourceIds"`
}

// The set of arguments for constructing a DefenderSetting resource.
type DefenderSettingArgs struct {
	// Size of the device quota (as a opposed to a Pay as You Go billing model). Value is required to be in multiples of 1000.
	DeviceQuota pulumi.IntInput
	// The kind of onboarding for the subscription
	OnboardingKind pulumi.StringInput
	// Sentinel Workspace Resource Ids
	SentinelWorkspaceResourceIds pulumi.StringArrayInput
}

func (DefenderSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defenderSettingArgs)(nil)).Elem()
}

type DefenderSettingInput interface {
	pulumi.Input

	ToDefenderSettingOutput() DefenderSettingOutput
	ToDefenderSettingOutputWithContext(ctx context.Context) DefenderSettingOutput
}

func (*DefenderSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderSetting)(nil))
}

func (i *DefenderSetting) ToDefenderSettingOutput() DefenderSettingOutput {
	return i.ToDefenderSettingOutputWithContext(context.Background())
}

func (i *DefenderSetting) ToDefenderSettingOutputWithContext(ctx context.Context) DefenderSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderSettingOutput)
}

type DefenderSettingOutput struct {
	*pulumi.OutputState
}

func (DefenderSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderSetting)(nil))
}

func (o DefenderSettingOutput) ToDefenderSettingOutput() DefenderSettingOutput {
	return o
}

func (o DefenderSettingOutput) ToDefenderSettingOutputWithContext(ctx context.Context) DefenderSettingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DefenderSettingOutput{})
}
