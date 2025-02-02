// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Role Assignments
func LookupRoleAssignment(ctx *pulumi.Context, args *LookupRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentResult, error) {
	var rv LookupRoleAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20200801preview:getRoleAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleAssignmentArgs struct {
	// The name of the role assignment. It can be any valid GUID.
	RoleAssignmentName string `pulumi:"roleAssignmentName"`
	// The scope of the operation or resource. Valid scopes are: subscription (format: '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}', or resource (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
	Scope string `pulumi:"scope"`
}

// Role Assignments
type LookupRoleAssignmentResult struct {
	// The conditions on the role assignment. This limits the resources it can be assigned to. e.g.: @Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'
	Condition *string `pulumi:"condition"`
	// Version of the condition. Currently accepted value is '2.0'
	ConditionVersion *string `pulumi:"conditionVersion"`
	// Id of the user who created the assignment
	CreatedBy string `pulumi:"createdBy"`
	// Time it was created
	CreatedOn string `pulumi:"createdOn"`
	// Id of the delegated managed identity resource
	DelegatedManagedIdentityResourceId *string `pulumi:"delegatedManagedIdentityResourceId"`
	// Description of role assignment
	Description *string `pulumi:"description"`
	// The role assignment ID.
	Id string `pulumi:"id"`
	// The role assignment name.
	Name string `pulumi:"name"`
	// The principal ID.
	PrincipalId string `pulumi:"principalId"`
	// The principal type of the assigned principal ID.
	PrincipalType *string `pulumi:"principalType"`
	// The role definition ID.
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
	// The role assignment scope.
	Scope string `pulumi:"scope"`
	// The role assignment type.
	Type string `pulumi:"type"`
	// Id of the user who updated the assignment
	UpdatedBy string `pulumi:"updatedBy"`
	// Time it was updated
	UpdatedOn string `pulumi:"updatedOn"`
}
