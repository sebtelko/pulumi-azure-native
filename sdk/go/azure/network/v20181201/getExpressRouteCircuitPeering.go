// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Peering in an ExpressRouteCircuit resource.
func LookupExpressRouteCircuitPeering(ctx *pulumi.Context, args *LookupExpressRouteCircuitPeeringArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitPeeringResult, error) {
	var rv LookupExpressRouteCircuitPeeringResult
	err := ctx.Invoke("azure-native:network/v20181201:getExpressRouteCircuitPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitPeeringArgs struct {
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Peering in an ExpressRouteCircuit resource.
type LookupExpressRouteCircuitPeeringResult struct {
	// The Azure ASN.
	AzureASN *int `pulumi:"azureASN"`
	// The list of circuit connections associated with Azure Private Peering for this circuit.
	Connections []ExpressRouteCircuitConnectionResponse `pulumi:"connections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The ExpressRoute connection.
	ExpressRouteConnection *ExpressRouteConnectionIdResponse `pulumi:"expressRouteConnection"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig *Ipv6ExpressRouteCircuitPeeringConfigResponse `pulumi:"ipv6PeeringConfig"`
	// Gets whether the provider or the customer last modified the peering.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringConfigResponse `pulumi:"microsoftPeeringConfig"`
	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The peer ASN.
	PeerASN *float64 `pulumi:"peerASN"`
	// The list of peered circuit connections associated with Azure Private Peering for this circuit.
	PeeredConnections []PeerExpressRouteCircuitConnectionResponse `pulumi:"peeredConnections"`
	// The peering type.
	PeeringType *string `pulumi:"peeringType"`
	// The primary port.
	PrimaryAzurePort *string `pulumi:"primaryAzurePort"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix *string `pulumi:"primaryPeerAddressPrefix"`
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The reference of the RouteFilter resource.
	RouteFilter *RouteFilterResponse `pulumi:"routeFilter"`
	// The secondary port.
	SecondaryAzurePort *string `pulumi:"secondaryAzurePort"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix *string `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey *string `pulumi:"sharedKey"`
	// The peering state.
	State *string `pulumi:"state"`
	// Gets peering stats.
	Stats *ExpressRouteCircuitStatsResponse `pulumi:"stats"`
	// The VLAN ID.
	VlanId *int `pulumi:"vlanId"`
}
