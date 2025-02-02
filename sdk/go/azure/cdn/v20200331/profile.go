// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200331

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CDN profile is a logical grouping of endpoints that share the same settings, such as CDN provider and pricing tier.
type Profile struct {
	pulumi.CustomResourceState

	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning status of the profile.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource status of the profile.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200331:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20150601:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20150601:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20160402:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20160402:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20161002:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20161002:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20170402:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20170402:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20171012:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20171012:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190415:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20190415:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20190615:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615preview:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20190615preview:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20191231:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200415:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Profile"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200901:Profile"),
		},
	})
	opts = append(opts, aliases)
	var resource Profile
	err := ctx.RegisterResource("azure-native:cdn/v20200331:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azure-native:cdn/v20200331:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Provisioning status of the profile.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource status of the profile.
	ResourceState *string `pulumi:"resourceState"`
	// The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ProfileState struct {
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Provisioning status of the profile.
	ProvisioningState pulumi.StringPtrInput
	// Resource status of the profile.
	ResourceState pulumi.StringPtrInput
	// The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
	Sku SkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// Resource location.
	Location *string `pulumi:"location"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName *string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
	Sku Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// Resource location.
	Location pulumi.StringPtrInput
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
	Sku SkuInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((*Profile)(nil))
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

type ProfileOutput struct {
	*pulumi.OutputState
}

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Profile)(nil))
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProfileOutput{})
}
