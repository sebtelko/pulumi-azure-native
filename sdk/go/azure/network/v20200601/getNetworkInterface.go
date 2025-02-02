// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A network interface in a resource group.
func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:network/v20200601:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInterfaceArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the network interface.
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A network interface in a resource group.
type LookupNetworkInterfaceResult struct {
	// The DNS settings in network interface.
	DnsSettings *NetworkInterfaceDnsSettingsResponse `pulumi:"dnsSettings"`
	// A reference to the dscp configuration to which the network interface is linked.
	DscpConfiguration SubResourceResponse `pulumi:"dscpConfiguration"`
	// If the network interface is accelerated networking enabled.
	EnableAcceleratedNetworking *bool `pulumi:"enableAcceleratedNetworking"`
	// Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding *bool `pulumi:"enableIPForwarding"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// A list of references to linked BareMetal resources.
	HostedWorkloads []string `pulumi:"hostedWorkloads"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// A list of IPConfigurations of the network interface.
	IpConfigurations []NetworkInterfaceIPConfigurationResponse `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The MAC address of the network interface.
	MacAddress string `pulumi:"macAddress"`
	// Resource name.
	Name string `pulumi:"name"`
	// The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupResponse `pulumi:"networkSecurityGroup"`
	// Whether this is a primary network interface on a virtual machine.
	Primary bool `pulumi:"primary"`
	// A reference to the private endpoint to which the network interface is linked.
	PrivateEndpoint PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// The provisioning state of the network interface resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the network interface resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A list of TapConfigurations of the network interface.
	TapConfigurations []NetworkInterfaceTapConfigurationResponse `pulumi:"tapConfigurations"`
	// Resource type.
	Type string `pulumi:"type"`
	// The reference to a virtual machine.
	VirtualMachine SubResourceResponse `pulumi:"virtualMachine"`
}
