// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of topic resource.
type Topic struct {
	pulumi.CustomResourceState

	// Last time the message was sent, or a request was received, for this topic.
	AccessedAt pulumi.StringOutput `pulumi:"accessedAt"`
	// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrOutput `pulumi:"autoDeleteOnIdle"`
	// Message count details
	CountDetails MessageCountDetailsResponseOutput `pulumi:"countDetails"`
	// Exact time the message was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive pulumi.StringPtrOutput `pulumi:"defaultMessageTimeToLive"`
	// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrOutput `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations pulumi.BoolPtrOutput `pulumi:"enableBatchedOperations"`
	// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
	EnableExpress pulumi.BoolPtrOutput `pulumi:"enableExpress"`
	// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning pulumi.BoolPtrOutput `pulumi:"enablePartitioning"`
	// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
	MaxSizeInMegabytes pulumi.IntPtrOutput `pulumi:"maxSizeInMegabytes"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection pulumi.BoolPtrOutput `pulumi:"requiresDuplicateDetection"`
	// Size of the topic, in bytes.
	SizeInBytes pulumi.Float64Output `pulumi:"sizeInBytes"`
	// Enumerates the possible values for the status of a messaging entity.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Number of subscriptions.
	SubscriptionCount pulumi.IntOutput `pulumi:"subscriptionCount"`
	// Value that indicates whether the topic supports ordering.
	SupportOrdering pulumi.BoolPtrOutput `pulumi:"supportOrdering"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20180101preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20140901:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20140901:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20150801:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20170401:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20210101preview:Topic"),
		},
	})
	opts = append(opts, aliases)
	var resource Topic
	err := ctx.RegisterResource("azure-native:servicebus/v20180101preview:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azure-native:servicebus/v20180101preview:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// Last time the message was sent, or a request was received, for this topic.
	AccessedAt *string `pulumi:"accessedAt"`
	// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// Message count details
	CountDetails *MessageCountDetailsResponse `pulumi:"countDetails"`
	// Exact time the message was created.
	CreatedAt *string `pulumi:"createdAt"`
	// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive *string `pulumi:"defaultMessageTimeToLive"`
	// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// Resource name
	Name *string `pulumi:"name"`
	// Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// Size of the topic, in bytes.
	SizeInBytes *float64 `pulumi:"sizeInBytes"`
	// Enumerates the possible values for the status of a messaging entity.
	Status *string `pulumi:"status"`
	// Number of subscriptions.
	SubscriptionCount *int `pulumi:"subscriptionCount"`
	// Value that indicates whether the topic supports ordering.
	SupportOrdering *bool `pulumi:"supportOrdering"`
	// Resource type
	Type *string `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type TopicState struct {
	// Last time the message was sent, or a request was received, for this topic.
	AccessedAt pulumi.StringPtrInput
	// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// Message count details
	CountDetails MessageCountDetailsResponsePtrInput
	// Exact time the message was created.
	CreatedAt pulumi.StringPtrInput
	// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive pulumi.StringPtrInput
	// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations pulumi.BoolPtrInput
	// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
	EnableExpress pulumi.BoolPtrInput
	// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning pulumi.BoolPtrInput
	// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
	MaxSizeInMegabytes pulumi.IntPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection pulumi.BoolPtrInput
	// Size of the topic, in bytes.
	SizeInBytes pulumi.Float64PtrInput
	// Enumerates the possible values for the status of a messaging entity.
	Status pulumi.StringPtrInput
	// Number of subscriptions.
	SubscriptionCount pulumi.IntPtrInput
	// Value that indicates whether the topic supports ordering.
	SupportOrdering pulumi.BoolPtrInput
	// Resource type
	Type pulumi.StringPtrInput
	// The exact time the message was updated.
	UpdatedAt pulumi.StringPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive *string `pulumi:"defaultMessageTimeToLive"`
	// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Enumerates the possible values for the status of a messaging entity.
	Status *string `pulumi:"status"`
	// Value that indicates whether the topic supports ordering.
	SupportOrdering *bool `pulumi:"supportOrdering"`
	// The topic name.
	TopicName *string `pulumi:"topicName"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
	DefaultMessageTimeToLive pulumi.StringPtrInput
	// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations pulumi.BoolPtrInput
	// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
	EnableExpress pulumi.BoolPtrInput
	// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning pulumi.BoolPtrInput
	// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
	MaxSizeInMegabytes pulumi.IntPtrInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection pulumi.BoolPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Enumerates the possible values for the status of a messaging entity.
	Status *EntityStatus
	// Value that indicates whether the topic supports ordering.
	SupportOrdering pulumi.BoolPtrInput
	// The topic name.
	TopicName pulumi.StringPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

type TopicOutput struct {
	*pulumi.OutputState
}

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TopicOutput{})
}
