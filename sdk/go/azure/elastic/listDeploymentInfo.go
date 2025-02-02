// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elastic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The properties of deployment in Elastic cloud corresponding to the Elastic monitor resource.
// API Version: 2020-07-01-preview.
func ListDeploymentInfo(ctx *pulumi.Context, args *ListDeploymentInfoArgs, opts ...pulumi.InvokeOption) (*ListDeploymentInfoResult, error) {
	var rv ListDeploymentInfoResult
	err := ctx.Invoke("azure-native:elastic:listDeploymentInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDeploymentInfoArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The properties of deployment in Elastic cloud corresponding to the Elastic monitor resource.
type ListDeploymentInfoResult struct {
	// Disk capacity of the elasticsearch in Elastic cloud deployment.
	DiskCapacity string `pulumi:"diskCapacity"`
	// RAM capacity of the elasticsearch in Elastic cloud deployment.
	MemoryCapacity string `pulumi:"memoryCapacity"`
	// The Elastic deployment status.
	Status string `pulumi:"status"`
	// Version of the elasticsearch in Elastic cloud deployment.
	Version string `pulumi:"version"`
}
