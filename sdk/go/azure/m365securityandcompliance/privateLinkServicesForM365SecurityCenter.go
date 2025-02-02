// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package m365securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The description of the service.
// API Version: 2021-03-25-preview.
type PrivateLinkServicesForM365SecurityCenter struct {
	pulumi.CustomResourceState

	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServicesResourceResponseIdentityPtrOutput `pulumi:"identity"`
	// The kind of the service.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The common properties of a service.
	Properties ServicesPropertiesResponseOutput `pulumi:"properties"`
	// Required property for system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateLinkServicesForM365SecurityCenter registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkServicesForM365SecurityCenter(ctx *pulumi.Context,
	name string, args *PrivateLinkServicesForM365SecurityCenterArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForM365SecurityCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:m365securityandcompliance:privateLinkServicesForM365SecurityCenter"),
		},
		{
			Type: pulumi.String("azure-native:m365securityandcompliance/v20210325preview:privateLinkServicesForM365SecurityCenter"),
		},
		{
			Type: pulumi.String("azure-nextgen:m365securityandcompliance/v20210325preview:privateLinkServicesForM365SecurityCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicesForM365SecurityCenter
	err := ctx.RegisterResource("azure-native:m365securityandcompliance:privateLinkServicesForM365SecurityCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkServicesForM365SecurityCenter gets an existing PrivateLinkServicesForM365SecurityCenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkServicesForM365SecurityCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForM365SecurityCenterState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForM365SecurityCenter, error) {
	var resource PrivateLinkServicesForM365SecurityCenter
	err := ctx.ReadResource("azure-native:m365securityandcompliance:privateLinkServicesForM365SecurityCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkServicesForM365SecurityCenter resources.
type privateLinkServicesForM365SecurityCenterState struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `pulumi:"etag"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServicesResourceResponseIdentity `pulumi:"identity"`
	// The kind of the service.
	Kind *string `pulumi:"kind"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The common properties of a service.
	Properties *ServicesPropertiesResponse `pulumi:"properties"`
	// Required property for system data
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type PrivateLinkServicesForM365SecurityCenterState struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag pulumi.StringPtrInput
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServicesResourceResponseIdentityPtrInput
	// The kind of the service.
	Kind pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The common properties of a service.
	Properties ServicesPropertiesResponsePtrInput
	// Required property for system data
	SystemData SystemDataResponsePtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (PrivateLinkServicesForM365SecurityCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForM365SecurityCenterState)(nil)).Elem()
}

type privateLinkServicesForM365SecurityCenterArgs struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `pulumi:"etag"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServicesResourceIdentity `pulumi:"identity"`
	// The kind of the service.
	Kind string `pulumi:"kind"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The common properties of a service.
	Properties *ServicesProperties `pulumi:"properties"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName *string `pulumi:"resourceName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateLinkServicesForM365SecurityCenter resource.
type PrivateLinkServicesForM365SecurityCenterArgs struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag pulumi.StringPtrInput
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServicesResourceIdentityPtrInput
	// The kind of the service.
	Kind Kind
	// The resource location.
	Location pulumi.StringPtrInput
	// The common properties of a service.
	Properties ServicesPropertiesPtrInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// The name of the service instance.
	ResourceName pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (PrivateLinkServicesForM365SecurityCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForM365SecurityCenterArgs)(nil)).Elem()
}

type PrivateLinkServicesForM365SecurityCenterInput interface {
	pulumi.Input

	ToPrivateLinkServicesForM365SecurityCenterOutput() PrivateLinkServicesForM365SecurityCenterOutput
	ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365SecurityCenterOutput
}

func (*PrivateLinkServicesForM365SecurityCenter) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServicesForM365SecurityCenter)(nil))
}

func (i *PrivateLinkServicesForM365SecurityCenter) ToPrivateLinkServicesForM365SecurityCenterOutput() PrivateLinkServicesForM365SecurityCenterOutput {
	return i.ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForM365SecurityCenter) ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365SecurityCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForM365SecurityCenterOutput)
}

type PrivateLinkServicesForM365SecurityCenterOutput struct {
	*pulumi.OutputState
}

func (PrivateLinkServicesForM365SecurityCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServicesForM365SecurityCenter)(nil))
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) ToPrivateLinkServicesForM365SecurityCenterOutput() PrivateLinkServicesForM365SecurityCenterOutput {
	return o
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365SecurityCenterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForM365SecurityCenterOutput{})
}
