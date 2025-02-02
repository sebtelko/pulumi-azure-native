// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Dedicated cloud node model
func LookupDedicatedCloudNode(ctx *pulumi.Context, args *LookupDedicatedCloudNodeArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedCloudNodeResult, error) {
	var rv LookupDedicatedCloudNodeResult
	err := ctx.Invoke("azure-native:vmwarecloudsimple/v20190401:getDedicatedCloudNode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedCloudNodeArgs struct {
	// dedicated cloud node name
	DedicatedCloudNodeName string `pulumi:"dedicatedCloudNodeName"`
	// The name of the resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Dedicated cloud node model
type LookupDedicatedCloudNodeResult struct {
	// Availability Zone id, e.g. "az1"
	AvailabilityZoneId string `pulumi:"availabilityZoneId"`
	// Availability Zone name, e.g. "Availability Zone 1"
	AvailabilityZoneName string `pulumi:"availabilityZoneName"`
	// VMWare Cloud Rack Name
	CloudRackName string `pulumi:"cloudRackName"`
	// date time the resource was created
	Created interface{} `pulumi:"created"`
	// SKU's id
	Id string `pulumi:"id"`
	// Azure region
	Location string `pulumi:"location"`
	// SKU's name
	Name string `pulumi:"name"`
	// count of nodes to create
	NodesCount int `pulumi:"nodesCount"`
	// Placement Group id, e.g. "n1"
	PlacementGroupId string `pulumi:"placementGroupId"`
	// Placement Name, e.g. "Placement Group 1"
	PlacementGroupName string `pulumi:"placementGroupName"`
	// Private Cloud Id
	PrivateCloudId string `pulumi:"privateCloudId"`
	// Resource Pool Name
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The provisioning status of the resource
	ProvisioningState string `pulumi:"provisioningState"`
	// purchase id
	PurchaseId string `pulumi:"purchaseId"`
	// Dedicated Cloud Nodes SKU
	Sku *SkuResponse `pulumi:"sku"`
	// Node status, indicates is private cloud set up on this node or not
	Status string `pulumi:"status"`
	// Dedicated Cloud Nodes tags
	Tags map[string]string `pulumi:"tags"`
	// {resourceProviderNamespace}/{resourceType}
	Type string `pulumi:"type"`
	// VMWare Cluster Name
	VmwareClusterName string `pulumi:"vmwareClusterName"`
}
