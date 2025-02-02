// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of a namespace authorization rule.
func LookupQueueAuthorizationRule(ctx *pulumi.Context, args *LookupQueueAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupQueueAuthorizationRuleResult, error) {
	var rv LookupQueueAuthorizationRuleResult
	err := ctx.Invoke("azure-native:servicebus/v20140901:getQueueAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The queue name.
	QueueName string `pulumi:"queueName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a namespace authorization rule.
type LookupQueueAuthorizationRuleResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
	// Resource type
	Type string `pulumi:"type"`
}
