// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// User credentials used for publishing activity.
func LookupWebAppDeploymentSlot(ctx *pulumi.Context, args *LookupWebAppDeploymentSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDeploymentSlotResult, error) {
	var rv LookupWebAppDeploymentSlotResult
	err := ctx.Invoke("azure-native:web/v20200601:getWebAppDeploymentSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDeploymentSlotArgs struct {
	// Deployment ID.
	Id string `pulumi:"id"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API gets a deployment for the production slot.
	Slot string `pulumi:"slot"`
}

// User credentials used for publishing activity.
type LookupWebAppDeploymentSlotResult struct {
	// True if deployment is currently active, false if completed and null if not started.
	Active *bool `pulumi:"active"`
	// Who authored the deployment.
	Author *string `pulumi:"author"`
	// Author email.
	AuthorEmail *string `pulumi:"authorEmail"`
	// Who performed the deployment.
	Deployer *string `pulumi:"deployer"`
	// Details on deployment.
	Details *string `pulumi:"details"`
	// End time.
	EndTime *string `pulumi:"endTime"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Details about deployment status.
	Message *string `pulumi:"message"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Start time.
	StartTime *string `pulumi:"startTime"`
	// Deployment status.
	Status *int `pulumi:"status"`
	// Resource type.
	Type string `pulumi:"type"`
}
