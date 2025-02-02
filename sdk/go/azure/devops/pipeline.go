// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Pipeline used to configure Continuous Integration (CI) & Continuous Delivery (CD) for Azure resources.
// API Version: 2020-07-13-preview.
type Pipeline struct {
	pulumi.CustomResourceState

	// Configuration used to bootstrap the Pipeline.
	BootstrapConfiguration BootstrapConfigurationResponseOutput `pulumi:"bootstrapConfiguration"`
	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Unique identifier of the Pipeline
	PipelineId pulumi.IntOutput `pulumi:"pipelineId"`
	// Specifies which CI/CD provider to use. Valid options are 'azurePipeline', 'githubWorkflow'.
	PipelineType pulumi.StringOutput `pulumi:"pipelineType"`
	// The system metadata pertaining to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource Tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BootstrapConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'BootstrapConfiguration'")
	}
	if args.PipelineType == nil {
		return nil, errors.New("invalid value for required argument 'PipelineType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devops:Pipeline"),
		},
		{
			Type: pulumi.String("azure-native:devops/v20190701preview:Pipeline"),
		},
		{
			Type: pulumi.String("azure-nextgen:devops/v20190701preview:Pipeline"),
		},
		{
			Type: pulumi.String("azure-native:devops/v20200713preview:Pipeline"),
		},
		{
			Type: pulumi.String("azure-nextgen:devops/v20200713preview:Pipeline"),
		},
	})
	opts = append(opts, aliases)
	var resource Pipeline
	err := ctx.RegisterResource("azure-native:devops:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeline gets an existing Pipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("azure-native:devops:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
	// Configuration used to bootstrap the Pipeline.
	BootstrapConfiguration *BootstrapConfigurationResponse `pulumi:"bootstrapConfiguration"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Unique identifier of the Pipeline
	PipelineId *int `pulumi:"pipelineId"`
	// Specifies which CI/CD provider to use. Valid options are 'azurePipeline', 'githubWorkflow'.
	PipelineType *string `pulumi:"pipelineType"`
	// The system metadata pertaining to this resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource Tags
	Tags map[string]string `pulumi:"tags"`
	// Resource Type
	Type *string `pulumi:"type"`
}

type PipelineState struct {
	// Configuration used to bootstrap the Pipeline.
	BootstrapConfiguration BootstrapConfigurationResponsePtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringPtrInput
	// Unique identifier of the Pipeline
	PipelineId pulumi.IntPtrInput
	// Specifies which CI/CD provider to use. Valid options are 'azurePipeline', 'githubWorkflow'.
	PipelineType pulumi.StringPtrInput
	// The system metadata pertaining to this resource.
	SystemData SystemDataResponsePtrInput
	// Resource Tags
	Tags pulumi.StringMapInput
	// Resource Type
	Type pulumi.StringPtrInput
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	// Configuration used to bootstrap the Pipeline.
	BootstrapConfiguration BootstrapConfiguration `pulumi:"bootstrapConfiguration"`
	// Resource Location
	Location *string `pulumi:"location"`
	// The name of the Pipeline resource in ARM.
	PipelineName *string `pulumi:"pipelineName"`
	// Specifies which CI/CD provider to use. Valid options are 'azurePipeline', 'githubWorkflow'.
	PipelineType string `pulumi:"pipelineType"`
	// Name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource Tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	// Configuration used to bootstrap the Pipeline.
	BootstrapConfiguration BootstrapConfigurationInput
	// Resource Location
	Location pulumi.StringPtrInput
	// The name of the Pipeline resource in ARM.
	PipelineName pulumi.StringPtrInput
	// Specifies which CI/CD provider to use. Valid options are 'azurePipeline', 'githubWorkflow'.
	PipelineType pulumi.StringInput
	// Name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Resource Tags
	Tags pulumi.StringMapInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

type PipelineInput interface {
	pulumi.Input

	ToPipelineOutput() PipelineOutput
	ToPipelineOutputWithContext(ctx context.Context) PipelineOutput
}

func (*Pipeline) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipeline)(nil))
}

func (i *Pipeline) ToPipelineOutput() PipelineOutput {
	return i.ToPipelineOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput)
}

type PipelineOutput struct {
	*pulumi.OutputState
}

func (PipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipeline)(nil))
}

func (o PipelineOutput) ToPipelineOutput() PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PipelineOutput{})
}
