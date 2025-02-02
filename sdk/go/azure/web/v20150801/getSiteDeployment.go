// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents user credentials used for publishing activity
func LookupSiteDeployment(ctx *pulumi.Context, args *LookupSiteDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupSiteDeploymentResult, error) {
	var rv LookupSiteDeploymentResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteDeploymentArgs struct {
	// Id of the deployment
	Id string `pulumi:"id"`
	// Name of web app
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents user credentials used for publishing activity
type LookupSiteDeploymentResult struct {
	// Active
	Active *bool `pulumi:"active"`
	// Author
	Author *string `pulumi:"author"`
	// AuthorEmail
	AuthorEmail *string `pulumi:"authorEmail"`
	// Deployer
	Deployer *string `pulumi:"deployer"`
	// Detail
	Details *string `pulumi:"details"`
	// EndTime
	EndTime *string `pulumi:"endTime"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Message
	Message *string `pulumi:"message"`
	// Resource Name
	Name *string `pulumi:"name"`
	// StartTime
	StartTime *string `pulumi:"startTime"`
	// Status
	Status *int `pulumi:"status"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}
