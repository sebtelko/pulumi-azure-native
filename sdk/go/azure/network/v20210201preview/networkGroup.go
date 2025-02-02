// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The network group resource
type NetworkGroup struct {
	pulumi.CustomResourceState

	// Network group conditional filter.
	ConditionalMembership pulumi.StringPtrOutput `pulumi:"conditionalMembership"`
	// A description of the network group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A friendly name for the network group.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Group members of network group.
	GroupMembers GroupMembersItemResponseArrayOutput `pulumi:"groupMembers"`
	// Group member type.
	MemberType pulumi.StringPtrOutput `pulumi:"memberType"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the scope assignment resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkGroup(ctx *pulumi.Context,
	name string, args *NetworkGroupArgs, opts ...pulumi.ResourceOption) (*NetworkGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20210201preview:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-native:network:NetworkGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:NetworkGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkGroup
	err := ctx.RegisterResource("azure-native:network/v20210201preview:NetworkGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkGroup gets an existing NetworkGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkGroupState, opts ...pulumi.ResourceOption) (*NetworkGroup, error) {
	var resource NetworkGroup
	err := ctx.ReadResource("azure-native:network/v20210201preview:NetworkGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkGroup resources.
type networkGroupState struct {
	// Network group conditional filter.
	ConditionalMembership *string `pulumi:"conditionalMembership"`
	// A description of the network group.
	Description *string `pulumi:"description"`
	// A friendly name for the network group.
	DisplayName *string `pulumi:"displayName"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Group members of network group.
	GroupMembers []GroupMembersItemResponse `pulumi:"groupMembers"`
	// Group member type.
	MemberType *string `pulumi:"memberType"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The provisioning state of the scope assignment resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The system metadata related to this resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type NetworkGroupState struct {
	// Network group conditional filter.
	ConditionalMembership pulumi.StringPtrInput
	// A description of the network group.
	Description pulumi.StringPtrInput
	// A friendly name for the network group.
	DisplayName pulumi.StringPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Group members of network group.
	GroupMembers GroupMembersItemResponseArrayInput
	// Group member type.
	MemberType pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The provisioning state of the scope assignment resource.
	ProvisioningState pulumi.StringPtrInput
	// The system metadata related to this resource.
	SystemData SystemDataResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (NetworkGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkGroupState)(nil)).Elem()
}

type networkGroupArgs struct {
	// Network group conditional filter.
	ConditionalMembership *string `pulumi:"conditionalMembership"`
	// A description of the network group.
	Description *string `pulumi:"description"`
	// A friendly name for the network group.
	DisplayName *string `pulumi:"displayName"`
	// Group members of network group.
	GroupMembers []GroupMembersItem `pulumi:"groupMembers"`
	// Group member type.
	MemberType *string `pulumi:"memberType"`
	// The name of the network group to get.
	NetworkGroupName *string `pulumi:"networkGroupName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NetworkGroup resource.
type NetworkGroupArgs struct {
	// Network group conditional filter.
	ConditionalMembership pulumi.StringPtrInput
	// A description of the network group.
	Description pulumi.StringPtrInput
	// A friendly name for the network group.
	DisplayName pulumi.StringPtrInput
	// Group members of network group.
	GroupMembers GroupMembersItemArrayInput
	// Group member type.
	MemberType pulumi.StringPtrInput
	// The name of the network group to get.
	NetworkGroupName pulumi.StringPtrInput
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (NetworkGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkGroupArgs)(nil)).Elem()
}

type NetworkGroupInput interface {
	pulumi.Input

	ToNetworkGroupOutput() NetworkGroupOutput
	ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput
}

func (*NetworkGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkGroup)(nil))
}

func (i *NetworkGroup) ToNetworkGroupOutput() NetworkGroupOutput {
	return i.ToNetworkGroupOutputWithContext(context.Background())
}

func (i *NetworkGroup) ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkGroupOutput)
}

type NetworkGroupOutput struct {
	*pulumi.OutputState
}

func (NetworkGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkGroup)(nil))
}

func (o NetworkGroupOutput) ToNetworkGroupOutput() NetworkGroupOutput {
	return o
}

func (o NetworkGroupOutput) ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkGroupOutput{})
}
