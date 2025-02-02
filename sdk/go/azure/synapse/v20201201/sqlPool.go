// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A SQL Analytics pool
type SqlPool struct {
	pulumi.CustomResourceState

	// Collation mode
	Collation pulumi.StringPtrOutput `pulumi:"collation"`
	// What is this?
	CreateMode pulumi.StringPtrOutput `pulumi:"createMode"`
	// Date the SQL pool was created
	CreationDate pulumi.StringPtrOutput `pulumi:"creationDate"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Maximum size in bytes
	MaxSizeBytes pulumi.Float64PtrOutput `pulumi:"maxSizeBytes"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource state
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Backup database to restore from
	RecoverableDatabaseId pulumi.StringPtrOutput `pulumi:"recoverableDatabaseId"`
	// Snapshot time to restore
	RestorePointInTime pulumi.StringPtrOutput `pulumi:"restorePointInTime"`
	// SQL pool SKU
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Source database to create from
	SourceDatabaseId pulumi.StringPtrOutput `pulumi:"sourceDatabaseId"`
	// Resource status
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The storage account type used to store backups for this sql pool.
	StorageAccountType pulumi.StringPtrOutput `pulumi:"storageAccountType"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlPool registers a new resource with the given unique name, arguments, and options.
func NewSqlPool(ctx *pulumi.Context,
	name string, args *SqlPoolArgs, opts ...pulumi.ResourceOption) (*SqlPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:synapse/v20201201:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse:SqlPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:SqlPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20190601preview:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20200401preview:SqlPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20200401preview:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:SqlPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210301:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:SqlPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210401preview:SqlPool"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlPool
	err := ctx.RegisterResource("azure-native:synapse/v20201201:SqlPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlPool gets an existing SqlPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlPoolState, opts ...pulumi.ResourceOption) (*SqlPool, error) {
	var resource SqlPool
	err := ctx.ReadResource("azure-native:synapse/v20201201:SqlPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlPool resources.
type sqlPoolState struct {
	// Collation mode
	Collation *string `pulumi:"collation"`
	// What is this?
	CreateMode *string `pulumi:"createMode"`
	// Date the SQL pool was created
	CreationDate *string `pulumi:"creationDate"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Maximum size in bytes
	MaxSizeBytes *float64 `pulumi:"maxSizeBytes"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Resource state
	ProvisioningState *string `pulumi:"provisioningState"`
	// Backup database to restore from
	RecoverableDatabaseId *string `pulumi:"recoverableDatabaseId"`
	// Snapshot time to restore
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// SQL pool SKU
	Sku *SkuResponse `pulumi:"sku"`
	// Source database to create from
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// Resource status
	Status *string `pulumi:"status"`
	// The storage account type used to store backups for this sql pool.
	StorageAccountType *string `pulumi:"storageAccountType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type SqlPoolState struct {
	// Collation mode
	Collation pulumi.StringPtrInput
	// What is this?
	CreateMode pulumi.StringPtrInput
	// Date the SQL pool was created
	CreationDate pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Maximum size in bytes
	MaxSizeBytes pulumi.Float64PtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Resource state
	ProvisioningState pulumi.StringPtrInput
	// Backup database to restore from
	RecoverableDatabaseId pulumi.StringPtrInput
	// Snapshot time to restore
	RestorePointInTime pulumi.StringPtrInput
	// SQL pool SKU
	Sku SkuResponsePtrInput
	// Source database to create from
	SourceDatabaseId pulumi.StringPtrInput
	// Resource status
	Status pulumi.StringPtrInput
	// The storage account type used to store backups for this sql pool.
	StorageAccountType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (SqlPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolState)(nil)).Elem()
}

type sqlPoolArgs struct {
	// Collation mode
	Collation *string `pulumi:"collation"`
	// What is this?
	CreateMode *string `pulumi:"createMode"`
	// Date the SQL pool was created
	CreationDate *string `pulumi:"creationDate"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Maximum size in bytes
	MaxSizeBytes *float64 `pulumi:"maxSizeBytes"`
	// Resource state
	ProvisioningState *string `pulumi:"provisioningState"`
	// Backup database to restore from
	RecoverableDatabaseId *string `pulumi:"recoverableDatabaseId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Snapshot time to restore
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// SQL pool SKU
	Sku *Sku `pulumi:"sku"`
	// Source database to create from
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// SQL pool name
	SqlPoolName *string `pulumi:"sqlPoolName"`
	// Resource status
	Status *string `pulumi:"status"`
	// The storage account type used to store backups for this sql pool.
	StorageAccountType *string `pulumi:"storageAccountType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a SqlPool resource.
type SqlPoolArgs struct {
	// Collation mode
	Collation pulumi.StringPtrInput
	// What is this?
	CreateMode pulumi.StringPtrInput
	// Date the SQL pool was created
	CreationDate pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Maximum size in bytes
	MaxSizeBytes pulumi.Float64PtrInput
	// Resource state
	ProvisioningState pulumi.StringPtrInput
	// Backup database to restore from
	RecoverableDatabaseId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Snapshot time to restore
	RestorePointInTime pulumi.StringPtrInput
	// SQL pool SKU
	Sku SkuPtrInput
	// Source database to create from
	SourceDatabaseId pulumi.StringPtrInput
	// SQL pool name
	SqlPoolName pulumi.StringPtrInput
	// Resource status
	Status pulumi.StringPtrInput
	// The storage account type used to store backups for this sql pool.
	StorageAccountType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the workspace
	WorkspaceName pulumi.StringInput
}

func (SqlPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolArgs)(nil)).Elem()
}

type SqlPoolInput interface {
	pulumi.Input

	ToSqlPoolOutput() SqlPoolOutput
	ToSqlPoolOutputWithContext(ctx context.Context) SqlPoolOutput
}

func (*SqlPool) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPool)(nil))
}

func (i *SqlPool) ToSqlPoolOutput() SqlPoolOutput {
	return i.ToSqlPoolOutputWithContext(context.Background())
}

func (i *SqlPool) ToSqlPoolOutputWithContext(ctx context.Context) SqlPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolOutput)
}

type SqlPoolOutput struct {
	*pulumi.OutputState
}

func (SqlPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPool)(nil))
}

func (o SqlPoolOutput) ToSqlPoolOutput() SqlPoolOutput {
	return o
}

func (o SqlPoolOutput) ToSqlPoolOutputWithContext(ctx context.Context) SqlPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SqlPoolOutput{})
}
