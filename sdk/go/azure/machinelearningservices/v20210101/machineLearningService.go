// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Machine Learning service object wrapped into ARM resource envelope.
type MachineLearningService struct {
	pulumi.CustomResourceState

	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Service properties
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The sku of the workspace.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMachineLearningService registers a new resource with the given unique name, arguments, and options.
func NewMachineLearningService(ctx *pulumi.Context,
	name string, args *MachineLearningServiceArgs, opts ...pulumi.ResourceOption) (*MachineLearningService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeType == nil {
		return nil, errors.New("invalid value for required argument 'ComputeType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210101:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200501preview:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200515preview:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200901preview:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210401:MachineLearningService"),
		},
	})
	opts = append(opts, aliases)
	var resource MachineLearningService
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210101:MachineLearningService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineLearningService gets an existing MachineLearningService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineLearningService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineLearningServiceState, opts ...pulumi.ResourceOption) (*MachineLearningService, error) {
	var resource MachineLearningService
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210101:MachineLearningService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineLearningService resources.
type machineLearningServiceState struct {
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name *string `pulumi:"name"`
	// Service properties
	Properties interface{} `pulumi:"properties"`
	// The sku of the workspace.
	Sku *SkuResponse `pulumi:"sku"`
	// Read only system data
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type *string `pulumi:"type"`
}

type MachineLearningServiceState struct {
	// The identity of the resource.
	Identity IdentityResponsePtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Specifies the name of the resource.
	Name pulumi.StringPtrInput
	// Service properties
	Properties pulumi.Input
	// The sku of the workspace.
	Sku SkuResponsePtrInput
	// Read only system data
	SystemData SystemDataResponsePtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Specifies the type of the resource.
	Type pulumi.StringPtrInput
}

func (MachineLearningServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningServiceState)(nil)).Elem()
}

type machineLearningServiceArgs struct {
	// The compute environment type for the service.
	ComputeType string `pulumi:"computeType"`
	// The description of the service.
	Description *string `pulumi:"description"`
	// The Environment, models and assets needed for inferencing.
	EnvironmentImageRequest *CreateServiceRequestEnvironmentImageRequest `pulumi:"environmentImageRequest"`
	// The authentication keys.
	Keys *CreateServiceRequestKeys `pulumi:"keys"`
	// The service tag dictionary. Tags are mutable.
	KvTags map[string]string `pulumi:"kvTags"`
	// The name of the Azure location/region.
	Location *string `pulumi:"location"`
	// The service properties dictionary. Properties are immutable.
	Properties map[string]string `pulumi:"properties"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Azure Machine Learning service.
	ServiceName *string `pulumi:"serviceName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MachineLearningService resource.
type MachineLearningServiceArgs struct {
	// The compute environment type for the service.
	ComputeType pulumi.StringInput
	// The description of the service.
	Description pulumi.StringPtrInput
	// The Environment, models and assets needed for inferencing.
	EnvironmentImageRequest CreateServiceRequestEnvironmentImageRequestPtrInput
	// The authentication keys.
	Keys CreateServiceRequestKeysPtrInput
	// The service tag dictionary. Tags are mutable.
	KvTags pulumi.StringMapInput
	// The name of the Azure location/region.
	Location pulumi.StringPtrInput
	// The service properties dictionary. Properties are immutable.
	Properties pulumi.StringMapInput
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput
	// Name of the Azure Machine Learning service.
	ServiceName pulumi.StringPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (MachineLearningServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningServiceArgs)(nil)).Elem()
}

type MachineLearningServiceInput interface {
	pulumi.Input

	ToMachineLearningServiceOutput() MachineLearningServiceOutput
	ToMachineLearningServiceOutputWithContext(ctx context.Context) MachineLearningServiceOutput
}

func (*MachineLearningService) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningService)(nil))
}

func (i *MachineLearningService) ToMachineLearningServiceOutput() MachineLearningServiceOutput {
	return i.ToMachineLearningServiceOutputWithContext(context.Background())
}

func (i *MachineLearningService) ToMachineLearningServiceOutputWithContext(ctx context.Context) MachineLearningServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningServiceOutput)
}

type MachineLearningServiceOutput struct {
	*pulumi.OutputState
}

func (MachineLearningServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningService)(nil))
}

func (o MachineLearningServiceOutput) ToMachineLearningServiceOutput() MachineLearningServiceOutput {
	return o
}

func (o MachineLearningServiceOutput) ToMachineLearningServiceOutputWithContext(ctx context.Context) MachineLearningServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MachineLearningServiceOutput{})
}
