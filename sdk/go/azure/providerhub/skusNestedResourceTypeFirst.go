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
type SkusNestedResourceTypeFirst struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties SkuResourceResponsePropertiesOutput `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSkusNestedResourceTypeFirst registers a new resource with the given unique name, arguments, and options.
func NewSkusNestedResourceTypeFirst(ctx *pulumi.Context,
	name string, args *SkusNestedResourceTypeFirstArgs, opts ...pulumi.ResourceOption) (*SkusNestedResourceTypeFirst, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NestedResourceTypeFirst == nil {
		return nil, errors.New("invalid value for required argument 'NestedResourceTypeFirst'")
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
			Type: pulumi.String("azure-nextgen:providerhub:SkusNestedResourceTypeFirst"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:SkusNestedResourceTypeFirst"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20201120:SkusNestedResourceTypeFirst"),
		},
	})
	opts = append(opts, aliases)
	var resource SkusNestedResourceTypeFirst
	err := ctx.RegisterResource("azure-native:providerhub:SkusNestedResourceTypeFirst", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSkusNestedResourceTypeFirst gets an existing SkusNestedResourceTypeFirst resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSkusNestedResourceTypeFirst(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SkusNestedResourceTypeFirstState, opts ...pulumi.ResourceOption) (*SkusNestedResourceTypeFirst, error) {
	var resource SkusNestedResourceTypeFirst
	err := ctx.ReadResource("azure-native:providerhub:SkusNestedResourceTypeFirst", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SkusNestedResourceTypeFirst resources.
type skusNestedResourceTypeFirstState struct {
	// The name of the resource
	Name       *string                        `pulumi:"name"`
	Properties *SkuResourceResponseProperties `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type SkusNestedResourceTypeFirstState struct {
	// The name of the resource
	Name       pulumi.StringPtrInput
	Properties SkuResourceResponsePropertiesPtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (SkusNestedResourceTypeFirstState) ElementType() reflect.Type {
	return reflect.TypeOf((*skusNestedResourceTypeFirstState)(nil)).Elem()
}

type skusNestedResourceTypeFirstArgs struct {
	// The first child resource type.
	NestedResourceTypeFirst string `pulumi:"nestedResourceTypeFirst"`
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace string `pulumi:"providerNamespace"`
	// The resource type.
	ResourceType string `pulumi:"resourceType"`
	// The SKU.
	Sku         *string      `pulumi:"sku"`
	SkuSettings []SkuSetting `pulumi:"skuSettings"`
}

// The set of arguments for constructing a SkusNestedResourceTypeFirst resource.
type SkusNestedResourceTypeFirstArgs struct {
	// The first child resource type.
	NestedResourceTypeFirst pulumi.StringInput
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringInput
	// The resource type.
	ResourceType pulumi.StringInput
	// The SKU.
	Sku         pulumi.StringPtrInput
	SkuSettings SkuSettingArrayInput
}

func (SkusNestedResourceTypeFirstArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*skusNestedResourceTypeFirstArgs)(nil)).Elem()
}

type SkusNestedResourceTypeFirstInput interface {
	pulumi.Input

	ToSkusNestedResourceTypeFirstOutput() SkusNestedResourceTypeFirstOutput
	ToSkusNestedResourceTypeFirstOutputWithContext(ctx context.Context) SkusNestedResourceTypeFirstOutput
}

func (*SkusNestedResourceTypeFirst) ElementType() reflect.Type {
	return reflect.TypeOf((*SkusNestedResourceTypeFirst)(nil))
}

func (i *SkusNestedResourceTypeFirst) ToSkusNestedResourceTypeFirstOutput() SkusNestedResourceTypeFirstOutput {
	return i.ToSkusNestedResourceTypeFirstOutputWithContext(context.Background())
}

func (i *SkusNestedResourceTypeFirst) ToSkusNestedResourceTypeFirstOutputWithContext(ctx context.Context) SkusNestedResourceTypeFirstOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkusNestedResourceTypeFirstOutput)
}

type SkusNestedResourceTypeFirstOutput struct {
	*pulumi.OutputState
}

func (SkusNestedResourceTypeFirstOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkusNestedResourceTypeFirst)(nil))
}

func (o SkusNestedResourceTypeFirstOutput) ToSkusNestedResourceTypeFirstOutput() SkusNestedResourceTypeFirstOutput {
	return o
}

func (o SkusNestedResourceTypeFirstOutput) ToSkusNestedResourceTypeFirstOutputWithContext(ctx context.Context) SkusNestedResourceTypeFirstOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SkusNestedResourceTypeFirstOutput{})
}
