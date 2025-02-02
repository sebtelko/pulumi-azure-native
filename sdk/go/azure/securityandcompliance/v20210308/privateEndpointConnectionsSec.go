// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210308

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Private Endpoint Connection resource.
type PrivateEndpointConnectionsSec struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Required property for system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnectionsSec registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnectionsSec(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionsSecArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsSec, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance/v20210308:PrivateEndpointConnectionsSec"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance:PrivateEndpointConnectionsSec"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance:PrivateEndpointConnectionsSec"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210111:PrivateEndpointConnectionsSec"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance/v20210111:PrivateEndpointConnectionsSec"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionsSec
	err := ctx.RegisterResource("azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsSec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnectionsSec gets an existing PrivateEndpointConnectionsSec resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnectionsSec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionsSecState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsSec, error) {
	var resource PrivateEndpointConnectionsSec
	err := ctx.ReadResource("azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsSec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnectionsSec resources.
type privateEndpointConnectionsSecState struct {
	// The name of the resource
	Name *string `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Required property for system data
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type PrivateEndpointConnectionsSecState struct {
	// The name of the resource
	Name pulumi.StringPtrInput
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponsePtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrInput
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringPtrInput
	// Required property for system data
	SystemData SystemDataResponsePtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (PrivateEndpointConnectionsSecState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsSecState)(nil)).Elem()
}

type privateEndpointConnectionsSecArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a PrivateEndpointConnectionsSec resource.
type PrivateEndpointConnectionsSecArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// The name of the service instance.
	ResourceName pulumi.StringInput
}

func (PrivateEndpointConnectionsSecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsSecArgs)(nil)).Elem()
}

type PrivateEndpointConnectionsSecInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionsSecOutput() PrivateEndpointConnectionsSecOutput
	ToPrivateEndpointConnectionsSecOutputWithContext(ctx context.Context) PrivateEndpointConnectionsSecOutput
}

func (*PrivateEndpointConnectionsSec) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionsSec)(nil))
}

func (i *PrivateEndpointConnectionsSec) ToPrivateEndpointConnectionsSecOutput() PrivateEndpointConnectionsSecOutput {
	return i.ToPrivateEndpointConnectionsSecOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionsSec) ToPrivateEndpointConnectionsSecOutputWithContext(ctx context.Context) PrivateEndpointConnectionsSecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionsSecOutput)
}

type PrivateEndpointConnectionsSecOutput struct {
	*pulumi.OutputState
}

func (PrivateEndpointConnectionsSecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionsSec)(nil))
}

func (o PrivateEndpointConnectionsSecOutput) ToPrivateEndpointConnectionsSecOutput() PrivateEndpointConnectionsSecOutput {
	return o
}

func (o PrivateEndpointConnectionsSecOutput) ToPrivateEndpointConnectionsSecOutputWithContext(ctx context.Context) PrivateEndpointConnectionsSecOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionsSecOutput{})
}
