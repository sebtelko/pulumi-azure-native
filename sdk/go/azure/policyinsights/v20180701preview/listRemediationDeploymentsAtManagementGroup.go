// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List of deployments for a remediation.
func ListRemediationDeploymentsAtManagementGroup(ctx *pulumi.Context, args *ListRemediationDeploymentsAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtManagementGroupResult, error) {
	var rv ListRemediationDeploymentsAtManagementGroupResult
	err := ctx.Invoke("azure-native:policyinsights/v20180701preview:listRemediationDeploymentsAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtManagementGroupArgs struct {
	// Management group ID.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The namespace for Microsoft Management RP; only "Microsoft.Management" is allowed.
	ManagementGroupsNamespace string `pulumi:"managementGroupsNamespace"`
	// The name of the remediation.
	RemediationName string `pulumi:"remediationName"`
	// Maximum number of records to return.
	Top *int `pulumi:"top"`
}

// List of deployments for a remediation.
type ListRemediationDeploymentsAtManagementGroupResult struct {
	// The URL to get the next set of results.
	NextLink string `pulumi:"nextLink"`
	// Array of deployments for the remediation.
	Value []RemediationDeploymentResponse `pulumi:"value"`
}
