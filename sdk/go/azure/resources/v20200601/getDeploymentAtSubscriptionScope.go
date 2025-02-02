// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deployment information.
func LookupDeploymentAtSubscriptionScope(ctx *pulumi.Context, args *LookupDeploymentAtSubscriptionScopeArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentAtSubscriptionScopeResult, error) {
	var rv LookupDeploymentAtSubscriptionScopeResult
	err := ctx.Invoke("azure-native:resources/v20200601:getDeploymentAtSubscriptionScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentAtSubscriptionScopeArgs struct {
	// The name of the deployment.
	DeploymentName string `pulumi:"deploymentName"`
}

// Deployment information.
type LookupDeploymentAtSubscriptionScopeResult struct {
	// The ID of the deployment.
	Id string `pulumi:"id"`
	// the location of the deployment.
	Location *string `pulumi:"location"`
	// The name of the deployment.
	Name string `pulumi:"name"`
	// Deployment properties.
	Properties DeploymentPropertiesExtendedResponse `pulumi:"properties"`
	// Deployment tags
	Tags map[string]string `pulumi:"tags"`
	// The type of the deployment.
	Type string `pulumi:"type"`
}
