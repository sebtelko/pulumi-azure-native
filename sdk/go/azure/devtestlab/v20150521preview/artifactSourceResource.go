// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Properties of an artifact source.
type ArtifactSourceResource struct {
	pulumi.CustomResourceState

	// The branch reference of the artifact source.
	BranchRef pulumi.StringPtrOutput `pulumi:"branchRef"`
	// The display name of the artifact source.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The folder path of the artifact source.
	FolderPath pulumi.StringPtrOutput `pulumi:"folderPath"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The security token of the artifact source.
	SecurityToken pulumi.StringPtrOutput `pulumi:"securityToken"`
	// The type of the artifact source.
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	// The status of the artifact source.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The URI of the artifact source.
	Uri pulumi.StringPtrOutput `pulumi:"uri"`
}

// NewArtifactSourceResource registers a new resource with the given unique name, arguments, and options.
func NewArtifactSourceResource(ctx *pulumi.Context,
	name string, args *ArtifactSourceResourceArgs, opts ...pulumi.ResourceOption) (*ArtifactSourceResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20150521preview:ArtifactSourceResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab:ArtifactSourceResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab:ArtifactSourceResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:ArtifactSourceResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20160515:ArtifactSourceResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:ArtifactSourceResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20180915:ArtifactSourceResource"),
		},
	})
	opts = append(opts, aliases)
	var resource ArtifactSourceResource
	err := ctx.RegisterResource("azure-native:devtestlab/v20150521preview:ArtifactSourceResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArtifactSourceResource gets an existing ArtifactSourceResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArtifactSourceResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactSourceResourceState, opts ...pulumi.ResourceOption) (*ArtifactSourceResource, error) {
	var resource ArtifactSourceResource
	err := ctx.ReadResource("azure-native:devtestlab/v20150521preview:ArtifactSourceResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArtifactSourceResource resources.
type artifactSourceResourceState struct {
	// The branch reference of the artifact source.
	BranchRef *string `pulumi:"branchRef"`
	// The display name of the artifact source.
	DisplayName *string `pulumi:"displayName"`
	// The folder path of the artifact source.
	FolderPath *string `pulumi:"folderPath"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The security token of the artifact source.
	SecurityToken *string `pulumi:"securityToken"`
	// The type of the artifact source.
	SourceType *string `pulumi:"sourceType"`
	// The status of the artifact source.
	Status *string `pulumi:"status"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// The URI of the artifact source.
	Uri *string `pulumi:"uri"`
}

type ArtifactSourceResourceState struct {
	// The branch reference of the artifact source.
	BranchRef pulumi.StringPtrInput
	// The display name of the artifact source.
	DisplayName pulumi.StringPtrInput
	// The folder path of the artifact source.
	FolderPath pulumi.StringPtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The security token of the artifact source.
	SecurityToken pulumi.StringPtrInput
	// The type of the artifact source.
	SourceType pulumi.StringPtrInput
	// The status of the artifact source.
	Status pulumi.StringPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
	// The URI of the artifact source.
	Uri pulumi.StringPtrInput
}

func (ArtifactSourceResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactSourceResourceState)(nil)).Elem()
}

type artifactSourceResourceArgs struct {
	// The branch reference of the artifact source.
	BranchRef *string `pulumi:"branchRef"`
	// The display name of the artifact source.
	DisplayName *string `pulumi:"displayName"`
	// The folder path of the artifact source.
	FolderPath *string `pulumi:"folderPath"`
	// The identifier of the resource.
	Id *string `pulumi:"id"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security token of the artifact source.
	SecurityToken *string `pulumi:"securityToken"`
	// The type of the artifact source.
	SourceType *string `pulumi:"sourceType"`
	// The status of the artifact source.
	Status *string `pulumi:"status"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// The URI of the artifact source.
	Uri *string `pulumi:"uri"`
}

// The set of arguments for constructing a ArtifactSourceResource resource.
type ArtifactSourceResourceArgs struct {
	// The branch reference of the artifact source.
	BranchRef pulumi.StringPtrInput
	// The display name of the artifact source.
	DisplayName pulumi.StringPtrInput
	// The folder path of the artifact source.
	FolderPath pulumi.StringPtrInput
	// The identifier of the resource.
	Id pulumi.StringPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The security token of the artifact source.
	SecurityToken pulumi.StringPtrInput
	// The type of the artifact source.
	SourceType pulumi.StringPtrInput
	// The status of the artifact source.
	Status pulumi.StringPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
	// The URI of the artifact source.
	Uri pulumi.StringPtrInput
}

func (ArtifactSourceResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactSourceResourceArgs)(nil)).Elem()
}

type ArtifactSourceResourceInput interface {
	pulumi.Input

	ToArtifactSourceResourceOutput() ArtifactSourceResourceOutput
	ToArtifactSourceResourceOutputWithContext(ctx context.Context) ArtifactSourceResourceOutput
}

func (*ArtifactSourceResource) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactSourceResource)(nil))
}

func (i *ArtifactSourceResource) ToArtifactSourceResourceOutput() ArtifactSourceResourceOutput {
	return i.ToArtifactSourceResourceOutputWithContext(context.Background())
}

func (i *ArtifactSourceResource) ToArtifactSourceResourceOutputWithContext(ctx context.Context) ArtifactSourceResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactSourceResourceOutput)
}

type ArtifactSourceResourceOutput struct {
	*pulumi.OutputState
}

func (ArtifactSourceResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactSourceResource)(nil))
}

func (o ArtifactSourceResourceOutput) ToArtifactSourceResourceOutput() ArtifactSourceResourceOutput {
	return o
}

func (o ArtifactSourceResourceOutput) ToArtifactSourceResourceOutputWithContext(ctx context.Context) ArtifactSourceResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ArtifactSourceResourceOutput{})
}
