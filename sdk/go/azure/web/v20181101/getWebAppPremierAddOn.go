// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Premier add-on.
func LookupWebAppPremierAddOn(ctx *pulumi.Context, args *LookupWebAppPremierAddOnArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPremierAddOnResult, error) {
	var rv LookupWebAppPremierAddOnResult
	err := ctx.Invoke("azure-native:web/v20181101:getWebAppPremierAddOn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPremierAddOnArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Add-on name.
	PremierAddOnName string `pulumi:"premierAddOnName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Premier add-on.
type LookupWebAppPremierAddOnResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Premier add on Marketplace offer.
	MarketplaceOffer *string `pulumi:"marketplaceOffer"`
	// Premier add on Marketplace publisher.
	MarketplacePublisher *string `pulumi:"marketplacePublisher"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Premier add on Product.
	Product *string `pulumi:"product"`
	// Premier add on SKU.
	Sku *string `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// Premier add on Vendor.
	Vendor *string `pulumi:"vendor"`
}
