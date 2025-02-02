// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configures where to store the OMS agent data for workspaces under a scope
type WorkspaceSetting struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// All the VMs in this scope will send their security data to the mentioned workspace unless overridden by a setting with more specific scope
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The full Azure ID of the workspace to save the data in
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewWorkspaceSetting registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceSetting(ctx *pulumi.Context,
	name string, args *WorkspaceSettingArgs, opts ...pulumi.ResourceOption) (*WorkspaceSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:security/v20170801preview:WorkspaceSetting"),
		},
		{
			Type: pulumi.String("azure-native:security:WorkspaceSetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:security:WorkspaceSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkspaceSetting
	err := ctx.RegisterResource("azure-native:security/v20170801preview:WorkspaceSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceSetting gets an existing WorkspaceSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceSettingState, opts ...pulumi.ResourceOption) (*WorkspaceSetting, error) {
	var resource WorkspaceSetting
	err := ctx.ReadResource("azure-native:security/v20170801preview:WorkspaceSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceSetting resources.
type workspaceSettingState struct {
	// Resource name
	Name *string `pulumi:"name"`
	// All the VMs in this scope will send their security data to the mentioned workspace unless overridden by a setting with more specific scope
	Scope *string `pulumi:"scope"`
	// Resource type
	Type *string `pulumi:"type"`
	// The full Azure ID of the workspace to save the data in
	WorkspaceId *string `pulumi:"workspaceId"`
}

type WorkspaceSettingState struct {
	// Resource name
	Name pulumi.StringPtrInput
	// All the VMs in this scope will send their security data to the mentioned workspace unless overridden by a setting with more specific scope
	Scope pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
	// The full Azure ID of the workspace to save the data in
	WorkspaceId pulumi.StringPtrInput
}

func (WorkspaceSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSettingState)(nil)).Elem()
}

type workspaceSettingArgs struct {
	// All the VMs in this scope will send their security data to the mentioned workspace unless overridden by a setting with more specific scope
	Scope string `pulumi:"scope"`
	// The full Azure ID of the workspace to save the data in
	WorkspaceId string `pulumi:"workspaceId"`
	// Name of the security setting
	WorkspaceSettingName *string `pulumi:"workspaceSettingName"`
}

// The set of arguments for constructing a WorkspaceSetting resource.
type WorkspaceSettingArgs struct {
	// All the VMs in this scope will send their security data to the mentioned workspace unless overridden by a setting with more specific scope
	Scope pulumi.StringInput
	// The full Azure ID of the workspace to save the data in
	WorkspaceId pulumi.StringInput
	// Name of the security setting
	WorkspaceSettingName pulumi.StringPtrInput
}

func (WorkspaceSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSettingArgs)(nil)).Elem()
}

type WorkspaceSettingInput interface {
	pulumi.Input

	ToWorkspaceSettingOutput() WorkspaceSettingOutput
	ToWorkspaceSettingOutputWithContext(ctx context.Context) WorkspaceSettingOutput
}

func (*WorkspaceSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSetting)(nil))
}

func (i *WorkspaceSetting) ToWorkspaceSettingOutput() WorkspaceSettingOutput {
	return i.ToWorkspaceSettingOutputWithContext(context.Background())
}

func (i *WorkspaceSetting) ToWorkspaceSettingOutputWithContext(ctx context.Context) WorkspaceSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSettingOutput)
}

type WorkspaceSettingOutput struct {
	*pulumi.OutputState
}

func (WorkspaceSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSetting)(nil))
}

func (o WorkspaceSettingOutput) ToWorkspaceSettingOutput() WorkspaceSettingOutput {
	return o
}

func (o WorkspaceSettingOutput) ToWorkspaceSettingOutputWithContext(ctx context.Context) WorkspaceSettingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkspaceSettingOutput{})
}
