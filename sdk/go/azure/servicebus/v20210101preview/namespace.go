// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of a namespace resource.
type Namespace struct {
	pulumi.CustomResourceState

	// The time the namespace was created
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Properties of BYOK Encryption description
	Encryption EncryptionResponsePtrOutput `pulumi:"encryption"`
	// Properties of BYOK Identity description
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The Geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Identifier for Azure Insights metrics
	MetricId pulumi.StringOutput `pulumi:"metricId"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the namespace.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint pulumi.StringOutput `pulumi:"serviceBusEndpoint"`
	// Properties of SKU
	Sku SBSkuResponsePtrOutput `pulumi:"sku"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20210101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20140901:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20140901:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20150801:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20170401:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:Namespace"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20180101preview:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:servicebus/v20210101preview:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:servicebus/v20210101preview:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// The time the namespace was created
	CreatedAt *string `pulumi:"createdAt"`
	// Properties of BYOK Encryption description
	Encryption *EncryptionResponse `pulumi:"encryption"`
	// Properties of BYOK Identity description
	Identity *IdentityResponse `pulumi:"identity"`
	// The Geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Identifier for Azure Insights metrics
	MetricId *string `pulumi:"metricId"`
	// Resource name
	Name *string `pulumi:"name"`
	// List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Provisioning state of the namespace.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint *string `pulumi:"serviceBusEndpoint"`
	// Properties of SKU
	Sku *SBSkuResponse `pulumi:"sku"`
	// The system meta data relating to this resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

type NamespaceState struct {
	// The time the namespace was created
	CreatedAt pulumi.StringPtrInput
	// Properties of BYOK Encryption description
	Encryption EncryptionResponsePtrInput
	// Properties of BYOK Identity description
	Identity IdentityResponsePtrInput
	// The Geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Identifier for Azure Insights metrics
	MetricId pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// List of private endpoint connections.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayInput
	// Provisioning state of the namespace.
	ProvisioningState pulumi.StringPtrInput
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint pulumi.StringPtrInput
	// Properties of SKU
	Sku SBSkuResponsePtrInput
	// The system meta data relating to this resource.
	SystemData SystemDataResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// The time the namespace was updated.
	UpdatedAt pulumi.StringPtrInput
	// Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Properties of BYOK Encryption description
	Encryption *Encryption `pulumi:"encryption"`
	// Properties of BYOK Identity description
	Identity *Identity `pulumi:"identity"`
	// The Geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The namespace name.
	NamespaceName *string `pulumi:"namespaceName"`
	// List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnectionType `pulumi:"privateEndpointConnections"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Properties of SKU
	Sku *SBSku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Properties of BYOK Encryption description
	Encryption EncryptionPtrInput
	// Properties of BYOK Identity description
	Identity IdentityPtrInput
	// The Geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The namespace name.
	NamespaceName pulumi.StringPtrInput
	// List of private endpoint connections.
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Properties of SKU
	Sku SBSkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil))
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

type NamespaceOutput struct {
	*pulumi.OutputState
}

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil))
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
}
