// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// App Service Environment ARM resource.
type AppServiceEnvironment struct {
	pulumi.CustomResourceState

	// List of comma separated strings describing which VM sizes are allowed for front-ends.
	AllowedMultiSizes pulumi.StringOutput `pulumi:"allowedMultiSizes"`
	// List of comma separated strings describing which VM sizes are allowed for workers.
	AllowedWorkerSizes pulumi.StringOutput `pulumi:"allowedWorkerSizes"`
	// API Management Account associated with the App Service Environment.
	ApiManagementAccountId pulumi.StringPtrOutput `pulumi:"apiManagementAccountId"`
	// Custom settings for changing the behavior of the App Service Environment.
	ClusterSettings NameValuePairResponseArrayOutput `pulumi:"clusterSettings"`
	// Edition of the metadata database for the App Service Environment, e.g. "Standard".
	DatabaseEdition pulumi.StringOutput `pulumi:"databaseEdition"`
	// Service objective of the metadata database for the App Service Environment, e.g. "S0".
	DatabaseServiceObjective pulumi.StringOutput `pulumi:"databaseServiceObjective"`
	// Default Scale Factor for FrontEnds.
	DefaultFrontEndScaleFactor pulumi.IntOutput `pulumi:"defaultFrontEndScaleFactor"`
	// DNS suffix of the App Service Environment.
	DnsSuffix pulumi.StringPtrOutput `pulumi:"dnsSuffix"`
	// True/false indicating whether the App Service Environment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
	// (most likely because NSG blocked the incoming traffic).
	DynamicCacheEnabled pulumi.BoolPtrOutput `pulumi:"dynamicCacheEnabled"`
	// Current total, used, and available worker capacities.
	EnvironmentCapacities StampCapacityResponseArrayOutput `pulumi:"environmentCapacities"`
	// True/false indicating whether the App Service Environment is healthy.
	EnvironmentIsHealthy pulumi.BoolOutput `pulumi:"environmentIsHealthy"`
	// Detailed message about with results of the last check of the App Service Environment.
	EnvironmentStatus pulumi.StringOutput `pulumi:"environmentStatus"`
	// Scale factor for front-ends.
	FrontEndScaleFactor pulumi.IntPtrOutput `pulumi:"frontEndScaleFactor"`
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
	InternalLoadBalancingMode pulumi.StringPtrOutput `pulumi:"internalLoadBalancingMode"`
	// Number of IP SSL addresses reserved for the App Service Environment.
	IpsslAddressCount pulumi.IntPtrOutput `pulumi:"ipsslAddressCount"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Last deployment action on the App Service Environment.
	LastAction pulumi.StringOutput `pulumi:"lastAction"`
	// Result of the last deployment action on the App Service Environment.
	LastActionResult pulumi.StringOutput `pulumi:"lastActionResult"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Maximum number of VMs in the App Service Environment.
	MaximumNumberOfMachines pulumi.IntOutput `pulumi:"maximumNumberOfMachines"`
	// Number of front-end instances.
	MultiRoleCount pulumi.IntPtrOutput `pulumi:"multiRoleCount"`
	// Front-end VM size, e.g. "Medium", "Large".
	MultiSize pulumi.StringPtrOutput `pulumi:"multiSize"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Access control list for controlling traffic to the App Service Environment.
	NetworkAccessControlList NetworkAccessControlEntryResponseArrayOutput `pulumi:"networkAccessControlList"`
	// Provisioning state of the App Service Environment.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource group of the App Service Environment.
	ResourceGroup pulumi.StringOutput `pulumi:"resourceGroup"`
	// Current status of the App Service Environment.
	Status pulumi.StringOutput `pulumi:"status"`
	// Subscription of the App Service Environment.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// <code>true</code> if the App Service Environment is suspended; otherwise, <code>false</code>. The environment can be suspended, e.g. when the management endpoint is no longer available
	//  (most likely because NSG blocked the incoming traffic).
	Suspended pulumi.BoolPtrOutput `pulumi:"suspended"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Number of upgrade domains of the App Service Environment.
	UpgradeDomains pulumi.IntOutput `pulumi:"upgradeDomains"`
	// User added ip ranges to whitelist on ASE db
	UserWhitelistedIpRanges pulumi.StringArrayOutput `pulumi:"userWhitelistedIpRanges"`
	// Description of IP SSL mapping for the App Service Environment.
	VipMappings VirtualIPMappingResponseArrayOutput `pulumi:"vipMappings"`
	// Description of the Virtual Network.
	VirtualNetwork VirtualNetworkProfileResponseOutput `pulumi:"virtualNetwork"`
	// Name of the Virtual Network for the App Service Environment.
	VnetName pulumi.StringPtrOutput `pulumi:"vnetName"`
	// Resource group of the Virtual Network.
	VnetResourceGroupName pulumi.StringPtrOutput `pulumi:"vnetResourceGroupName"`
	// Subnet of the Virtual Network.
	VnetSubnetName pulumi.StringPtrOutput `pulumi:"vnetSubnetName"`
	// Description of worker pools with worker size IDs, VM sizes, and number of workers in each pool.
	WorkerPools WorkerPoolResponseArrayOutput `pulumi:"workerPools"`
}

// NewAppServiceEnvironment registers a new resource with the given unique name, arguments, and options.
func NewAppServiceEnvironment(ctx *pulumi.Context,
	name string, args *AppServiceEnvironmentArgs, opts ...pulumi.ResourceOption) (*AppServiceEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetwork == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetwork'")
	}
	if args.WorkerPools == nil {
		return nil, errors.New("invalid value for required argument 'WorkerPools'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20160901:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:AppServiceEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:AppServiceEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServiceEnvironment
	err := ctx.RegisterResource("azure-native:web/v20160901:AppServiceEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppServiceEnvironment gets an existing AppServiceEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppServiceEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceEnvironmentState, opts ...pulumi.ResourceOption) (*AppServiceEnvironment, error) {
	var resource AppServiceEnvironment
	err := ctx.ReadResource("azure-native:web/v20160901:AppServiceEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppServiceEnvironment resources.
type appServiceEnvironmentState struct {
	// List of comma separated strings describing which VM sizes are allowed for front-ends.
	AllowedMultiSizes *string `pulumi:"allowedMultiSizes"`
	// List of comma separated strings describing which VM sizes are allowed for workers.
	AllowedWorkerSizes *string `pulumi:"allowedWorkerSizes"`
	// API Management Account associated with the App Service Environment.
	ApiManagementAccountId *string `pulumi:"apiManagementAccountId"`
	// Custom settings for changing the behavior of the App Service Environment.
	ClusterSettings []NameValuePairResponse `pulumi:"clusterSettings"`
	// Edition of the metadata database for the App Service Environment, e.g. "Standard".
	DatabaseEdition *string `pulumi:"databaseEdition"`
	// Service objective of the metadata database for the App Service Environment, e.g. "S0".
	DatabaseServiceObjective *string `pulumi:"databaseServiceObjective"`
	// Default Scale Factor for FrontEnds.
	DefaultFrontEndScaleFactor *int `pulumi:"defaultFrontEndScaleFactor"`
	// DNS suffix of the App Service Environment.
	DnsSuffix *string `pulumi:"dnsSuffix"`
	// True/false indicating whether the App Service Environment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
	// (most likely because NSG blocked the incoming traffic).
	DynamicCacheEnabled *bool `pulumi:"dynamicCacheEnabled"`
	// Current total, used, and available worker capacities.
	EnvironmentCapacities []StampCapacityResponse `pulumi:"environmentCapacities"`
	// True/false indicating whether the App Service Environment is healthy.
	EnvironmentIsHealthy *bool `pulumi:"environmentIsHealthy"`
	// Detailed message about with results of the last check of the App Service Environment.
	EnvironmentStatus *string `pulumi:"environmentStatus"`
	// Scale factor for front-ends.
	FrontEndScaleFactor *int `pulumi:"frontEndScaleFactor"`
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
	InternalLoadBalancingMode *string `pulumi:"internalLoadBalancingMode"`
	// Number of IP SSL addresses reserved for the App Service Environment.
	IpsslAddressCount *int `pulumi:"ipsslAddressCount"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Last deployment action on the App Service Environment.
	LastAction *string `pulumi:"lastAction"`
	// Result of the last deployment action on the App Service Environment.
	LastActionResult *string `pulumi:"lastActionResult"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Maximum number of VMs in the App Service Environment.
	MaximumNumberOfMachines *int `pulumi:"maximumNumberOfMachines"`
	// Number of front-end instances.
	MultiRoleCount *int `pulumi:"multiRoleCount"`
	// Front-end VM size, e.g. "Medium", "Large".
	MultiSize *string `pulumi:"multiSize"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// Access control list for controlling traffic to the App Service Environment.
	NetworkAccessControlList []NetworkAccessControlEntryResponse `pulumi:"networkAccessControlList"`
	// Provisioning state of the App Service Environment.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource group of the App Service Environment.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Current status of the App Service Environment.
	Status *string `pulumi:"status"`
	// Subscription of the App Service Environment.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// <code>true</code> if the App Service Environment is suspended; otherwise, <code>false</code>. The environment can be suspended, e.g. when the management endpoint is no longer available
	//  (most likely because NSG blocked the incoming traffic).
	Suspended *bool `pulumi:"suspended"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// Number of upgrade domains of the App Service Environment.
	UpgradeDomains *int `pulumi:"upgradeDomains"`
	// User added ip ranges to whitelist on ASE db
	UserWhitelistedIpRanges []string `pulumi:"userWhitelistedIpRanges"`
	// Description of IP SSL mapping for the App Service Environment.
	VipMappings []VirtualIPMappingResponse `pulumi:"vipMappings"`
	// Description of the Virtual Network.
	VirtualNetwork *VirtualNetworkProfileResponse `pulumi:"virtualNetwork"`
	// Name of the Virtual Network for the App Service Environment.
	VnetName *string `pulumi:"vnetName"`
	// Resource group of the Virtual Network.
	VnetResourceGroupName *string `pulumi:"vnetResourceGroupName"`
	// Subnet of the Virtual Network.
	VnetSubnetName *string `pulumi:"vnetSubnetName"`
	// Description of worker pools with worker size IDs, VM sizes, and number of workers in each pool.
	WorkerPools []WorkerPoolResponse `pulumi:"workerPools"`
}

type AppServiceEnvironmentState struct {
	// List of comma separated strings describing which VM sizes are allowed for front-ends.
	AllowedMultiSizes pulumi.StringPtrInput
	// List of comma separated strings describing which VM sizes are allowed for workers.
	AllowedWorkerSizes pulumi.StringPtrInput
	// API Management Account associated with the App Service Environment.
	ApiManagementAccountId pulumi.StringPtrInput
	// Custom settings for changing the behavior of the App Service Environment.
	ClusterSettings NameValuePairResponseArrayInput
	// Edition of the metadata database for the App Service Environment, e.g. "Standard".
	DatabaseEdition pulumi.StringPtrInput
	// Service objective of the metadata database for the App Service Environment, e.g. "S0".
	DatabaseServiceObjective pulumi.StringPtrInput
	// Default Scale Factor for FrontEnds.
	DefaultFrontEndScaleFactor pulumi.IntPtrInput
	// DNS suffix of the App Service Environment.
	DnsSuffix pulumi.StringPtrInput
	// True/false indicating whether the App Service Environment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
	// (most likely because NSG blocked the incoming traffic).
	DynamicCacheEnabled pulumi.BoolPtrInput
	// Current total, used, and available worker capacities.
	EnvironmentCapacities StampCapacityResponseArrayInput
	// True/false indicating whether the App Service Environment is healthy.
	EnvironmentIsHealthy pulumi.BoolPtrInput
	// Detailed message about with results of the last check of the App Service Environment.
	EnvironmentStatus pulumi.StringPtrInput
	// Scale factor for front-ends.
	FrontEndScaleFactor pulumi.IntPtrInput
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
	InternalLoadBalancingMode pulumi.StringPtrInput
	// Number of IP SSL addresses reserved for the App Service Environment.
	IpsslAddressCount pulumi.IntPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Last deployment action on the App Service Environment.
	LastAction pulumi.StringPtrInput
	// Result of the last deployment action on the App Service Environment.
	LastActionResult pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Maximum number of VMs in the App Service Environment.
	MaximumNumberOfMachines pulumi.IntPtrInput
	// Number of front-end instances.
	MultiRoleCount pulumi.IntPtrInput
	// Front-end VM size, e.g. "Medium", "Large".
	MultiSize pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// Access control list for controlling traffic to the App Service Environment.
	NetworkAccessControlList NetworkAccessControlEntryResponseArrayInput
	// Provisioning state of the App Service Environment.
	ProvisioningState pulumi.StringPtrInput
	// Resource group of the App Service Environment.
	ResourceGroup pulumi.StringPtrInput
	// Current status of the App Service Environment.
	Status pulumi.StringPtrInput
	// Subscription of the App Service Environment.
	SubscriptionId pulumi.StringPtrInput
	// <code>true</code> if the App Service Environment is suspended; otherwise, <code>false</code>. The environment can be suspended, e.g. when the management endpoint is no longer available
	//  (most likely because NSG blocked the incoming traffic).
	Suspended pulumi.BoolPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// Number of upgrade domains of the App Service Environment.
	UpgradeDomains pulumi.IntPtrInput
	// User added ip ranges to whitelist on ASE db
	UserWhitelistedIpRanges pulumi.StringArrayInput
	// Description of IP SSL mapping for the App Service Environment.
	VipMappings VirtualIPMappingResponseArrayInput
	// Description of the Virtual Network.
	VirtualNetwork VirtualNetworkProfileResponsePtrInput
	// Name of the Virtual Network for the App Service Environment.
	VnetName pulumi.StringPtrInput
	// Resource group of the Virtual Network.
	VnetResourceGroupName pulumi.StringPtrInput
	// Subnet of the Virtual Network.
	VnetSubnetName pulumi.StringPtrInput
	// Description of worker pools with worker size IDs, VM sizes, and number of workers in each pool.
	WorkerPools WorkerPoolResponseArrayInput
}

func (AppServiceEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentState)(nil)).Elem()
}

type appServiceEnvironmentArgs struct {
	// API Management Account associated with the App Service Environment.
	ApiManagementAccountId *string `pulumi:"apiManagementAccountId"`
	// Custom settings for changing the behavior of the App Service Environment.
	ClusterSettings []NameValuePair `pulumi:"clusterSettings"`
	// DNS suffix of the App Service Environment.
	DnsSuffix *string `pulumi:"dnsSuffix"`
	// True/false indicating whether the App Service Environment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
	// (most likely because NSG blocked the incoming traffic).
	DynamicCacheEnabled *bool `pulumi:"dynamicCacheEnabled"`
	// Scale factor for front-ends.
	FrontEndScaleFactor *int `pulumi:"frontEndScaleFactor"`
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
	InternalLoadBalancingMode *string `pulumi:"internalLoadBalancingMode"`
	// Number of IP SSL addresses reserved for the App Service Environment.
	IpsslAddressCount *int `pulumi:"ipsslAddressCount"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Number of front-end instances.
	MultiRoleCount *int `pulumi:"multiRoleCount"`
	// Front-end VM size, e.g. "Medium", "Large".
	MultiSize *string `pulumi:"multiSize"`
	// Name of the App Service Environment.
	Name *string `pulumi:"name"`
	// Access control list for controlling traffic to the App Service Environment.
	NetworkAccessControlList []NetworkAccessControlEntry `pulumi:"networkAccessControlList"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// <code>true</code> if the App Service Environment is suspended; otherwise, <code>false</code>. The environment can be suspended, e.g. when the management endpoint is no longer available
	//  (most likely because NSG blocked the incoming traffic).
	Suspended *bool `pulumi:"suspended"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// User added ip ranges to whitelist on ASE db
	UserWhitelistedIpRanges []string `pulumi:"userWhitelistedIpRanges"`
	// Description of the Virtual Network.
	VirtualNetwork VirtualNetworkProfile `pulumi:"virtualNetwork"`
	// Name of the Virtual Network for the App Service Environment.
	VnetName *string `pulumi:"vnetName"`
	// Resource group of the Virtual Network.
	VnetResourceGroupName *string `pulumi:"vnetResourceGroupName"`
	// Subnet of the Virtual Network.
	VnetSubnetName *string `pulumi:"vnetSubnetName"`
	// Description of worker pools with worker size IDs, VM sizes, and number of workers in each pool.
	WorkerPools []WorkerPool `pulumi:"workerPools"`
}

// The set of arguments for constructing a AppServiceEnvironment resource.
type AppServiceEnvironmentArgs struct {
	// API Management Account associated with the App Service Environment.
	ApiManagementAccountId pulumi.StringPtrInput
	// Custom settings for changing the behavior of the App Service Environment.
	ClusterSettings NameValuePairArrayInput
	// DNS suffix of the App Service Environment.
	DnsSuffix pulumi.StringPtrInput
	// True/false indicating whether the App Service Environment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
	// (most likely because NSG blocked the incoming traffic).
	DynamicCacheEnabled pulumi.BoolPtrInput
	// Scale factor for front-ends.
	FrontEndScaleFactor pulumi.IntPtrInput
	// Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
	InternalLoadBalancingMode *InternalLoadBalancingMode
	// Number of IP SSL addresses reserved for the App Service Environment.
	IpsslAddressCount pulumi.IntPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Number of front-end instances.
	MultiRoleCount pulumi.IntPtrInput
	// Front-end VM size, e.g. "Medium", "Large".
	MultiSize pulumi.StringPtrInput
	// Name of the App Service Environment.
	Name pulumi.StringPtrInput
	// Access control list for controlling traffic to the App Service Environment.
	NetworkAccessControlList NetworkAccessControlEntryArrayInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// <code>true</code> if the App Service Environment is suspended; otherwise, <code>false</code>. The environment can be suspended, e.g. when the management endpoint is no longer available
	//  (most likely because NSG blocked the incoming traffic).
	Suspended pulumi.BoolPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// User added ip ranges to whitelist on ASE db
	UserWhitelistedIpRanges pulumi.StringArrayInput
	// Description of the Virtual Network.
	VirtualNetwork VirtualNetworkProfileInput
	// Name of the Virtual Network for the App Service Environment.
	VnetName pulumi.StringPtrInput
	// Resource group of the Virtual Network.
	VnetResourceGroupName pulumi.StringPtrInput
	// Subnet of the Virtual Network.
	VnetSubnetName pulumi.StringPtrInput
	// Description of worker pools with worker size IDs, VM sizes, and number of workers in each pool.
	WorkerPools WorkerPoolArrayInput
}

func (AppServiceEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentArgs)(nil)).Elem()
}

type AppServiceEnvironmentInput interface {
	pulumi.Input

	ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput
	ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput
}

func (*AppServiceEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServiceEnvironment)(nil))
}

func (i *AppServiceEnvironment) ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput {
	return i.ToAppServiceEnvironmentOutputWithContext(context.Background())
}

func (i *AppServiceEnvironment) ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceEnvironmentOutput)
}

type AppServiceEnvironmentOutput struct {
	*pulumi.OutputState
}

func (AppServiceEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServiceEnvironment)(nil))
}

func (o AppServiceEnvironmentOutput) ToAppServiceEnvironmentOutput() AppServiceEnvironmentOutput {
	return o
}

func (o AppServiceEnvironmentOutput) ToAppServiceEnvironmentOutputWithContext(ctx context.Context) AppServiceEnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AppServiceEnvironmentOutput{})
}
