// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Properties of the blob container, including Id, resource name, resource type, Etag.
func LookupBlobContainer(ctx *pulumi.Context, args *LookupBlobContainerArgs, opts ...pulumi.InvokeOption) (*LookupBlobContainerResult, error) {
	var rv LookupBlobContainerResult
	err := ctx.Invoke("azure-native:storage/v20190401:getBlobContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobContainerArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ContainerName string `pulumi:"containerName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Properties of the blob container, including Id, resource name, resource type, Etag.
type LookupBlobContainerResult struct {
	// Resource Etag.
	Etag string `pulumi:"etag"`
	// The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container. The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container.
	HasImmutabilityPolicy bool `pulumi:"hasImmutabilityPolicy"`
	// The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.
	HasLegalHold bool `pulumi:"hasLegalHold"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The ImmutabilityPolicy property of the container.
	ImmutabilityPolicy ImmutabilityPolicyPropertiesResponse `pulumi:"immutabilityPolicy"`
	// Returns the date and time the container was last modified.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased.
	LeaseDuration string `pulumi:"leaseDuration"`
	// Lease state of the container.
	LeaseState string `pulumi:"leaseState"`
	// The lease status of the container.
	LeaseStatus string `pulumi:"leaseStatus"`
	// The LegalHold property of the container.
	LegalHold LegalHoldPropertiesResponse `pulumi:"legalHold"`
	// A name-value pair to associate with the container as metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Specifies whether data in the container may be accessed publicly and the level of access.
	PublicAccess *string `pulumi:"publicAccess"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
