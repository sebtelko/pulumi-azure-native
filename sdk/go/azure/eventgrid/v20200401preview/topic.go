// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EventGrid Topic
type Topic struct {
	pulumi.CustomResourceState

	// Endpoint for the topic.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Identity information for the resource.
	Identity IdentityInfoResponsePtrOutput `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleResponseArrayOutput `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema pulumi.StringPtrOutput `pulumi:"inputSchema"`
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping JsonInputSchemaMappingResponsePtrOutput `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Metric resource id for the topic.
	MetricResourceId pulumi.StringOutput `pulumi:"metricResourceId"`
	// Name of the resource.
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the topic.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The Sku pricing tier for the topic.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.InputSchema == nil {
		args.InputSchema = pulumi.StringPtr("EventGridSchema")
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200401preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170615preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20170615preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20170915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180101:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20180101:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180501preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20180501preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20180915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190101:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20190101:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190201preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20190201preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190601:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20190601:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200101preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200101preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200601:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20201015preview:Topic"),
		},
	})
	opts = append(opts, aliases)
	var resource Topic
	err := ctx.RegisterResource("azure-native:eventgrid/v20200401preview:Topic", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:eventgrid/v20200401preview:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// Endpoint for the topic.
	Endpoint *string `pulumi:"endpoint"`
	// Identity information for the resource.
	Identity *IdentityInfoResponse `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRuleResponse `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema *string `pulumi:"inputSchema"`
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping *JsonInputSchemaMappingResponse `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// Metric resource id for the topic.
	MetricResourceId *string `pulumi:"metricResourceId"`
	// Name of the resource.
	Name                       *string                             `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Provisioning state of the topic.
	ProvisioningState *string `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The Sku pricing tier for the topic.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource.
	Type *string `pulumi:"type"`
}

type TopicState struct {
	// Endpoint for the topic.
	Endpoint pulumi.StringPtrInput
	// Identity information for the resource.
	Identity IdentityInfoResponsePtrInput
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleResponseArrayInput
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema pulumi.StringPtrInput
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping JsonInputSchemaMappingResponsePtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// Metric resource id for the topic.
	MetricResourceId pulumi.StringPtrInput
	// Name of the resource.
	Name                       pulumi.StringPtrInput
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayInput
	// Provisioning state of the topic.
	ProvisioningState pulumi.StringPtrInput
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrInput
	// The Sku pricing tier for the topic.
	Sku ResourceSkuResponsePtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
	// Type of the resource.
	Type pulumi.StringPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// Identity information for the resource.
	Identity *IdentityInfo `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRule `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema *string `pulumi:"inputSchema"`
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping *JsonInputSchemaMapping `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location                   *string                         `pulumi:"location"`
	PrivateEndpointConnections []PrivateEndpointConnectionType `pulumi:"privateEndpointConnections"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Sku pricing tier for the topic.
	Sku *ResourceSku `pulumi:"sku"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Name of the topic.
	TopicName *string `pulumi:"topicName"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// Identity information for the resource.
	Identity IdentityInfoPtrInput
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleArrayInput
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema pulumi.StringPtrInput
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping JsonInputSchemaMappingPtrInput
	// Location of the resource.
	Location                   pulumi.StringPtrInput
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// The Sku pricing tier for the topic.
	Sku ResourceSkuPtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
	// Name of the topic.
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
