// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Live Output.
type LiveOutput struct {
	pulumi.CustomResourceState

	// ISO 8601 timespan duration of the archive window length. This is duration that customer want to retain the recorded content.
	ArchiveWindowLength pulumi.StringOutput `pulumi:"archiveWindowLength"`
	// The asset name.
	AssetName pulumi.StringOutput `pulumi:"assetName"`
	// The exact time the Live Output was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// The description of the Live Output.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The HLS configuration.
	Hls HlsResponsePtrOutput `pulumi:"hls"`
	// The exact time the Live Output was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The manifest file name.
	ManifestName pulumi.StringPtrOutput `pulumi:"manifestName"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The output snapshot time.
	OutputSnapTime pulumi.Float64PtrOutput `pulumi:"outputSnapTime"`
	// The provisioning state of the Live Output.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource state of the Live Output.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLiveOutput registers a new resource with the given unique name, arguments, and options.
func NewLiveOutput(ctx *pulumi.Context,
	name string, args *LiveOutputArgs, opts ...pulumi.ResourceOption) (*LiveOutput, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ArchiveWindowLength == nil {
		return nil, errors.New("invalid value for required argument 'ArchiveWindowLength'")
	}
	if args.AssetName == nil {
		return nil, errors.New("invalid value for required argument 'AssetName'")
	}
	if args.LiveEventName == nil {
		return nil, errors.New("invalid value for required argument 'LiveEventName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:media/v20180601preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-nextgen:media:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180330preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180701:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20190501preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20190501preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20200501:LiveOutput"),
		},
	})
	opts = append(opts, aliases)
	var resource LiveOutput
	err := ctx.RegisterResource("azure-native:media/v20180601preview:LiveOutput", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLiveOutput gets an existing LiveOutput resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLiveOutput(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiveOutputState, opts ...pulumi.ResourceOption) (*LiveOutput, error) {
	var resource LiveOutput
	err := ctx.ReadResource("azure-native:media/v20180601preview:LiveOutput", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LiveOutput resources.
type liveOutputState struct {
	// ISO 8601 timespan duration of the archive window length. This is duration that customer want to retain the recorded content.
	ArchiveWindowLength *string `pulumi:"archiveWindowLength"`
	// The asset name.
	AssetName *string `pulumi:"assetName"`
	// The exact time the Live Output was created.
	Created *string `pulumi:"created"`
	// The description of the Live Output.
	Description *string `pulumi:"description"`
	// The HLS configuration.
	Hls *HlsResponse `pulumi:"hls"`
	// The exact time the Live Output was last modified.
	LastModified *string `pulumi:"lastModified"`
	// The manifest file name.
	ManifestName *string `pulumi:"manifestName"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The output snapshot time.
	OutputSnapTime *float64 `pulumi:"outputSnapTime"`
	// The provisioning state of the Live Output.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource state of the Live Output.
	ResourceState *string `pulumi:"resourceState"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type LiveOutputState struct {
	// ISO 8601 timespan duration of the archive window length. This is duration that customer want to retain the recorded content.
	ArchiveWindowLength pulumi.StringPtrInput
	// The asset name.
	AssetName pulumi.StringPtrInput
	// The exact time the Live Output was created.
	Created pulumi.StringPtrInput
	// The description of the Live Output.
	Description pulumi.StringPtrInput
	// The HLS configuration.
	Hls HlsResponsePtrInput
	// The exact time the Live Output was last modified.
	LastModified pulumi.StringPtrInput
	// The manifest file name.
	ManifestName pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The output snapshot time.
	OutputSnapTime pulumi.Float64PtrInput
	// The provisioning state of the Live Output.
	ProvisioningState pulumi.StringPtrInput
	// The resource state of the Live Output.
	ResourceState pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (LiveOutputState) ElementType() reflect.Type {
	return reflect.TypeOf((*liveOutputState)(nil)).Elem()
}

type liveOutputArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// ISO 8601 timespan duration of the archive window length. This is duration that customer want to retain the recorded content.
	ArchiveWindowLength string `pulumi:"archiveWindowLength"`
	// The asset name.
	AssetName string `pulumi:"assetName"`
	// The description of the Live Output.
	Description *string `pulumi:"description"`
	// The HLS configuration.
	Hls *Hls `pulumi:"hls"`
	// The name of the Live Event.
	LiveEventName string `pulumi:"liveEventName"`
	// The name of the Live Output.
	LiveOutputName *string `pulumi:"liveOutputName"`
	// The manifest file name.
	ManifestName *string `pulumi:"manifestName"`
	// The output snapshot time.
	OutputSnapTime *float64 `pulumi:"outputSnapTime"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LiveOutput resource.
type LiveOutputArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// ISO 8601 timespan duration of the archive window length. This is duration that customer want to retain the recorded content.
	ArchiveWindowLength pulumi.StringInput
	// The asset name.
	AssetName pulumi.StringInput
	// The description of the Live Output.
	Description pulumi.StringPtrInput
	// The HLS configuration.
	Hls HlsPtrInput
	// The name of the Live Event.
	LiveEventName pulumi.StringInput
	// The name of the Live Output.
	LiveOutputName pulumi.StringPtrInput
	// The manifest file name.
	ManifestName pulumi.StringPtrInput
	// The output snapshot time.
	OutputSnapTime pulumi.Float64PtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (LiveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liveOutputArgs)(nil)).Elem()
}

type LiveOutputInput interface {
	pulumi.Input

	ToLiveOutputOutput() LiveOutputOutput
	ToLiveOutputOutputWithContext(ctx context.Context) LiveOutputOutput
}

func (*LiveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveOutput)(nil))
}

func (i *LiveOutput) ToLiveOutputOutput() LiveOutputOutput {
	return i.ToLiveOutputOutputWithContext(context.Background())
}

func (i *LiveOutput) ToLiveOutputOutputWithContext(ctx context.Context) LiveOutputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveOutputOutput)
}

type LiveOutputOutput struct {
	*pulumi.OutputState
}

func (LiveOutputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveOutput)(nil))
}

func (o LiveOutputOutput) ToLiveOutputOutput() LiveOutputOutput {
	return o
}

func (o LiveOutputOutput) ToLiveOutputOutputWithContext(ctx context.Context) LiveOutputOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LiveOutputOutput{})
}
