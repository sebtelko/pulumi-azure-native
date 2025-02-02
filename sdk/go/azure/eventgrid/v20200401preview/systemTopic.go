// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EventGrid System Topic.
type SystemTopic struct {
	pulumi.CustomResourceState

	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Metric resource id for the system topic.
	MetricResourceId pulumi.StringOutput `pulumi:"metricResourceId"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the system topic.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Source for the system topic.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// The system metadata relating to System Topic resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// TopicType for the system topic.
	TopicType pulumi.StringPtrOutput `pulumi:"topicType"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSystemTopic registers a new resource with the given unique name, arguments, and options.
func NewSystemTopic(ctx *pulumi.Context,
	name string, args *SystemTopicArgs, opts ...pulumi.ResourceOption) (*SystemTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200401preview:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:SystemTopic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20201015preview:SystemTopic"),
		},
	})
	opts = append(opts, aliases)
	var resource SystemTopic
	err := ctx.RegisterResource("azure-native:eventgrid/v20200401preview:SystemTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemTopic gets an existing SystemTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemTopicState, opts ...pulumi.ResourceOption) (*SystemTopic, error) {
	var resource SystemTopic
	err := ctx.ReadResource("azure-native:eventgrid/v20200401preview:SystemTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemTopic resources.
type systemTopicState struct {
	// Location of the resource.
	Location *string `pulumi:"location"`
	// Metric resource id for the system topic.
	MetricResourceId *string `pulumi:"metricResourceId"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Provisioning state of the system topic.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Source for the system topic.
	Source *string `pulumi:"source"`
	// The system metadata relating to System Topic resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// TopicType for the system topic.
	TopicType *string `pulumi:"topicType"`
	// Type of the resource.
	Type *string `pulumi:"type"`
}

type SystemTopicState struct {
	// Location of the resource.
	Location pulumi.StringPtrInput
	// Metric resource id for the system topic.
	MetricResourceId pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Provisioning state of the system topic.
	ProvisioningState pulumi.StringPtrInput
	// Source for the system topic.
	Source pulumi.StringPtrInput
	// The system metadata relating to System Topic resource.
	SystemData SystemDataResponsePtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
	// TopicType for the system topic.
	TopicType pulumi.StringPtrInput
	// Type of the resource.
	Type pulumi.StringPtrInput
}

func (SystemTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicState)(nil)).Elem()
}

type systemTopicArgs struct {
	// Location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Source for the system topic.
	Source *string `pulumi:"source"`
	// Name of the system topic.
	SystemTopicName *string `pulumi:"systemTopicName"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// TopicType for the system topic.
	TopicType *string `pulumi:"topicType"`
}

// The set of arguments for constructing a SystemTopic resource.
type SystemTopicArgs struct {
	// Location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Source for the system topic.
	Source pulumi.StringPtrInput
	// Name of the system topic.
	SystemTopicName pulumi.StringPtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
	// TopicType for the system topic.
	TopicType pulumi.StringPtrInput
}

func (SystemTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicArgs)(nil)).Elem()
}

type SystemTopicInput interface {
	pulumi.Input

	ToSystemTopicOutput() SystemTopicOutput
	ToSystemTopicOutputWithContext(ctx context.Context) SystemTopicOutput
}

func (*SystemTopic) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemTopic)(nil))
}

func (i *SystemTopic) ToSystemTopicOutput() SystemTopicOutput {
	return i.ToSystemTopicOutputWithContext(context.Background())
}

func (i *SystemTopic) ToSystemTopicOutputWithContext(ctx context.Context) SystemTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicOutput)
}

type SystemTopicOutput struct {
	*pulumi.OutputState
}

func (SystemTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemTopic)(nil))
}

func (o SystemTopicOutput) ToSystemTopicOutput() SystemTopicOutput {
	return o
}

func (o SystemTopicOutput) ToSystemTopicOutputWithContext(ctx context.Context) SystemTopicOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SystemTopicOutput{})
}
