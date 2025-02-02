// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datalakestore

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data Lake Store trusted identity provider information.
// API Version: 2016-11-01.
type TrustedIdProvider struct {
	pulumi.CustomResourceState

	// The URL of this trusted identity provider.
	IdProvider pulumi.StringOutput `pulumi:"idProvider"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTrustedIdProvider registers a new resource with the given unique name, arguments, and options.
func NewTrustedIdProvider(ctx *pulumi.Context,
	name string, args *TrustedIdProviderArgs, opts ...pulumi.ResourceOption) (*TrustedIdProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.IdProvider == nil {
		return nil, errors.New("invalid value for required argument 'IdProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datalakestore:TrustedIdProvider"),
		},
		{
			Type: pulumi.String("azure-native:datalakestore/v20161101:TrustedIdProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:datalakestore/v20161101:TrustedIdProvider"),
		},
	})
	opts = append(opts, aliases)
	var resource TrustedIdProvider
	err := ctx.RegisterResource("azure-native:datalakestore:TrustedIdProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrustedIdProvider gets an existing TrustedIdProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrustedIdProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrustedIdProviderState, opts ...pulumi.ResourceOption) (*TrustedIdProvider, error) {
	var resource TrustedIdProvider
	err := ctx.ReadResource("azure-native:datalakestore:TrustedIdProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrustedIdProvider resources.
type trustedIdProviderState struct {
	// The URL of this trusted identity provider.
	IdProvider *string `pulumi:"idProvider"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type TrustedIdProviderState struct {
	// The URL of this trusted identity provider.
	IdProvider pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (TrustedIdProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*trustedIdProviderState)(nil)).Elem()
}

type trustedIdProviderArgs struct {
	// The name of the Data Lake Store account.
	AccountName string `pulumi:"accountName"`
	// The URL of this trusted identity provider.
	IdProvider string `pulumi:"idProvider"`
	// The name of the Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the trusted identity provider. This is used for differentiation of providers in the account.
	TrustedIdProviderName *string `pulumi:"trustedIdProviderName"`
}

// The set of arguments for constructing a TrustedIdProvider resource.
type TrustedIdProviderArgs struct {
	// The name of the Data Lake Store account.
	AccountName pulumi.StringInput
	// The URL of this trusted identity provider.
	IdProvider pulumi.StringInput
	// The name of the Azure resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the trusted identity provider. This is used for differentiation of providers in the account.
	TrustedIdProviderName pulumi.StringPtrInput
}

func (TrustedIdProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trustedIdProviderArgs)(nil)).Elem()
}

type TrustedIdProviderInput interface {
	pulumi.Input

	ToTrustedIdProviderOutput() TrustedIdProviderOutput
	ToTrustedIdProviderOutputWithContext(ctx context.Context) TrustedIdProviderOutput
}

func (*TrustedIdProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedIdProvider)(nil))
}

func (i *TrustedIdProvider) ToTrustedIdProviderOutput() TrustedIdProviderOutput {
	return i.ToTrustedIdProviderOutputWithContext(context.Background())
}

func (i *TrustedIdProvider) ToTrustedIdProviderOutputWithContext(ctx context.Context) TrustedIdProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedIdProviderOutput)
}

type TrustedIdProviderOutput struct {
	*pulumi.OutputState
}

func (TrustedIdProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedIdProvider)(nil))
}

func (o TrustedIdProviderOutput) ToTrustedIdProviderOutput() TrustedIdProviderOutput {
	return o
}

func (o TrustedIdProviderOutput) ToTrustedIdProviderOutputWithContext(ctx context.Context) TrustedIdProviderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TrustedIdProviderOutput{})
}
