// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets or sets the type.
type ParameterType pulumi.String

const (
	ParameterTypeNotSpecified = ParameterType("NotSpecified")
	ParameterTypeString       = ParameterType("String")
	ParameterTypeSecureString = ParameterType("SecureString")
	ParameterTypeInt          = ParameterType("Int")
	ParameterTypeFloat        = ParameterType("Float")
	ParameterTypeBool         = ParameterType("Bool")
	ParameterTypeArray        = ParameterType("Array")
	ParameterTypeObject       = ParameterType("Object")
	ParameterTypeSecureObject = ParameterType("SecureObject")
)

func (ParameterType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ParameterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ParameterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Gets or sets the name.
type SkuName pulumi.String

const (
	SkuNameNotSpecified = SkuName("NotSpecified")
	SkuNameFree         = SkuName("Free")
	SkuNameShared       = SkuName("Shared")
	SkuNameBasic        = SkuName("Basic")
	SkuNameStandard     = SkuName("Standard")
	SkuNamePremium      = SkuName("Premium")
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

// Gets or sets the state.
type WorkflowStateEnum pulumi.String

const (
	WorkflowStateEnumNotSpecified = WorkflowStateEnum("NotSpecified")
	WorkflowStateEnumEnabled      = WorkflowStateEnum("Enabled")
	WorkflowStateEnumDisabled     = WorkflowStateEnum("Disabled")
	WorkflowStateEnumDeleted      = WorkflowStateEnum("Deleted")
	WorkflowStateEnumSuspended    = WorkflowStateEnum("Suspended")
)

func (WorkflowStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e WorkflowStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkflowStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkflowStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkflowStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
