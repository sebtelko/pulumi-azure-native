// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A virtual network rule.
type VirtualNetworkRule struct {
	pulumi.CustomResourceState

	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrOutput `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Virtual Network Rule State
	State pulumi.StringOutput `pulumi:"state"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId pulumi.StringOutput `pulumi:"virtualNetworkSubnetId"`
}

// NewVirtualNetworkRule registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkRule(ctx *pulumi.Context,
	name string, args *VirtualNetworkRuleArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.VirtualNetworkSubnetId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetworkSubnetId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20150501preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:VirtualNetworkRule"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkRule
	err := ctx.RegisterResource("azure-native:sql/v20201101preview:VirtualNetworkRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkRule gets an existing VirtualNetworkRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkRuleState, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	var resource VirtualNetworkRule
	err := ctx.ReadResource("azure-native:sql/v20201101preview:VirtualNetworkRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkRule resources.
type virtualNetworkRuleState struct {
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint *bool `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Virtual Network Rule State
	State *string `pulumi:"state"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId *string `pulumi:"virtualNetworkSubnetId"`
}

type VirtualNetworkRuleState struct {
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Virtual Network Rule State
	State pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId pulumi.StringPtrInput
}

func (VirtualNetworkRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleState)(nil)).Elem()
}

type virtualNetworkRuleArgs struct {
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint *bool `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the virtual network rule.
	VirtualNetworkRuleName *string `pulumi:"virtualNetworkRuleName"`
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId string `pulumi:"virtualNetworkSubnetId"`
}

// The set of arguments for constructing a VirtualNetworkRule resource.
type VirtualNetworkRuleArgs struct {
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The name of the virtual network rule.
	VirtualNetworkRuleName pulumi.StringPtrInput
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId pulumi.StringInput
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleArgs)(nil)).Elem()
}

type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput
}

func (*VirtualNetworkRule) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil))
}

func (i *VirtualNetworkRule) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i *VirtualNetworkRule) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleOutput struct {
	*pulumi.OutputState
}

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil))
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
}
