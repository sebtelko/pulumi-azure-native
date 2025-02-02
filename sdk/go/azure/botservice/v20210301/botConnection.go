// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Bot channel resource definition
type BotConnection struct {
	pulumi.CustomResourceState

	// Entity Tag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of properties specific to bot channel resource
	Properties ConnectionSettingPropertiesResponseOutput `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBotConnection registers a new resource with the given unique name, arguments, and options.
func NewBotConnection(ctx *pulumi.Context,
	name string, args *BotConnectionArgs, opts ...pulumi.ResourceOption) (*BotConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:botservice/v20210301:BotConnection"),
		},
		{
			Type: pulumi.String("azure-native:botservice:BotConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:botservice:BotConnection"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20171201:BotConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:botservice/v20171201:BotConnection"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20180712:BotConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:botservice/v20180712:BotConnection"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20200602:BotConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:botservice/v20200602:BotConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource BotConnection
	err := ctx.RegisterResource("azure-native:botservice/v20210301:BotConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBotConnection gets an existing BotConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBotConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BotConnectionState, opts ...pulumi.ResourceOption) (*BotConnection, error) {
	var resource BotConnection
	err := ctx.ReadResource("azure-native:botservice/v20210301:BotConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BotConnection resources.
type botConnectionState struct {
	// Entity Tag
	Etag *string `pulumi:"etag"`
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name *string `pulumi:"name"`
	// The set of properties specific to bot channel resource
	Properties *ConnectionSettingPropertiesResponse `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type *string `pulumi:"type"`
}

type BotConnectionState struct {
	// Entity Tag
	Etag pulumi.StringPtrInput
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Specifies the name of the resource.
	Name pulumi.StringPtrInput
	// The set of properties specific to bot channel resource
	Properties ConnectionSettingPropertiesResponsePtrInput
	// Gets or sets the SKU of the resource.
	Sku SkuResponsePtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Specifies the type of the resource.
	Type pulumi.StringPtrInput
}

func (BotConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*botConnectionState)(nil)).Elem()
}

type botConnectionArgs struct {
	// The name of the Bot Service Connection Setting resource.
	ConnectionName *string `pulumi:"connectionName"`
	// Entity Tag
	Etag *string `pulumi:"etag"`
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// The set of properties specific to bot channel resource
	Properties *ConnectionSettingProperties `pulumi:"properties"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName string `pulumi:"resourceName"`
	// Gets or sets the SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a BotConnection resource.
type BotConnectionArgs struct {
	// The name of the Bot Service Connection Setting resource.
	ConnectionName pulumi.StringPtrInput
	// Entity Tag
	Etag pulumi.StringPtrInput
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// The set of properties specific to bot channel resource
	Properties ConnectionSettingPropertiesPtrInput
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName pulumi.StringInput
	// The name of the Bot resource.
	ResourceName pulumi.StringInput
	// Gets or sets the SKU of the resource.
	Sku SkuPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
}

func (BotConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*botConnectionArgs)(nil)).Elem()
}

type BotConnectionInput interface {
	pulumi.Input

	ToBotConnectionOutput() BotConnectionOutput
	ToBotConnectionOutputWithContext(ctx context.Context) BotConnectionOutput
}

func (*BotConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*BotConnection)(nil))
}

func (i *BotConnection) ToBotConnectionOutput() BotConnectionOutput {
	return i.ToBotConnectionOutputWithContext(context.Background())
}

func (i *BotConnection) ToBotConnectionOutputWithContext(ctx context.Context) BotConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotConnectionOutput)
}

type BotConnectionOutput struct {
	*pulumi.OutputState
}

func (BotConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BotConnection)(nil))
}

func (o BotConnectionOutput) ToBotConnectionOutput() BotConnectionOutput {
	return o
}

func (o BotConnectionOutput) ToBotConnectionOutputWithContext(ctx context.Context) BotConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BotConnectionOutput{})
}
