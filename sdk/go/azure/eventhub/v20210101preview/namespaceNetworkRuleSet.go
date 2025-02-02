// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of topic resource.
type NamespaceNetworkRuleSet struct {
	pulumi.CustomResourceState

	// Default Action for Network Rule Set
	DefaultAction pulumi.StringPtrOutput `pulumi:"defaultAction"`
	// List of IpRules
	IpRules NWRuleSetIpRulesResponseArrayOutput `pulumi:"ipRules"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Value that indicates whether Trusted Service Access is Enabled or not.
	TrustedServiceAccessEnabled pulumi.BoolPtrOutput `pulumi:"trustedServiceAccessEnabled"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// List VirtualNetwork Rules
	VirtualNetworkRules NWRuleSetVirtualNetworkRulesResponseArrayOutput `pulumi:"virtualNetworkRules"`
}

// NewNamespaceNetworkRuleSet registers a new resource with the given unique name, arguments, and options.
func NewNamespaceNetworkRuleSet(ctx *pulumi.Context,
	name string, args *NamespaceNetworkRuleSetArgs, opts ...pulumi.ResourceOption) (*NamespaceNetworkRuleSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20210101preview:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-native:eventhub:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20170401:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20180101preview:NamespaceNetworkRuleSet"),
		},
	})
	opts = append(opts, aliases)
	var resource NamespaceNetworkRuleSet
	err := ctx.RegisterResource("azure-native:eventhub/v20210101preview:NamespaceNetworkRuleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceNetworkRuleSet gets an existing NamespaceNetworkRuleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceNetworkRuleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceNetworkRuleSetState, opts ...pulumi.ResourceOption) (*NamespaceNetworkRuleSet, error) {
	var resource NamespaceNetworkRuleSet
	err := ctx.ReadResource("azure-native:eventhub/v20210101preview:NamespaceNetworkRuleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceNetworkRuleSet resources.
type namespaceNetworkRuleSetState struct {
	// Default Action for Network Rule Set
	DefaultAction *string `pulumi:"defaultAction"`
	// List of IpRules
	IpRules []NWRuleSetIpRulesResponse `pulumi:"ipRules"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The system meta data relating to this resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Value that indicates whether Trusted Service Access is Enabled or not.
	TrustedServiceAccessEnabled *bool `pulumi:"trustedServiceAccessEnabled"`
	// Resource type.
	Type *string `pulumi:"type"`
	// List VirtualNetwork Rules
	VirtualNetworkRules []NWRuleSetVirtualNetworkRulesResponse `pulumi:"virtualNetworkRules"`
}

type NamespaceNetworkRuleSetState struct {
	// Default Action for Network Rule Set
	DefaultAction pulumi.StringPtrInput
	// List of IpRules
	IpRules NWRuleSetIpRulesResponseArrayInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The system meta data relating to this resource.
	SystemData SystemDataResponsePtrInput
	// Value that indicates whether Trusted Service Access is Enabled or not.
	TrustedServiceAccessEnabled pulumi.BoolPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// List VirtualNetwork Rules
	VirtualNetworkRules NWRuleSetVirtualNetworkRulesResponseArrayInput
}

func (NamespaceNetworkRuleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceNetworkRuleSetState)(nil)).Elem()
}

type namespaceNetworkRuleSetArgs struct {
	// Default Action for Network Rule Set
	DefaultAction *string `pulumi:"defaultAction"`
	// List of IpRules
	IpRules []NWRuleSetIpRules `pulumi:"ipRules"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Value that indicates whether Trusted Service Access is Enabled or not.
	TrustedServiceAccessEnabled *bool `pulumi:"trustedServiceAccessEnabled"`
	// List VirtualNetwork Rules
	VirtualNetworkRules []NWRuleSetVirtualNetworkRules `pulumi:"virtualNetworkRules"`
}

// The set of arguments for constructing a NamespaceNetworkRuleSet resource.
type NamespaceNetworkRuleSetArgs struct {
	// Default Action for Network Rule Set
	DefaultAction pulumi.StringPtrInput
	// List of IpRules
	IpRules NWRuleSetIpRulesArrayInput
	// The Namespace name
	NamespaceName pulumi.StringInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
	// Value that indicates whether Trusted Service Access is Enabled or not.
	TrustedServiceAccessEnabled pulumi.BoolPtrInput
	// List VirtualNetwork Rules
	VirtualNetworkRules NWRuleSetVirtualNetworkRulesArrayInput
}

func (NamespaceNetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceNetworkRuleSetArgs)(nil)).Elem()
}

type NamespaceNetworkRuleSetInput interface {
	pulumi.Input

	ToNamespaceNetworkRuleSetOutput() NamespaceNetworkRuleSetOutput
	ToNamespaceNetworkRuleSetOutputWithContext(ctx context.Context) NamespaceNetworkRuleSetOutput
}

func (*NamespaceNetworkRuleSet) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceNetworkRuleSet)(nil))
}

func (i *NamespaceNetworkRuleSet) ToNamespaceNetworkRuleSetOutput() NamespaceNetworkRuleSetOutput {
	return i.ToNamespaceNetworkRuleSetOutputWithContext(context.Background())
}

func (i *NamespaceNetworkRuleSet) ToNamespaceNetworkRuleSetOutputWithContext(ctx context.Context) NamespaceNetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceNetworkRuleSetOutput)
}

type NamespaceNetworkRuleSetOutput struct {
	*pulumi.OutputState
}

func (NamespaceNetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceNetworkRuleSet)(nil))
}

func (o NamespaceNetworkRuleSetOutput) ToNamespaceNetworkRuleSetOutput() NamespaceNetworkRuleSetOutput {
	return o
}

func (o NamespaceNetworkRuleSetOutput) ToNamespaceNetworkRuleSetOutputWithContext(ctx context.Context) NamespaceNetworkRuleSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamespaceNetworkRuleSetOutput{})
}
