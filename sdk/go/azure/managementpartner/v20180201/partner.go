// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// this is the management partner operations response
type Partner struct {
	pulumi.CustomResourceState

	// This is the DateTime when the partner was created.
	CreatedTime pulumi.StringPtrOutput `pulumi:"createdTime"`
	// Type of the partner
	Etag pulumi.IntPtrOutput `pulumi:"etag"`
	// Name of the partner
	Name pulumi.StringOutput `pulumi:"name"`
	// This is the object id.
	ObjectId pulumi.StringPtrOutput `pulumi:"objectId"`
	// This is the partner id
	PartnerId pulumi.StringPtrOutput `pulumi:"partnerId"`
	// This is the partner name
	PartnerName pulumi.StringPtrOutput `pulumi:"partnerName"`
	// This is the tenant id.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// Type of resource. "Microsoft.ManagementPartner/partners"
	Type pulumi.StringOutput `pulumi:"type"`
	// This is the DateTime when the partner was updated.
	UpdatedTime pulumi.StringPtrOutput `pulumi:"updatedTime"`
	// This is the version.
	Version pulumi.IntPtrOutput `pulumi:"version"`
}

// NewPartner registers a new resource with the given unique name, arguments, and options.
func NewPartner(ctx *pulumi.Context,
	name string, args *PartnerArgs, opts ...pulumi.ResourceOption) (*Partner, error) {
	if args == nil {
		args = &PartnerArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:managementpartner/v20180201:Partner"),
		},
		{
			Type: pulumi.String("azure-native:managementpartner:Partner"),
		},
		{
			Type: pulumi.String("azure-nextgen:managementpartner:Partner"),
		},
	})
	opts = append(opts, aliases)
	var resource Partner
	err := ctx.RegisterResource("azure-native:managementpartner/v20180201:Partner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPartner gets an existing Partner resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPartner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerState, opts ...pulumi.ResourceOption) (*Partner, error) {
	var resource Partner
	err := ctx.ReadResource("azure-native:managementpartner/v20180201:Partner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Partner resources.
type partnerState struct {
	// This is the DateTime when the partner was created.
	CreatedTime *string `pulumi:"createdTime"`
	// Type of the partner
	Etag *int `pulumi:"etag"`
	// Name of the partner
	Name *string `pulumi:"name"`
	// This is the object id.
	ObjectId *string `pulumi:"objectId"`
	// This is the partner id
	PartnerId *string `pulumi:"partnerId"`
	// This is the partner name
	PartnerName *string `pulumi:"partnerName"`
	// This is the tenant id.
	TenantId *string `pulumi:"tenantId"`
	// Type of resource. "Microsoft.ManagementPartner/partners"
	Type *string `pulumi:"type"`
	// This is the DateTime when the partner was updated.
	UpdatedTime *string `pulumi:"updatedTime"`
	// This is the version.
	Version *int `pulumi:"version"`
}

type PartnerState struct {
	// This is the DateTime when the partner was created.
	CreatedTime pulumi.StringPtrInput
	// Type of the partner
	Etag pulumi.IntPtrInput
	// Name of the partner
	Name pulumi.StringPtrInput
	// This is the object id.
	ObjectId pulumi.StringPtrInput
	// This is the partner id
	PartnerId pulumi.StringPtrInput
	// This is the partner name
	PartnerName pulumi.StringPtrInput
	// This is the tenant id.
	TenantId pulumi.StringPtrInput
	// Type of resource. "Microsoft.ManagementPartner/partners"
	Type pulumi.StringPtrInput
	// This is the DateTime when the partner was updated.
	UpdatedTime pulumi.StringPtrInput
	// This is the version.
	Version pulumi.IntPtrInput
}

func (PartnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerState)(nil)).Elem()
}

type partnerArgs struct {
	// Id of the Partner
	PartnerId *string `pulumi:"partnerId"`
}

// The set of arguments for constructing a Partner resource.
type PartnerArgs struct {
	// Id of the Partner
	PartnerId pulumi.StringPtrInput
}

func (PartnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerArgs)(nil)).Elem()
}

type PartnerInput interface {
	pulumi.Input

	ToPartnerOutput() PartnerOutput
	ToPartnerOutputWithContext(ctx context.Context) PartnerOutput
}

func (*Partner) ElementType() reflect.Type {
	return reflect.TypeOf((*Partner)(nil))
}

func (i *Partner) ToPartnerOutput() PartnerOutput {
	return i.ToPartnerOutputWithContext(context.Background())
}

func (i *Partner) ToPartnerOutputWithContext(ctx context.Context) PartnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerOutput)
}

type PartnerOutput struct {
	*pulumi.OutputState
}

func (PartnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Partner)(nil))
}

func (o PartnerOutput) ToPartnerOutput() PartnerOutput {
	return o
}

func (o PartnerOutput) ToPartnerOutputWithContext(ctx context.Context) PartnerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PartnerOutput{})
}
