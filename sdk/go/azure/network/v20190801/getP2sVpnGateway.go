// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// P2SVpnGateway Resource.
func LookupP2sVpnGateway(ctx *pulumi.Context, args *LookupP2sVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupP2sVpnGatewayResult, error) {
	var rv LookupP2sVpnGatewayResult
	err := ctx.Invoke("azure-native:network/v20190801:getP2sVpnGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupP2sVpnGatewayArgs struct {
	// The name of the gateway.
	GatewayName string `pulumi:"gatewayName"`
	// The resource group name of the P2SVpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// P2SVpnGateway Resource.
type LookupP2sVpnGatewayResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// List of all p2s connection configurations of the gateway.
	P2SConnectionConfigurations []P2SConnectionConfigurationResponse `pulumi:"p2SConnectionConfigurations"`
	// The provisioning state of the P2S VPN gateway resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The VirtualHub to which the gateway belongs.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
	// All P2S VPN clients' connection health status.
	VpnClientConnectionHealth VpnClientConnectionHealthResponse `pulumi:"vpnClientConnectionHealth"`
	// The scale unit for this p2s vpn gateway.
	VpnGatewayScaleUnit *int `pulumi:"vpnGatewayScaleUnit"`
	// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
	VpnServerConfiguration *SubResourceResponse `pulumi:"vpnServerConfiguration"`
}
