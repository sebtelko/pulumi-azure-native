// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Result of the request to list networkManagerEffectiveConfiguration. It contains a list of groups and a URL link to get the next set of results.
// API Version: 2021-02-01-preview.
func ListEffectiveConfiguration(ctx *pulumi.Context, args *ListEffectiveConfigurationArgs, opts ...pulumi.InvokeOption) (*ListEffectiveConfigurationResult, error) {
	var rv ListEffectiveConfigurationResult
	err := ctx.Invoke("azure-native:network:listEffectiveConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEffectiveConfigurationArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// The name of the virtual network.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// Result of the request to list networkManagerEffectiveConfiguration. It contains a list of groups and a URL link to get the next set of results.
type ListEffectiveConfigurationResult struct {
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// Gets a page of NetworkManagerEffectiveConfiguration
	Value []interface{} `pulumi:"value"`
}
