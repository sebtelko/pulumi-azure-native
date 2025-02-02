// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package providerhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API Version: 2020-11-20.
type Skus struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties SkuResourceResponsePropertiesOutput `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSkus registers a new resource with the given unique name, arguments, and options.
func NewSkus(ctx *pulumi.Context,
	name string, args *SkusArgs, opts ...pulumi.ResourceOption) (*Skus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.SkuSettings == nil {
		return nil, errors.New("invalid value for required argument 'SkuSettings'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:providerhub:Skus"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:Skus"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20201120:Skus"),
		},
	})
	opts = append(opts, aliases)
	var resource Skus
	err := ctx.RegisterResource("azure-native:providerhub:Skus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSkus gets an existing Skus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSkus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SkusState, opts ...pulumi.ResourceOption) (*Skus, error) {
	var resource Skus
	err := ctx.ReadResource("azure-native:providerhub:Skus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Skus resources.
type skusState struct {
	// The name of the resource
	Name       *string                        `pulumi:"name"`
	Properties *SkuResourceResponseProperties `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type SkusState struct {
	// The name of the resource
	Name       pulumi.StringPtrInput
	Properties SkuResourceResponsePropertiesPtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (SkusState) ElementType() reflect.Type {
	return reflect.TypeOf((*skusState)(nil)).Elem()
}

type skusArgs struct {
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace string `pulumi:"providerNamespace"`
	// The resource type.
	ResourceType string `pulumi:"resourceType"`
	// The SKU.
	Sku         *string      `pulumi:"sku"`
	SkuSettings []SkuSetting `pulumi:"skuSettings"`
}

// The set of arguments for constructing a Skus resource.
type SkusArgs struct {
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringInput
	// The resource type.
	ResourceType pulumi.StringInput
	// The SKU.
	Sku         pulumi.StringPtrInput
	SkuSettings SkuSettingArrayInput
}

func (SkusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*skusArgs)(nil)).Elem()
}

type SkusInput interface {
	pulumi.Input

	ToSkusOutput() SkusOutput
	ToSkusOutputWithContext(ctx context.Context) SkusOutput
}

func (*Skus) ElementType() reflect.Type {
	return reflect.TypeOf((*Skus)(nil))
}

func (i *Skus) ToSkusOutput() SkusOutput {
	return i.ToSkusOutputWithContext(context.Background())
}

func (i *Skus) ToSkusOutputWithContext(ctx context.Context) SkusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkusOutput)
}

type SkusOutput struct {
	*pulumi.OutputState
}

func (SkusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Skus)(nil))
}

func (o SkusOutput) ToSkusOutput() SkusOutput {
	return o
}

func (o SkusOutput) ToSkusOutputWithContext(ctx context.Context) SkusOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SkusOutput{})
}
