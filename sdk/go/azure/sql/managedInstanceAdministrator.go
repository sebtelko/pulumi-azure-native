// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure SQL managed instance administrator.
// API Version: 2020-11-01-preview.
type ManagedInstanceAdministrator struct {
	pulumi.CustomResourceState

	// Type of the managed instance administrator.
	AdministratorType pulumi.StringOutput `pulumi:"administratorType"`
	// Login name of the managed instance administrator.
	Login pulumi.StringOutput `pulumi:"login"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// SID (object ID) of the managed instance administrator.
	Sid pulumi.StringOutput `pulumi:"sid"`
	// Tenant ID of the managed instance administrator.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedInstanceAdministrator registers a new resource with the given unique name, arguments, and options.
func NewManagedInstanceAdministrator(ctx *pulumi.Context,
	name string, args *ManagedInstanceAdministratorArgs, opts ...pulumi.ResourceOption) (*ManagedInstanceAdministrator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdministratorType == nil {
		return nil, errors.New("invalid value for required argument 'AdministratorType'")
	}
	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sid == nil {
		return nil, errors.New("invalid value for required argument 'Sid'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20170301preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:ManagedInstanceAdministrator"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedInstanceAdministrator
	err := ctx.RegisterResource("azure-native:sql:ManagedInstanceAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedInstanceAdministrator gets an existing ManagedInstanceAdministrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedInstanceAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceAdministratorState, opts ...pulumi.ResourceOption) (*ManagedInstanceAdministrator, error) {
	var resource ManagedInstanceAdministrator
	err := ctx.ReadResource("azure-native:sql:ManagedInstanceAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedInstanceAdministrator resources.
type managedInstanceAdministratorState struct {
	// Type of the managed instance administrator.
	AdministratorType *string `pulumi:"administratorType"`
	// Login name of the managed instance administrator.
	Login *string `pulumi:"login"`
	// Resource name.
	Name *string `pulumi:"name"`
	// SID (object ID) of the managed instance administrator.
	Sid *string `pulumi:"sid"`
	// Tenant ID of the managed instance administrator.
	TenantId *string `pulumi:"tenantId"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ManagedInstanceAdministratorState struct {
	// Type of the managed instance administrator.
	AdministratorType pulumi.StringPtrInput
	// Login name of the managed instance administrator.
	Login pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// SID (object ID) of the managed instance administrator.
	Sid pulumi.StringPtrInput
	// Tenant ID of the managed instance administrator.
	TenantId pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ManagedInstanceAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceAdministratorState)(nil)).Elem()
}

type managedInstanceAdministratorArgs struct {
	AdministratorName *string `pulumi:"administratorName"`
	// Type of the managed instance administrator.
	AdministratorType string `pulumi:"administratorType"`
	// Login name of the managed instance administrator.
	Login string `pulumi:"login"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SID (object ID) of the managed instance administrator.
	Sid string `pulumi:"sid"`
	// Tenant ID of the managed instance administrator.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ManagedInstanceAdministrator resource.
type ManagedInstanceAdministratorArgs struct {
	AdministratorName pulumi.StringPtrInput
	// Type of the managed instance administrator.
	AdministratorType pulumi.StringInput
	// Login name of the managed instance administrator.
	Login pulumi.StringInput
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// SID (object ID) of the managed instance administrator.
	Sid pulumi.StringInput
	// Tenant ID of the managed instance administrator.
	TenantId pulumi.StringPtrInput
}

func (ManagedInstanceAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceAdministratorArgs)(nil)).Elem()
}

type ManagedInstanceAdministratorInput interface {
	pulumi.Input

	ToManagedInstanceAdministratorOutput() ManagedInstanceAdministratorOutput
	ToManagedInstanceAdministratorOutputWithContext(ctx context.Context) ManagedInstanceAdministratorOutput
}

func (*ManagedInstanceAdministrator) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceAdministrator)(nil))
}

func (i *ManagedInstanceAdministrator) ToManagedInstanceAdministratorOutput() ManagedInstanceAdministratorOutput {
	return i.ToManagedInstanceAdministratorOutputWithContext(context.Background())
}

func (i *ManagedInstanceAdministrator) ToManagedInstanceAdministratorOutputWithContext(ctx context.Context) ManagedInstanceAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceAdministratorOutput)
}

type ManagedInstanceAdministratorOutput struct {
	*pulumi.OutputState
}

func (ManagedInstanceAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceAdministrator)(nil))
}

func (o ManagedInstanceAdministratorOutput) ToManagedInstanceAdministratorOutput() ManagedInstanceAdministratorOutput {
	return o
}

func (o ManagedInstanceAdministratorOutput) ToManagedInstanceAdministratorOutputWithContext(ctx context.Context) ManagedInstanceAdministratorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedInstanceAdministratorOutput{})
}
