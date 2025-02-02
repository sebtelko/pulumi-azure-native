// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies the state of the policy. If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required.
type BlobAuditingPolicyState pulumi.String

const (
	BlobAuditingPolicyStateEnabled  = BlobAuditingPolicyState("Enabled")
	BlobAuditingPolicyStateDisabled = BlobAuditingPolicyState("Disabled")
)

func (BlobAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e BlobAuditingPolicyState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobAuditingPolicyState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobAuditingPolicyState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BlobAuditingPolicyState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Collation of the metadata catalog.
type CatalogCollationType pulumi.String

const (
	CatalogCollationType_DATABASE_DEFAULT             = CatalogCollationType("DATABASE_DEFAULT")
	CatalogCollationType_SQL_Latin1_General_CP1_CI_AS = CatalogCollationType("SQL_Latin1_General_CP1_CI_AS")
)

func (CatalogCollationType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e CatalogCollationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CatalogCollationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CatalogCollationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CatalogCollationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

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
type CreateMode pulumi.String

const (
	CreateModeDefault                        = CreateMode("Default")
	CreateModeCopy                           = CreateMode("Copy")
	CreateModeSecondary                      = CreateMode("Secondary")
	CreateModeOnlineSecondary                = CreateMode("OnlineSecondary")
	CreateModePointInTimeRestore             = CreateMode("PointInTimeRestore")
	CreateModeRestore                        = CreateMode("Restore")
	CreateModeRecovery                       = CreateMode("Recovery")
	CreateModeRestoreExternalBackup          = CreateMode("RestoreExternalBackup")
	CreateModeRestoreExternalBackupSecondary = CreateMode("RestoreExternalBackupSecondary")
	CreateModeRestoreLongTermRetentionBackup = CreateMode("RestoreLongTermRetentionBackup")
)

func (CreateMode) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e CreateMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CreateMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Schedule interval type
type JobScheduleType pulumi.String

const (
	JobScheduleTypeOnce      = JobScheduleType("Once")
	JobScheduleTypeRecurring = JobScheduleType("Recurring")
)

func (JobScheduleType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e JobScheduleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobScheduleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobScheduleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobScheduleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The source of the action to execute.
type JobStepActionSource pulumi.String

const (
	JobStepActionSourceInline = JobStepActionSource("Inline")
)

func (JobStepActionSource) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e JobStepActionSource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepActionSource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepActionSource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobStepActionSource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Type of action being executed by the job step.
type JobStepActionType pulumi.String

const (
	JobStepActionTypeTSql = JobStepActionType("TSql")
)

func (JobStepActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e JobStepActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobStepActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The output destination type.
type JobStepOutputTypeEnum pulumi.String

const (
	JobStepOutputTypeEnumSqlDatabase = JobStepOutputTypeEnum("SqlDatabase")
)

func (JobStepOutputTypeEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e JobStepOutputTypeEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepOutputTypeEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepOutputTypeEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobStepOutputTypeEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Whether the target is included or excluded from the group.
type JobTargetGroupMembershipType pulumi.String

const (
	JobTargetGroupMembershipTypeInclude = JobTargetGroupMembershipType("Include")
	JobTargetGroupMembershipTypeExclude = JobTargetGroupMembershipType("Exclude")
)

func (JobTargetGroupMembershipType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e JobTargetGroupMembershipType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobTargetGroupMembershipType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobTargetGroupMembershipType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobTargetGroupMembershipType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The target type.
type JobTargetType pulumi.String

const (
	JobTargetTypeTargetGroup    = JobTargetType("TargetGroup")
	JobTargetTypeSqlDatabase    = JobTargetType("SqlDatabase")
	JobTargetTypeSqlElasticPool = JobTargetType("SqlElasticPool")
	JobTargetTypeSqlShardMap    = JobTargetType("SqlShardMap")
	JobTargetTypeSqlServer      = JobTargetType("SqlServer")
)

func (JobTargetType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e JobTargetType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobTargetType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobTargetType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobTargetType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Managed database create mode. PointInTimeRestore: Create a database by restoring a point in time backup of an existing database. SourceDatabaseName, SourceManagedInstanceName and PointInTime must be specified. RestoreExternalBackup: Create a database by restoring from external backup files. Collation, StorageContainerUri and StorageContainerSasToken must be specified. Recovery: Creates a database by restoring a geo-replicated backup. RecoverableDatabaseId must be specified as the recoverable database resource ID to restore.
type ManagedDatabaseCreateMode pulumi.String

const (
	ManagedDatabaseCreateModeDefault                        = ManagedDatabaseCreateMode("Default")
	ManagedDatabaseCreateModeRestoreExternalBackup          = ManagedDatabaseCreateMode("RestoreExternalBackup")
	ManagedDatabaseCreateModePointInTimeRestore             = ManagedDatabaseCreateMode("PointInTimeRestore")
	ManagedDatabaseCreateModeRecovery                       = ManagedDatabaseCreateMode("Recovery")
	ManagedDatabaseCreateModeRestoreLongTermRetentionBackup = ManagedDatabaseCreateMode("RestoreLongTermRetentionBackup")
)

func (ManagedDatabaseCreateMode) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ManagedDatabaseCreateMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedDatabaseCreateMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedDatabaseCreateMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedDatabaseCreateMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Type of the managed instance administrator.
type ManagedInstanceAdministratorType pulumi.String

const (
	ManagedInstanceAdministratorTypeActiveDirectory = ManagedInstanceAdministratorType("ActiveDirectory")
)

func (ManagedInstanceAdministratorType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ManagedInstanceAdministratorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedInstanceAdministratorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedInstanceAdministratorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedInstanceAdministratorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The name of the sample schema to apply when creating this database.
type SampleName pulumi.String

const (
	SampleNameAdventureWorksLT       = SampleName("AdventureWorksLT")
	SampleNameWideWorldImportersStd  = SampleName("WideWorldImportersStd")
	SampleNameWideWorldImportersFull = SampleName("WideWorldImportersFull")
)

func (SampleName) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e SampleName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SampleName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SampleName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SampleName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
type SecurityAlertPolicyState pulumi.String

const (
	SecurityAlertPolicyStateNew      = SecurityAlertPolicyState("New")
	SecurityAlertPolicyStateEnabled  = SecurityAlertPolicyState("Enabled")
	SecurityAlertPolicyStateDisabled = SecurityAlertPolicyState("Disabled")
)

func (SecurityAlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e SecurityAlertPolicyState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityAlertPolicyState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityAlertPolicyState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityAlertPolicyState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SensitivityLabelRank pulumi.String

const (
	SensitivityLabelRankNone     = SensitivityLabelRank("None")
	SensitivityLabelRankLow      = SensitivityLabelRank("Low")
	SensitivityLabelRankMedium   = SensitivityLabelRank("Medium")
	SensitivityLabelRankHigh     = SensitivityLabelRank("High")
	SensitivityLabelRankCritical = SensitivityLabelRank("Critical")
)

func (SensitivityLabelRank) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e SensitivityLabelRank) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SensitivityLabelRank) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SensitivityLabelRank) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SensitivityLabelRank) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
