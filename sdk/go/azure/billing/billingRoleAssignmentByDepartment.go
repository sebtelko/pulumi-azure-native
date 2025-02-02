// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The role assignment
// API Version: 2019-10-01-preview.
type BillingRoleAssignmentByDepartment struct {
	pulumi.CustomResourceState

	// The principal Id of the user who created the role assignment.
	CreatedByPrincipalId pulumi.StringOutput `pulumi:"createdByPrincipalId"`
	// The tenant Id of the user who created the role assignment.
	CreatedByPrincipalTenantId pulumi.StringOutput `pulumi:"createdByPrincipalTenantId"`
	// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
	CreatedByUserEmailAddress pulumi.StringOutput `pulumi:"createdByUserEmailAddress"`
	// The date the role assignment was created.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The principal id of the user to whom the role was assigned.
	PrincipalId pulumi.StringPtrOutput `pulumi:"principalId"`
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId pulumi.StringPtrOutput `pulumi:"principalTenantId"`
	// The ID of the role definition.
	RoleDefinitionId pulumi.StringPtrOutput `pulumi:"roleDefinitionId"`
	// The scope at which the role was assigned.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType pulumi.StringPtrOutput `pulumi:"userAuthenticationType"`
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress pulumi.StringPtrOutput `pulumi:"userEmailAddress"`
}

// NewBillingRoleAssignmentByDepartment registers a new resource with the given unique name, arguments, and options.
func NewBillingRoleAssignmentByDepartment(ctx *pulumi.Context,
	name string, args *BillingRoleAssignmentByDepartmentArgs, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByDepartment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountName == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountName'")
	}
	if args.DepartmentName == nil {
		return nil, errors.New("invalid value for required argument 'DepartmentName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:billing:BillingRoleAssignmentByDepartment"),
		},
		{
			Type: pulumi.String("azure-native:billing/v20191001preview:BillingRoleAssignmentByDepartment"),
		},
		{
			Type: pulumi.String("azure-nextgen:billing/v20191001preview:BillingRoleAssignmentByDepartment"),
		},
	})
	opts = append(opts, aliases)
	var resource BillingRoleAssignmentByDepartment
	err := ctx.RegisterResource("azure-native:billing:BillingRoleAssignmentByDepartment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingRoleAssignmentByDepartment gets an existing BillingRoleAssignmentByDepartment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingRoleAssignmentByDepartment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingRoleAssignmentByDepartmentState, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByDepartment, error) {
	var resource BillingRoleAssignmentByDepartment
	err := ctx.ReadResource("azure-native:billing:BillingRoleAssignmentByDepartment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingRoleAssignmentByDepartment resources.
type billingRoleAssignmentByDepartmentState struct {
	// The principal Id of the user who created the role assignment.
	CreatedByPrincipalId *string `pulumi:"createdByPrincipalId"`
	// The tenant Id of the user who created the role assignment.
	CreatedByPrincipalTenantId *string `pulumi:"createdByPrincipalTenantId"`
	// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
	CreatedByUserEmailAddress *string `pulumi:"createdByUserEmailAddress"`
	// The date the role assignment was created.
	CreatedOn *string `pulumi:"createdOn"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The principal id of the user to whom the role was assigned.
	PrincipalId *string `pulumi:"principalId"`
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId *string `pulumi:"principalTenantId"`
	// The ID of the role definition.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The scope at which the role was assigned.
	Scope *string `pulumi:"scope"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType *string `pulumi:"userAuthenticationType"`
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress *string `pulumi:"userEmailAddress"`
}

type BillingRoleAssignmentByDepartmentState struct {
	// The principal Id of the user who created the role assignment.
	CreatedByPrincipalId pulumi.StringPtrInput
	// The tenant Id of the user who created the role assignment.
	CreatedByPrincipalTenantId pulumi.StringPtrInput
	// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
	CreatedByUserEmailAddress pulumi.StringPtrInput
	// The date the role assignment was created.
	CreatedOn pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The principal id of the user to whom the role was assigned.
	PrincipalId pulumi.StringPtrInput
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId pulumi.StringPtrInput
	// The ID of the role definition.
	RoleDefinitionId pulumi.StringPtrInput
	// The scope at which the role was assigned.
	Scope pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType pulumi.StringPtrInput
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress pulumi.StringPtrInput
}

func (BillingRoleAssignmentByDepartmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByDepartmentState)(nil)).Elem()
}

type billingRoleAssignmentByDepartmentArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName string `pulumi:"billingAccountName"`
	// The ID that uniquely identifies a role assignment.
	BillingRoleAssignmentName *string `pulumi:"billingRoleAssignmentName"`
	// The ID that uniquely identifies a department.
	DepartmentName string `pulumi:"departmentName"`
	// The principal id of the user to whom the role was assigned.
	PrincipalId *string `pulumi:"principalId"`
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId *string `pulumi:"principalTenantId"`
	// The ID of the role definition.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType *string `pulumi:"userAuthenticationType"`
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress *string `pulumi:"userEmailAddress"`
}

// The set of arguments for constructing a BillingRoleAssignmentByDepartment resource.
type BillingRoleAssignmentByDepartmentArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName pulumi.StringInput
	// The ID that uniquely identifies a role assignment.
	BillingRoleAssignmentName pulumi.StringPtrInput
	// The ID that uniquely identifies a department.
	DepartmentName pulumi.StringInput
	// The principal id of the user to whom the role was assigned.
	PrincipalId pulumi.StringPtrInput
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId pulumi.StringPtrInput
	// The ID of the role definition.
	RoleDefinitionId pulumi.StringPtrInput
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType pulumi.StringPtrInput
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress pulumi.StringPtrInput
}

func (BillingRoleAssignmentByDepartmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByDepartmentArgs)(nil)).Elem()
}

type BillingRoleAssignmentByDepartmentInput interface {
	pulumi.Input

	ToBillingRoleAssignmentByDepartmentOutput() BillingRoleAssignmentByDepartmentOutput
	ToBillingRoleAssignmentByDepartmentOutputWithContext(ctx context.Context) BillingRoleAssignmentByDepartmentOutput
}

func (*BillingRoleAssignmentByDepartment) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingRoleAssignmentByDepartment)(nil))
}

func (i *BillingRoleAssignmentByDepartment) ToBillingRoleAssignmentByDepartmentOutput() BillingRoleAssignmentByDepartmentOutput {
	return i.ToBillingRoleAssignmentByDepartmentOutputWithContext(context.Background())
}

func (i *BillingRoleAssignmentByDepartment) ToBillingRoleAssignmentByDepartmentOutputWithContext(ctx context.Context) BillingRoleAssignmentByDepartmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingRoleAssignmentByDepartmentOutput)
}

type BillingRoleAssignmentByDepartmentOutput struct {
	*pulumi.OutputState
}

func (BillingRoleAssignmentByDepartmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingRoleAssignmentByDepartment)(nil))
}

func (o BillingRoleAssignmentByDepartmentOutput) ToBillingRoleAssignmentByDepartmentOutput() BillingRoleAssignmentByDepartmentOutput {
	return o
}

func (o BillingRoleAssignmentByDepartmentOutput) ToBillingRoleAssignmentByDepartmentOutputWithContext(ctx context.Context) BillingRoleAssignmentByDepartmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BillingRoleAssignmentByDepartmentOutput{})
}
