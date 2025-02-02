// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fabric definition.
func LookupReplicationFabric(ctx *pulumi.Context, args *LookupReplicationFabricArgs, opts ...pulumi.InvokeOption) (*LookupReplicationFabricResult, error) {
	var rv LookupReplicationFabricResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210301:getReplicationFabric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationFabricArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// Fabric definition.
type LookupReplicationFabricResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// Fabric related data.
	Properties FabricPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type string `pulumi:"type"`
}
