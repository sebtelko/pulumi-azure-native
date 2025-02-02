// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:scheduler/v20140801preview:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	// The job collection name.
	JobCollectionName string `pulumi:"jobCollectionName"`
	// The job name.
	JobName string `pulumi:"jobName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupJobResult struct {
	// Gets the job resource identifier.
	Id string `pulumi:"id"`
	// Gets the job resource name.
	Name string `pulumi:"name"`
	// Gets or sets the job properties.
	Properties JobPropertiesResponse `pulumi:"properties"`
	// Gets the job resource type.
	Type string `pulumi:"type"`
}
