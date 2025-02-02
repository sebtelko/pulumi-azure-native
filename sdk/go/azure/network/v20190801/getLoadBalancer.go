// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// LoadBalancer resource.
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("azure-native:network/v20190801:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// LoadBalancer resource.
type LookupLoadBalancerResult struct {
	// Collection of backend address pools used by a load balancer.
	BackendAddressPools []BackendAddressPoolResponse `pulumi:"backendAddressPools"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Object representing the frontend IPs to be used for the load balancer.
	FrontendIPConfigurations []FrontendIPConfigurationResponse `pulumi:"frontendIPConfigurations"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools []InboundNatPoolResponse `pulumi:"inboundNatPools"`
	// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatRules []InboundNatRuleResponse `pulumi:"inboundNatRules"`
	// Object collection representing the load balancing rules Gets the provisioning.
	LoadBalancingRules []LoadBalancingRuleResponse `pulumi:"loadBalancingRules"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The outbound rules.
	OutboundRules []OutboundRuleResponse `pulumi:"outboundRules"`
	// Collection of probe objects used in the load balancer.
	Probes []ProbeResponse `pulumi:"probes"`
	// The provisioning state of the load balancer resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the load balancer resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// The load balancer SKU.
	Sku *LoadBalancerSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}
