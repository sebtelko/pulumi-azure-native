// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Job resource type. The progress and state can be obtained by polling a Job or subscribing to events using EventGrid.
// API Version: 2020-05-01.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:media:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Job name.
	JobName string `pulumi:"jobName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Transform name.
	TransformName string `pulumi:"transformName"`
}

// A Job resource type. The progress and state can be obtained by polling a Job or subscribing to events using EventGrid.
type LookupJobResult struct {
	// Customer provided key, value pairs that will be returned in Job and JobOutput state events.
	CorrelationData map[string]string `pulumi:"correlationData"`
	// The UTC date and time when the customer has created the Job, in 'YYYY-MM-DDThh:mm:ssZ' format.
	Created string `pulumi:"created"`
	// Optional customer supplied description of the Job.
	Description *string `pulumi:"description"`
	// The UTC date and time at which this Job finished processing.
	EndTime string `pulumi:"endTime"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The inputs for the Job.
	Input interface{} `pulumi:"input"`
	// The UTC date and time when the customer has last updated the Job, in 'YYYY-MM-DDThh:mm:ssZ' format.
	LastModified string `pulumi:"lastModified"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The outputs for the Job.
	Outputs []JobOutputAssetResponse `pulumi:"outputs"`
	// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
	Priority *string `pulumi:"priority"`
	// The UTC date and time at which this Job began processing.
	StartTime string `pulumi:"startTime"`
	// The current state of the job.
	State string `pulumi:"state"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
