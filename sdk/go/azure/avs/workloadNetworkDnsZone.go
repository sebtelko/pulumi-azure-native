// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NSX DNS Zone
// API Version: 2020-07-17-preview.
type WorkloadNetworkDnsZone struct {
	pulumi.CustomResourceState

	// Display name of the DNS Zone.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// DNS Server IP array of the DNS Zone.
	DnsServerIps pulumi.StringArrayOutput `pulumi:"dnsServerIps"`
	// Number of DNS Services using the DNS zone.
	DnsServices pulumi.Float64PtrOutput `pulumi:"dnsServices"`
	// Domain names of the DNS Zone.
	Domain pulumi.StringArrayOutput `pulumi:"domain"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// NSX revision number.
	Revision pulumi.Float64PtrOutput `pulumi:"revision"`
	// Source IP of the DNS Zone.
	SourceIp pulumi.StringPtrOutput `pulumi:"sourceIp"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadNetworkDnsZone registers a new resource with the given unique name, arguments, and options.
func NewWorkloadNetworkDnsZone(ctx *pulumi.Context,
	name string, args *WorkloadNetworkDnsZoneArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:avs:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20200717preview:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkDnsZone"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210101preview:WorkloadNetworkDnsZone"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkDnsZone
	err := ctx.RegisterResource("azure-native:avs:WorkloadNetworkDnsZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadNetworkDnsZone gets an existing WorkloadNetworkDnsZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadNetworkDnsZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkDnsZoneState, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsZone, error) {
	var resource WorkloadNetworkDnsZone
	err := ctx.ReadResource("azure-native:avs:WorkloadNetworkDnsZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadNetworkDnsZone resources.
type workloadNetworkDnsZoneState struct {
	// Display name of the DNS Zone.
	DisplayName *string `pulumi:"displayName"`
	// DNS Server IP array of the DNS Zone.
	DnsServerIps []string `pulumi:"dnsServerIps"`
	// Number of DNS Services using the DNS zone.
	DnsServices *float64 `pulumi:"dnsServices"`
	// Domain names of the DNS Zone.
	Domain []string `pulumi:"domain"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The provisioning state
	ProvisioningState *string `pulumi:"provisioningState"`
	// NSX revision number.
	Revision *float64 `pulumi:"revision"`
	// Source IP of the DNS Zone.
	SourceIp *string `pulumi:"sourceIp"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WorkloadNetworkDnsZoneState struct {
	// Display name of the DNS Zone.
	DisplayName pulumi.StringPtrInput
	// DNS Server IP array of the DNS Zone.
	DnsServerIps pulumi.StringArrayInput
	// Number of DNS Services using the DNS zone.
	DnsServices pulumi.Float64PtrInput
	// Domain names of the DNS Zone.
	Domain pulumi.StringArrayInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The provisioning state
	ProvisioningState pulumi.StringPtrInput
	// NSX revision number.
	Revision pulumi.Float64PtrInput
	// Source IP of the DNS Zone.
	SourceIp pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WorkloadNetworkDnsZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDnsZoneState)(nil)).Elem()
}

type workloadNetworkDnsZoneArgs struct {
	// Display name of the DNS Zone.
	DisplayName *string `pulumi:"displayName"`
	// DNS Server IP array of the DNS Zone.
	DnsServerIps []string `pulumi:"dnsServerIps"`
	// Number of DNS Services using the DNS zone.
	DnsServices *float64 `pulumi:"dnsServices"`
	// NSX DNS Zone identifier. Generally the same as the DNS Zone's display name
	DnsZoneId *string `pulumi:"dnsZoneId"`
	// Domain names of the DNS Zone.
	Domain []string `pulumi:"domain"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// NSX revision number.
	Revision *float64 `pulumi:"revision"`
	// Source IP of the DNS Zone.
	SourceIp *string `pulumi:"sourceIp"`
}

// The set of arguments for constructing a WorkloadNetworkDnsZone resource.
type WorkloadNetworkDnsZoneArgs struct {
	// Display name of the DNS Zone.
	DisplayName pulumi.StringPtrInput
	// DNS Server IP array of the DNS Zone.
	DnsServerIps pulumi.StringArrayInput
	// Number of DNS Services using the DNS zone.
	DnsServices pulumi.Float64PtrInput
	// NSX DNS Zone identifier. Generally the same as the DNS Zone's display name
	DnsZoneId pulumi.StringPtrInput
	// Domain names of the DNS Zone.
	Domain pulumi.StringArrayInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// NSX revision number.
	Revision pulumi.Float64PtrInput
	// Source IP of the DNS Zone.
	SourceIp pulumi.StringPtrInput
}

func (WorkloadNetworkDnsZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDnsZoneArgs)(nil)).Elem()
}

type WorkloadNetworkDnsZoneInput interface {
	pulumi.Input

	ToWorkloadNetworkDnsZoneOutput() WorkloadNetworkDnsZoneOutput
	ToWorkloadNetworkDnsZoneOutputWithContext(ctx context.Context) WorkloadNetworkDnsZoneOutput
}

func (*WorkloadNetworkDnsZone) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDnsZone)(nil))
}

func (i *WorkloadNetworkDnsZone) ToWorkloadNetworkDnsZoneOutput() WorkloadNetworkDnsZoneOutput {
	return i.ToWorkloadNetworkDnsZoneOutputWithContext(context.Background())
}

func (i *WorkloadNetworkDnsZone) ToWorkloadNetworkDnsZoneOutputWithContext(ctx context.Context) WorkloadNetworkDnsZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDnsZoneOutput)
}

type WorkloadNetworkDnsZoneOutput struct {
	*pulumi.OutputState
}

func (WorkloadNetworkDnsZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDnsZone)(nil))
}

func (o WorkloadNetworkDnsZoneOutput) ToWorkloadNetworkDnsZoneOutput() WorkloadNetworkDnsZoneOutput {
	return o
}

func (o WorkloadNetworkDnsZoneOutput) ToWorkloadNetworkDnsZoneOutputWithContext(ctx context.Context) WorkloadNetworkDnsZoneOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkDnsZoneOutput{})
}
