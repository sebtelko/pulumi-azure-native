// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notebooks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A NotebookProxy resource.
// API Version: 2019-10-11-preview.
type NotebookProxy struct {
	pulumi.CustomResourceState

	// The friendly string identifier of the creator of the NotebookProxy resource.
	Hostname pulumi.StringPtrOutput `pulumi:"hostname"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The unique identifier (a GUID) generated for every resource.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The type of the resource. Ex- Microsoft.Storage/storageAccounts or Microsoft.Notebooks/notebookProxies.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNotebookProxy registers a new resource with the given unique name, arguments, and options.
func NewNotebookProxy(ctx *pulumi.Context,
	name string, args *NotebookProxyArgs, opts ...pulumi.ResourceOption) (*NotebookProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:notebooks:NotebookProxy"),
		},
		{
			Type: pulumi.String("azure-native:notebooks/v20191011preview:NotebookProxy"),
		},
		{
			Type: pulumi.String("azure-nextgen:notebooks/v20191011preview:NotebookProxy"),
		},
	})
	opts = append(opts, aliases)
	var resource NotebookProxy
	err := ctx.RegisterResource("azure-native:notebooks:NotebookProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebookProxy gets an existing NotebookProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebookProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookProxyState, opts ...pulumi.ResourceOption) (*NotebookProxy, error) {
	var resource NotebookProxy
	err := ctx.ReadResource("azure-native:notebooks:NotebookProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebookProxy resources.
type notebookProxyState struct {
	// The friendly string identifier of the creator of the NotebookProxy resource.
	Hostname *string `pulumi:"hostname"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The unique identifier (a GUID) generated for every resource.
	ResourceId *string `pulumi:"resourceId"`
	// The type of the resource. Ex- Microsoft.Storage/storageAccounts or Microsoft.Notebooks/notebookProxies.
	Type *string `pulumi:"type"`
}

type NotebookProxyState struct {
	// The friendly string identifier of the creator of the NotebookProxy resource.
	Hostname pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The unique identifier (a GUID) generated for every resource.
	ResourceId pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Storage/storageAccounts or Microsoft.Notebooks/notebookProxies.
	Type pulumi.StringPtrInput
}

func (NotebookProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookProxyState)(nil)).Elem()
}

type notebookProxyArgs struct {
	// The friendly string identifier of the creator of the NotebookProxy resource.
	Hostname *string `pulumi:"hostname"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName *string `pulumi:"resourceName"`
}

// The set of arguments for constructing a NotebookProxy resource.
type NotebookProxyArgs struct {
	// The friendly string identifier of the creator of the NotebookProxy resource.
	Hostname pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the resource.
	ResourceName pulumi.StringPtrInput
}

func (NotebookProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookProxyArgs)(nil)).Elem()
}

type NotebookProxyInput interface {
	pulumi.Input

	ToNotebookProxyOutput() NotebookProxyOutput
	ToNotebookProxyOutputWithContext(ctx context.Context) NotebookProxyOutput
}

func (*NotebookProxy) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookProxy)(nil))
}

func (i *NotebookProxy) ToNotebookProxyOutput() NotebookProxyOutput {
	return i.ToNotebookProxyOutputWithContext(context.Background())
}

func (i *NotebookProxy) ToNotebookProxyOutputWithContext(ctx context.Context) NotebookProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookProxyOutput)
}

type NotebookProxyOutput struct {
	*pulumi.OutputState
}

func (NotebookProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookProxy)(nil))
}

func (o NotebookProxyOutput) ToNotebookProxyOutput() NotebookProxyOutput {
	return o
}

func (o NotebookProxyOutput) ToNotebookProxyOutputWithContext(ctx context.Context) NotebookProxyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NotebookProxyOutput{})
}
