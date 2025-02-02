// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Single item in List or Get Consumer group operation
func LookupConsumerGroup(ctx *pulumi.Context, args *LookupConsumerGroupArgs, opts ...pulumi.InvokeOption) (*LookupConsumerGroupResult, error) {
	var rv LookupConsumerGroupResult
	err := ctx.Invoke("azure-native:eventhub/v20170401:getConsumerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConsumerGroupArgs struct {
	// The consumer group name
	ConsumerGroupName string `pulumi:"consumerGroupName"`
	// The Event Hub name
	EventHubName string `pulumi:"eventHubName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Single item in List or Get Consumer group operation
type LookupConsumerGroupResult struct {
	// Exact time the message was created.
	CreatedAt string `pulumi:"createdAt"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt string `pulumi:"updatedAt"`
	// User Metadata is a placeholder to store user-defined string data with maximum length 1024. e.g. it can be used to store descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
	UserMetadata *string `pulumi:"userMetadata"`
}
