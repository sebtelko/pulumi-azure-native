// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NetApp account resource
type Account struct {
	pulumi.CustomResourceState

	// Active Directories
	ActiveDirectories ActiveDirectoryResponseArrayOutput `pulumi:"activeDirectories"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190501:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20170815:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20170815:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190601:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190601:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190701:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190701:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190801:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190801:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191001:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20191001:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191101:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20191101:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200201:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200201:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200301:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200301:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200501:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200601:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200701:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200801:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200901:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20201101:Account"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20201201:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:netapp/v20190501:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:netapp/v20190501:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// Active Directories
	ActiveDirectories []ActiveDirectoryResponse `pulumi:"activeDirectories"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags
	Tags interface{} `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type AccountState struct {
	// Active Directories
	ActiveDirectories ActiveDirectoryResponseArrayInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Azure lifecycle management
	ProvisioningState pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.Input
	// Resource type
	Type pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the NetApp account
	AccountName *string `pulumi:"accountName"`
	// Active Directories
	ActiveDirectories []ActiveDirectory `pulumi:"activeDirectories"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringPtrInput
	// Active Directories
	ActiveDirectories ActiveDirectoryArrayInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.Input
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct {
	*pulumi.OutputState
}

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
