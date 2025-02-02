// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181120

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Guest configuration assignment is an association between a machine and guest configuration.
type GuestConfigurationAssignment struct {
	pulumi.CustomResourceState

	// Region where the VM is located.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGuestConfigurationAssignment registers a new resource with the given unique name, arguments, and options.
func NewGuestConfigurationAssignment(ctx *pulumi.Context,
	name string, args *GuestConfigurationAssignmentArgs, opts ...pulumi.ResourceOption) (*GuestConfigurationAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmName == nil {
		return nil, errors.New("invalid value for required argument 'VmName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:guestconfiguration/v20181120:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:guestconfiguration:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20180630preview:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:guestconfiguration/v20180630preview:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20200625:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:guestconfiguration/v20200625:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20210125:GuestConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:guestconfiguration/v20210125:GuestConfigurationAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource GuestConfigurationAssignment
	err := ctx.RegisterResource("azure-native:guestconfiguration/v20181120:GuestConfigurationAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGuestConfigurationAssignment gets an existing GuestConfigurationAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGuestConfigurationAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestConfigurationAssignmentState, opts ...pulumi.ResourceOption) (*GuestConfigurationAssignment, error) {
	var resource GuestConfigurationAssignment
	err := ctx.ReadResource("azure-native:guestconfiguration/v20181120:GuestConfigurationAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GuestConfigurationAssignment resources.
type guestConfigurationAssignmentState struct {
	// Region where the VM is located.
	Location *string `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name *string `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties *GuestConfigurationAssignmentPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type GuestConfigurationAssignmentState struct {
	// Region where the VM is located.
	Location pulumi.StringPtrInput
	// Name of the guest configuration assignment.
	Name pulumi.StringPtrInput
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (GuestConfigurationAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationAssignmentState)(nil)).Elem()
}

type guestConfigurationAssignmentArgs struct {
	// Name of the guest configuration assignment.
	GuestConfigurationAssignmentName *string `pulumi:"guestConfigurationAssignmentName"`
	// Region where the VM is located.
	Location *string `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name *string `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties *GuestConfigurationAssignmentProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual machine.
	VmName string `pulumi:"vmName"`
}

// The set of arguments for constructing a GuestConfigurationAssignment resource.
type GuestConfigurationAssignmentArgs struct {
	// Name of the guest configuration assignment.
	GuestConfigurationAssignmentName pulumi.StringPtrInput
	// Region where the VM is located.
	Location pulumi.StringPtrInput
	// Name of the guest configuration assignment.
	Name pulumi.StringPtrInput
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the virtual machine.
	VmName pulumi.StringInput
}

func (GuestConfigurationAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationAssignmentArgs)(nil)).Elem()
}

type GuestConfigurationAssignmentInput interface {
	pulumi.Input

	ToGuestConfigurationAssignmentOutput() GuestConfigurationAssignmentOutput
	ToGuestConfigurationAssignmentOutputWithContext(ctx context.Context) GuestConfigurationAssignmentOutput
}

func (*GuestConfigurationAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationAssignment)(nil))
}

func (i *GuestConfigurationAssignment) ToGuestConfigurationAssignmentOutput() GuestConfigurationAssignmentOutput {
	return i.ToGuestConfigurationAssignmentOutputWithContext(context.Background())
}

func (i *GuestConfigurationAssignment) ToGuestConfigurationAssignmentOutputWithContext(ctx context.Context) GuestConfigurationAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentOutput)
}

type GuestConfigurationAssignmentOutput struct {
	*pulumi.OutputState
}

func (GuestConfigurationAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationAssignment)(nil))
}

func (o GuestConfigurationAssignmentOutput) ToGuestConfigurationAssignmentOutput() GuestConfigurationAssignmentOutput {
	return o
}

func (o GuestConfigurationAssignmentOutput) ToGuestConfigurationAssignmentOutputWithContext(ctx context.Context) GuestConfigurationAssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GuestConfigurationAssignmentOutput{})
}
