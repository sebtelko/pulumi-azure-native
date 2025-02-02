// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A NetworkInterface in a resource group
func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:network/v20150501preview:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInterfaceArgs struct {
	// The name of the network interface.
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A NetworkInterface in a resource group
type LookupNetworkInterfaceResult struct {
	// Gets or sets DNS Settings in  NetworkInterface
	DnsSettings *NetworkInterfaceDnsSettingsResponse `pulumi:"dnsSettings"`
	// Gets or sets whether IPForwarding is enabled on the NIC
	EnableIPForwarding *bool `pulumi:"enableIPForwarding"`
	// Gets a unique read-only string that changes whenever the resource is updated
	Etag *string `pulumi:"etag"`
	// Resource Id
	Id string `pulumi:"id"`
	// Gets or sets list of IPConfigurations of the NetworkInterface
	IpConfigurations []NetworkInterfaceIpConfigurationResponse `pulumi:"ipConfigurations"`
	// Resource location
	Location string `pulumi:"location"`
	// Gets the MAC Address of the network interface
	MacAddress *string `pulumi:"macAddress"`
	// Resource name
	Name string `pulumi:"name"`
	// Gets or sets the reference of the NetworkSecurityGroup resource
	NetworkSecurityGroup *SubResourceResponse `pulumi:"networkSecurityGroup"`
	// Gets whether this is a primary NIC on a virtual machine
	Primary *bool `pulumi:"primary"`
	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets or sets resource guid property of the network interface resource
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// Gets or sets the reference of a VirtualMachine
	VirtualMachine *SubResourceResponse `pulumi:"virtualMachine"`
}
