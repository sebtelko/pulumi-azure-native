// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package providerhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API Version: 2020-11-20.
type ProviderRegistration struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name       pulumi.StringOutput                          `pulumi:"name"`
	Properties ProviderRegistrationResponsePropertiesOutput `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProviderRegistration registers a new resource with the given unique name, arguments, and options.
func NewProviderRegistration(ctx *pulumi.Context,
	name string, args *ProviderRegistrationArgs, opts ...pulumi.ResourceOption) (*ProviderRegistration, error) {
	if args == nil {
		args = &ProviderRegistrationArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:providerhub:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20201120:ProviderRegistration"),
		},
	})
	opts = append(opts, aliases)
	var resource ProviderRegistration
	err := ctx.RegisterResource("azure-native:providerhub:ProviderRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderRegistration gets an existing ProviderRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderRegistrationState, opts ...pulumi.ResourceOption) (*ProviderRegistration, error) {
	var resource ProviderRegistration
	err := ctx.ReadResource("azure-native:providerhub:ProviderRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderRegistration resources.
type providerRegistrationState struct {
	// The name of the resource
	Name       *string                                 `pulumi:"name"`
	Properties *ProviderRegistrationResponseProperties `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type ProviderRegistrationState struct {
	// The name of the resource
	Name       pulumi.StringPtrInput
	Properties ProviderRegistrationResponsePropertiesPtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (ProviderRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerRegistrationState)(nil)).Elem()
}

type providerRegistrationArgs struct {
	Properties *ProviderRegistrationProperties `pulumi:"properties"`
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace *string `pulumi:"providerNamespace"`
}

// The set of arguments for constructing a ProviderRegistration resource.
type ProviderRegistrationArgs struct {
	Properties ProviderRegistrationPropertiesPtrInput
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringPtrInput
}

func (ProviderRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerRegistrationArgs)(nil)).Elem()
}

type ProviderRegistrationInput interface {
	pulumi.Input

	ToProviderRegistrationOutput() ProviderRegistrationOutput
	ToProviderRegistrationOutputWithContext(ctx context.Context) ProviderRegistrationOutput
}

func (*ProviderRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderRegistration)(nil))
}

func (i *ProviderRegistration) ToProviderRegistrationOutput() ProviderRegistrationOutput {
	return i.ToProviderRegistrationOutputWithContext(context.Background())
}

func (i *ProviderRegistration) ToProviderRegistrationOutputWithContext(ctx context.Context) ProviderRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderRegistrationOutput)
}

type ProviderRegistrationOutput struct {
	*pulumi.OutputState
}

func (ProviderRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderRegistration)(nil))
}

func (o ProviderRegistrationOutput) ToProviderRegistrationOutput() ProviderRegistrationOutput {
	return o
}

func (o ProviderRegistrationOutput) ToProviderRegistrationOutputWithContext(ctx context.Context) ProviderRegistrationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderRegistrationOutput{})
}
