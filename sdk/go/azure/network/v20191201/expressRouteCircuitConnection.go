// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.
type ExpressRouteCircuitConnection struct {
	pulumi.CustomResourceState

	// /29 IP address space to carve out Customer addresses for tunnels.
	AddressPrefix pulumi.StringPtrOutput `pulumi:"addressPrefix"`
	// The authorization key.
	AuthorizationKey pulumi.StringPtrOutput `pulumi:"authorizationKey"`
	// Express Route Circuit connection state.
	CircuitConnectionStatus pulumi.StringOutput `pulumi:"circuitConnectionStatus"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
	ExpressRouteCircuitPeering SubResourceResponsePtrOutput `pulumi:"expressRouteCircuitPeering"`
	// IPv6 Address PrefixProperties of the express route circuit connection.
	Ipv6CircuitConnectionConfig Ipv6CircuitConnectionConfigResponsePtrOutput `pulumi:"ipv6CircuitConnectionConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Reference to Express Route Circuit Private Peering Resource of the peered circuit.
	PeerExpressRouteCircuitPeering SubResourceResponsePtrOutput `pulumi:"peerExpressRouteCircuitPeering"`
	// The provisioning state of the express route circuit connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExpressRouteCircuitConnection registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitConnection(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitConnectionArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CircuitName == nil {
		return nil, errors.New("invalid value for required argument 'CircuitName'")
	}
	if args.PeeringName == nil {
		return nil, errors.New("invalid value for required argument 'PeeringName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:ExpressRouteCircuitConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteCircuitConnection
	err := ctx.RegisterResource("azure-native:network/v20191201:ExpressRouteCircuitConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuitConnection gets an existing ExpressRouteCircuitConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuitConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitConnectionState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitConnection, error) {
	var resource ExpressRouteCircuitConnection
	err := ctx.ReadResource("azure-native:network/v20191201:ExpressRouteCircuitConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuitConnection resources.
type expressRouteCircuitConnectionState struct {
	// /29 IP address space to carve out Customer addresses for tunnels.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// The authorization key.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// Express Route Circuit connection state.
	CircuitConnectionStatus *string `pulumi:"circuitConnectionStatus"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
	ExpressRouteCircuitPeering *SubResourceResponse `pulumi:"expressRouteCircuitPeering"`
	// IPv6 Address PrefixProperties of the express route circuit connection.
	Ipv6CircuitConnectionConfig *Ipv6CircuitConnectionConfigResponse `pulumi:"ipv6CircuitConnectionConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Reference to Express Route Circuit Private Peering Resource of the peered circuit.
	PeerExpressRouteCircuitPeering *SubResourceResponse `pulumi:"peerExpressRouteCircuitPeering"`
	// The provisioning state of the express route circuit connection resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Type of the resource.
	Type *string `pulumi:"type"`
}

type ExpressRouteCircuitConnectionState struct {
	// /29 IP address space to carve out Customer addresses for tunnels.
	AddressPrefix pulumi.StringPtrInput
	// The authorization key.
	AuthorizationKey pulumi.StringPtrInput
	// Express Route Circuit connection state.
	CircuitConnectionStatus pulumi.StringPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
	ExpressRouteCircuitPeering SubResourceResponsePtrInput
	// IPv6 Address PrefixProperties of the express route circuit connection.
	Ipv6CircuitConnectionConfig Ipv6CircuitConnectionConfigResponsePtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Reference to Express Route Circuit Private Peering Resource of the peered circuit.
	PeerExpressRouteCircuitPeering SubResourceResponsePtrInput
	// The provisioning state of the express route circuit connection resource.
	ProvisioningState pulumi.StringPtrInput
	// Type of the resource.
	Type pulumi.StringPtrInput
}

func (ExpressRouteCircuitConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitConnectionState)(nil)).Elem()
}

type expressRouteCircuitConnectionArgs struct {
	// /29 IP address space to carve out Customer addresses for tunnels.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// The authorization key.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// The name of the express route circuit connection.
	ConnectionName *string `pulumi:"connectionName"`
	// Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
	ExpressRouteCircuitPeering *SubResource `pulumi:"expressRouteCircuitPeering"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// IPv6 Address PrefixProperties of the express route circuit connection.
	Ipv6CircuitConnectionConfig *Ipv6CircuitConnectionConfig `pulumi:"ipv6CircuitConnectionConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Reference to Express Route Circuit Private Peering Resource of the peered circuit.
	PeerExpressRouteCircuitPeering *SubResource `pulumi:"peerExpressRouteCircuitPeering"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ExpressRouteCircuitConnection resource.
type ExpressRouteCircuitConnectionArgs struct {
	// /29 IP address space to carve out Customer addresses for tunnels.
	AddressPrefix pulumi.StringPtrInput
	// The authorization key.
	AuthorizationKey pulumi.StringPtrInput
	// The name of the express route circuit.
	CircuitName pulumi.StringInput
	// The name of the express route circuit connection.
	ConnectionName pulumi.StringPtrInput
	// Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
	ExpressRouteCircuitPeering SubResourcePtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// IPv6 Address PrefixProperties of the express route circuit connection.
	Ipv6CircuitConnectionConfig Ipv6CircuitConnectionConfigPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Reference to Express Route Circuit Private Peering Resource of the peered circuit.
	PeerExpressRouteCircuitPeering SubResourcePtrInput
	// The name of the peering.
	PeeringName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ExpressRouteCircuitConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitConnectionArgs)(nil)).Elem()
}

type ExpressRouteCircuitConnectionInput interface {
	pulumi.Input

	ToExpressRouteCircuitConnectionOutput() ExpressRouteCircuitConnectionOutput
	ToExpressRouteCircuitConnectionOutputWithContext(ctx context.Context) ExpressRouteCircuitConnectionOutput
}

func (*ExpressRouteCircuitConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitConnection)(nil))
}

func (i *ExpressRouteCircuitConnection) ToExpressRouteCircuitConnectionOutput() ExpressRouteCircuitConnectionOutput {
	return i.ToExpressRouteCircuitConnectionOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuitConnection) ToExpressRouteCircuitConnectionOutputWithContext(ctx context.Context) ExpressRouteCircuitConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitConnectionOutput)
}

type ExpressRouteCircuitConnectionOutput struct {
	*pulumi.OutputState
}

func (ExpressRouteCircuitConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitConnection)(nil))
}

func (o ExpressRouteCircuitConnectionOutput) ToExpressRouteCircuitConnectionOutput() ExpressRouteCircuitConnectionOutput {
	return o
}

func (o ExpressRouteCircuitConnectionOutput) ToExpressRouteCircuitConnectionOutputWithContext(ctx context.Context) ExpressRouteCircuitConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteCircuitConnectionOutput{})
}
