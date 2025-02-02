// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package labservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a lab.
// API Version: 2018-10-15.
type Lab struct {
	pulumi.CustomResourceState

	// Object id of the user that created the lab.
	CreatedByObjectId pulumi.StringOutput `pulumi:"createdByObjectId"`
	// Lab creator name
	CreatedByUserPrincipalName pulumi.StringOutput `pulumi:"createdByUserPrincipalName"`
	// Creation date for the lab
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// Invitation code that users can use to join a lab.
	InvitationCode pulumi.StringOutput `pulumi:"invitationCode"`
	// The details of the latest operation. ex: status, error
	LatestOperationResult LatestOperationResultResponseOutput `pulumi:"latestOperationResult"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Maximum number of users allowed in the lab.
	MaxUsersInLab pulumi.IntPtrOutput `pulumi:"maxUsersInLab"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrOutput `pulumi:"uniqueIdentifier"`
	// Maximum duration a user can use an environment for in the lab.
	UsageQuota pulumi.StringPtrOutput `pulumi:"usageQuota"`
	// Lab user access mode (open to all vs. restricted to those listed on the lab).
	UserAccessMode pulumi.StringPtrOutput `pulumi:"userAccessMode"`
	// Maximum value MaxUsersInLab can be set to, as specified by the service
	UserQuota pulumi.IntOutput `pulumi:"userQuota"`
}

// NewLab registers a new resource with the given unique name, arguments, and options.
func NewLab(ctx *pulumi.Context,
	name string, args *LabArgs, opts ...pulumi.ResourceOption) (*Lab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabAccountName == nil {
		return nil, errors.New("invalid value for required argument 'LabAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:labservices:Lab"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20181015:Lab"),
		},
		{
			Type: pulumi.String("azure-nextgen:labservices/v20181015:Lab"),
		},
	})
	opts = append(opts, aliases)
	var resource Lab
	err := ctx.RegisterResource("azure-native:labservices:Lab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLab gets an existing Lab resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabState, opts ...pulumi.ResourceOption) (*Lab, error) {
	var resource Lab
	err := ctx.ReadResource("azure-native:labservices:Lab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Lab resources.
type labState struct {
	// Object id of the user that created the lab.
	CreatedByObjectId *string `pulumi:"createdByObjectId"`
	// Lab creator name
	CreatedByUserPrincipalName *string `pulumi:"createdByUserPrincipalName"`
	// Creation date for the lab
	CreatedDate *string `pulumi:"createdDate"`
	// Invitation code that users can use to join a lab.
	InvitationCode *string `pulumi:"invitationCode"`
	// The details of the latest operation. ex: status, error
	LatestOperationResult *LatestOperationResultResponse `pulumi:"latestOperationResult"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// Maximum number of users allowed in the lab.
	MaxUsersInLab *int `pulumi:"maxUsersInLab"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
	// Maximum duration a user can use an environment for in the lab.
	UsageQuota *string `pulumi:"usageQuota"`
	// Lab user access mode (open to all vs. restricted to those listed on the lab).
	UserAccessMode *string `pulumi:"userAccessMode"`
	// Maximum value MaxUsersInLab can be set to, as specified by the service
	UserQuota *int `pulumi:"userQuota"`
}

type LabState struct {
	// Object id of the user that created the lab.
	CreatedByObjectId pulumi.StringPtrInput
	// Lab creator name
	CreatedByUserPrincipalName pulumi.StringPtrInput
	// Creation date for the lab
	CreatedDate pulumi.StringPtrInput
	// Invitation code that users can use to join a lab.
	InvitationCode pulumi.StringPtrInput
	// The details of the latest operation. ex: status, error
	LatestOperationResult LatestOperationResultResponsePtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// Maximum number of users allowed in the lab.
	MaxUsersInLab pulumi.IntPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrInput
	// Maximum duration a user can use an environment for in the lab.
	UsageQuota pulumi.StringPtrInput
	// Lab user access mode (open to all vs. restricted to those listed on the lab).
	UserAccessMode pulumi.StringPtrInput
	// Maximum value MaxUsersInLab can be set to, as specified by the service
	UserQuota pulumi.IntPtrInput
}

func (LabState) ElementType() reflect.Type {
	return reflect.TypeOf((*labState)(nil)).Elem()
}

type labArgs struct {
	// The name of the lab Account.
	LabAccountName string `pulumi:"labAccountName"`
	// The name of the lab.
	LabName *string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// Maximum number of users allowed in the lab.
	MaxUsersInLab *int `pulumi:"maxUsersInLab"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
	// Maximum duration a user can use an environment for in the lab.
	UsageQuota *string `pulumi:"usageQuota"`
	// Lab user access mode (open to all vs. restricted to those listed on the lab).
	UserAccessMode *string `pulumi:"userAccessMode"`
}

// The set of arguments for constructing a Lab resource.
type LabArgs struct {
	// The name of the lab Account.
	LabAccountName pulumi.StringInput
	// The name of the lab.
	LabName pulumi.StringPtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// Maximum number of users allowed in the lab.
	MaxUsersInLab pulumi.IntPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrInput
	// Maximum duration a user can use an environment for in the lab.
	UsageQuota pulumi.StringPtrInput
	// Lab user access mode (open to all vs. restricted to those listed on the lab).
	UserAccessMode pulumi.StringPtrInput
}

func (LabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labArgs)(nil)).Elem()
}

type LabInput interface {
	pulumi.Input

	ToLabOutput() LabOutput
	ToLabOutputWithContext(ctx context.Context) LabOutput
}

func (*Lab) ElementType() reflect.Type {
	return reflect.TypeOf((*Lab)(nil))
}

func (i *Lab) ToLabOutput() LabOutput {
	return i.ToLabOutputWithContext(context.Background())
}

func (i *Lab) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabOutput)
}

type LabOutput struct {
	*pulumi.OutputState
}

func (LabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Lab)(nil))
}

func (o LabOutput) ToLabOutput() LabOutput {
	return o
}

func (o LabOutput) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LabOutput{})
}
