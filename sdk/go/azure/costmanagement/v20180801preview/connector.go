// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Connector model definition
type Connector struct {
	pulumi.CustomResourceState

	// Collection information
	Collection ConnectorCollectionInfoResponseOutput `pulumi:"collection"`
	// Connector definition creation datetime
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// Credentials authentication key (eg AWS ARN)
	CredentialsKey pulumi.StringPtrOutput `pulumi:"credentialsKey"`
	// Connector DisplayName (defaults to Name)
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Connector kind (eg aws)
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Connector location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Connector last modified datetime
	ModifiedOn pulumi.StringOutput `pulumi:"modifiedOn"`
	// Connector name
	Name pulumi.StringOutput `pulumi:"name"`
	// Connector providerAccountId (determined from credentials)
	ProviderAccountId pulumi.StringOutput `pulumi:"providerAccountId"`
	// Identifying source report. (For AWS this is a CUR report name, defined with Daily and with Resources)
	ReportId pulumi.StringPtrOutput `pulumi:"reportId"`
	// Connector status
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Connector type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnector registers a new resource with the given unique name, arguments, and options.
func NewConnector(ctx *pulumi.Context,
	name string, args *ConnectorArgs, opts ...pulumi.ResourceOption) (*Connector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:costmanagement/v20180801preview:Connector"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement:Connector"),
		},
		{
			Type: pulumi.String("azure-nextgen:costmanagement:Connector"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20190301preview:Connector"),
		},
		{
			Type: pulumi.String("azure-nextgen:costmanagement/v20190301preview:Connector"),
		},
	})
	opts = append(opts, aliases)
	var resource Connector
	err := ctx.RegisterResource("azure-native:costmanagement/v20180801preview:Connector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnector gets an existing Connector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorState, opts ...pulumi.ResourceOption) (*Connector, error) {
	var resource Connector
	err := ctx.ReadResource("azure-native:costmanagement/v20180801preview:Connector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connector resources.
type connectorState struct {
	// Collection information
	Collection *ConnectorCollectionInfoResponse `pulumi:"collection"`
	// Connector definition creation datetime
	CreatedOn *string `pulumi:"createdOn"`
	// Credentials authentication key (eg AWS ARN)
	CredentialsKey *string `pulumi:"credentialsKey"`
	// Connector DisplayName (defaults to Name)
	DisplayName *string `pulumi:"displayName"`
	// Connector kind (eg aws)
	Kind *string `pulumi:"kind"`
	// Connector location
	Location *string `pulumi:"location"`
	// Connector last modified datetime
	ModifiedOn *string `pulumi:"modifiedOn"`
	// Connector name
	Name *string `pulumi:"name"`
	// Connector providerAccountId (determined from credentials)
	ProviderAccountId *string `pulumi:"providerAccountId"`
	// Identifying source report. (For AWS this is a CUR report name, defined with Daily and with Resources)
	ReportId *string `pulumi:"reportId"`
	// Connector status
	Status *string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Connector type
	Type *string `pulumi:"type"`
}

type ConnectorState struct {
	// Collection information
	Collection ConnectorCollectionInfoResponsePtrInput
	// Connector definition creation datetime
	CreatedOn pulumi.StringPtrInput
	// Credentials authentication key (eg AWS ARN)
	CredentialsKey pulumi.StringPtrInput
	// Connector DisplayName (defaults to Name)
	DisplayName pulumi.StringPtrInput
	// Connector kind (eg aws)
	Kind pulumi.StringPtrInput
	// Connector location
	Location pulumi.StringPtrInput
	// Connector last modified datetime
	ModifiedOn pulumi.StringPtrInput
	// Connector name
	Name pulumi.StringPtrInput
	// Connector providerAccountId (determined from credentials)
	ProviderAccountId pulumi.StringPtrInput
	// Identifying source report. (For AWS this is a CUR report name, defined with Daily and with Resources)
	ReportId pulumi.StringPtrInput
	// Connector status
	Status pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Connector type
	Type pulumi.StringPtrInput
}

func (ConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorState)(nil)).Elem()
}

type connectorArgs struct {
	// Connector Name.
	ConnectorName *string `pulumi:"connectorName"`
	// Credentials authentication key (eg AWS ARN)
	CredentialsKey *string `pulumi:"credentialsKey"`
	// Credentials secret (eg AWS ExternalId)
	CredentialsSecret *string `pulumi:"credentialsSecret"`
	// Connector DisplayName (defaults to Name)
	DisplayName *string `pulumi:"displayName"`
	// Connector kind (eg aws)
	Kind *string `pulumi:"kind"`
	// Connector location
	Location *string `pulumi:"location"`
	// Identifying source report. (For AWS this is a CUR report name, defined with Daily and with Resources)
	ReportId *string `pulumi:"reportId"`
	// Azure Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Connector status
	Status *string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Connector resource.
type ConnectorArgs struct {
	// Connector Name.
	ConnectorName pulumi.StringPtrInput
	// Credentials authentication key (eg AWS ARN)
	CredentialsKey pulumi.StringPtrInput
	// Credentials secret (eg AWS ExternalId)
	CredentialsSecret pulumi.StringPtrInput
	// Connector DisplayName (defaults to Name)
	DisplayName pulumi.StringPtrInput
	// Connector kind (eg aws)
	Kind pulumi.StringPtrInput
	// Connector location
	Location pulumi.StringPtrInput
	// Identifying source report. (For AWS this is a CUR report name, defined with Daily and with Resources)
	ReportId pulumi.StringPtrInput
	// Azure Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// Connector status
	Status pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorArgs)(nil)).Elem()
}

type ConnectorInput interface {
	pulumi.Input

	ToConnectorOutput() ConnectorOutput
	ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput
}

func (*Connector) ElementType() reflect.Type {
	return reflect.TypeOf((*Connector)(nil))
}

func (i *Connector) ToConnectorOutput() ConnectorOutput {
	return i.ToConnectorOutputWithContext(context.Background())
}

func (i *Connector) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorOutput)
}

type ConnectorOutput struct {
	*pulumi.OutputState
}

func (ConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Connector)(nil))
}

func (o ConnectorOutput) ToConnectorOutput() ConnectorOutput {
	return o
}

func (o ConnectorOutput) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectorOutput{})
}
