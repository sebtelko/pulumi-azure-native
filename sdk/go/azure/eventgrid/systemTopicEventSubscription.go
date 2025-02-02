// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Event Subscription
// API Version: 2020-04-01-preview.
type SystemTopicEventSubscription struct {
	pulumi.CustomResourceState

	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterDestination StorageBlobDeadLetterDestinationResponsePtrOutput `pulumi:"deadLetterDestination"`
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterWithResourceIdentity DeadLetterWithResourceIdentityResponsePtrOutput `pulumi:"deadLetterWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeliveryWithResourceIdentity DeliveryWithResourceIdentityResponsePtrOutput `pulumi:"deliveryWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	Destination pulumi.AnyOutput `pulumi:"destination"`
	// The event delivery schema for the event subscription.
	EventDeliverySchema pulumi.StringPtrOutput `pulumi:"eventDeliverySchema"`
	// Expiration time of the event subscription.
	ExpirationTimeUtc pulumi.StringPtrOutput `pulumi:"expirationTimeUtc"`
	// Information about the filter for the event subscription.
	Filter EventSubscriptionFilterResponsePtrOutput `pulumi:"filter"`
	// List of user defined labels.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the event subscription.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy RetryPolicyResponsePtrOutput `pulumi:"retryPolicy"`
	// The system metadata relating to Event Subscription resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Name of the topic of the event subscription.
	Topic pulumi.StringOutput `pulumi:"topic"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSystemTopicEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewSystemTopicEventSubscription(ctx *pulumi.Context,
	name string, args *SystemTopicEventSubscriptionArgs, opts ...pulumi.ResourceOption) (*SystemTopicEventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SystemTopicName == nil {
		return nil, errors.New("invalid value for required argument 'SystemTopicName'")
	}
	if args.EventDeliverySchema == nil {
		args.EventDeliverySchema = pulumi.StringPtr("EventGridSchema")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:eventgrid:SystemTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:SystemTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200401preview:SystemTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:SystemTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20201015preview:SystemTopicEventSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource SystemTopicEventSubscription
	err := ctx.RegisterResource("azure-native:eventgrid:SystemTopicEventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemTopicEventSubscription gets an existing SystemTopicEventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemTopicEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemTopicEventSubscriptionState, opts ...pulumi.ResourceOption) (*SystemTopicEventSubscription, error) {
	var resource SystemTopicEventSubscription
	err := ctx.ReadResource("azure-native:eventgrid:SystemTopicEventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemTopicEventSubscription resources.
type systemTopicEventSubscriptionState struct {
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterDestination *StorageBlobDeadLetterDestinationResponse `pulumi:"deadLetterDestination"`
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterWithResourceIdentity *DeadLetterWithResourceIdentityResponse `pulumi:"deadLetterWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeliveryWithResourceIdentity *DeliveryWithResourceIdentityResponse `pulumi:"deliveryWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	Destination interface{} `pulumi:"destination"`
	// The event delivery schema for the event subscription.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// Expiration time of the event subscription.
	ExpirationTimeUtc *string `pulumi:"expirationTimeUtc"`
	// Information about the filter for the event subscription.
	Filter *EventSubscriptionFilterResponse `pulumi:"filter"`
	// List of user defined labels.
	Labels []string `pulumi:"labels"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Provisioning state of the event subscription.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy *RetryPolicyResponse `pulumi:"retryPolicy"`
	// The system metadata relating to Event Subscription resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Name of the topic of the event subscription.
	Topic *string `pulumi:"topic"`
	// Type of the resource.
	Type *string `pulumi:"type"`
}

type SystemTopicEventSubscriptionState struct {
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterDestination StorageBlobDeadLetterDestinationResponsePtrInput
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterWithResourceIdentity DeadLetterWithResourceIdentityResponsePtrInput
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeliveryWithResourceIdentity DeliveryWithResourceIdentityResponsePtrInput
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	Destination pulumi.Input
	// The event delivery schema for the event subscription.
	EventDeliverySchema pulumi.StringPtrInput
	// Expiration time of the event subscription.
	ExpirationTimeUtc pulumi.StringPtrInput
	// Information about the filter for the event subscription.
	Filter EventSubscriptionFilterResponsePtrInput
	// List of user defined labels.
	Labels pulumi.StringArrayInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Provisioning state of the event subscription.
	ProvisioningState pulumi.StringPtrInput
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy RetryPolicyResponsePtrInput
	// The system metadata relating to Event Subscription resource.
	SystemData SystemDataResponsePtrInput
	// Name of the topic of the event subscription.
	Topic pulumi.StringPtrInput
	// Type of the resource.
	Type pulumi.StringPtrInput
}

func (SystemTopicEventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicEventSubscriptionState)(nil)).Elem()
}

type systemTopicEventSubscriptionArgs struct {
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterDestination *StorageBlobDeadLetterDestination `pulumi:"deadLetterDestination"`
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterWithResourceIdentity *DeadLetterWithResourceIdentity `pulumi:"deadLetterWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeliveryWithResourceIdentity *DeliveryWithResourceIdentity `pulumi:"deliveryWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	Destination interface{} `pulumi:"destination"`
	// The event delivery schema for the event subscription.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName *string `pulumi:"eventSubscriptionName"`
	// Expiration time of the event subscription.
	ExpirationTimeUtc *string `pulumi:"expirationTimeUtc"`
	// Information about the filter for the event subscription.
	Filter *EventSubscriptionFilter `pulumi:"filter"`
	// List of user defined labels.
	Labels []string `pulumi:"labels"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy *RetryPolicy `pulumi:"retryPolicy"`
	// Name of the system topic.
	SystemTopicName string `pulumi:"systemTopicName"`
}

// The set of arguments for constructing a SystemTopicEventSubscription resource.
type SystemTopicEventSubscriptionArgs struct {
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterDestination StorageBlobDeadLetterDestinationPtrInput
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterWithResourceIdentity DeadLetterWithResourceIdentityPtrInput
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeliveryWithResourceIdentity DeliveryWithResourceIdentityPtrInput
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	Destination pulumi.Input
	// The event delivery schema for the event subscription.
	EventDeliverySchema pulumi.StringPtrInput
	// Name of the event subscription to be created. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName pulumi.StringPtrInput
	// Expiration time of the event subscription.
	ExpirationTimeUtc pulumi.StringPtrInput
	// Information about the filter for the event subscription.
	Filter EventSubscriptionFilterPtrInput
	// List of user defined labels.
	Labels pulumi.StringArrayInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy RetryPolicyPtrInput
	// Name of the system topic.
	SystemTopicName pulumi.StringInput
}

func (SystemTopicEventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicEventSubscriptionArgs)(nil)).Elem()
}

type SystemTopicEventSubscriptionInput interface {
	pulumi.Input

	ToSystemTopicEventSubscriptionOutput() SystemTopicEventSubscriptionOutput
	ToSystemTopicEventSubscriptionOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionOutput
}

func (*SystemTopicEventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemTopicEventSubscription)(nil))
}

func (i *SystemTopicEventSubscription) ToSystemTopicEventSubscriptionOutput() SystemTopicEventSubscriptionOutput {
	return i.ToSystemTopicEventSubscriptionOutputWithContext(context.Background())
}

func (i *SystemTopicEventSubscription) ToSystemTopicEventSubscriptionOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicEventSubscriptionOutput)
}

type SystemTopicEventSubscriptionOutput struct {
	*pulumi.OutputState
}

func (SystemTopicEventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemTopicEventSubscription)(nil))
}

func (o SystemTopicEventSubscriptionOutput) ToSystemTopicEventSubscriptionOutput() SystemTopicEventSubscriptionOutput {
	return o
}

func (o SystemTopicEventSubscriptionOutput) ToSystemTopicEventSubscriptionOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SystemTopicEventSubscriptionOutput{})
}
