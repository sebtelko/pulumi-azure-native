// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The access control record.
func LookupAccessControlRecord(ctx *pulumi.Context, args *LookupAccessControlRecordArgs, opts ...pulumi.InvokeOption) (*LookupAccessControlRecordResult, error) {
	var rv LookupAccessControlRecordResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:getAccessControlRecord", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessControlRecordArgs struct {
	// Name of access control record to be fetched.
	AccessControlRecordName string `pulumi:"accessControlRecordName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The access control record.
type LookupAccessControlRecordResult struct {
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// The iSCSI initiator name (IQN).
	InitiatorName string `pulumi:"initiatorName"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The name of the object.
	Name string `pulumi:"name"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
	// The number of volumes using the access control record.
	VolumeCount int `pulumi:"volumeCount"`
}
