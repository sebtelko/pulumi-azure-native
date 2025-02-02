// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data connector.
type DataConnector struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The data connector kind
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataConnector registers a new resource with the given unique name, arguments, and options.
func NewDataConnector(ctx *pulumi.Context,
	name string, args *DataConnectorArgs, opts ...pulumi.ResourceOption) (*DataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20200101:DataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:DataConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights:DataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource DataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20200101:DataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataConnector gets an existing DataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataConnectorState, opts ...pulumi.ResourceOption) (*DataConnector, error) {
	var resource DataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20200101:DataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataConnector resources.
type dataConnectorState struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The data connector kind
	Kind *string `pulumi:"kind"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type DataConnectorState struct {
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// The data connector kind
	Kind pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (DataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorState)(nil)).Elem()
}

type dataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The data connector kind
	Kind string `pulumi:"kind"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DataConnector resource.
type DataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// Etag of the azure resource
	Etag pulumi.StringPtrInput
	// The data connector kind
	Kind pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (DataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataConnectorArgs)(nil)).Elem()
}

type DataConnectorInput interface {
	pulumi.Input

	ToDataConnectorOutput() DataConnectorOutput
	ToDataConnectorOutputWithContext(ctx context.Context) DataConnectorOutput
}

func (*DataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnector)(nil))
}

func (i *DataConnector) ToDataConnectorOutput() DataConnectorOutput {
	return i.ToDataConnectorOutputWithContext(context.Background())
}

func (i *DataConnector) ToDataConnectorOutputWithContext(ctx context.Context) DataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorOutput)
}

type DataConnectorOutput struct {
	*pulumi.OutputState
}

func (DataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnector)(nil))
}

func (o DataConnectorOutput) ToDataConnectorOutput() DataConnectorOutput {
	return o
}

func (o DataConnectorOutput) ToDataConnectorOutputWithContext(ctx context.Context) DataConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataConnectorOutput{})
}
