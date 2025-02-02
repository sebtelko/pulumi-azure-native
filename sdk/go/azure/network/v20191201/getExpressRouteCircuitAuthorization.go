// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Authorization in an ExpressRouteCircuit resource.
func LookupExpressRouteCircuitAuthorization(ctx *pulumi.Context, args *LookupExpressRouteCircuitAuthorizationArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitAuthorizationResult, error) {
	var rv LookupExpressRouteCircuitAuthorizationResult
	err := ctx.Invoke("azure-native:network/v20191201:getExpressRouteCircuitAuthorization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitAuthorizationArgs struct {
	// The name of the authorization.
	AuthorizationName string `pulumi:"authorizationName"`
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Authorization in an ExpressRouteCircuit resource.
type LookupExpressRouteCircuitAuthorizationResult struct {
	// The authorization key.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The authorization use status.
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the authorization resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Type of the resource.
	Type string `pulumi:"type"`
}
