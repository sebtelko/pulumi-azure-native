// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JitNetworkAccessPolicy struct {
	pulumi.CustomResourceState

	// Kind of the resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Location where the resource is stored
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the provisioning state of the Just-in-Time policy.
	ProvisioningState pulumi.StringOutput                        `pulumi:"provisioningState"`
	Requests          JitNetworkAccessRequestResponseArrayOutput `pulumi:"requests"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines JitNetworkAccessPolicyVirtualMachineResponseArrayOutput `pulumi:"virtualMachines"`
}

// NewJitNetworkAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewJitNetworkAccessPolicy(ctx *pulumi.Context,
	name string, args *JitNetworkAccessPolicyArgs, opts ...pulumi.ResourceOption) (*JitNetworkAccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AscLocation == nil {
		return nil, errors.New("invalid value for required argument 'AscLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualMachines == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachines'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:security/v20200101:JitNetworkAccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:security:JitNetworkAccessPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:security:JitNetworkAccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:security/v20150601preview:JitNetworkAccessPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:security/v20150601preview:JitNetworkAccessPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource JitNetworkAccessPolicy
	err := ctx.RegisterResource("azure-native:security/v20200101:JitNetworkAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJitNetworkAccessPolicy gets an existing JitNetworkAccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJitNetworkAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JitNetworkAccessPolicyState, opts ...pulumi.ResourceOption) (*JitNetworkAccessPolicy, error) {
	var resource JitNetworkAccessPolicy
	err := ctx.ReadResource("azure-native:security/v20200101:JitNetworkAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JitNetworkAccessPolicy resources.
type jitNetworkAccessPolicyState struct {
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Location where the resource is stored
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Gets the provisioning state of the Just-in-Time policy.
	ProvisioningState *string                           `pulumi:"provisioningState"`
	Requests          []JitNetworkAccessRequestResponse `pulumi:"requests"`
	// Resource type
	Type *string `pulumi:"type"`
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines []JitNetworkAccessPolicyVirtualMachineResponse `pulumi:"virtualMachines"`
}

type JitNetworkAccessPolicyState struct {
	// Kind of the resource
	Kind pulumi.StringPtrInput
	// Location where the resource is stored
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Gets the provisioning state of the Just-in-Time policy.
	ProvisioningState pulumi.StringPtrInput
	Requests          JitNetworkAccessRequestResponseArrayInput
	// Resource type
	Type pulumi.StringPtrInput
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines JitNetworkAccessPolicyVirtualMachineResponseArrayInput
}

func (JitNetworkAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*jitNetworkAccessPolicyState)(nil)).Elem()
}

type jitNetworkAccessPolicyArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation string `pulumi:"ascLocation"`
	// Name of a Just-in-Time access configuration policy.
	JitNetworkAccessPolicyName *string `pulumi:"jitNetworkAccessPolicyName"`
	// Kind of the resource
	Kind     *string                   `pulumi:"kind"`
	Requests []JitNetworkAccessRequest `pulumi:"requests"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines []JitNetworkAccessPolicyVirtualMachine `pulumi:"virtualMachines"`
}

// The set of arguments for constructing a JitNetworkAccessPolicy resource.
type JitNetworkAccessPolicyArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation pulumi.StringInput
	// Name of a Just-in-Time access configuration policy.
	JitNetworkAccessPolicyName pulumi.StringPtrInput
	// Kind of the resource
	Kind     pulumi.StringPtrInput
	Requests JitNetworkAccessRequestArrayInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines JitNetworkAccessPolicyVirtualMachineArrayInput
}

func (JitNetworkAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jitNetworkAccessPolicyArgs)(nil)).Elem()
}

type JitNetworkAccessPolicyInput interface {
	pulumi.Input

	ToJitNetworkAccessPolicyOutput() JitNetworkAccessPolicyOutput
	ToJitNetworkAccessPolicyOutputWithContext(ctx context.Context) JitNetworkAccessPolicyOutput
}

func (*JitNetworkAccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPolicy)(nil))
}

func (i *JitNetworkAccessPolicy) ToJitNetworkAccessPolicyOutput() JitNetworkAccessPolicyOutput {
	return i.ToJitNetworkAccessPolicyOutputWithContext(context.Background())
}

func (i *JitNetworkAccessPolicy) ToJitNetworkAccessPolicyOutputWithContext(ctx context.Context) JitNetworkAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPolicyOutput)
}

type JitNetworkAccessPolicyOutput struct {
	*pulumi.OutputState
}

func (JitNetworkAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPolicy)(nil))
}

func (o JitNetworkAccessPolicyOutput) ToJitNetworkAccessPolicyOutput() JitNetworkAccessPolicyOutput {
	return o
}

func (o JitNetworkAccessPolicyOutput) ToJitNetworkAccessPolicyOutputWithContext(ctx context.Context) JitNetworkAccessPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JitNetworkAccessPolicyOutput{})
}
