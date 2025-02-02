// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The backup management type associated with the backup item.
type BackupManagementType pulumi.String

const (
	BackupManagementTypeInvalid           = BackupManagementType("Invalid")
	BackupManagementTypeAzureIaasVM       = BackupManagementType("AzureIaasVM")
	BackupManagementTypeMAB               = BackupManagementType("MAB")
	BackupManagementTypeDPM               = BackupManagementType("DPM")
	BackupManagementTypeAzureBackupServer = BackupManagementType("AzureBackupServer")
	BackupManagementTypeAzureSql          = BackupManagementType("AzureSql")
)

func (BackupManagementType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e BackupManagementType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupManagementType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupManagementType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BackupManagementType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The workload type for this item.
type DataSourceType pulumi.String

const (
	DataSourceTypeInvalid    = DataSourceType("Invalid")
	DataSourceTypeVM         = DataSourceType("VM")
	DataSourceTypeFileFolder = DataSourceType("FileFolder")
	DataSourceTypeAzureSqlDb = DataSourceType("AzureSqlDb")
	DataSourceTypeSQLDB      = DataSourceType("SQLDB")
	DataSourceTypeExchange   = DataSourceType("Exchange")
	DataSourceTypeSharepoint = DataSourceType("Sharepoint")
	DataSourceTypeDPMUnknown = DataSourceType("DPMUnknown")
)

func (DataSourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e DataSourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataSourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DayOfWeek pulumi.String

const (
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
)

func (DayOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e DayOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DayOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// DayOfWeekArrayInput is an input type that accepts DayOfWeekArray and DayOfWeekArrayOutput values.
// You can construct a concrete instance of `DayOfWeekArrayInput` via:
//
//          DayOfWeekArray{ DayOfWeekArgs{...} }
type DayOfWeekArrayInput interface {
	pulumi.Input

	ToDayOfWeekArrayOutput() DayOfWeekArrayOutput
	ToDayOfWeekArrayOutputWithContext(context.Context) DayOfWeekArrayOutput
}

type DayOfWeekArray []DayOfWeek

func (DayOfWeekArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DayOfWeek)(nil)).Elem()
}

func (i DayOfWeekArray) ToDayOfWeekArrayOutput() DayOfWeekArrayOutput {
	return i.ToDayOfWeekArrayOutputWithContext(context.Background())
}

func (i DayOfWeekArray) ToDayOfWeekArrayOutputWithContext(ctx context.Context) DayOfWeekArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayOfWeekArrayOutput)
}

type DayOfWeekArrayOutput struct{ *pulumi.OutputState }

func (DayOfWeekArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekArrayOutput) ToDayOfWeekArrayOutput() DayOfWeekArrayOutput {
	return o
}

func (o DayOfWeekArrayOutput) ToDayOfWeekArrayOutputWithContext(ctx context.Context) DayOfWeekArrayOutput {
	return o
}

func (o DayOfWeekArrayOutput) Index(i pulumi.IntInput) pulumi.StringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) pulumi.StringOutput {
		return vs[0].([]DayOfWeek)[vs[1].(int)].ToStringOutput()
	}).(pulumi.StringOutput)
}

type MonthOfYear pulumi.String

const (
	MonthOfYearInvalid   = MonthOfYear("Invalid")
	MonthOfYearJanuary   = MonthOfYear("January")
	MonthOfYearFebruary  = MonthOfYear("February")
	MonthOfYearMarch     = MonthOfYear("March")
	MonthOfYearApril     = MonthOfYear("April")
	MonthOfYearMay       = MonthOfYear("May")
	MonthOfYearJune      = MonthOfYear("June")
	MonthOfYearJuly      = MonthOfYear("July")
	MonthOfYearAugust    = MonthOfYear("August")
	MonthOfYearSeptember = MonthOfYear("September")
	MonthOfYearOctober   = MonthOfYear("October")
	MonthOfYearNovember  = MonthOfYear("November")
	MonthOfYearDecember  = MonthOfYear("December")
)

func (MonthOfYear) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e MonthOfYear) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonthOfYear) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonthOfYear) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MonthOfYear) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// MonthOfYearArrayInput is an input type that accepts MonthOfYearArray and MonthOfYearArrayOutput values.
// You can construct a concrete instance of `MonthOfYearArrayInput` via:
//
//          MonthOfYearArray{ MonthOfYearArgs{...} }
type MonthOfYearArrayInput interface {
	pulumi.Input

	ToMonthOfYearArrayOutput() MonthOfYearArrayOutput
	ToMonthOfYearArrayOutputWithContext(context.Context) MonthOfYearArrayOutput
}

type MonthOfYearArray []MonthOfYear

func (MonthOfYearArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonthOfYear)(nil)).Elem()
}

func (i MonthOfYearArray) ToMonthOfYearArrayOutput() MonthOfYearArrayOutput {
	return i.ToMonthOfYearArrayOutputWithContext(context.Background())
}

func (i MonthOfYearArray) ToMonthOfYearArrayOutputWithContext(ctx context.Context) MonthOfYearArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthOfYearArrayOutput)
}

type MonthOfYearArrayOutput struct{ *pulumi.OutputState }

func (MonthOfYearArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonthOfYear)(nil)).Elem()
}

func (o MonthOfYearArrayOutput) ToMonthOfYearArrayOutput() MonthOfYearArrayOutput {
	return o
}

func (o MonthOfYearArrayOutput) ToMonthOfYearArrayOutputWithContext(ctx context.Context) MonthOfYearArrayOutput {
	return o
}

func (o MonthOfYearArrayOutput) Index(i pulumi.IntInput) pulumi.StringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) pulumi.StringOutput {
		return vs[0].([]MonthOfYear)[vs[1].(int)].ToStringOutput()
	}).(pulumi.StringOutput)
}

// The backup state of the backup item.
type ProtectedItemState pulumi.String

const (
	ProtectedItemStateInvalid           = ProtectedItemState("Invalid")
	ProtectedItemStateIRPending         = ProtectedItemState("IRPending")
	ProtectedItemStateProtected         = ProtectedItemState("Protected")
	ProtectedItemStateProtectionError   = ProtectedItemState("ProtectionError")
	ProtectedItemStateProtectionStopped = ProtectedItemState("ProtectionStopped")
	ProtectedItemStateProtectionPaused  = ProtectedItemState("ProtectionPaused")
)

func (ProtectedItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ProtectedItemState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectedItemState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectedItemState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProtectedItemState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The backup state of this backup item.
type ProtectionState pulumi.String

const (
	ProtectionStateInvalid           = ProtectionState("Invalid")
	ProtectionStateIRPending         = ProtectionState("IRPending")
	ProtectionStateProtected         = ProtectionState("Protected")
	ProtectionStateProtectionError   = ProtectionState("ProtectionError")
	ProtectionStateProtectionStopped = ProtectionState("ProtectionStopped")
	ProtectionStateProtectionPaused  = ProtectionState("ProtectionPaused")
)

func (ProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ProtectionState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectionState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectionState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProtectionState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The identity type.
type ResourceIdentityType pulumi.String

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The retention duration type of the retention policy.
type RetentionDurationType pulumi.String

const (
	RetentionDurationTypeInvalid = RetentionDurationType("Invalid")
	RetentionDurationTypeDays    = RetentionDurationType("Days")
	RetentionDurationTypeWeeks   = RetentionDurationType("Weeks")
	RetentionDurationTypeMonths  = RetentionDurationType("Months")
	RetentionDurationTypeYears   = RetentionDurationType("Years")
)

func (RetentionDurationType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e RetentionDurationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionDurationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionDurationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RetentionDurationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Retention schedule format for the yearly retention policy.
type RetentionScheduleFormat pulumi.String

const (
	RetentionScheduleFormatInvalid = RetentionScheduleFormat("Invalid")
	RetentionScheduleFormatDaily   = RetentionScheduleFormat("Daily")
	RetentionScheduleFormatWeekly  = RetentionScheduleFormat("Weekly")
)

func (RetentionScheduleFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e RetentionScheduleFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionScheduleFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionScheduleFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RetentionScheduleFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Defines the frequency interval (daily or weekly) for the schedule policy.
type ScheduleRunType pulumi.String

const (
	ScheduleRunTypeInvalid = ScheduleRunType("Invalid")
	ScheduleRunTypeDaily   = ScheduleRunType("Daily")
	ScheduleRunTypeWeekly  = ScheduleRunType("Weekly")
)

func (ScheduleRunType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ScheduleRunType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleRunType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleRunType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScheduleRunType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The Sku name.
type SkuName pulumi.String

const (
	SkuNameStandard = SkuName("Standard")
	SkuNameRS0      = SkuName("RS0")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WeekOfMonth pulumi.String

const (
	WeekOfMonthFirst  = WeekOfMonth("First")
	WeekOfMonthSecond = WeekOfMonth("Second")
	WeekOfMonthThird  = WeekOfMonth("Third")
	WeekOfMonthFourth = WeekOfMonth("Fourth")
	WeekOfMonthLast   = WeekOfMonth("Last")
)

func (WeekOfMonth) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e WeekOfMonth) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeekOfMonth) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeekOfMonth) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WeekOfMonth) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// WeekOfMonthArrayInput is an input type that accepts WeekOfMonthArray and WeekOfMonthArrayOutput values.
// You can construct a concrete instance of `WeekOfMonthArrayInput` via:
//
//          WeekOfMonthArray{ WeekOfMonthArgs{...} }
type WeekOfMonthArrayInput interface {
	pulumi.Input

	ToWeekOfMonthArrayOutput() WeekOfMonthArrayOutput
	ToWeekOfMonthArrayOutputWithContext(context.Context) WeekOfMonthArrayOutput
}

type WeekOfMonthArray []WeekOfMonth

func (WeekOfMonthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WeekOfMonth)(nil)).Elem()
}

func (i WeekOfMonthArray) ToWeekOfMonthArrayOutput() WeekOfMonthArrayOutput {
	return i.ToWeekOfMonthArrayOutputWithContext(context.Background())
}

func (i WeekOfMonthArray) ToWeekOfMonthArrayOutputWithContext(ctx context.Context) WeekOfMonthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekOfMonthArrayOutput)
}

type WeekOfMonthArrayOutput struct{ *pulumi.OutputState }

func (WeekOfMonthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WeekOfMonth)(nil)).Elem()
}

func (o WeekOfMonthArrayOutput) ToWeekOfMonthArrayOutput() WeekOfMonthArrayOutput {
	return o
}

func (o WeekOfMonthArrayOutput) ToWeekOfMonthArrayOutputWithContext(ctx context.Context) WeekOfMonthArrayOutput {
	return o
}

func (o WeekOfMonthArrayOutput) Index(i pulumi.IntInput) pulumi.StringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) pulumi.StringOutput {
		return vs[0].([]WeekOfMonth)[vs[1].(int)].ToStringOutput()
	}).(pulumi.StringOutput)
}
