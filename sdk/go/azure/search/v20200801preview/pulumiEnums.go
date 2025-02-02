// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
type HostingMode pulumi.String

const (
	HostingModeDefault     = HostingMode("default")
	HostingModeHighDensity = HostingMode("highDensity")
)

func (HostingMode) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e HostingMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostingMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostingMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostingMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The identity type.
type IdentityType pulumi.String

const (
	IdentityTypeNone           = IdentityType("None")
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
)

func (IdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e IdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Status of the the private link service connection. Can be Pending, Approved, Rejected, or Disconnected.
type PrivateLinkServiceConnectionStatus pulumi.String

const (
	PrivateLinkServiceConnectionStatusPending      = PrivateLinkServiceConnectionStatus("Pending")
	PrivateLinkServiceConnectionStatusApproved     = PrivateLinkServiceConnectionStatus("Approved")
	PrivateLinkServiceConnectionStatusRejected     = PrivateLinkServiceConnectionStatus("Rejected")
	PrivateLinkServiceConnectionStatusDisconnected = PrivateLinkServiceConnectionStatus("Disconnected")
)

func (PrivateLinkServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e PrivateLinkServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
type PublicNetworkAccess pulumi.String

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("disabled")
)

func (PublicNetworkAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e PublicNetworkAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The provisioning state of the shared private link resource. Can be Updating, Deleting, Failed, Succeeded or Incomplete.
type SharedPrivateLinkResourceProvisioningState pulumi.String

const (
	SharedPrivateLinkResourceProvisioningStateUpdating   = SharedPrivateLinkResourceProvisioningState("Updating")
	SharedPrivateLinkResourceProvisioningStateDeleting   = SharedPrivateLinkResourceProvisioningState("Deleting")
	SharedPrivateLinkResourceProvisioningStateFailed     = SharedPrivateLinkResourceProvisioningState("Failed")
	SharedPrivateLinkResourceProvisioningStateSucceeded  = SharedPrivateLinkResourceProvisioningState("Succeeded")
	SharedPrivateLinkResourceProvisioningStateIncomplete = SharedPrivateLinkResourceProvisioningState("Incomplete")
)

func (SharedPrivateLinkResourceProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e SharedPrivateLinkResourceProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedPrivateLinkResourceProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedPrivateLinkResourceProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SharedPrivateLinkResourceProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Status of the shared private link resource. Can be Pending, Approved, Rejected or Disconnected.
type SharedPrivateLinkResourceStatus pulumi.String

const (
	SharedPrivateLinkResourceStatusPending      = SharedPrivateLinkResourceStatus("Pending")
	SharedPrivateLinkResourceStatusApproved     = SharedPrivateLinkResourceStatus("Approved")
	SharedPrivateLinkResourceStatusRejected     = SharedPrivateLinkResourceStatus("Rejected")
	SharedPrivateLinkResourceStatusDisconnected = SharedPrivateLinkResourceStatus("Disconnected")
)

func (SharedPrivateLinkResourceStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e SharedPrivateLinkResourceStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedPrivateLinkResourceStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedPrivateLinkResourceStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SharedPrivateLinkResourceStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The SKU of the search service. Valid values include: 'free': Shared service. 'basic': Dedicated service with up to 3 replicas. 'standard': Dedicated service with up to 12 partitions and 12 replicas. 'standard2': Similar to standard, but with more capacity per search unit. 'standard3': The largest Standard offering with up to 12 partitions and 12 replicas (or up to 3 partitions with more indexes if you also set the hostingMode property to 'highDensity'). 'storage_optimized_l1': Supports 1TB per partition, up to 12 partitions. 'storage_optimized_l2': Supports 2TB per partition, up to 12 partitions.'
type SkuName pulumi.String

const (
	SkuNameFree                  = SkuName("free")
	SkuNameBasic                 = SkuName("basic")
	SkuNameStandard              = SkuName("standard")
	SkuNameStandard2             = SkuName("standard2")
	SkuNameStandard3             = SkuName("standard3")
	SkuName_Storage_optimized_l1 = SkuName("storage_optimized_l1")
	SkuName_Storage_optimized_l2 = SkuName("storage_optimized_l2")
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
