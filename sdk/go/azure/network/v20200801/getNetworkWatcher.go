// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network watcher in a resource group.
func LookupNetworkWatcher(ctx *pulumi.Context, args *LookupNetworkWatcherArgs, opts ...pulumi.InvokeOption) (*LookupNetworkWatcherResult, error) {
	var rv LookupNetworkWatcherResult
	err := ctx.Invoke("azure-native:network/v20200801:getNetworkWatcher", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkWatcherArgs struct {
	// The name of the network watcher.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Network watcher in a resource group.
type LookupNetworkWatcherResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the network watcher resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}
