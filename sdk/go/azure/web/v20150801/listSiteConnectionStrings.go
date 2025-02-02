// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// String dictionary resource
func ListSiteConnectionStrings(ctx *pulumi.Context, args *ListSiteConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListSiteConnectionStringsResult, error) {
	var rv ListSiteConnectionStringsResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteConnectionStringsArgs struct {
	// Name of web app
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource
type ListSiteConnectionStringsResult struct {
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Connection strings
	Properties map[string]ConnStringValueTypePairResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}
