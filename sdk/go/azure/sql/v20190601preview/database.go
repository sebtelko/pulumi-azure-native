// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A database resource.
type Database struct {
	pulumi.CustomResourceState

	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
	AutoPauseDelay pulumi.IntPtrOutput `pulumi:"autoPauseDelay"`
	// Collation of the metadata catalog.
	CatalogCollation pulumi.StringPtrOutput `pulumi:"catalogCollation"`
	// The collation of the database.
	Collation pulumi.StringPtrOutput `pulumi:"collation"`
	// The creation date of the database (ISO8601 format).
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The current service level objective name of the database.
	CurrentServiceObjectiveName pulumi.StringOutput `pulumi:"currentServiceObjectiveName"`
	// The name and tier of the SKU.
	CurrentSku SkuResponseOutput `pulumi:"currentSku"`
	// The ID of the database.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// The default secondary region for this database.
	DefaultSecondaryLocation pulumi.StringOutput `pulumi:"defaultSecondaryLocation"`
	// This records the earliest start date and time that restore is available for this database (ISO8601 format).
	EarliestRestoreDate pulumi.StringOutput `pulumi:"earliestRestoreDate"`
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId pulumi.StringPtrOutput `pulumi:"elasticPoolId"`
	// Failover Group resource identifier that this database belongs to.
	FailoverGroupId pulumi.StringOutput `pulumi:"failoverGroupId"`
	// Kind of database. This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource that manages the database.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// The max log size for this database.
	MaxLogSizeBytes pulumi.Float64Output `pulumi:"maxLogSizeBytes"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes pulumi.Float64PtrOutput `pulumi:"maxSizeBytes"`
	// Minimal capacity that database will always have allocated, if not paused
	MinCapacity pulumi.Float64PtrOutput `pulumi:"minCapacity"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The date when database was paused by user configuration or action(ISO8601 format). Null if the database is ready.
	PausedDate pulumi.StringOutput `pulumi:"pausedDate"`
	// The number of readonly secondary replicas associated with the database.
	ReadReplicaCount pulumi.IntPtrOutput `pulumi:"readReplicaCount"`
	// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region.
	ReadScale pulumi.StringPtrOutput `pulumi:"readScale"`
	// The requested service level objective name of the database.
	RequestedServiceObjectiveName pulumi.StringOutput `pulumi:"requestedServiceObjectiveName"`
	// The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
	ResumedDate pulumi.StringOutput `pulumi:"resumedDate"`
	// The database SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The status of the database.
	Status pulumi.StringOutput `pulumi:"status"`
	// The storage account type used to store backups for this database.
	StorageAccountType pulumi.StringPtrOutput `pulumi:"storageAccountType"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
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
			Type: pulumi.String("azure-nextgen:sql/v20190601preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20140401:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20170301preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20171001preview:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20171001preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:Database"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:Database"),
		},
	})
	opts = append(opts, aliases)
	var resource Database
	err := ctx.RegisterResource("azure-native:sql/v20190601preview:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure-native:sql/v20190601preview:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
	AutoPauseDelay *int `pulumi:"autoPauseDelay"`
	// Collation of the metadata catalog.
	CatalogCollation *string `pulumi:"catalogCollation"`
	// The collation of the database.
	Collation *string `pulumi:"collation"`
	// The creation date of the database (ISO8601 format).
	CreationDate *string `pulumi:"creationDate"`
	// The current service level objective name of the database.
	CurrentServiceObjectiveName *string `pulumi:"currentServiceObjectiveName"`
	// The name and tier of the SKU.
	CurrentSku *SkuResponse `pulumi:"currentSku"`
	// The ID of the database.
	DatabaseId *string `pulumi:"databaseId"`
	// The default secondary region for this database.
	DefaultSecondaryLocation *string `pulumi:"defaultSecondaryLocation"`
	// This records the earliest start date and time that restore is available for this database (ISO8601 format).
	EarliestRestoreDate *string `pulumi:"earliestRestoreDate"`
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId *string `pulumi:"elasticPoolId"`
	// Failover Group resource identifier that this database belongs to.
	FailoverGroupId *string `pulumi:"failoverGroupId"`
	// Kind of database. This is metadata used for the Azure portal experience.
	Kind *string `pulumi:"kind"`
	// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType *string `pulumi:"licenseType"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource that manages the database.
	ManagedBy *string `pulumi:"managedBy"`
	// The max log size for this database.
	MaxLogSizeBytes *float64 `pulumi:"maxLogSizeBytes"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes *float64 `pulumi:"maxSizeBytes"`
	// Minimal capacity that database will always have allocated, if not paused
	MinCapacity *float64 `pulumi:"minCapacity"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The date when database was paused by user configuration or action(ISO8601 format). Null if the database is ready.
	PausedDate *string `pulumi:"pausedDate"`
	// The number of readonly secondary replicas associated with the database.
	ReadReplicaCount *int `pulumi:"readReplicaCount"`
	// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region.
	ReadScale *string `pulumi:"readScale"`
	// The requested service level objective name of the database.
	RequestedServiceObjectiveName *string `pulumi:"requestedServiceObjectiveName"`
	// The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
	ResumedDate *string `pulumi:"resumedDate"`
	// The database SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Sku *SkuResponse `pulumi:"sku"`
	// The status of the database.
	Status *string `pulumi:"status"`
	// The storage account type used to store backups for this database.
	StorageAccountType *string `pulumi:"storageAccountType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

type DatabaseState struct {
	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
	AutoPauseDelay pulumi.IntPtrInput
	// Collation of the metadata catalog.
	CatalogCollation pulumi.StringPtrInput
	// The collation of the database.
	Collation pulumi.StringPtrInput
	// The creation date of the database (ISO8601 format).
	CreationDate pulumi.StringPtrInput
	// The current service level objective name of the database.
	CurrentServiceObjectiveName pulumi.StringPtrInput
	// The name and tier of the SKU.
	CurrentSku SkuResponsePtrInput
	// The ID of the database.
	DatabaseId pulumi.StringPtrInput
	// The default secondary region for this database.
	DefaultSecondaryLocation pulumi.StringPtrInput
	// This records the earliest start date and time that restore is available for this database (ISO8601 format).
	EarliestRestoreDate pulumi.StringPtrInput
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId pulumi.StringPtrInput
	// Failover Group resource identifier that this database belongs to.
	FailoverGroupId pulumi.StringPtrInput
	// Kind of database. This is metadata used for the Azure portal experience.
	Kind pulumi.StringPtrInput
	// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource that manages the database.
	ManagedBy pulumi.StringPtrInput
	// The max log size for this database.
	MaxLogSizeBytes pulumi.Float64PtrInput
	// The max size of the database expressed in bytes.
	MaxSizeBytes pulumi.Float64PtrInput
	// Minimal capacity that database will always have allocated, if not paused
	MinCapacity pulumi.Float64PtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The date when database was paused by user configuration or action(ISO8601 format). Null if the database is ready.
	PausedDate pulumi.StringPtrInput
	// The number of readonly secondary replicas associated with the database.
	ReadReplicaCount pulumi.IntPtrInput
	// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region.
	ReadScale pulumi.StringPtrInput
	// The requested service level objective name of the database.
	RequestedServiceObjectiveName pulumi.StringPtrInput
	// The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
	ResumedDate pulumi.StringPtrInput
	// The database SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Sku SkuResponsePtrInput
	// The status of the database.
	Status pulumi.StringPtrInput
	// The storage account type used to store backups for this database.
	StorageAccountType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
	AutoPauseDelay *int `pulumi:"autoPauseDelay"`
	// Collation of the metadata catalog.
	CatalogCollation *string `pulumi:"catalogCollation"`
	// The collation of the database.
	Collation *string `pulumi:"collation"`
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode *string `pulumi:"createMode"`
	// The name of the database.
	DatabaseName *string `pulumi:"databaseName"`
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId *string `pulumi:"elasticPoolId"`
	// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType *string `pulumi:"licenseType"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The resource identifier of the long term retention backup associated with create operation of this database.
	LongTermRetentionBackupResourceId *string `pulumi:"longTermRetentionBackupResourceId"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes *float64 `pulumi:"maxSizeBytes"`
	// Minimal capacity that database will always have allocated, if not paused
	MinCapacity *float64 `pulumi:"minCapacity"`
	// The number of readonly secondary replicas associated with the database.
	ReadReplicaCount *int `pulumi:"readReplicaCount"`
	// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region.
	ReadScale *string `pulumi:"readScale"`
	// The resource identifier of the recoverable database associated with create operation of this database.
	RecoverableDatabaseId *string `pulumi:"recoverableDatabaseId"`
	// The resource identifier of the recovery point associated with create operation of this database.
	RecoveryServicesRecoveryPointId *string `pulumi:"recoveryServicesRecoveryPointId"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource identifier of the restorable dropped database associated with create operation of this database.
	RestorableDroppedDatabaseId *string `pulumi:"restorableDroppedDatabaseId"`
	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The name of the sample schema to apply when creating this database.
	SampleName *string `pulumi:"sampleName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The database SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Sku *Sku `pulumi:"sku"`
	// Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate *string `pulumi:"sourceDatabaseDeletionDate"`
	// The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// The storage account type used to store backups for this database.
	StorageAccountType *string `pulumi:"storageAccountType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
	AutoPauseDelay pulumi.IntPtrInput
	// Collation of the metadata catalog.
	CatalogCollation pulumi.StringPtrInput
	// The collation of the database.
	Collation pulumi.StringPtrInput
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode pulumi.StringPtrInput
	// The name of the database.
	DatabaseName pulumi.StringPtrInput
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId pulumi.StringPtrInput
	// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The resource identifier of the long term retention backup associated with create operation of this database.
	LongTermRetentionBackupResourceId pulumi.StringPtrInput
	// The max size of the database expressed in bytes.
	MaxSizeBytes pulumi.Float64PtrInput
	// Minimal capacity that database will always have allocated, if not paused
	MinCapacity pulumi.Float64PtrInput
	// The number of readonly secondary replicas associated with the database.
	ReadReplicaCount pulumi.IntPtrInput
	// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region.
	ReadScale pulumi.StringPtrInput
	// The resource identifier of the recoverable database associated with create operation of this database.
	RecoverableDatabaseId pulumi.StringPtrInput
	// The resource identifier of the recovery point associated with create operation of this database.
	RecoveryServicesRecoveryPointId pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The resource identifier of the restorable dropped database associated with create operation of this database.
	RestorableDroppedDatabaseId pulumi.StringPtrInput
	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime pulumi.StringPtrInput
	// The name of the sample schema to apply when creating this database.
	SampleName pulumi.StringPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The database SKU.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Sku SkuPtrInput
	// Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate pulumi.StringPtrInput
	// The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId pulumi.StringPtrInput
	// The storage account type used to store backups for this database.
	StorageAccountType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((*Database)(nil))
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

type DatabaseOutput struct {
	*pulumi.OutputState
}

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Database)(nil))
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseOutput{})
}
