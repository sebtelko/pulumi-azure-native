// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A streaming job object, containing all information associated with the named streaming job.
func LookupStreamingJob(ctx *pulumi.Context, args *LookupStreamingJobArgs, opts ...pulumi.InvokeOption) (*LookupStreamingJobResult, error) {
	var rv LookupStreamingJobResult
	err := ctx.Invoke("azure-native:streamanalytics/v20160301:getStreamingJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingJobArgs struct {
	// The $expand OData query parameter. This is a comma-separated list of additional streaming job properties to include in the response, beyond the default set returned when this parameter is absent. The default set is all streaming job properties other than 'inputs', 'transformation', 'outputs', and 'functions'.
	Expand *string `pulumi:"expand"`
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A streaming job object, containing all information associated with the named streaming job.
type LookupStreamingJobResult struct {
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel *string `pulumi:"compatibilityLevel"`
	// Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
	CreatedDate string `pulumi:"createdDate"`
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale *string `pulumi:"dataLocale"`
	// The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag string `pulumi:"etag"`
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds *int `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds *int `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy *string `pulumi:"eventsOutOfOrderPolicy"`
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions []FunctionResponse `pulumi:"functions"`
	// Resource Id
	Id string `pulumi:"id"`
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs []InputResponse `pulumi:"inputs"`
	// A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
	JobId string `pulumi:"jobId"`
	// Describes the state of the streaming job.
	JobState string `pulumi:"jobState"`
	// Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
	LastOutputEventTime string `pulumi:"lastOutputEventTime"`
	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy *string `pulumi:"outputErrorPolicy"`
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode *string `pulumi:"outputStartMode"`
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime *string `pulumi:"outputStartTime"`
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs []OutputResponse `pulumi:"outputs"`
	// Describes the provisioning status of the streaming job.
	ProvisioningState string `pulumi:"provisioningState"`
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation *TransformationResponse `pulumi:"transformation"`
	// Resource type
	Type string `pulumi:"type"`
}
