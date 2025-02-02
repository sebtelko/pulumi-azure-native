// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Virtual Network Tap resource.
func LookupVirtualNetworkTap(ctx *pulumi.Context, args *LookupVirtualNetworkTapArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkTapResult, error) {
	var rv LookupVirtualNetworkTapResult
	err := ctx.Invoke("azure-native:network/v20191101:getVirtualNetworkTap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkTapArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of virtual network tap.
	TapName string `pulumi:"tapName"`
}

// Virtual Network Tap resource.
type LookupVirtualNetworkTapResult struct {
	// The reference to the private IP address on the internal Load Balancer that will receive the tap.
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfigurationResponse `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	// The reference to the private IP Address of the collector nic that will receive the tap.
	DestinationNetworkInterfaceIPConfiguration *NetworkInterfaceIPConfigurationResponse `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	// The VXLAN destination port that will receive the tapped traffic.
	DestinationPort *int `pulumi:"destinationPort"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Specifies the list of resource IDs for the network interface IP configuration that needs to be tapped.
	NetworkInterfaceTapConfigurations []NetworkInterfaceTapConfigurationResponse `pulumi:"networkInterfaceTapConfigurations"`
	// The provisioning state of the virtual network tap resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the virtual network tap resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}
