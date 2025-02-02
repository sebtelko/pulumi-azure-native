// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure SQL Database sync agent.
type SyncAgent struct {
	pulumi.CustomResourceState

	// Expiration time of the sync agent version.
	ExpiryTime pulumi.StringOutput `pulumi:"expiryTime"`
	// If the sync agent version is up to date.
	IsUpToDate pulumi.BoolOutput `pulumi:"isUpToDate"`
	// Last alive time of the sync agent.
	LastAliveTime pulumi.StringOutput `pulumi:"lastAliveTime"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the sync agent.
	State pulumi.StringOutput `pulumi:"state"`
	// ARM resource id of the sync database in the sync agent.
	SyncDatabaseId pulumi.StringPtrOutput `pulumi:"syncDatabaseId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the sync agent.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewSyncAgent registers a new resource with the given unique name, arguments, and options.
func NewSyncAgent(ctx *pulumi.Context,
	name string, args *SyncAgentArgs, opts ...pulumi.ResourceOption) (*SyncAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20150501preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:SyncAgent"),
		},
	})
	opts = append(opts, aliases)
	var resource SyncAgent
	err := ctx.RegisterResource("azure-native:sql/v20200801preview:SyncAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncAgent gets an existing SyncAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncAgentState, opts ...pulumi.ResourceOption) (*SyncAgent, error) {
	var resource SyncAgent
	err := ctx.ReadResource("azure-native:sql/v20200801preview:SyncAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncAgent resources.
type syncAgentState struct {
	// Expiration time of the sync agent version.
	ExpiryTime *string `pulumi:"expiryTime"`
	// If the sync agent version is up to date.
	IsUpToDate *bool `pulumi:"isUpToDate"`
	// Last alive time of the sync agent.
	LastAliveTime *string `pulumi:"lastAliveTime"`
	// Resource name.
	Name *string `pulumi:"name"`
	// State of the sync agent.
	State *string `pulumi:"state"`
	// ARM resource id of the sync database in the sync agent.
	SyncDatabaseId *string `pulumi:"syncDatabaseId"`
	// Resource type.
	Type *string `pulumi:"type"`
	// Version of the sync agent.
	Version *string `pulumi:"version"`
}

type SyncAgentState struct {
	// Expiration time of the sync agent version.
	ExpiryTime pulumi.StringPtrInput
	// If the sync agent version is up to date.
	IsUpToDate pulumi.BoolPtrInput
	// Last alive time of the sync agent.
	LastAliveTime pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// State of the sync agent.
	State pulumi.StringPtrInput
	// ARM resource id of the sync database in the sync agent.
	SyncDatabaseId pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// Version of the sync agent.
	Version pulumi.StringPtrInput
}

func (SyncAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAgentState)(nil)).Elem()
}

type syncAgentArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server on which the sync agent is hosted.
	ServerName string `pulumi:"serverName"`
	// The name of the sync agent.
	SyncAgentName *string `pulumi:"syncAgentName"`
	// ARM resource id of the sync database in the sync agent.
	SyncDatabaseId *string `pulumi:"syncDatabaseId"`
}

// The set of arguments for constructing a SyncAgent resource.
type SyncAgentArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server on which the sync agent is hosted.
	ServerName pulumi.StringInput
	// The name of the sync agent.
	SyncAgentName pulumi.StringPtrInput
	// ARM resource id of the sync database in the sync agent.
	SyncDatabaseId pulumi.StringPtrInput
}

func (SyncAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAgentArgs)(nil)).Elem()
}

type SyncAgentInput interface {
	pulumi.Input

	ToSyncAgentOutput() SyncAgentOutput
	ToSyncAgentOutputWithContext(ctx context.Context) SyncAgentOutput
}

func (*SyncAgent) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncAgent)(nil))
}

func (i *SyncAgent) ToSyncAgentOutput() SyncAgentOutput {
	return i.ToSyncAgentOutputWithContext(context.Background())
}

func (i *SyncAgent) ToSyncAgentOutputWithContext(ctx context.Context) SyncAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAgentOutput)
}

type SyncAgentOutput struct {
	*pulumi.OutputState
}

func (SyncAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncAgent)(nil))
}

func (o SyncAgentOutput) ToSyncAgentOutput() SyncAgentOutput {
	return o
}

func (o SyncAgentOutput) ToSyncAgentOutputWithContext(ctx context.Context) SyncAgentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SyncAgentOutput{})
}
