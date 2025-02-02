// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The peering service prefix class.
type Prefix struct {
	pulumi.CustomResourceState

	// The error message for validation state
	ErrorMessage pulumi.StringOutput `pulumi:"errorMessage"`
	// The list of events for peering service prefix
	Events PeeringServicePrefixEventResponseArrayOutput `pulumi:"events"`
	// The prefix learned type
	LearnedType pulumi.StringOutput `pulumi:"learnedType"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The peering service prefix key
	PeeringServicePrefixKey pulumi.StringPtrOutput `pulumi:"peeringServicePrefixKey"`
	// The prefix from which your traffic originates.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// The prefix validation state
	PrefixValidationState pulumi.StringOutput `pulumi:"prefixValidationState"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrefix registers a new resource with the given unique name, arguments, and options.
func NewPrefix(ctx *pulumi.Context,
	name string, args *PrefixArgs, opts ...pulumi.ResourceOption) (*Prefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeeringServiceName == nil {
		return nil, errors.New("invalid value for required argument 'PeeringServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:peering/v20200401:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering:Prefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:peering:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190801preview:Prefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:peering/v20190801preview:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190901preview:Prefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:peering/v20190901preview:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:Prefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:peering/v20200101preview:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20201001:Prefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:peering/v20201001:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:Prefix"),
		},
		{
			Type: pulumi.String("azure-nextgen:peering/v20210101:Prefix"),
		},
	})
	opts = append(opts, aliases)
	var resource Prefix
	err := ctx.RegisterResource("azure-native:peering/v20200401:Prefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrefix gets an existing Prefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrefixState, opts ...pulumi.ResourceOption) (*Prefix, error) {
	var resource Prefix
	err := ctx.ReadResource("azure-native:peering/v20200401:Prefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Prefix resources.
type prefixState struct {
	// The error message for validation state
	ErrorMessage *string `pulumi:"errorMessage"`
	// The list of events for peering service prefix
	Events []PeeringServicePrefixEventResponse `pulumi:"events"`
	// The prefix learned type
	LearnedType *string `pulumi:"learnedType"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The peering service prefix key
	PeeringServicePrefixKey *string `pulumi:"peeringServicePrefixKey"`
	// The prefix from which your traffic originates.
	Prefix *string `pulumi:"prefix"`
	// The prefix validation state
	PrefixValidationState *string `pulumi:"prefixValidationState"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type PrefixState struct {
	// The error message for validation state
	ErrorMessage pulumi.StringPtrInput
	// The list of events for peering service prefix
	Events PeeringServicePrefixEventResponseArrayInput
	// The prefix learned type
	LearnedType pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The peering service prefix key
	PeeringServicePrefixKey pulumi.StringPtrInput
	// The prefix from which your traffic originates.
	Prefix pulumi.StringPtrInput
	// The prefix validation state
	PrefixValidationState pulumi.StringPtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (PrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*prefixState)(nil)).Elem()
}

type prefixArgs struct {
	// The name of the peering service.
	PeeringServiceName string `pulumi:"peeringServiceName"`
	// The peering service prefix key
	PeeringServicePrefixKey *string `pulumi:"peeringServicePrefixKey"`
	// The prefix from which your traffic originates.
	Prefix *string `pulumi:"prefix"`
	// The name of the prefix.
	PrefixName *string `pulumi:"prefixName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Prefix resource.
type PrefixArgs struct {
	// The name of the peering service.
	PeeringServiceName pulumi.StringInput
	// The peering service prefix key
	PeeringServicePrefixKey pulumi.StringPtrInput
	// The prefix from which your traffic originates.
	Prefix pulumi.StringPtrInput
	// The name of the prefix.
	PrefixName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (PrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prefixArgs)(nil)).Elem()
}

type PrefixInput interface {
	pulumi.Input

	ToPrefixOutput() PrefixOutput
	ToPrefixOutputWithContext(ctx context.Context) PrefixOutput
}

func (*Prefix) ElementType() reflect.Type {
	return reflect.TypeOf((*Prefix)(nil))
}

func (i *Prefix) ToPrefixOutput() PrefixOutput {
	return i.ToPrefixOutputWithContext(context.Background())
}

func (i *Prefix) ToPrefixOutputWithContext(ctx context.Context) PrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrefixOutput)
}

type PrefixOutput struct {
	*pulumi.OutputState
}

func (PrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Prefix)(nil))
}

func (o PrefixOutput) ToPrefixOutput() PrefixOutput {
	return o
}

func (o PrefixOutput) ToPrefixOutputWithContext(ctx context.Context) PrefixOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrefixOutput{})
}
