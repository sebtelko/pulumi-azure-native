// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Domain ownership Identifier.
type DomainOwnershipIdentifier struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Ownership Id.
	OwnershipId pulumi.StringPtrOutput `pulumi:"ownershipId"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDomainOwnershipIdentifier registers a new resource with the given unique name, arguments, and options.
func NewDomainOwnershipIdentifier(ctx *pulumi.Context,
	name string, args *DomainOwnershipIdentifierArgs, opts ...pulumi.ResourceOption) (*DomainOwnershipIdentifier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:domainregistration/v20200901:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:domainregistration:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20150401:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:domainregistration/v20150401:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20180201:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:domainregistration/v20180201:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20190801:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:domainregistration/v20190801:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20200601:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:domainregistration/v20200601:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20201001:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:domainregistration/v20201001:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20201201:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-nextgen:domainregistration/v20201201:DomainOwnershipIdentifier"),
		},
	})
	opts = append(opts, aliases)
	var resource DomainOwnershipIdentifier
	err := ctx.RegisterResource("azure-native:domainregistration/v20200901:DomainOwnershipIdentifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainOwnershipIdentifier gets an existing DomainOwnershipIdentifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainOwnershipIdentifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainOwnershipIdentifierState, opts ...pulumi.ResourceOption) (*DomainOwnershipIdentifier, error) {
	var resource DomainOwnershipIdentifier
	err := ctx.ReadResource("azure-native:domainregistration/v20200901:DomainOwnershipIdentifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainOwnershipIdentifier resources.
type domainOwnershipIdentifierState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// Ownership Id.
	OwnershipId *string `pulumi:"ownershipId"`
	// The system metadata relating to this resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type DomainOwnershipIdentifierState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// Ownership Id.
	OwnershipId pulumi.StringPtrInput
	// The system metadata relating to this resource.
	SystemData SystemDataResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (DomainOwnershipIdentifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainOwnershipIdentifierState)(nil)).Elem()
}

type domainOwnershipIdentifierArgs struct {
	// Name of domain.
	DomainName string `pulumi:"domainName"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of identifier.
	Name *string `pulumi:"name"`
	// Ownership Id.
	OwnershipId *string `pulumi:"ownershipId"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DomainOwnershipIdentifier resource.
type DomainOwnershipIdentifierArgs struct {
	// Name of domain.
	DomainName pulumi.StringInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of identifier.
	Name pulumi.StringPtrInput
	// Ownership Id.
	OwnershipId pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (DomainOwnershipIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainOwnershipIdentifierArgs)(nil)).Elem()
}

type DomainOwnershipIdentifierInput interface {
	pulumi.Input

	ToDomainOwnershipIdentifierOutput() DomainOwnershipIdentifierOutput
	ToDomainOwnershipIdentifierOutputWithContext(ctx context.Context) DomainOwnershipIdentifierOutput
}

func (*DomainOwnershipIdentifier) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainOwnershipIdentifier)(nil))
}

func (i *DomainOwnershipIdentifier) ToDomainOwnershipIdentifierOutput() DomainOwnershipIdentifierOutput {
	return i.ToDomainOwnershipIdentifierOutputWithContext(context.Background())
}

func (i *DomainOwnershipIdentifier) ToDomainOwnershipIdentifierOutputWithContext(ctx context.Context) DomainOwnershipIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOwnershipIdentifierOutput)
}

type DomainOwnershipIdentifierOutput struct {
	*pulumi.OutputState
}

func (DomainOwnershipIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainOwnershipIdentifier)(nil))
}

func (o DomainOwnershipIdentifierOutput) ToDomainOwnershipIdentifierOutput() DomainOwnershipIdentifierOutput {
	return o
}

func (o DomainOwnershipIdentifierOutput) ToDomainOwnershipIdentifierOutputWithContext(ctx context.Context) DomainOwnershipIdentifierOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainOwnershipIdentifierOutput{})
}
