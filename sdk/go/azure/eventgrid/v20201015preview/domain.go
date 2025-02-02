// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EventGrid Domain.
type Domain struct {
	pulumi.CustomResourceState

	// Endpoint for the domain.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Identity information for the resource.
	Identity IdentityInfoResponsePtrOutput `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleResponseArrayOutput `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema pulumi.StringPtrOutput `pulumi:"inputSchema"`
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping JsonInputSchemaMappingResponsePtrOutput `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Metric resource id for the domain.
	MetricResourceId pulumi.StringOutput `pulumi:"metricResourceId"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the domain.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The Sku pricing tier for the domain.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// The system metadata relating to Domain resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.InputSchema == nil {
		args.InputSchema = pulumi.StringPtr("EventGridSchema")
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20201015preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180915preview:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20180915preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190201preview:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20190201preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190601:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20190601:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200101preview:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200101preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200401preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:Domain"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200601:Domain"),
		},
	})
	opts = append(opts, aliases)
	var resource Domain
	err := ctx.RegisterResource("azure-native:eventgrid/v20201015preview:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("azure-native:eventgrid/v20201015preview:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// Endpoint for the domain.
	Endpoint *string `pulumi:"endpoint"`
	// Identity information for the resource.
	Identity *IdentityInfoResponse `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRuleResponse `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema *string `pulumi:"inputSchema"`
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping *JsonInputSchemaMappingResponse `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// Metric resource id for the domain.
	MetricResourceId *string `pulumi:"metricResourceId"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Provisioning state of the domain.
	ProvisioningState *string `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The Sku pricing tier for the domain.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// The system metadata relating to Domain resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource.
	Type *string `pulumi:"type"`
}

type DomainState struct {
	// Endpoint for the domain.
	Endpoint pulumi.StringPtrInput
	// Identity information for the resource.
	Identity IdentityInfoResponsePtrInput
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleResponseArrayInput
	// This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema pulumi.StringPtrInput
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping JsonInputSchemaMappingResponsePtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// Metric resource id for the domain.
	MetricResourceId pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// List of private endpoint connections.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayInput
	// Provisioning state of the domain.
	ProvisioningState pulumi.StringPtrInput
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrInput
	// The Sku pricing tier for the domain.
	Sku ResourceSkuResponsePtrInput
	// The system metadata relating to Domain resource.
	SystemData SystemDataResponsePtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
	// Type of the resource.
	Type pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Name of the domain.
	DomainName *string `pulumi:"domainName"`
	// Identity information for the resource.
	Identity *IdentityInfo `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRule `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema *string `pulumi:"inputSchema"`
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping *JsonInputSchemaMapping `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Sku pricing tier for the domain.
	Sku *ResourceSku `pulumi:"sku"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Name of the domain.
	DomainName pulumi.StringPtrInput
	// Identity information for the resource.
	Identity IdentityInfoPtrInput
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleArrayInput
	// This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema pulumi.StringPtrInput
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping JsonInputSchemaMappingPtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// The Sku pricing tier for the domain.
	Sku ResourceSkuPtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct {
	*pulumi.OutputState
}

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
