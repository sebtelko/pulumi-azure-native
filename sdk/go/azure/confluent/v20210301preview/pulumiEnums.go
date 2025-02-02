// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SaaS Offer Status
type SaaSOfferStatus pulumi.String

const (
	SaaSOfferStatusStarted                 = SaaSOfferStatus("Started")
	SaaSOfferStatusPendingFulfillmentStart = SaaSOfferStatus("PendingFulfillmentStart")
	SaaSOfferStatusInProgress              = SaaSOfferStatus("InProgress")
	SaaSOfferStatusSubscribed              = SaaSOfferStatus("Subscribed")
	SaaSOfferStatusSuspended               = SaaSOfferStatus("Suspended")
	SaaSOfferStatusReinstated              = SaaSOfferStatus("Reinstated")
	SaaSOfferStatusSucceeded               = SaaSOfferStatus("Succeeded")
	SaaSOfferStatusFailed                  = SaaSOfferStatus("Failed")
	SaaSOfferStatusUnsubscribed            = SaaSOfferStatus("Unsubscribed")
	SaaSOfferStatusUpdating                = SaaSOfferStatus("Updating")
)

func (SaaSOfferStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e SaaSOfferStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SaaSOfferStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SaaSOfferStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SaaSOfferStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
