// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a lab.
// API Version: 2018-10-15.
func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	var rv LookupLabResult
	err := ctx.Invoke("azure-native:labservices:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabArgs struct {
	// Specify the $expand query. Example: 'properties($select=maxUsersInLab)'
	Expand *string `pulumi:"expand"`
	// The name of the lab Account.
	LabAccountName string `pulumi:"labAccountName"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a lab.
type LookupLabResult struct {
	// Object id of the user that created the lab.
	CreatedByObjectId string `pulumi:"createdByObjectId"`
	// Lab creator name
	CreatedByUserPrincipalName string `pulumi:"createdByUserPrincipalName"`
	// Creation date for the lab
	CreatedDate string `pulumi:"createdDate"`
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// Invitation code that users can use to join a lab.
	InvitationCode string `pulumi:"invitationCode"`
	// The details of the latest operation. ex: status, error
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// Maximum number of users allowed in the lab.
	MaxUsersInLab *int `pulumi:"maxUsersInLab"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
	// Maximum duration a user can use an environment for in the lab.
	UsageQuota *string `pulumi:"usageQuota"`
	// Lab user access mode (open to all vs. restricted to those listed on the lab).
	UserAccessMode *string `pulumi:"userAccessMode"`
	// Maximum value MaxUsersInLab can be set to, as specified by the service
	UserQuota int `pulumi:"userQuota"`
}
