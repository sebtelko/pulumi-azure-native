// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Hybrid Connection for an App Service app.
func LookupWebAppRelayServiceConnection(ctx *pulumi.Context, args *LookupWebAppRelayServiceConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppRelayServiceConnectionResult, error) {
	var rv LookupWebAppRelayServiceConnectionResult
	err := ctx.Invoke("azure-native:web/v20160801:getWebAppRelayServiceConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppRelayServiceConnectionArgs struct {
	// Name of the hybrid connection.
	EntityName string `pulumi:"entityName"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Hybrid Connection for an App Service app.
type LookupWebAppRelayServiceConnectionResult struct {
	BiztalkUri             *string `pulumi:"biztalkUri"`
	EntityConnectionString *string `pulumi:"entityConnectionString"`
	EntityName             *string `pulumi:"entityName"`
	Hostname               *string `pulumi:"hostname"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name                     string  `pulumi:"name"`
	Port                     *int    `pulumi:"port"`
	ResourceConnectionString *string `pulumi:"resourceConnectionString"`
	ResourceType             *string `pulumi:"resourceType"`
	// Resource type.
	Type string `pulumi:"type"`
}
