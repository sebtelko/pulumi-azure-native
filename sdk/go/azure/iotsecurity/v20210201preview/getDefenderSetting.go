// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IoT Defender settings
func LookupDefenderSetting(ctx *pulumi.Context, args *LookupDefenderSettingArgs, opts ...pulumi.InvokeOption) (*LookupDefenderSettingResult, error) {
	var rv LookupDefenderSettingResult
	err := ctx.Invoke("azure-native:iotsecurity/v20210201preview:getDefenderSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDefenderSettingArgs struct {
}

// IoT Defender settings
type LookupDefenderSettingResult struct {
	// Size of the device quota (as a opposed to a Pay as You Go billing model). Value is required to be in multiples of 1000.
	DeviceQuota int `pulumi:"deviceQuota"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The kind of onboarding for the subscription
	OnboardingKind string `pulumi:"onboardingKind"`
	// Sentinel Workspace Resource Ids
	SentinelWorkspaceResourceIds []string `pulumi:"sentinelWorkspaceResourceIds"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
