// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Management lock information.
func LookupManagementLockAtResourceGroupLevel(ctx *pulumi.Context, args *LookupManagementLockAtResourceGroupLevelArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockAtResourceGroupLevelResult, error) {
	var rv LookupManagementLockAtResourceGroupLevelResult
	err := ctx.Invoke("azure-native:authorization/v20150101:getManagementLockAtResourceGroupLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockAtResourceGroupLevelArgs struct {
	// The lock name.
	LockName string `pulumi:"lockName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Management lock information.
type LookupManagementLockAtResourceGroupLevelResult struct {
	// The Id of the lock.
	Id string `pulumi:"id"`
	// The lock level of the management lock.
	Level *string `pulumi:"level"`
	// The name of the lock.
	Name *string `pulumi:"name"`
	// The notes of the management lock.
	Notes *string `pulumi:"notes"`
	// The type of the lock.
	Type string `pulumi:"type"`
}
