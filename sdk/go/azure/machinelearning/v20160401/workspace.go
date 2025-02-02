// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An object that represents a machine learning workspace.
type Workspace struct {
	pulumi.CustomResourceState

	// The creation time for this workspace resource.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The key vault identifier used for encrypted workspaces.
	KeyVaultIdentifierId pulumi.StringPtrOutput `pulumi:"keyVaultIdentifierId"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The email id of the owner for this workspace.
	OwnerEmail pulumi.StringOutput `pulumi:"ownerEmail"`
	// The regional endpoint for the machine learning studio service which hosts this workspace.
	StudioEndpoint pulumi.StringOutput `pulumi:"studioEndpoint"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The fully qualified arm id of the storage account associated with this workspace.
	UserStorageAccountId pulumi.StringOutput `pulumi:"userStorageAccountId"`
	// The immutable id associated with this workspace.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
	// The current state of workspace resource.
	WorkspaceState pulumi.StringOutput `pulumi:"workspaceState"`
	// The type of this workspace.
	WorkspaceType pulumi.StringOutput `pulumi:"workspaceType"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OwnerEmail == nil {
		return nil, errors.New("invalid value for required argument 'OwnerEmail'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserStorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'UserStorageAccountId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearning/v20160401:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearning:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearning:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearning/v20191001:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearning/v20191001:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:machinelearning/v20160401:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:machinelearning/v20160401:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
	// The creation time for this workspace resource.
	CreationTime *string `pulumi:"creationTime"`
	// The key vault identifier used for encrypted workspaces.
	KeyVaultIdentifierId *string `pulumi:"keyVaultIdentifierId"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The email id of the owner for this workspace.
	OwnerEmail *string `pulumi:"ownerEmail"`
	// The regional endpoint for the machine learning studio service which hosts this workspace.
	StudioEndpoint *string `pulumi:"studioEndpoint"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// The fully qualified arm id of the storage account associated with this workspace.
	UserStorageAccountId *string `pulumi:"userStorageAccountId"`
	// The immutable id associated with this workspace.
	WorkspaceId *string `pulumi:"workspaceId"`
	// The current state of workspace resource.
	WorkspaceState *string `pulumi:"workspaceState"`
	// The type of this workspace.
	WorkspaceType *string `pulumi:"workspaceType"`
}

type WorkspaceState struct {
	// The creation time for this workspace resource.
	CreationTime pulumi.StringPtrInput
	// The key vault identifier used for encrypted workspaces.
	KeyVaultIdentifierId pulumi.StringPtrInput
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The email id of the owner for this workspace.
	OwnerEmail pulumi.StringPtrInput
	// The regional endpoint for the machine learning studio service which hosts this workspace.
	StudioEndpoint pulumi.StringPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
	// The fully qualified arm id of the storage account associated with this workspace.
	UserStorageAccountId pulumi.StringPtrInput
	// The immutable id associated with this workspace.
	WorkspaceId pulumi.StringPtrInput
	// The current state of workspace resource.
	WorkspaceState pulumi.StringPtrInput
	// The type of this workspace.
	WorkspaceType pulumi.StringPtrInput
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The key vault identifier used for encrypted workspaces.
	KeyVaultIdentifierId *string `pulumi:"keyVaultIdentifierId"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The email id of the owner for this workspace.
	OwnerEmail string `pulumi:"ownerEmail"`
	// The name of the resource group to which the machine learning workspace belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The fully qualified arm id of the storage account associated with this workspace.
	UserStorageAccountId string `pulumi:"userStorageAccountId"`
	// The name of the machine learning workspace.
	WorkspaceName *string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The key vault identifier used for encrypted workspaces.
	KeyVaultIdentifierId pulumi.StringPtrInput
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The email id of the owner for this workspace.
	OwnerEmail pulumi.StringInput
	// The name of the resource group to which the machine learning workspace belongs.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The fully qualified arm id of the storage account associated with this workspace.
	UserStorageAccountId pulumi.StringInput
	// The name of the machine learning workspace.
	WorkspaceName pulumi.StringPtrInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((*Workspace)(nil))
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

type WorkspaceOutput struct {
	*pulumi.OutputState
}

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Workspace)(nil))
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
