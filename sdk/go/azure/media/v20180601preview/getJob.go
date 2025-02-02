// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Job resource type. The progress and state can be obtained by polling a Job or subscribing to events using EventGrid.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:media/v20180601preview:getJob", args, &rv, opts...)
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
	// Customer provided correlation data that will be returned in Job completed events.
	CorrelationData map[string]string `pulumi:"correlationData"`
	// The UTC date and time when the Job was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
	Created string `pulumi:"created"`
	// Optional customer supplied description of the Job.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource.
	Id string `pulumi:"id"`
	// The inputs for the Job.
	Input interface{} `pulumi:"input"`
	// The UTC date and time when the Job was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
	LastModified string `pulumi:"lastModified"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The outputs for the Job.
	Outputs []JobOutputAssetResponse `pulumi:"outputs"`
	// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
	Priority *string `pulumi:"priority"`
	// The current state of the job.
	State string `pulumi:"state"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
