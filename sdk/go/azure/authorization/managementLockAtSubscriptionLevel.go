// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The lock information.
// API Version: 2016-09-01.
type ManagementLockAtSubscriptionLevel struct {
	pulumi.CustomResourceState

	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level pulumi.StringOutput `pulumi:"level"`
	// The name of the lock.
	Name pulumi.StringOutput `pulumi:"name"`
	// Notes about the lock. Maximum of 512 characters.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The owners of the lock.
	Owners ManagementLockOwnerResponseArrayOutput `pulumi:"owners"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagementLockAtSubscriptionLevel registers a new resource with the given unique name, arguments, and options.
func NewManagementLockAtSubscriptionLevel(ctx *pulumi.Context,
	name string, args *ManagementLockAtSubscriptionLevelArgs, opts ...pulumi.ResourceOption) (*ManagementLockAtSubscriptionLevel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Level == nil {
		return nil, errors.New("invalid value for required argument 'Level'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:authorization:ManagementLockAtSubscriptionLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20150101:ManagementLockAtSubscriptionLevel"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20150101:ManagementLockAtSubscriptionLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20160901:ManagementLockAtSubscriptionLevel"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20160901:ManagementLockAtSubscriptionLevel"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementLockAtSubscriptionLevel
	err := ctx.RegisterResource("azure-native:authorization:ManagementLockAtSubscriptionLevel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementLockAtSubscriptionLevel gets an existing ManagementLockAtSubscriptionLevel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementLockAtSubscriptionLevel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementLockAtSubscriptionLevelState, opts ...pulumi.ResourceOption) (*ManagementLockAtSubscriptionLevel, error) {
	var resource ManagementLockAtSubscriptionLevel
	err := ctx.ReadResource("azure-native:authorization:ManagementLockAtSubscriptionLevel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementLockAtSubscriptionLevel resources.
type managementLockAtSubscriptionLevelState struct {
	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level *string `pulumi:"level"`
	// The name of the lock.
	Name *string `pulumi:"name"`
	// Notes about the lock. Maximum of 512 characters.
	Notes *string `pulumi:"notes"`
	// The owners of the lock.
	Owners []ManagementLockOwnerResponse `pulumi:"owners"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type *string `pulumi:"type"`
}

type ManagementLockAtSubscriptionLevelState struct {
	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level pulumi.StringPtrInput
	// The name of the lock.
	Name pulumi.StringPtrInput
	// Notes about the lock. Maximum of 512 characters.
	Notes pulumi.StringPtrInput
	// The owners of the lock.
	Owners ManagementLockOwnerResponseArrayInput
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type pulumi.StringPtrInput
}

func (ManagementLockAtSubscriptionLevelState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtSubscriptionLevelState)(nil)).Elem()
}

type managementLockAtSubscriptionLevelArgs struct {
	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level string `pulumi:"level"`
	// The name of lock. The lock name can be a maximum of 260 characters. It cannot contain <, > %, &, :, \, ?, /, or any control characters.
	LockName *string `pulumi:"lockName"`
	// Notes about the lock. Maximum of 512 characters.
	Notes *string `pulumi:"notes"`
	// The owners of the lock.
	Owners []ManagementLockOwner `pulumi:"owners"`
}

// The set of arguments for constructing a ManagementLockAtSubscriptionLevel resource.
type ManagementLockAtSubscriptionLevelArgs struct {
	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level pulumi.StringInput
	// The name of lock. The lock name can be a maximum of 260 characters. It cannot contain <, > %, &, :, \, ?, /, or any control characters.
	LockName pulumi.StringPtrInput
	// Notes about the lock. Maximum of 512 characters.
	Notes pulumi.StringPtrInput
	// The owners of the lock.
	Owners ManagementLockOwnerArrayInput
}

func (ManagementLockAtSubscriptionLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtSubscriptionLevelArgs)(nil)).Elem()
}

type ManagementLockAtSubscriptionLevelInput interface {
	pulumi.Input

	ToManagementLockAtSubscriptionLevelOutput() ManagementLockAtSubscriptionLevelOutput
	ToManagementLockAtSubscriptionLevelOutputWithContext(ctx context.Context) ManagementLockAtSubscriptionLevelOutput
}

func (*ManagementLockAtSubscriptionLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtSubscriptionLevel)(nil))
}

func (i *ManagementLockAtSubscriptionLevel) ToManagementLockAtSubscriptionLevelOutput() ManagementLockAtSubscriptionLevelOutput {
	return i.ToManagementLockAtSubscriptionLevelOutputWithContext(context.Background())
}

func (i *ManagementLockAtSubscriptionLevel) ToManagementLockAtSubscriptionLevelOutputWithContext(ctx context.Context) ManagementLockAtSubscriptionLevelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockAtSubscriptionLevelOutput)
}

type ManagementLockAtSubscriptionLevelOutput struct {
	*pulumi.OutputState
}

func (ManagementLockAtSubscriptionLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockAtSubscriptionLevel)(nil))
}

func (o ManagementLockAtSubscriptionLevelOutput) ToManagementLockAtSubscriptionLevelOutput() ManagementLockAtSubscriptionLevelOutput {
	return o
}

func (o ManagementLockAtSubscriptionLevelOutput) ToManagementLockAtSubscriptionLevelOutputWithContext(ctx context.Context) ManagementLockAtSubscriptionLevelOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagementLockAtSubscriptionLevelOutput{})
}
