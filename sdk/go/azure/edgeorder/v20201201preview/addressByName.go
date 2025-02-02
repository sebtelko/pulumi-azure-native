// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Address Resource.
type AddressByName struct {
	pulumi.CustomResourceState

	// Contact details for the address
	ContactDetails ContactDetailsResponseOutput `pulumi:"contactDetails"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Shipping details for the address
	ShippingAddress ShippingAddressResponsePtrOutput `pulumi:"shippingAddress"`
	// Represents resource creation and update time
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAddressByName registers a new resource with the given unique name, arguments, and options.
func NewAddressByName(ctx *pulumi.Context,
	name string, args *AddressByNameArgs, opts ...pulumi.ResourceOption) (*AddressByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContactDetails == nil {
		return nil, errors.New("invalid value for required argument 'ContactDetails'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:edgeorder/v20201201preview:AddressByName"),
		},
		{
			Type: pulumi.String("azure-native:edgeorder:AddressByName"),
		},
		{
			Type: pulumi.String("azure-nextgen:edgeorder:AddressByName"),
		},
	})
	opts = append(opts, aliases)
	var resource AddressByName
	err := ctx.RegisterResource("azure-native:edgeorder/v20201201preview:AddressByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddressByName gets an existing AddressByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddressByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressByNameState, opts ...pulumi.ResourceOption) (*AddressByName, error) {
	var resource AddressByName
	err := ctx.ReadResource("azure-native:edgeorder/v20201201preview:AddressByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AddressByName resources.
type addressByNameState struct {
	// Contact details for the address
	ContactDetails *ContactDetailsResponse `pulumi:"contactDetails"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Shipping details for the address
	ShippingAddress *ShippingAddressResponse `pulumi:"shippingAddress"`
	// Represents resource creation and update time
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type AddressByNameState struct {
	// Contact details for the address
	ContactDetails ContactDetailsResponsePtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Shipping details for the address
	ShippingAddress ShippingAddressResponsePtrInput
	// Represents resource creation and update time
	SystemData SystemDataResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (AddressByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressByNameState)(nil)).Elem()
}

type addressByNameArgs struct {
	// The name of the address Resource within the specified resource group. address names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	AddressName *string `pulumi:"addressName"`
	// Contact details for the address
	ContactDetails ContactDetails `pulumi:"contactDetails"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Shipping details for the address
	ShippingAddress *ShippingAddress `pulumi:"shippingAddress"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AddressByName resource.
type AddressByNameArgs struct {
	// The name of the address Resource within the specified resource group. address names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	AddressName pulumi.StringPtrInput
	// Contact details for the address
	ContactDetails ContactDetailsInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Shipping details for the address
	ShippingAddress ShippingAddressPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (AddressByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressByNameArgs)(nil)).Elem()
}

type AddressByNameInput interface {
	pulumi.Input

	ToAddressByNameOutput() AddressByNameOutput
	ToAddressByNameOutputWithContext(ctx context.Context) AddressByNameOutput
}

func (*AddressByName) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressByName)(nil))
}

func (i *AddressByName) ToAddressByNameOutput() AddressByNameOutput {
	return i.ToAddressByNameOutputWithContext(context.Background())
}

func (i *AddressByName) ToAddressByNameOutputWithContext(ctx context.Context) AddressByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressByNameOutput)
}

type AddressByNameOutput struct {
	*pulumi.OutputState
}

func (AddressByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressByName)(nil))
}

func (o AddressByNameOutput) ToAddressByNameOutput() AddressByNameOutput {
	return o
}

func (o AddressByNameOutput) ToAddressByNameOutputWithContext(ctx context.Context) AddressByNameOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AddressByNameOutput{})
}
