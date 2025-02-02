// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scheduler

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API Version: 2016-03-01.
type Job struct {
	pulumi.CustomResourceState

	// Gets the job resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the job properties.
	Properties JobPropertiesResponseOutput `pulumi:"properties"`
	// Gets the job resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobCollectionName == nil {
		return nil, errors.New("invalid value for required argument 'JobCollectionName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:scheduler:Job"),
		},
		{
			Type: pulumi.String("azure-native:scheduler/v20140801preview:Job"),
		},
		{
			Type: pulumi.String("azure-nextgen:scheduler/v20140801preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:scheduler/v20160101:Job"),
		},
		{
			Type: pulumi.String("azure-nextgen:scheduler/v20160101:Job"),
		},
		{
			Type: pulumi.String("azure-native:scheduler/v20160301:Job"),
		},
		{
			Type: pulumi.String("azure-nextgen:scheduler/v20160301:Job"),
		},
	})
	opts = append(opts, aliases)
	var resource Job
	err := ctx.RegisterResource("azure-native:scheduler:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure-native:scheduler:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// Gets the job resource name.
	Name *string `pulumi:"name"`
	// Gets or sets the job properties.
	Properties *JobPropertiesResponse `pulumi:"properties"`
	// Gets the job resource type.
	Type *string `pulumi:"type"`
}

type JobState struct {
	// Gets the job resource name.
	Name pulumi.StringPtrInput
	// Gets or sets the job properties.
	Properties JobPropertiesResponsePtrInput
	// Gets the job resource type.
	Type pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// The job collection name.
	JobCollectionName string `pulumi:"jobCollectionName"`
	// The job name.
	JobName *string `pulumi:"jobName"`
	// Gets or sets the job properties.
	Properties *JobProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// The job collection name.
	JobCollectionName pulumi.StringInput
	// The job name.
	JobName pulumi.StringPtrInput
	// Gets or sets the job properties.
	Properties JobPropertiesPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((*Job)(nil))
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct {
	*pulumi.OutputState
}

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Job)(nil))
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}
