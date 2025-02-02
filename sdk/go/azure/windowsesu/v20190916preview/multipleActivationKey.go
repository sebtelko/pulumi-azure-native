// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190916preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// MAK key details.
type MultipleActivationKey struct {
	pulumi.CustomResourceState

	// Agreement number under which the key is requested.
	AgreementNumber pulumi.StringPtrOutput `pulumi:"agreementNumber"`
	// End of support of security updates activated by the MAK key.
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// Number of activations/servers using the MAK key.
	InstalledServerNumber pulumi.IntPtrOutput `pulumi:"installedServerNumber"`
	// <code> true </code> if user has eligible on-premises Windows physical or virtual machines, and that the requested key will only be used in their organization; <code> false </code> otherwise.
	IsEligible pulumi.BoolPtrOutput `pulumi:"isEligible"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// MAK 5x5 key.
	MultipleActivationKey pulumi.StringOutput `pulumi:"multipleActivationKey"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Type of OS for which the key is requested.
	OsType            pulumi.StringPtrOutput `pulumi:"osType"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	// Type of support
	SupportType pulumi.StringPtrOutput `pulumi:"supportType"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMultipleActivationKey registers a new resource with the given unique name, arguments, and options.
func NewMultipleActivationKey(ctx *pulumi.Context,
	name string, args *MultipleActivationKeyArgs, opts ...pulumi.ResourceOption) (*MultipleActivationKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SupportType == nil {
		args.SupportType = pulumi.StringPtr("SupplementalServicing")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:windowsesu/v20190916preview:MultipleActivationKey"),
		},
		{
			Type: pulumi.String("azure-native:windowsesu:MultipleActivationKey"),
		},
		{
			Type: pulumi.String("azure-nextgen:windowsesu:MultipleActivationKey"),
		},
	})
	opts = append(opts, aliases)
	var resource MultipleActivationKey
	err := ctx.RegisterResource("azure-native:windowsesu/v20190916preview:MultipleActivationKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMultipleActivationKey gets an existing MultipleActivationKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMultipleActivationKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MultipleActivationKeyState, opts ...pulumi.ResourceOption) (*MultipleActivationKey, error) {
	var resource MultipleActivationKey
	err := ctx.ReadResource("azure-native:windowsesu/v20190916preview:MultipleActivationKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MultipleActivationKey resources.
type multipleActivationKeyState struct {
	// Agreement number under which the key is requested.
	AgreementNumber *string `pulumi:"agreementNumber"`
	// End of support of security updates activated by the MAK key.
	ExpirationDate *string `pulumi:"expirationDate"`
	// Number of activations/servers using the MAK key.
	InstalledServerNumber *int `pulumi:"installedServerNumber"`
	// <code> true </code> if user has eligible on-premises Windows physical or virtual machines, and that the requested key will only be used in their organization; <code> false </code> otherwise.
	IsEligible *bool `pulumi:"isEligible"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// MAK 5x5 key.
	MultipleActivationKey *string `pulumi:"multipleActivationKey"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Type of OS for which the key is requested.
	OsType            *string `pulumi:"osType"`
	ProvisioningState *string `pulumi:"provisioningState"`
	// Type of support
	SupportType *string `pulumi:"supportType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type MultipleActivationKeyState struct {
	// Agreement number under which the key is requested.
	AgreementNumber pulumi.StringPtrInput
	// End of support of security updates activated by the MAK key.
	ExpirationDate pulumi.StringPtrInput
	// Number of activations/servers using the MAK key.
	InstalledServerNumber pulumi.IntPtrInput
	// <code> true </code> if user has eligible on-premises Windows physical or virtual machines, and that the requested key will only be used in their organization; <code> false </code> otherwise.
	IsEligible pulumi.BoolPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// MAK 5x5 key.
	MultipleActivationKey pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Type of OS for which the key is requested.
	OsType            pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	// Type of support
	SupportType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (MultipleActivationKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*multipleActivationKeyState)(nil)).Elem()
}

type multipleActivationKeyArgs struct {
	// Agreement number under which the key is requested.
	AgreementNumber *string `pulumi:"agreementNumber"`
	// Number of activations/servers using the MAK key.
	InstalledServerNumber *int `pulumi:"installedServerNumber"`
	// <code> true </code> if user has eligible on-premises Windows physical or virtual machines, and that the requested key will only be used in their organization; <code> false </code> otherwise.
	IsEligible *bool `pulumi:"isEligible"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the MAK key.
	MultipleActivationKeyName *string `pulumi:"multipleActivationKeyName"`
	// Type of OS for which the key is requested.
	OsType *string `pulumi:"osType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Type of support
	SupportType *string `pulumi:"supportType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MultipleActivationKey resource.
type MultipleActivationKeyArgs struct {
	// Agreement number under which the key is requested.
	AgreementNumber pulumi.StringPtrInput
	// Number of activations/servers using the MAK key.
	InstalledServerNumber pulumi.IntPtrInput
	// <code> true </code> if user has eligible on-premises Windows physical or virtual machines, and that the requested key will only be used in their organization; <code> false </code> otherwise.
	IsEligible pulumi.BoolPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the MAK key.
	MultipleActivationKeyName pulumi.StringPtrInput
	// Type of OS for which the key is requested.
	OsType pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Type of support
	SupportType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (MultipleActivationKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*multipleActivationKeyArgs)(nil)).Elem()
}

type MultipleActivationKeyInput interface {
	pulumi.Input

	ToMultipleActivationKeyOutput() MultipleActivationKeyOutput
	ToMultipleActivationKeyOutputWithContext(ctx context.Context) MultipleActivationKeyOutput
}

func (*MultipleActivationKey) ElementType() reflect.Type {
	return reflect.TypeOf((*MultipleActivationKey)(nil))
}

func (i *MultipleActivationKey) ToMultipleActivationKeyOutput() MultipleActivationKeyOutput {
	return i.ToMultipleActivationKeyOutputWithContext(context.Background())
}

func (i *MultipleActivationKey) ToMultipleActivationKeyOutputWithContext(ctx context.Context) MultipleActivationKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultipleActivationKeyOutput)
}

type MultipleActivationKeyOutput struct {
	*pulumi.OutputState
}

func (MultipleActivationKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MultipleActivationKey)(nil))
}

func (o MultipleActivationKeyOutput) ToMultipleActivationKeyOutput() MultipleActivationKeyOutput {
	return o
}

func (o MultipleActivationKeyOutput) ToMultipleActivationKeyOutputWithContext(ctx context.Context) MultipleActivationKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MultipleActivationKeyOutput{})
}
