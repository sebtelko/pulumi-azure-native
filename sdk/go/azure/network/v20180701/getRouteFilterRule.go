// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Route Filter Rule Resource
func LookupRouteFilterRule(ctx *pulumi.Context, args *LookupRouteFilterRuleArgs, opts ...pulumi.InvokeOption) (*LookupRouteFilterRuleResult, error) {
	var rv LookupRouteFilterRuleResult
	err := ctx.Invoke("azure-native:network/v20180701:getRouteFilterRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteFilterRuleArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route filter.
	RouteFilterName string `pulumi:"routeFilterName"`
	// The name of the rule.
	RuleName string `pulumi:"ruleName"`
}

// Route Filter Rule Resource
type LookupRouteFilterRuleResult struct {
	// The access type of the rule. Valid values are: 'Allow', 'Deny'
	Access string `pulumi:"access"`
	// The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020']
	Communities []string `pulumi:"communities"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
	ProvisioningState string `pulumi:"provisioningState"`
	// The rule type of the rule. Valid value is: 'Community'
	RouteFilterRuleType string `pulumi:"routeFilterRuleType"`
}
