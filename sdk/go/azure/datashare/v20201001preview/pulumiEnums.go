// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Kind of data set.
type DataSetKind pulumi.String

const (
	DataSetKindBlob                         = DataSetKind("Blob")
	DataSetKindContainer                    = DataSetKind("Container")
	DataSetKindBlobFolder                   = DataSetKind("BlobFolder")
	DataSetKindAdlsGen2FileSystem           = DataSetKind("AdlsGen2FileSystem")
	DataSetKindAdlsGen2Folder               = DataSetKind("AdlsGen2Folder")
	DataSetKindAdlsGen2File                 = DataSetKind("AdlsGen2File")
	DataSetKindAdlsGen1Folder               = DataSetKind("AdlsGen1Folder")
	DataSetKindAdlsGen1File                 = DataSetKind("AdlsGen1File")
	DataSetKindAdlsGen2StorageAccount       = DataSetKind("AdlsGen2StorageAccount")
	DataSetKindKustoCluster                 = DataSetKind("KustoCluster")
	DataSetKindKustoDatabase                = DataSetKind("KustoDatabase")
	DataSetKindSqlDBTable                   = DataSetKind("SqlDBTable")
	DataSetKindSqlDWTable                   = DataSetKind("SqlDWTable")
	DataSetKindSynapseWorkspaceSqlPoolTable = DataSetKind("SynapseWorkspaceSqlPoolTable")
	DataSetKindBlobStorageAccount           = DataSetKind("BlobStorageAccount")
)

func (DataSetKind) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e DataSetKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSetKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSetKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataSetKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Kind of data set mapping.
type DataSetMappingKind pulumi.String

const (
	DataSetMappingKindBlob                         = DataSetMappingKind("Blob")
	DataSetMappingKindContainer                    = DataSetMappingKind("Container")
	DataSetMappingKindBlobFolder                   = DataSetMappingKind("BlobFolder")
	DataSetMappingKindAdlsGen2FileSystem           = DataSetMappingKind("AdlsGen2FileSystem")
	DataSetMappingKindAdlsGen2Folder               = DataSetMappingKind("AdlsGen2Folder")
	DataSetMappingKindAdlsGen2File                 = DataSetMappingKind("AdlsGen2File")
	DataSetMappingKindAdlsGen2StorageAccount       = DataSetMappingKind("AdlsGen2StorageAccount")
	DataSetMappingKindKustoCluster                 = DataSetMappingKind("KustoCluster")
	DataSetMappingKindKustoDatabase                = DataSetMappingKind("KustoDatabase")
	DataSetMappingKindSqlDBTable                   = DataSetMappingKind("SqlDBTable")
	DataSetMappingKindSqlDWTable                   = DataSetMappingKind("SqlDWTable")
	DataSetMappingKindSynapseWorkspaceSqlPoolTable = DataSetMappingKind("SynapseWorkspaceSqlPoolTable")
	DataSetMappingKindBlobStorageAccount           = DataSetMappingKind("BlobStorageAccount")
)

func (DataSetMappingKind) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e DataSetMappingKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSetMappingKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSetMappingKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataSetMappingKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Share kind.
type ShareKind pulumi.String

const (
	ShareKindCopyBased = ShareKind("CopyBased")
	ShareKindInPlace   = ShareKind("InPlace")
)

func (ShareKind) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ShareKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ShareKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Kind of synchronization setting.
type SynchronizationSettingKind pulumi.String

const (
	SynchronizationSettingKindScheduleBased = SynchronizationSettingKind("ScheduleBased")
)

func (SynchronizationSettingKind) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e SynchronizationSettingKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SynchronizationSettingKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SynchronizationSettingKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SynchronizationSettingKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Kind of synchronization on trigger.
type TriggerKind pulumi.String

const (
	TriggerKindScheduleBased = TriggerKind("ScheduleBased")
)

func (TriggerKind) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e TriggerKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggerKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Identity Type
type Type pulumi.String

const (
	TypeSystemAssigned = Type("SystemAssigned")
)

func (Type) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Type) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Type) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Type) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Type) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
