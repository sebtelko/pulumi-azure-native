// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150408

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB MongoDB collection.
type DatabaseAccountMongoDBCollection struct {
	pulumi.CustomResourceState

	// List of index keys
	Indexes MongoIndexResponseArrayOutput `pulumi:"indexes"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// A key-value pair of shard keys to be applied for the request.
	ShardKey pulumi.StringMapOutput `pulumi:"shardKey"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseAccountMongoDBCollection registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccountMongoDBCollection(ctx *pulumi.Context,
	name string, args *DatabaseAccountMongoDBCollectionArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150408:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150401:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20151106:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160319:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160331:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20190801:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20191212:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200301:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200401:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200601preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200901:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210115:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210301preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210315:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210401preview:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210415:DatabaseAccountMongoDBCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountMongoDBCollection
	err := ctx.RegisterResource("azure-native:documentdb/v20150408:DatabaseAccountMongoDBCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccountMongoDBCollection gets an existing DatabaseAccountMongoDBCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccountMongoDBCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountMongoDBCollectionState, opts ...pulumi.ResourceOption) (*DatabaseAccountMongoDBCollection, error) {
	var resource DatabaseAccountMongoDBCollection
	err := ctx.ReadResource("azure-native:documentdb/v20150408:DatabaseAccountMongoDBCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccountMongoDBCollection resources.
type databaseAccountMongoDBCollectionState struct {
	// List of index keys
	Indexes []MongoIndexResponse `pulumi:"indexes"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the database account.
	Name *string `pulumi:"name"`
	// A key-value pair of shard keys to be applied for the request.
	ShardKey map[string]string `pulumi:"shardKey"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type DatabaseAccountMongoDBCollectionState struct {
	// List of index keys
	Indexes MongoIndexResponseArrayInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the database account.
	Name pulumi.StringPtrInput
	// A key-value pair of shard keys to be applied for the request.
	ShardKey pulumi.StringMapInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (DatabaseAccountMongoDBCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBCollectionState)(nil)).Elem()
}

type databaseAccountMongoDBCollectionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB collection name.
	CollectionName *string `pulumi:"collectionName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a MongoDB collection
	Resource MongoDBCollectionResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DatabaseAccountMongoDBCollection resource.
type DatabaseAccountMongoDBCollectionArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB collection name.
	CollectionName pulumi.StringPtrInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a MongoDB collection
	Resource MongoDBCollectionResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountMongoDBCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountMongoDBCollectionArgs)(nil)).Elem()
}

type DatabaseAccountMongoDBCollectionInput interface {
	pulumi.Input

	ToDatabaseAccountMongoDBCollectionOutput() DatabaseAccountMongoDBCollectionOutput
	ToDatabaseAccountMongoDBCollectionOutputWithContext(ctx context.Context) DatabaseAccountMongoDBCollectionOutput
}

func (*DatabaseAccountMongoDBCollection) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountMongoDBCollection)(nil))
}

func (i *DatabaseAccountMongoDBCollection) ToDatabaseAccountMongoDBCollectionOutput() DatabaseAccountMongoDBCollectionOutput {
	return i.ToDatabaseAccountMongoDBCollectionOutputWithContext(context.Background())
}

func (i *DatabaseAccountMongoDBCollection) ToDatabaseAccountMongoDBCollectionOutputWithContext(ctx context.Context) DatabaseAccountMongoDBCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountMongoDBCollectionOutput)
}

type DatabaseAccountMongoDBCollectionOutput struct {
	*pulumi.OutputState
}

func (DatabaseAccountMongoDBCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountMongoDBCollection)(nil))
}

func (o DatabaseAccountMongoDBCollectionOutput) ToDatabaseAccountMongoDBCollectionOutput() DatabaseAccountMongoDBCollectionOutput {
	return o
}

func (o DatabaseAccountMongoDBCollectionOutput) ToDatabaseAccountMongoDBCollectionOutputWithContext(ctx context.Context) DatabaseAccountMongoDBCollectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountMongoDBCollectionOutput{})
}
