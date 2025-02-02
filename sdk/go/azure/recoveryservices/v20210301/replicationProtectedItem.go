// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replication protected item.
type ReplicationProtectedItem struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// The custom data.
	Properties ReplicationProtectedItemPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationProtectedItem registers a new resource with the given unique name, arguments, and options.
func NewReplicationProtectedItem(ctx *pulumi.Context,
	name string, args *ReplicationProtectedItemArgs, opts ...pulumi.ResourceOption) (*ReplicationProtectedItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.ProtectionContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ProtectionContainerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210301:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20160810:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20180110:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20180710:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:recoveryservices/v20210210:ReplicationProtectedItem"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationProtectedItem
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210301:ReplicationProtectedItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationProtectedItem gets an existing ReplicationProtectedItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationProtectedItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationProtectedItemState, opts ...pulumi.ResourceOption) (*ReplicationProtectedItem, error) {
	var resource ReplicationProtectedItem
	err := ctx.ReadResource("azure-native:recoveryservices/v20210301:ReplicationProtectedItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationProtectedItem resources.
type replicationProtectedItemState struct {
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// The custom data.
	Properties *ReplicationProtectedItemPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type *string `pulumi:"type"`
}

type ReplicationProtectedItemState struct {
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringPtrInput
	// The custom data.
	Properties ReplicationProtectedItemPropertiesResponsePtrInput
	// Resource Type
	Type pulumi.StringPtrInput
}

func (ReplicationProtectedItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationProtectedItemState)(nil)).Elem()
}

type replicationProtectedItemArgs struct {
	// Name of the fabric.
	FabricName string `pulumi:"fabricName"`
	// Enable protection input properties.
	Properties *EnableProtectionInputProperties `pulumi:"properties"`
	// Protection container name.
	ProtectionContainerName string `pulumi:"protectionContainerName"`
	// A name for the replication protected item.
	ReplicatedProtectedItemName *string `pulumi:"replicatedProtectedItemName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationProtectedItem resource.
type ReplicationProtectedItemArgs struct {
	// Name of the fabric.
	FabricName pulumi.StringInput
	// Enable protection input properties.
	Properties EnableProtectionInputPropertiesPtrInput
	// Protection container name.
	ProtectionContainerName pulumi.StringInput
	// A name for the replication protected item.
	ReplicatedProtectedItemName pulumi.StringPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationProtectedItemArgs)(nil)).Elem()
}

type ReplicationProtectedItemInput interface {
	pulumi.Input

	ToReplicationProtectedItemOutput() ReplicationProtectedItemOutput
	ToReplicationProtectedItemOutputWithContext(ctx context.Context) ReplicationProtectedItemOutput
}

func (*ReplicationProtectedItem) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationProtectedItem)(nil))
}

func (i *ReplicationProtectedItem) ToReplicationProtectedItemOutput() ReplicationProtectedItemOutput {
	return i.ToReplicationProtectedItemOutputWithContext(context.Background())
}

func (i *ReplicationProtectedItem) ToReplicationProtectedItemOutputWithContext(ctx context.Context) ReplicationProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationProtectedItemOutput)
}

type ReplicationProtectedItemOutput struct {
	*pulumi.OutputState
}

func (ReplicationProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationProtectedItem)(nil))
}

func (o ReplicationProtectedItemOutput) ToReplicationProtectedItemOutput() ReplicationProtectedItemOutput {
	return o
}

func (o ReplicationProtectedItemOutput) ToReplicationProtectedItemOutputWithContext(ctx context.Context) ReplicationProtectedItemOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationProtectedItemOutput{})
}
