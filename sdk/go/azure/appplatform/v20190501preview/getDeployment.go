// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deployment resource payload
func LookupDeployment(ctx *pulumi.Context, args *LookupDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentResult, error) {
	var rv LookupDeploymentResult
	err := ctx.Invoke("azure-native:appplatform/v20190501preview:getDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the Deployment resource.
	DeploymentName string `pulumi:"deploymentName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Deployment resource payload
type LookupDeploymentResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Properties of the Deployment resource
	Properties DeploymentResourcePropertiesResponse `pulumi:"properties"`
	// Sku of the Deployment resource
	Sku *SkuResponse `pulumi:"sku"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
