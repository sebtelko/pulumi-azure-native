// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package relay

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of a namespace authorization rule.
// API Version: 2017-04-01.
func LookupNamespaceAuthorizationRule(ctx *pulumi.Context, args *LookupNamespaceAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceAuthorizationRuleResult, error) {
	var rv LookupNamespaceAuthorizationRuleResult
	err := ctx.Invoke("azure-native:relay:getNamespaceAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a namespace authorization rule.
type LookupNamespaceAuthorizationRuleResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
	// Resource type.
	Type string `pulumi:"type"`
}
