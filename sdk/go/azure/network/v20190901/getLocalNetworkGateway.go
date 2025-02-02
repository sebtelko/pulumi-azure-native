// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A common class for general resource information.
func LookupLocalNetworkGateway(ctx *pulumi.Context, args *LookupLocalNetworkGatewayArgs, opts ...pulumi.InvokeOption) (*LookupLocalNetworkGatewayResult, error) {
	var rv LookupLocalNetworkGatewayResult
	err := ctx.Invoke("azure-native:network/v20190901:getLocalNetworkGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocalNetworkGatewayArgs struct {
	// The name of the local network gateway.
	LocalNetworkGatewayName string `pulumi:"localNetworkGatewayName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A common class for general resource information.
type LookupLocalNetworkGatewayResult struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings *BgpSettingsResponse `pulumi:"bgpSettings"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// IP address of local network gateway.
	GatewayIpAddress *string `pulumi:"gatewayIpAddress"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Local network site address space.
	LocalNetworkAddressSpace *AddressSpaceResponse `pulumi:"localNetworkAddressSpace"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the local network gateway resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the local network gateway resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}
