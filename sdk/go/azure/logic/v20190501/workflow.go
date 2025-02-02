// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The workflow type.
type Workflow struct {
	pulumi.CustomResourceState

	// The access control configuration.
	AccessControl FlowAccessControlConfigurationResponsePtrOutput `pulumi:"accessControl"`
	// Gets the access endpoint.
	AccessEndpoint pulumi.StringOutput `pulumi:"accessEndpoint"`
	// Gets the changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// Gets the created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The definition.
	Definition pulumi.AnyOutput `pulumi:"definition"`
	// The endpoints configuration.
	EndpointsConfiguration FlowEndpointsConfigurationResponsePtrOutput `pulumi:"endpointsConfiguration"`
	// The integration account.
	IntegrationAccount ResourceReferenceResponsePtrOutput `pulumi:"integrationAccount"`
	// The integration service environment.
	IntegrationServiceEnvironment ResourceReferenceResponsePtrOutput `pulumi:"integrationServiceEnvironment"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters.
	Parameters WorkflowParameterResponseMapOutput `pulumi:"parameters"`
	// Gets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The sku.
	Sku SkuResponseOutput `pulumi:"sku"`
	// The state.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets the version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewWorkflow registers a new resource with the given unique name, arguments, and options.
func NewWorkflow(ctx *pulumi.Context,
	name string, args *WorkflowArgs, opts ...pulumi.ResourceOption) (*Workflow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic:Workflow"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20150201preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20150201preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:Workflow"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20160601:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:Workflow"),
		},
	})
	opts = append(opts, aliases)
	var resource Workflow
	err := ctx.RegisterResource("azure-native:logic/v20190501:Workflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflow gets an existing Workflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowState, opts ...pulumi.ResourceOption) (*Workflow, error) {
	var resource Workflow
	err := ctx.ReadResource("azure-native:logic/v20190501:Workflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workflow resources.
type workflowState struct {
	// The access control configuration.
	AccessControl *FlowAccessControlConfigurationResponse `pulumi:"accessControl"`
	// Gets the access endpoint.
	AccessEndpoint *string `pulumi:"accessEndpoint"`
	// Gets the changed time.
	ChangedTime *string `pulumi:"changedTime"`
	// Gets the created time.
	CreatedTime *string `pulumi:"createdTime"`
	// The definition.
	Definition interface{} `pulumi:"definition"`
	// The endpoints configuration.
	EndpointsConfiguration *FlowEndpointsConfigurationResponse `pulumi:"endpointsConfiguration"`
	// The integration account.
	IntegrationAccount *ResourceReferenceResponse `pulumi:"integrationAccount"`
	// The integration service environment.
	IntegrationServiceEnvironment *ResourceReferenceResponse `pulumi:"integrationServiceEnvironment"`
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The parameters.
	Parameters map[string]WorkflowParameterResponse `pulumi:"parameters"`
	// Gets the provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The sku.
	Sku *SkuResponse `pulumi:"sku"`
	// The state.
	State *string `pulumi:"state"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
	// Gets the version.
	Version *string `pulumi:"version"`
}

type WorkflowState struct {
	// The access control configuration.
	AccessControl FlowAccessControlConfigurationResponsePtrInput
	// Gets the access endpoint.
	AccessEndpoint pulumi.StringPtrInput
	// Gets the changed time.
	ChangedTime pulumi.StringPtrInput
	// Gets the created time.
	CreatedTime pulumi.StringPtrInput
	// The definition.
	Definition pulumi.Input
	// The endpoints configuration.
	EndpointsConfiguration FlowEndpointsConfigurationResponsePtrInput
	// The integration account.
	IntegrationAccount ResourceReferenceResponsePtrInput
	// The integration service environment.
	IntegrationServiceEnvironment ResourceReferenceResponsePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The parameters.
	Parameters WorkflowParameterResponseMapInput
	// Gets the provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// The sku.
	Sku SkuResponsePtrInput
	// The state.
	State pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
	// Gets the version.
	Version pulumi.StringPtrInput
}

func (WorkflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowState)(nil)).Elem()
}

type workflowArgs struct {
	// The access control configuration.
	AccessControl *FlowAccessControlConfiguration `pulumi:"accessControl"`
	// The definition.
	Definition interface{} `pulumi:"definition"`
	// The endpoints configuration.
	EndpointsConfiguration *FlowEndpointsConfiguration `pulumi:"endpointsConfiguration"`
	// The integration account.
	IntegrationAccount *ResourceReference `pulumi:"integrationAccount"`
	// The integration service environment.
	IntegrationServiceEnvironment *ResourceReference `pulumi:"integrationServiceEnvironment"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The parameters.
	Parameters map[string]WorkflowParameter `pulumi:"parameters"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The state.
	State *string `pulumi:"state"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The workflow name.
	WorkflowName *string `pulumi:"workflowName"`
}

// The set of arguments for constructing a Workflow resource.
type WorkflowArgs struct {
	// The access control configuration.
	AccessControl FlowAccessControlConfigurationPtrInput
	// The definition.
	Definition pulumi.Input
	// The endpoints configuration.
	EndpointsConfiguration FlowEndpointsConfigurationPtrInput
	// The integration account.
	IntegrationAccount ResourceReferencePtrInput
	// The integration service environment.
	IntegrationServiceEnvironment ResourceReferencePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The parameters.
	Parameters WorkflowParameterMapInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The state.
	State pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The workflow name.
	WorkflowName pulumi.StringPtrInput
}

func (WorkflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowArgs)(nil)).Elem()
}

type WorkflowInput interface {
	pulumi.Input

	ToWorkflowOutput() WorkflowOutput
	ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput
}

func (*Workflow) ElementType() reflect.Type {
	return reflect.TypeOf((*Workflow)(nil))
}

func (i *Workflow) ToWorkflowOutput() WorkflowOutput {
	return i.ToWorkflowOutputWithContext(context.Background())
}

func (i *Workflow) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowOutput)
}

type WorkflowOutput struct {
	*pulumi.OutputState
}

func (WorkflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Workflow)(nil))
}

func (o WorkflowOutput) ToWorkflowOutput() WorkflowOutput {
	return o
}

func (o WorkflowOutput) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkflowOutput{})
}
