// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A report resource.
func LookupReportByBillingAccount(ctx *pulumi.Context, args *LookupReportByBillingAccountArgs, opts ...pulumi.InvokeOption) (*LookupReportByBillingAccountResult, error) {
	var rv LookupReportByBillingAccountResult
	err := ctx.Invoke("azure-native:costmanagement/v20180801preview:getReportByBillingAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportByBillingAccountArgs struct {
	// BillingAccount ID
	BillingAccountId string `pulumi:"billingAccountId"`
	// Report Name.
	ReportName string `pulumi:"reportName"`
}

// A report resource.
type LookupReportByBillingAccountResult struct {
	// Has definition for the report.
	Definition ReportDefinitionResponse `pulumi:"definition"`
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	// The format of the report being delivered.
	Format *string `pulumi:"format"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Has schedule information for the report.
	Schedule *ReportScheduleResponse `pulumi:"schedule"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}
