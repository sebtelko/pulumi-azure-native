// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network function resource response.
type NetworkFunction struct {
	pulumi.CustomResourceState

	// The reference to the device resource.
	Device SubResourceResponsePtrOutput `pulumi:"device"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource URI of the managed application.
	ManagedApplication SubResourceResponseOutput `pulumi:"managedApplication"`
	// The parameters for the managed application.
	ManagedApplicationParameters pulumi.AnyOutput `pulumi:"managedApplicationParameters"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The network function configurations from the user.
	NetworkFunctionUserConfigurations NetworkFunctionUserConfigurationResponseArrayOutput `pulumi:"networkFunctionUserConfigurations"`
	// The provisioning state of the network function resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The service key for the network function resource.
	ServiceKey pulumi.StringOutput `pulumi:"serviceKey"`
	// The sku name for the network function.
	SkuName pulumi.StringPtrOutput `pulumi:"skuName"`
	// The sku type for the network function.
	SkuType pulumi.StringOutput `pulumi:"skuType"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The vendor name for the network function.
	VendorName pulumi.StringPtrOutput `pulumi:"vendorName"`
	// The vendor provisioning state for the network function resource.
	VendorProvisioningState pulumi.StringOutput `pulumi:"vendorProvisioningState"`
}

// NewNetworkFunction registers a new resource with the given unique name, arguments, and options.
func NewNetworkFunction(ctx *pulumi.Context,
	name string, args *NetworkFunctionArgs, opts ...pulumi.ResourceOption) (*NetworkFunction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:hybridnetwork/v20200101preview:NetworkFunction"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork:NetworkFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridnetwork:NetworkFunction"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkFunction
	err := ctx.RegisterResource("azure-native:hybridnetwork/v20200101preview:NetworkFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkFunction gets an existing NetworkFunction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkFunctionState, opts ...pulumi.ResourceOption) (*NetworkFunction, error) {
	var resource NetworkFunction
	err := ctx.ReadResource("azure-native:hybridnetwork/v20200101preview:NetworkFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkFunction resources.
type networkFunctionState struct {
	// The reference to the device resource.
	Device *SubResourceResponse `pulumi:"device"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The resource URI of the managed application.
	ManagedApplication *SubResourceResponse `pulumi:"managedApplication"`
	// The parameters for the managed application.
	ManagedApplicationParameters interface{} `pulumi:"managedApplicationParameters"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The network function configurations from the user.
	NetworkFunctionUserConfigurations []NetworkFunctionUserConfigurationResponse `pulumi:"networkFunctionUserConfigurations"`
	// The provisioning state of the network function resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The service key for the network function resource.
	ServiceKey *string `pulumi:"serviceKey"`
	// The sku name for the network function.
	SkuName *string `pulumi:"skuName"`
	// The sku type for the network function.
	SkuType *string `pulumi:"skuType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
	// The vendor name for the network function.
	VendorName *string `pulumi:"vendorName"`
	// The vendor provisioning state for the network function resource.
	VendorProvisioningState *string `pulumi:"vendorProvisioningState"`
}

type NetworkFunctionState struct {
	// The reference to the device resource.
	Device SubResourceResponsePtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The resource URI of the managed application.
	ManagedApplication SubResourceResponsePtrInput
	// The parameters for the managed application.
	ManagedApplicationParameters pulumi.Input
	// The name of the resource
	Name pulumi.StringPtrInput
	// The network function configurations from the user.
	NetworkFunctionUserConfigurations NetworkFunctionUserConfigurationResponseArrayInput
	// The provisioning state of the network function resource.
	ProvisioningState pulumi.StringPtrInput
	// The service key for the network function resource.
	ServiceKey pulumi.StringPtrInput
	// The sku name for the network function.
	SkuName pulumi.StringPtrInput
	// The sku type for the network function.
	SkuType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
	// The vendor name for the network function.
	VendorName pulumi.StringPtrInput
	// The vendor provisioning state for the network function resource.
	VendorProvisioningState pulumi.StringPtrInput
}

func (NetworkFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkFunctionState)(nil)).Elem()
}

type networkFunctionArgs struct {
	// The reference to the device resource.
	Device *SubResource `pulumi:"device"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The parameters for the managed application.
	ManagedApplicationParameters interface{} `pulumi:"managedApplicationParameters"`
	// Resource name for the network function resource.
	NetworkFunctionName *string `pulumi:"networkFunctionName"`
	// The network function configurations from the user.
	NetworkFunctionUserConfigurations []NetworkFunctionUserConfiguration `pulumi:"networkFunctionUserConfigurations"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku name for the network function.
	SkuName *string `pulumi:"skuName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The vendor name for the network function.
	VendorName *string `pulumi:"vendorName"`
}

// The set of arguments for constructing a NetworkFunction resource.
type NetworkFunctionArgs struct {
	// The reference to the device resource.
	Device SubResourcePtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The parameters for the managed application.
	ManagedApplicationParameters pulumi.Input
	// Resource name for the network function resource.
	NetworkFunctionName pulumi.StringPtrInput
	// The network function configurations from the user.
	NetworkFunctionUserConfigurations NetworkFunctionUserConfigurationArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The sku name for the network function.
	SkuName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The vendor name for the network function.
	VendorName pulumi.StringPtrInput
}

func (NetworkFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkFunctionArgs)(nil)).Elem()
}

type NetworkFunctionInput interface {
	pulumi.Input

	ToNetworkFunctionOutput() NetworkFunctionOutput
	ToNetworkFunctionOutputWithContext(ctx context.Context) NetworkFunctionOutput
}

func (*NetworkFunction) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunction)(nil))
}

func (i *NetworkFunction) ToNetworkFunctionOutput() NetworkFunctionOutput {
	return i.ToNetworkFunctionOutputWithContext(context.Background())
}

func (i *NetworkFunction) ToNetworkFunctionOutputWithContext(ctx context.Context) NetworkFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionOutput)
}

type NetworkFunctionOutput struct {
	*pulumi.OutputState
}

func (NetworkFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunction)(nil))
}

func (o NetworkFunctionOutput) ToNetworkFunctionOutput() NetworkFunctionOutput {
	return o
}

func (o NetworkFunctionOutput) ToNetworkFunctionOutputWithContext(ctx context.Context) NetworkFunctionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkFunctionOutput{})
}
