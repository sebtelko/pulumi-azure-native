// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210111

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The description of the service.
type PrivateLinkServicesForEDMUpload struct {
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

// NewPrivateLinkServicesForEDMUpload registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkServicesForEDMUpload(ctx *pulumi.Context,
	name string, args *PrivateLinkServicesForEDMUploadArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForEDMUpload, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance/v20210111:privateLinkServicesForEDMUpload"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance:privateLinkServicesForEDMUpload"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance:privateLinkServicesForEDMUpload"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210308:privateLinkServicesForEDMUpload"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance/v20210308:privateLinkServicesForEDMUpload"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicesForEDMUpload
	err := ctx.RegisterResource("azure-native:securityandcompliance/v20210111:privateLinkServicesForEDMUpload", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkServicesForEDMUpload gets an existing PrivateLinkServicesForEDMUpload resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkServicesForEDMUpload(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForEDMUploadState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForEDMUpload, error) {
	var resource PrivateLinkServicesForEDMUpload
	err := ctx.ReadResource("azure-native:securityandcompliance/v20210111:privateLinkServicesForEDMUpload", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkServicesForEDMUpload resources.
type privateLinkServicesForEDMUploadState struct {
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

type PrivateLinkServicesForEDMUploadState struct {
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

func (PrivateLinkServicesForEDMUploadState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForEDMUploadState)(nil)).Elem()
}

type privateLinkServicesForEDMUploadArgs struct {
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

// The set of arguments for constructing a PrivateLinkServicesForEDMUpload resource.
type PrivateLinkServicesForEDMUploadArgs struct {
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

func (PrivateLinkServicesForEDMUploadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForEDMUploadArgs)(nil)).Elem()
}

type PrivateLinkServicesForEDMUploadInput interface {
	pulumi.Input

	ToPrivateLinkServicesForEDMUploadOutput() PrivateLinkServicesForEDMUploadOutput
	ToPrivateLinkServicesForEDMUploadOutputWithContext(ctx context.Context) PrivateLinkServicesForEDMUploadOutput
}

func (*PrivateLinkServicesForEDMUpload) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServicesForEDMUpload)(nil))
}

func (i *PrivateLinkServicesForEDMUpload) ToPrivateLinkServicesForEDMUploadOutput() PrivateLinkServicesForEDMUploadOutput {
	return i.ToPrivateLinkServicesForEDMUploadOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForEDMUpload) ToPrivateLinkServicesForEDMUploadOutputWithContext(ctx context.Context) PrivateLinkServicesForEDMUploadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForEDMUploadOutput)
}

type PrivateLinkServicesForEDMUploadOutput struct {
	*pulumi.OutputState
}

func (PrivateLinkServicesForEDMUploadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServicesForEDMUpload)(nil))
}

func (o PrivateLinkServicesForEDMUploadOutput) ToPrivateLinkServicesForEDMUploadOutput() PrivateLinkServicesForEDMUploadOutput {
	return o
}

func (o PrivateLinkServicesForEDMUploadOutput) ToPrivateLinkServicesForEDMUploadOutputWithContext(ctx context.Context) PrivateLinkServicesForEDMUploadOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForEDMUploadOutput{})
}
