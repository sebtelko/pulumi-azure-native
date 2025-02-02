// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Role Addon
type Addon struct {
	pulumi.CustomResourceState

	// Addon type.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Addon type
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAddon registers a new resource with the given unique name, arguments, and options.
func NewAddon(ctx *pulumi.Context,
	name string, args *AddonArgs, opts ...pulumi.ResourceOption) (*Addon, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge:Addon"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Addon"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901preview:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Addon"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20201201:Addon"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Addon"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210201preview:Addon"),
		},
	})
	opts = append(opts, aliases)
	var resource Addon
	err := ctx.RegisterResource("azure-native:databoxedge/v20200901:Addon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddon gets an existing Addon resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddonState, opts ...pulumi.ResourceOption) (*Addon, error) {
	var resource Addon
	err := ctx.ReadResource("azure-native:databoxedge/v20200901:Addon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Addon resources.
type addonState struct {
	// Addon type.
	Kind *string `pulumi:"kind"`
	// The object name.
	Name *string `pulumi:"name"`
	// Addon type
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
}

type AddonState struct {
	// Addon type.
	Kind pulumi.StringPtrInput
	// The object name.
	Name pulumi.StringPtrInput
	// Addon type
	SystemData SystemDataResponsePtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
}

func (AddonState) ElementType() reflect.Type {
	return reflect.TypeOf((*addonState)(nil)).Elem()
}

type addonArgs struct {
	// The addon name.
	AddonName *string `pulumi:"addonName"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// Addon type.
	Kind string `pulumi:"kind"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The role name.
	RoleName string `pulumi:"roleName"`
}

// The set of arguments for constructing a Addon resource.
type AddonArgs struct {
	// The addon name.
	AddonName pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// Addon type.
	Kind pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The role name.
	RoleName pulumi.StringInput
}

func (AddonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addonArgs)(nil)).Elem()
}

type AddonInput interface {
	pulumi.Input

	ToAddonOutput() AddonOutput
	ToAddonOutputWithContext(ctx context.Context) AddonOutput
}

func (*Addon) ElementType() reflect.Type {
	return reflect.TypeOf((*Addon)(nil))
}

func (i *Addon) ToAddonOutput() AddonOutput {
	return i.ToAddonOutputWithContext(context.Background())
}

func (i *Addon) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonOutput)
}

type AddonOutput struct {
	*pulumi.OutputState
}

func (AddonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Addon)(nil))
}

func (o AddonOutput) ToAddonOutput() AddonOutput {
	return o
}

func (o AddonOutput) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AddonOutput{})
}
