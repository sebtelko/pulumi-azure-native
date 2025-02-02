// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Logger details.
func LookupLogger(ctx *pulumi.Context, args *LookupLoggerArgs, opts ...pulumi.InvokeOption) (*LookupLoggerResult, error) {
	var rv LookupLoggerResult
	err := ctx.Invoke("azure-native:apimanagement/v20170301:getLogger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoggerArgs struct {
	// Logger identifier. Must be unique in the API Management service instance.
	Loggerid string `pulumi:"loggerid"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Logger details.
type LookupLoggerResult struct {
	// The name and SendRule connection string of the event hub for azureEventHub logger.
	// Instrumentation key for applicationInsights logger.
	Credentials map[string]string `pulumi:"credentials"`
	// Logger description.
	Description *string `pulumi:"description"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Whether records are buffered in the logger before publishing. Default is assumed to be true.
	IsBuffered *bool `pulumi:"isBuffered"`
	// Logger type.
	LoggerType string `pulumi:"loggerType"`
	// Resource name.
	Name string `pulumi:"name"`
	// Sampling settings for an ApplicationInsights logger.
	Sampling *LoggerSamplingContractResponse `pulumi:"sampling"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}
