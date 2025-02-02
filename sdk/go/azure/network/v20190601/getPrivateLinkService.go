// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Private link service resource.
func LookupPrivateLinkService(ctx *pulumi.Context, args *LookupPrivateLinkServiceArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkServiceResult, error) {
	var rv LookupPrivateLinkServiceResult
	err := ctx.Invoke("azure-native:network/v20190601:getPrivateLinkService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkServiceArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the private link service.
	ServiceName string `pulumi:"serviceName"`
}

// Private link service resource.
type LookupPrivateLinkServiceResult struct {
	// The alias of the private link service.
	Alias string `pulumi:"alias"`
	// The auto-approval list of the private link service.
	AutoApproval *PrivateLinkServicePropertiesResponseAutoApproval `pulumi:"autoApproval"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The list of Fqdn.
	Fqdns []string `pulumi:"fqdns"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// An array of references to the private link service IP configuration.
	IpConfigurations []PrivateLinkServiceIpConfigurationResponse `pulumi:"ipConfigurations"`
	// An array of references to the load balancer IP configurations.
	LoadBalancerFrontendIpConfigurations []FrontendIPConfigurationResponse `pulumi:"loadBalancerFrontendIpConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Gets an array of references to the network interfaces created for this private link service.
	NetworkInterfaces []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	// An array of list about connections to the private endpoint.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The provisioning state of the private link service.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The visibility list of the private link service.
	Visibility *PrivateLinkServicePropertiesResponseVisibility `pulumi:"visibility"`
}
