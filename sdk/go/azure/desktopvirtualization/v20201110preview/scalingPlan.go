// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201110preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a scaling plan definition.
type ScalingPlan struct {
	pulumi.CustomResourceState

	// Description of scaling plan.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Exclusion tag for scaling plan.
	ExclusionTag pulumi.StringPtrOutput `pulumi:"exclusionTag"`
	// User friendly name of scaling plan.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// List of ScalingHostPoolReference definitions.
	HostPoolReferences ScalingHostPoolReferenceResponseArrayOutput `pulumi:"hostPoolReferences"`
	// HostPool type for scaling plan.
	HostPoolType pulumi.StringPtrOutput `pulumi:"hostPoolType"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of ScalingSchedule definitions.
	Schedules ScalingScheduleResponseArrayOutput `pulumi:"schedules"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Timezone of the scaling plan.
	TimeZone pulumi.StringPtrOutput `pulumi:"timeZone"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewScalingPlan registers a new resource with the given unique name, arguments, and options.
func NewScalingPlan(ctx *pulumi.Context,
	name string, args *ScalingPlanArgs, opts ...pulumi.ResourceOption) (*ScalingPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20201110preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210114preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210201preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210309preview:ScalingPlan"),
		},
	})
	opts = append(opts, aliases)
	var resource ScalingPlan
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20201110preview:ScalingPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingPlan gets an existing ScalingPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingPlanState, opts ...pulumi.ResourceOption) (*ScalingPlan, error) {
	var resource ScalingPlan
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20201110preview:ScalingPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingPlan resources.
type scalingPlanState struct {
	// Description of scaling plan.
	Description *string `pulumi:"description"`
	// Exclusion tag for scaling plan.
	ExclusionTag *string `pulumi:"exclusionTag"`
	// User friendly name of scaling plan.
	FriendlyName *string `pulumi:"friendlyName"`
	// List of ScalingHostPoolReference definitions.
	HostPoolReferences []ScalingHostPoolReferenceResponse `pulumi:"hostPoolReferences"`
	// HostPool type for scaling plan.
	HostPoolType *string `pulumi:"hostPoolType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// List of ScalingSchedule definitions.
	Schedules []ScalingScheduleResponse `pulumi:"schedules"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Timezone of the scaling plan.
	TimeZone *string `pulumi:"timeZone"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type ScalingPlanState struct {
	// Description of scaling plan.
	Description pulumi.StringPtrInput
	// Exclusion tag for scaling plan.
	ExclusionTag pulumi.StringPtrInput
	// User friendly name of scaling plan.
	FriendlyName pulumi.StringPtrInput
	// List of ScalingHostPoolReference definitions.
	HostPoolReferences ScalingHostPoolReferenceResponseArrayInput
	// HostPool type for scaling plan.
	HostPoolType pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// List of ScalingSchedule definitions.
	Schedules ScalingScheduleResponseArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Timezone of the scaling plan.
	TimeZone pulumi.StringPtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (ScalingPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanState)(nil)).Elem()
}

type scalingPlanArgs struct {
	// Description of scaling plan.
	Description *string `pulumi:"description"`
	// Exclusion tag for scaling plan.
	ExclusionTag *string `pulumi:"exclusionTag"`
	// User friendly name of scaling plan.
	FriendlyName *string `pulumi:"friendlyName"`
	// List of ScalingHostPoolReference definitions.
	HostPoolReferences []ScalingHostPoolReference `pulumi:"hostPoolReferences"`
	// HostPool type for scaling plan.
	HostPoolType *string `pulumi:"hostPoolType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the scaling plan.
	ScalingPlanName *string `pulumi:"scalingPlanName"`
	// List of ScalingSchedule definitions.
	Schedules []ScalingSchedule `pulumi:"schedules"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Timezone of the scaling plan.
	TimeZone *string `pulumi:"timeZone"`
}

// The set of arguments for constructing a ScalingPlan resource.
type ScalingPlanArgs struct {
	// Description of scaling plan.
	Description pulumi.StringPtrInput
	// Exclusion tag for scaling plan.
	ExclusionTag pulumi.StringPtrInput
	// User friendly name of scaling plan.
	FriendlyName pulumi.StringPtrInput
	// List of ScalingHostPoolReference definitions.
	HostPoolReferences ScalingHostPoolReferenceArrayInput
	// HostPool type for scaling plan.
	HostPoolType pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the scaling plan.
	ScalingPlanName pulumi.StringPtrInput
	// List of ScalingSchedule definitions.
	Schedules ScalingScheduleArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Timezone of the scaling plan.
	TimeZone pulumi.StringPtrInput
}

func (ScalingPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanArgs)(nil)).Elem()
}

type ScalingPlanInput interface {
	pulumi.Input

	ToScalingPlanOutput() ScalingPlanOutput
	ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput
}

func (*ScalingPlan) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingPlan)(nil))
}

func (i *ScalingPlan) ToScalingPlanOutput() ScalingPlanOutput {
	return i.ToScalingPlanOutputWithContext(context.Background())
}

func (i *ScalingPlan) ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanOutput)
}

type ScalingPlanOutput struct {
	*pulumi.OutputState
}

func (ScalingPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingPlan)(nil))
}

func (o ScalingPlanOutput) ToScalingPlanOutput() ScalingPlanOutput {
	return o
}

func (o ScalingPlanOutput) ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScalingPlanOutput{})
}
