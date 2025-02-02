// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Namespace/ServiceBus Connection String
// API Version: 2017-04-01.
func ListNamespaceKeys(ctx *pulumi.Context, args *ListNamespaceKeysArgs, opts ...pulumi.InvokeOption) (*ListNamespaceKeysResult, error) {
	var rv ListNamespaceKeysResult
	err := ctx.Invoke("azure-native:servicebus:listNamespaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNamespaceKeysArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Namespace/ServiceBus Connection String
type ListNamespaceKeysResult struct {
	// Primary connection string of the alias if GEO DR is enabled
	AliasPrimaryConnectionString string `pulumi:"aliasPrimaryConnectionString"`
	// Secondary  connection string of the alias if GEO DR is enabled
	AliasSecondaryConnectionString string `pulumi:"aliasSecondaryConnectionString"`
	// A string that describes the authorization rule.
	KeyName string `pulumi:"keyName"`
	// Primary connection string of the created namespace authorization rule.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey string `pulumi:"primaryKey"`
	// Secondary connection string of the created namespace authorization rule.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey string `pulumi:"secondaryKey"`
}
