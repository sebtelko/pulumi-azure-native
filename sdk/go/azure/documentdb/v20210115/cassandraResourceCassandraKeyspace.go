// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210115

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB Cassandra keyspace.
type CassandraResourceCassandraKeyspace struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                     `pulumi:"name"`
	Options  CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource CassandraKeyspaceGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCassandraResourceCassandraKeyspace registers a new resource with the given unique name, arguments, and options.
func NewCassandraResourceCassandraKeyspace(ctx *pulumi.Context,
	name string, args *CassandraResourceCassandraKeyspaceArgs, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraKeyspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210115:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150401:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150408:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20151106:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160319:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160331:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20190801:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20191212:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200301:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200401:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200601preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200901:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210301preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210315:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210401preview:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:CassandraResourceCassandraKeyspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210415:CassandraResourceCassandraKeyspace"),
		},
	})
	opts = append(opts, aliases)
	var resource CassandraResourceCassandraKeyspace
	err := ctx.RegisterResource("azure-native:documentdb/v20210115:CassandraResourceCassandraKeyspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCassandraResourceCassandraKeyspace gets an existing CassandraResourceCassandraKeyspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCassandraResourceCassandraKeyspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraResourceCassandraKeyspaceState, opts ...pulumi.ResourceOption) (*CassandraResourceCassandraKeyspace, error) {
	var resource CassandraResourceCassandraKeyspace
	err := ctx.ReadResource("azure-native:documentdb/v20210115:CassandraResourceCassandraKeyspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CassandraResourceCassandraKeyspace resources.
type cassandraResourceCassandraKeyspaceState struct {
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     *string                                         `pulumi:"name"`
	Options  *CassandraKeyspaceGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *CassandraKeyspaceGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type CassandraResourceCassandraKeyspaceState struct {
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the ARM resource.
	Name     pulumi.StringPtrInput
	Options  CassandraKeyspaceGetPropertiesResponseOptionsPtrInput
	Resource CassandraKeyspaceGetPropertiesResponseResourcePtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (CassandraResourceCassandraKeyspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraKeyspaceState)(nil)).Elem()
}

type cassandraResourceCassandraKeyspaceArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB keyspace name.
	KeyspaceName *string `pulumi:"keyspaceName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options *CreateUpdateOptions `pulumi:"options"`
	// The standard JSON format of a Cassandra keyspace
	Resource CassandraKeyspaceResource `pulumi:"resource"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CassandraResourceCassandraKeyspace resource.
type CassandraResourceCassandraKeyspaceArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB keyspace name.
	KeyspaceName pulumi.StringPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options CreateUpdateOptionsPtrInput
	// The standard JSON format of a Cassandra keyspace
	Resource CassandraKeyspaceResourceInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (CassandraResourceCassandraKeyspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraResourceCassandraKeyspaceArgs)(nil)).Elem()
}

type CassandraResourceCassandraKeyspaceInput interface {
	pulumi.Input

	ToCassandraResourceCassandraKeyspaceOutput() CassandraResourceCassandraKeyspaceOutput
	ToCassandraResourceCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraResourceCassandraKeyspaceOutput
}

func (*CassandraResourceCassandraKeyspace) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraResourceCassandraKeyspace)(nil))
}

func (i *CassandraResourceCassandraKeyspace) ToCassandraResourceCassandraKeyspaceOutput() CassandraResourceCassandraKeyspaceOutput {
	return i.ToCassandraResourceCassandraKeyspaceOutputWithContext(context.Background())
}

func (i *CassandraResourceCassandraKeyspace) ToCassandraResourceCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraResourceCassandraKeyspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraResourceCassandraKeyspaceOutput)
}

type CassandraResourceCassandraKeyspaceOutput struct {
	*pulumi.OutputState
}

func (CassandraResourceCassandraKeyspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraResourceCassandraKeyspace)(nil))
}

func (o CassandraResourceCassandraKeyspaceOutput) ToCassandraResourceCassandraKeyspaceOutput() CassandraResourceCassandraKeyspaceOutput {
	return o
}

func (o CassandraResourceCassandraKeyspaceOutput) ToCassandraResourceCassandraKeyspaceOutputWithContext(ctx context.Context) CassandraResourceCassandraKeyspaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CassandraResourceCassandraKeyspaceOutput{})
}
