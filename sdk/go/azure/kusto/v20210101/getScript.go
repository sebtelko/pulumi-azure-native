// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a database script.
func LookupScript(ctx *pulumi.Context, args *LookupScriptArgs, opts ...pulumi.InvokeOption) (*LookupScriptResult, error) {
	var rv LookupScriptResult
	err := ctx.Invoke("azure-native:kusto/v20210101:getScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScriptArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Kusto database script.
	ScriptName string `pulumi:"scriptName"`
}

// Class representing a database script.
type LookupScriptResult struct {
	// Flag that indicates whether to continue if one of the command fails.
	ContinueOnErrors *bool `pulumi:"continueOnErrors"`
	// A unique string. If changed the script will be applied again.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The url to the KQL script blob file.
	ScriptUrl string `pulumi:"scriptUrl"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
