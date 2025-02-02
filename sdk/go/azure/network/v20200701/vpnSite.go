// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VpnSite Resource.
type VpnSite struct {
	pulumi.CustomResourceState

	// The AddressSpace that contains an array of IP address ranges.
	AddressSpace AddressSpaceResponsePtrOutput `pulumi:"addressSpace"`
	// The set of bgp properties.
	BgpProperties BgpSettingsResponsePtrOutput `pulumi:"bgpProperties"`
	// The device properties.
	DeviceProperties DevicePropertiesResponsePtrOutput `pulumi:"deviceProperties"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The ip-address for the vpn-site.
	IpAddress pulumi.StringPtrOutput `pulumi:"ipAddress"`
	// IsSecuritySite flag.
	IsSecuritySite pulumi.BoolPtrOutput `pulumi:"isSecuritySite"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Office365 Policy.
	O365Policy O365PolicyPropertiesResponsePtrOutput `pulumi:"o365Policy"`
	// The provisioning state of the VPN site resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The key for vpn-site that can be used for connections.
	SiteKey pulumi.StringPtrOutput `pulumi:"siteKey"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The VirtualWAN to which the vpnSite belongs.
	VirtualWan SubResourceResponsePtrOutput `pulumi:"virtualWan"`
	// List of all vpn site links.
	VpnSiteLinks VpnSiteLinkResponseArrayOutput `pulumi:"vpnSiteLinks"`
}

// NewVpnSite registers a new resource with the given unique name, arguments, and options.
func NewVpnSite(ctx *pulumi.Context,
	name string, args *VpnSiteArgs, opts ...pulumi.ResourceOption) (*VpnSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:VpnSite"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VpnSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:VpnSite"),
		},
	})
	opts = append(opts, aliases)
	var resource VpnSite
	err := ctx.RegisterResource("azure-native:network/v20200701:VpnSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnSite gets an existing VpnSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnSiteState, opts ...pulumi.ResourceOption) (*VpnSite, error) {
	var resource VpnSite
	err := ctx.ReadResource("azure-native:network/v20200701:VpnSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnSite resources.
type vpnSiteState struct {
	// The AddressSpace that contains an array of IP address ranges.
	AddressSpace *AddressSpaceResponse `pulumi:"addressSpace"`
	// The set of bgp properties.
	BgpProperties *BgpSettingsResponse `pulumi:"bgpProperties"`
	// The device properties.
	DeviceProperties *DevicePropertiesResponse `pulumi:"deviceProperties"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The ip-address for the vpn-site.
	IpAddress *string `pulumi:"ipAddress"`
	// IsSecuritySite flag.
	IsSecuritySite *bool `pulumi:"isSecuritySite"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Office365 Policy.
	O365Policy *O365PolicyPropertiesResponse `pulumi:"o365Policy"`
	// The provisioning state of the VPN site resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The key for vpn-site that can be used for connections.
	SiteKey *string `pulumi:"siteKey"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The VirtualWAN to which the vpnSite belongs.
	VirtualWan *SubResourceResponse `pulumi:"virtualWan"`
	// List of all vpn site links.
	VpnSiteLinks []VpnSiteLinkResponse `pulumi:"vpnSiteLinks"`
}

type VpnSiteState struct {
	// The AddressSpace that contains an array of IP address ranges.
	AddressSpace AddressSpaceResponsePtrInput
	// The set of bgp properties.
	BgpProperties BgpSettingsResponsePtrInput
	// The device properties.
	DeviceProperties DevicePropertiesResponsePtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The ip-address for the vpn-site.
	IpAddress pulumi.StringPtrInput
	// IsSecuritySite flag.
	IsSecuritySite pulumi.BoolPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Office365 Policy.
	O365Policy O365PolicyPropertiesResponsePtrInput
	// The provisioning state of the VPN site resource.
	ProvisioningState pulumi.StringPtrInput
	// The key for vpn-site that can be used for connections.
	SiteKey pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The VirtualWAN to which the vpnSite belongs.
	VirtualWan SubResourceResponsePtrInput
	// List of all vpn site links.
	VpnSiteLinks VpnSiteLinkResponseArrayInput
}

func (VpnSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSiteState)(nil)).Elem()
}

type vpnSiteArgs struct {
	// The AddressSpace that contains an array of IP address ranges.
	AddressSpace *AddressSpace `pulumi:"addressSpace"`
	// The set of bgp properties.
	BgpProperties *BgpSettings `pulumi:"bgpProperties"`
	// The device properties.
	DeviceProperties *DeviceProperties `pulumi:"deviceProperties"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The ip-address for the vpn-site.
	IpAddress *string `pulumi:"ipAddress"`
	// IsSecuritySite flag.
	IsSecuritySite *bool `pulumi:"isSecuritySite"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Office365 Policy.
	O365Policy *O365PolicyProperties `pulumi:"o365Policy"`
	// The resource group name of the VpnSite.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The key for vpn-site that can be used for connections.
	SiteKey *string `pulumi:"siteKey"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The VirtualWAN to which the vpnSite belongs.
	VirtualWan *SubResource `pulumi:"virtualWan"`
	// List of all vpn site links.
	VpnSiteLinks []VpnSiteLink `pulumi:"vpnSiteLinks"`
	// The name of the VpnSite being created or updated.
	VpnSiteName *string `pulumi:"vpnSiteName"`
}

// The set of arguments for constructing a VpnSite resource.
type VpnSiteArgs struct {
	// The AddressSpace that contains an array of IP address ranges.
	AddressSpace AddressSpacePtrInput
	// The set of bgp properties.
	BgpProperties BgpSettingsPtrInput
	// The device properties.
	DeviceProperties DevicePropertiesPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The ip-address for the vpn-site.
	IpAddress pulumi.StringPtrInput
	// IsSecuritySite flag.
	IsSecuritySite pulumi.BoolPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Office365 Policy.
	O365Policy O365PolicyPropertiesPtrInput
	// The resource group name of the VpnSite.
	ResourceGroupName pulumi.StringInput
	// The key for vpn-site that can be used for connections.
	SiteKey pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The VirtualWAN to which the vpnSite belongs.
	VirtualWan SubResourcePtrInput
	// List of all vpn site links.
	VpnSiteLinks VpnSiteLinkArrayInput
	// The name of the VpnSite being created or updated.
	VpnSiteName pulumi.StringPtrInput
}

func (VpnSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSiteArgs)(nil)).Elem()
}

type VpnSiteInput interface {
	pulumi.Input

	ToVpnSiteOutput() VpnSiteOutput
	ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput
}

func (*VpnSite) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSite)(nil))
}

func (i *VpnSite) ToVpnSiteOutput() VpnSiteOutput {
	return i.ToVpnSiteOutputWithContext(context.Background())
}

func (i *VpnSite) ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteOutput)
}

type VpnSiteOutput struct {
	*pulumi.OutputState
}

func (VpnSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSite)(nil))
}

func (o VpnSiteOutput) ToVpnSiteOutput() VpnSiteOutput {
	return o
}

func (o VpnSiteOutput) ToVpnSiteOutputWithContext(ctx context.Context) VpnSiteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VpnSiteOutput{})
}
