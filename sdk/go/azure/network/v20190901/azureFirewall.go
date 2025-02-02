// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Firewall resource.
type AzureFirewall struct {
	pulumi.CustomResourceState

	// The additional properties used to further config this azure firewall.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections AzureFirewallApplicationRuleCollectionResponseArrayOutput `pulumi:"applicationRuleCollections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The firewallPolicy associated with this azure firewall.
	FirewallPolicy SubResourceResponsePtrOutput `pulumi:"firewallPolicy"`
	// IP addresses associated with AzureFirewall.
	HubIpAddresses HubIPAddressesResponseOutput `pulumi:"hubIpAddresses"`
	// IP configuration of the Azure Firewall resource.
	IpConfigurations AzureFirewallIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections AzureFirewallNatRuleCollectionResponseArrayOutput `pulumi:"natRuleCollections"`
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections AzureFirewallNetworkRuleCollectionResponseArrayOutput `pulumi:"networkRuleCollections"`
	// The provisioning state of the Azure firewall resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The Azure Firewall Resource SKU.
	Sku AzureFirewallSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrOutput `pulumi:"threatIntelMode"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The virtualHub to which the firewall belongs.
	VirtualHub SubResourceResponsePtrOutput `pulumi:"virtualHub"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewAzureFirewall registers a new resource with the given unique name, arguments, and options.
func NewAzureFirewall(ctx *pulumi.Context,
	name string, args *AzureFirewallArgs, opts ...pulumi.ResourceOption) (*AzureFirewall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:AzureFirewall"),
		},
	})
	opts = append(opts, aliases)
	var resource AzureFirewall
	err := ctx.RegisterResource("azure-native:network/v20190901:AzureFirewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureFirewall gets an existing AzureFirewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureFirewallState, opts ...pulumi.ResourceOption) (*AzureFirewall, error) {
	var resource AzureFirewall
	err := ctx.ReadResource("azure-native:network/v20190901:AzureFirewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureFirewall resources.
type azureFirewallState struct {
	// The additional properties used to further config this azure firewall.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollectionResponse `pulumi:"applicationRuleCollections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The firewallPolicy associated with this azure firewall.
	FirewallPolicy *SubResourceResponse `pulumi:"firewallPolicy"`
	// IP addresses associated with AzureFirewall.
	HubIpAddresses *HubIPAddressesResponse `pulumi:"hubIpAddresses"`
	// IP configuration of the Azure Firewall resource.
	IpConfigurations []AzureFirewallIPConfigurationResponse `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections []AzureFirewallNatRuleCollectionResponse `pulumi:"natRuleCollections"`
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections []AzureFirewallNetworkRuleCollectionResponse `pulumi:"networkRuleCollections"`
	// The provisioning state of the Azure firewall resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The Azure Firewall Resource SKU.
	Sku *AzureFirewallSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode *string `pulumi:"threatIntelMode"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The virtualHub to which the firewall belongs.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

type AzureFirewallState struct {
	// The additional properties used to further config this azure firewall.
	AdditionalProperties pulumi.StringMapInput
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections AzureFirewallApplicationRuleCollectionResponseArrayInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The firewallPolicy associated with this azure firewall.
	FirewallPolicy SubResourceResponsePtrInput
	// IP addresses associated with AzureFirewall.
	HubIpAddresses HubIPAddressesResponsePtrInput
	// IP configuration of the Azure Firewall resource.
	IpConfigurations AzureFirewallIPConfigurationResponseArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections AzureFirewallNatRuleCollectionResponseArrayInput
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections AzureFirewallNetworkRuleCollectionResponseArrayInput
	// The provisioning state of the Azure firewall resource.
	ProvisioningState pulumi.StringPtrInput
	// The Azure Firewall Resource SKU.
	Sku AzureFirewallSkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The virtualHub to which the firewall belongs.
	VirtualHub SubResourceResponsePtrInput
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (AzureFirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureFirewallState)(nil)).Elem()
}

type azureFirewallArgs struct {
	// The additional properties used to further config this azure firewall.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollection `pulumi:"applicationRuleCollections"`
	// The name of the Azure Firewall.
	AzureFirewallName *string `pulumi:"azureFirewallName"`
	// The firewallPolicy associated with this azure firewall.
	FirewallPolicy *SubResource `pulumi:"firewallPolicy"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// IP configuration of the Azure Firewall resource.
	IpConfigurations []AzureFirewallIPConfiguration `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections []AzureFirewallNatRuleCollection `pulumi:"natRuleCollections"`
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections []AzureFirewallNetworkRuleCollection `pulumi:"networkRuleCollections"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Azure Firewall Resource SKU.
	Sku *AzureFirewallSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode *string `pulumi:"threatIntelMode"`
	// The virtualHub to which the firewall belongs.
	VirtualHub *SubResource `pulumi:"virtualHub"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a AzureFirewall resource.
type AzureFirewallArgs struct {
	// The additional properties used to further config this azure firewall.
	AdditionalProperties pulumi.StringMapInput
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections AzureFirewallApplicationRuleCollectionArrayInput
	// The name of the Azure Firewall.
	AzureFirewallName pulumi.StringPtrInput
	// The firewallPolicy associated with this azure firewall.
	FirewallPolicy SubResourcePtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// IP configuration of the Azure Firewall resource.
	IpConfigurations AzureFirewallIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections AzureFirewallNatRuleCollectionArrayInput
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections AzureFirewallNetworkRuleCollectionArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The Azure Firewall Resource SKU.
	Sku AzureFirewallSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrInput
	// The virtualHub to which the firewall belongs.
	VirtualHub SubResourcePtrInput
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (AzureFirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureFirewallArgs)(nil)).Elem()
}

type AzureFirewallInput interface {
	pulumi.Input

	ToAzureFirewallOutput() AzureFirewallOutput
	ToAzureFirewallOutputWithContext(ctx context.Context) AzureFirewallOutput
}

func (*AzureFirewall) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewall)(nil))
}

func (i *AzureFirewall) ToAzureFirewallOutput() AzureFirewallOutput {
	return i.ToAzureFirewallOutputWithContext(context.Background())
}

func (i *AzureFirewall) ToAzureFirewallOutputWithContext(ctx context.Context) AzureFirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFirewallOutput)
}

type AzureFirewallOutput struct {
	*pulumi.OutputState
}

func (AzureFirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewall)(nil))
}

func (o AzureFirewallOutput) ToAzureFirewallOutput() AzureFirewallOutput {
	return o
}

func (o AzureFirewallOutput) ToAzureFirewallOutputWithContext(ctx context.Context) AzureFirewallOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AzureFirewallOutput{})
}
