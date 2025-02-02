// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// LogAnalytics operation status response
func GetLogAnalyticExportRequestRateByInterval(ctx *pulumi.Context, args *GetLogAnalyticExportRequestRateByIntervalArgs, opts ...pulumi.InvokeOption) (*GetLogAnalyticExportRequestRateByIntervalResult, error) {
	var rv GetLogAnalyticExportRequestRateByIntervalResult
	err := ctx.Invoke("azure-native:compute/v20201201:getLogAnalyticExportRequestRateByInterval", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLogAnalyticExportRequestRateByIntervalArgs struct {
	// SAS Uri of the logging blob container to which LogAnalytics Api writes output logs to.
	BlobContainerSasUri string `pulumi:"blobContainerSasUri"`
	// From time of the query
	FromTime string `pulumi:"fromTime"`
	// Group query result by Client Application ID.
	GroupByClientApplicationId *bool `pulumi:"groupByClientApplicationId"`
	// Group query result by Operation Name.
	GroupByOperationName *bool `pulumi:"groupByOperationName"`
	// Group query result by Resource Name.
	GroupByResourceName *bool `pulumi:"groupByResourceName"`
	// Group query result by Throttle Policy applied.
	GroupByThrottlePolicy *bool `pulumi:"groupByThrottlePolicy"`
	// Group query result by User Agent.
	GroupByUserAgent *bool `pulumi:"groupByUserAgent"`
	// Interval value in minutes used to create LogAnalytics call rate logs.
	IntervalLength string `pulumi:"intervalLength"`
	// The location upon which virtual-machine-sizes is queried.
	Location string `pulumi:"location"`
	// To time of the query
	ToTime string `pulumi:"toTime"`
}

// LogAnalytics operation status response
type GetLogAnalyticExportRequestRateByIntervalResult struct {
	// LogAnalyticsOutput
	Properties LogAnalyticsOutputResponse `pulumi:"properties"`
}
