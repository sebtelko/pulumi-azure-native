// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration Assignment
type ConfigurationAssignmentParent struct {
	pulumi.CustomResourceState

	// Location of the resource
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The maintenance configuration Id
	MaintenanceConfigurationId pulumi.StringPtrOutput `pulumi:"maintenanceConfigurationId"`
	// Name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The unique resourceId
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationAssignmentParent registers a new resource with the given unique name, arguments, and options.
func NewConfigurationAssignmentParent(ctx *pulumi.Context,
	name string, args *ConfigurationAssignmentParentArgs, opts ...pulumi.ResourceOption) (*ConfigurationAssignmentParent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.ResourceParentName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceParentName'")
	}
	if args.ResourceParentType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceParentType'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:maintenance/v20210401preview:ConfigurationAssignmentParent"),
		},
		{
			Type: pulumi.String("azure-native:maintenance:ConfigurationAssignmentParent"),
		},
		{
			Type: pulumi.String("azure-nextgen:maintenance:ConfigurationAssignmentParent"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationAssignmentParent
	err := ctx.RegisterResource("azure-native:maintenance/v20210401preview:ConfigurationAssignmentParent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationAssignmentParent gets an existing ConfigurationAssignmentParent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationAssignmentParent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationAssignmentParentState, opts ...pulumi.ResourceOption) (*ConfigurationAssignmentParent, error) {
	var resource ConfigurationAssignmentParent
	err := ctx.ReadResource("azure-native:maintenance/v20210401preview:ConfigurationAssignmentParent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationAssignmentParent resources.
type configurationAssignmentParentState struct {
	// Location of the resource
	Location *string `pulumi:"location"`
	// The maintenance configuration Id
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// Name of the resource
	Name *string `pulumi:"name"`
	// The unique resourceId
	ResourceId *string `pulumi:"resourceId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Type of the resource
	Type *string `pulumi:"type"`
}

type ConfigurationAssignmentParentState struct {
	// Location of the resource
	Location pulumi.StringPtrInput
	// The maintenance configuration Id
	MaintenanceConfigurationId pulumi.StringPtrInput
	// Name of the resource
	Name pulumi.StringPtrInput
	// The unique resourceId
	ResourceId pulumi.StringPtrInput
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponsePtrInput
	// Type of the resource
	Type pulumi.StringPtrInput
}

func (ConfigurationAssignmentParentState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentParentState)(nil)).Elem()
}

type configurationAssignmentParentArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName *string `pulumi:"configurationAssignmentName"`
	// Location of the resource
	Location *string `pulumi:"location"`
	// The maintenance configuration Id
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// Resource provider name
	ProviderName string `pulumi:"providerName"`
	// Resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The unique resourceId
	ResourceId *string `pulumi:"resourceId"`
	// Resource identifier
	ResourceName string `pulumi:"resourceName"`
	// Resource parent identifier
	ResourceParentName string `pulumi:"resourceParentName"`
	// Resource parent type
	ResourceParentType string `pulumi:"resourceParentType"`
	// Resource type
	ResourceType string `pulumi:"resourceType"`
}

// The set of arguments for constructing a ConfigurationAssignmentParent resource.
type ConfigurationAssignmentParentArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName pulumi.StringPtrInput
	// Location of the resource
	Location pulumi.StringPtrInput
	// The maintenance configuration Id
	MaintenanceConfigurationId pulumi.StringPtrInput
	// Resource provider name
	ProviderName pulumi.StringInput
	// Resource group name
	ResourceGroupName pulumi.StringInput
	// The unique resourceId
	ResourceId pulumi.StringPtrInput
	// Resource identifier
	ResourceName pulumi.StringInput
	// Resource parent identifier
	ResourceParentName pulumi.StringInput
	// Resource parent type
	ResourceParentType pulumi.StringInput
	// Resource type
	ResourceType pulumi.StringInput
}

func (ConfigurationAssignmentParentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentParentArgs)(nil)).Elem()
}

type ConfigurationAssignmentParentInput interface {
	pulumi.Input

	ToConfigurationAssignmentParentOutput() ConfigurationAssignmentParentOutput
	ToConfigurationAssignmentParentOutputWithContext(ctx context.Context) ConfigurationAssignmentParentOutput
}

func (*ConfigurationAssignmentParent) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationAssignmentParent)(nil))
}

func (i *ConfigurationAssignmentParent) ToConfigurationAssignmentParentOutput() ConfigurationAssignmentParentOutput {
	return i.ToConfigurationAssignmentParentOutputWithContext(context.Background())
}

func (i *ConfigurationAssignmentParent) ToConfigurationAssignmentParentOutputWithContext(ctx context.Context) ConfigurationAssignmentParentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationAssignmentParentOutput)
}

type ConfigurationAssignmentParentOutput struct {
	*pulumi.OutputState
}

func (ConfigurationAssignmentParentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationAssignmentParent)(nil))
}

func (o ConfigurationAssignmentParentOutput) ToConfigurationAssignmentParentOutput() ConfigurationAssignmentParentOutput {
	return o
}

func (o ConfigurationAssignmentParentOutput) ToConfigurationAssignmentParentOutputWithContext(ctx context.Context) ConfigurationAssignmentParentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigurationAssignmentParentOutput{})
}
