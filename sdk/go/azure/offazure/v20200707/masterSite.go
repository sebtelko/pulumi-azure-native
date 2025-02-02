// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200707

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Site REST Resource.
type MasterSite struct {
	pulumi.CustomResourceState

	// eTag for concurrency control.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Azure location in which Sites is created.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the Master site.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Nested properties of Master site.
	Properties MasterSitePropertiesResponseOutput `pulumi:"properties"`
	// Type of resource. Type = Microsoft.OffAzure/MasterSites.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMasterSite registers a new resource with the given unique name, arguments, and options.
func NewMasterSite(ctx *pulumi.Context,
	name string, args *MasterSiteArgs, opts ...pulumi.ResourceOption) (*MasterSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:offazure/v20200707:MasterSite"),
		},
		{
			Type: pulumi.String("azure-native:offazure:MasterSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:offazure:MasterSite"),
		},
	})
	opts = append(opts, aliases)
	var resource MasterSite
	err := ctx.RegisterResource("azure-native:offazure/v20200707:MasterSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMasterSite gets an existing MasterSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMasterSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MasterSiteState, opts ...pulumi.ResourceOption) (*MasterSite, error) {
	var resource MasterSite
	err := ctx.ReadResource("azure-native:offazure/v20200707:MasterSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MasterSite resources.
type masterSiteState struct {
	// eTag for concurrency control.
	ETag *string `pulumi:"eTag"`
	// Azure location in which Sites is created.
	Location *string `pulumi:"location"`
	// Name of the Master site.
	Name *string `pulumi:"name"`
	// Nested properties of Master site.
	Properties *MasterSitePropertiesResponse `pulumi:"properties"`
	// Type of resource. Type = Microsoft.OffAzure/MasterSites.
	Type *string `pulumi:"type"`
}

type MasterSiteState struct {
	// eTag for concurrency control.
	ETag pulumi.StringPtrInput
	// Azure location in which Sites is created.
	Location pulumi.StringPtrInput
	// Name of the Master site.
	Name pulumi.StringPtrInput
	// Nested properties of Master site.
	Properties MasterSitePropertiesResponsePtrInput
	// Type of resource. Type = Microsoft.OffAzure/MasterSites.
	Type pulumi.StringPtrInput
}

func (MasterSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*masterSiteState)(nil)).Elem()
}

type masterSiteArgs struct {
	// eTag for concurrency control.
	ETag *string `pulumi:"eTag"`
	// Azure location in which Sites is created.
	Location *string `pulumi:"location"`
	// Name of the Master site.
	Name *string `pulumi:"name"`
	// Nested properties of Master site.
	Properties *MasterSiteProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name.
	SiteName *string `pulumi:"siteName"`
}

// The set of arguments for constructing a MasterSite resource.
type MasterSiteArgs struct {
	// eTag for concurrency control.
	ETag pulumi.StringPtrInput
	// Azure location in which Sites is created.
	Location pulumi.StringPtrInput
	// Name of the Master site.
	Name pulumi.StringPtrInput
	// Nested properties of Master site.
	Properties MasterSitePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Site name.
	SiteName pulumi.StringPtrInput
}

func (MasterSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*masterSiteArgs)(nil)).Elem()
}

type MasterSiteInput interface {
	pulumi.Input

	ToMasterSiteOutput() MasterSiteOutput
	ToMasterSiteOutputWithContext(ctx context.Context) MasterSiteOutput
}

func (*MasterSite) ElementType() reflect.Type {
	return reflect.TypeOf((*MasterSite)(nil))
}

func (i *MasterSite) ToMasterSiteOutput() MasterSiteOutput {
	return i.ToMasterSiteOutputWithContext(context.Background())
}

func (i *MasterSite) ToMasterSiteOutputWithContext(ctx context.Context) MasterSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterSiteOutput)
}

type MasterSiteOutput struct {
	*pulumi.OutputState
}

func (MasterSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MasterSite)(nil))
}

func (o MasterSiteOutput) ToMasterSiteOutput() MasterSiteOutput {
	return o
}

func (o MasterSiteOutput) ToMasterSiteOutputWithContext(ctx context.Context) MasterSiteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MasterSiteOutput{})
}
