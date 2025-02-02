// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of HybridConnection Resource.
type HybridConnection struct {
	pulumi.CustomResourceState

	// The time the HybridConnection was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The number of listeners for this HybridConnection. min : 1 and max:25 supported
	ListenerCount pulumi.IntOutput `pulumi:"listenerCount"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// true if client authorization is needed for this HybridConnection; otherwise, false.
	RequiresClientAuthorization pulumi.BoolPtrOutput `pulumi:"requiresClientAuthorization"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrOutput `pulumi:"userMetadata"`
}

// NewHybridConnection registers a new resource with the given unique name, arguments, and options.
func NewHybridConnection(ctx *pulumi.Context,
	name string, args *HybridConnectionArgs, opts ...pulumi.ResourceOption) (*HybridConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:relay/v20160701:HybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:relay:HybridConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:relay:HybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20170401:HybridConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:relay/v20170401:HybridConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource HybridConnection
	err := ctx.RegisterResource("azure-native:relay/v20160701:HybridConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHybridConnection gets an existing HybridConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHybridConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridConnectionState, opts ...pulumi.ResourceOption) (*HybridConnection, error) {
	var resource HybridConnection
	err := ctx.ReadResource("azure-native:relay/v20160701:HybridConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HybridConnection resources.
type hybridConnectionState struct {
	// The time the HybridConnection was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The number of listeners for this HybridConnection. min : 1 and max:25 supported
	ListenerCount *int `pulumi:"listenerCount"`
	// Resource name
	Name *string `pulumi:"name"`
	// true if client authorization is needed for this HybridConnection; otherwise, false.
	RequiresClientAuthorization *bool `pulumi:"requiresClientAuthorization"`
	// Resource type
	Type *string `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
	// usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
	UserMetadata *string `pulumi:"userMetadata"`
}

type HybridConnectionState struct {
	// The time the HybridConnection was created.
	CreatedAt pulumi.StringPtrInput
	// The number of listeners for this HybridConnection. min : 1 and max:25 supported
	ListenerCount pulumi.IntPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// true if client authorization is needed for this HybridConnection; otherwise, false.
	RequiresClientAuthorization pulumi.BoolPtrInput
	// Resource type
	Type pulumi.StringPtrInput
	// The time the namespace was updated.
	UpdatedAt pulumi.StringPtrInput
	// usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrInput
}

func (HybridConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionState)(nil)).Elem()
}

type hybridConnectionArgs struct {
	// The hybrid connection name.
	HybridConnectionName *string `pulumi:"hybridConnectionName"`
	// The Namespace Name
	NamespaceName string `pulumi:"namespaceName"`
	// true if client authorization is needed for this HybridConnection; otherwise, false.
	RequiresClientAuthorization *bool `pulumi:"requiresClientAuthorization"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
	UserMetadata *string `pulumi:"userMetadata"`
}

// The set of arguments for constructing a HybridConnection resource.
type HybridConnectionArgs struct {
	// The hybrid connection name.
	HybridConnectionName pulumi.StringPtrInput
	// The Namespace Name
	NamespaceName pulumi.StringInput
	// true if client authorization is needed for this HybridConnection; otherwise, false.
	RequiresClientAuthorization pulumi.BoolPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrInput
}

func (HybridConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionArgs)(nil)).Elem()
}

type HybridConnectionInput interface {
	pulumi.Input

	ToHybridConnectionOutput() HybridConnectionOutput
	ToHybridConnectionOutputWithContext(ctx context.Context) HybridConnectionOutput
}

func (*HybridConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnection)(nil))
}

func (i *HybridConnection) ToHybridConnectionOutput() HybridConnectionOutput {
	return i.ToHybridConnectionOutputWithContext(context.Background())
}

func (i *HybridConnection) ToHybridConnectionOutputWithContext(ctx context.Context) HybridConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridConnectionOutput)
}

type HybridConnectionOutput struct {
	*pulumi.OutputState
}

func (HybridConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnection)(nil))
}

func (o HybridConnectionOutput) ToHybridConnectionOutput() HybridConnectionOutput {
	return o
}

func (o HybridConnectionOutput) ToHybridConnectionOutputWithContext(ctx context.Context) HybridConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HybridConnectionOutput{})
}
