// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Managed private endpoint resource type.
// API Version: 2018-06-01.
type ManagedPrivateEndpoint struct {
	pulumi.CustomResourceState

	// Etag identifies change in the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Managed private endpoint properties.
	Properties ManagedPrivateEndpointResponseOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedPrivateEndpoint registers a new resource with the given unique name, arguments, and options.
func NewManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, args *ManagedPrivateEndpointArgs, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.ManagedVirtualNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedVirtualNetworkName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datafactory:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:datafactory/v20180601:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:datafactory/v20180601:ManagedPrivateEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedPrivateEndpoint
	err := ctx.RegisterResource("azure-native:datafactory:ManagedPrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedPrivateEndpoint gets an existing ManagedPrivateEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPrivateEndpointState, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	var resource ManagedPrivateEndpoint
	err := ctx.ReadResource("azure-native:datafactory:ManagedPrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedPrivateEndpoint resources.
type managedPrivateEndpointState struct {
	// Etag identifies change in the resource.
	Etag *string `pulumi:"etag"`
	// The resource name.
	Name *string `pulumi:"name"`
	// Managed private endpoint properties.
	Properties *ManagedPrivateEndpointResponse `pulumi:"properties"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type ManagedPrivateEndpointState struct {
	// Etag identifies change in the resource.
	Etag pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// Managed private endpoint properties.
	Properties ManagedPrivateEndpointResponsePtrInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (ManagedPrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointState)(nil)).Elem()
}

type managedPrivateEndpointArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// Managed private endpoint name
	ManagedPrivateEndpointName *string `pulumi:"managedPrivateEndpointName"`
	// Managed virtual network name
	ManagedVirtualNetworkName string `pulumi:"managedVirtualNetworkName"`
	// Managed private endpoint properties.
	Properties ManagedPrivateEndpointType `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagedPrivateEndpoint resource.
type ManagedPrivateEndpointArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput
	// Managed private endpoint name
	ManagedPrivateEndpointName pulumi.StringPtrInput
	// Managed virtual network name
	ManagedVirtualNetworkName pulumi.StringInput
	// Managed private endpoint properties.
	Properties ManagedPrivateEndpointTypeInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
}

func (ManagedPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointArgs)(nil)).Elem()
}

type ManagedPrivateEndpointInput interface {
	pulumi.Input

	ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput
	ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput
}

func (*ManagedPrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpoint)(nil))
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return i.ToManagedPrivateEndpointOutputWithContext(context.Background())
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointOutput)
}

type ManagedPrivateEndpointOutput struct {
	*pulumi.OutputState
}

func (ManagedPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpoint)(nil))
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return o
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedPrivateEndpointOutput{})
}
