// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Domain service.
type DomainService struct {
	pulumi.CustomResourceState

	// Deployment Id
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// Domain Configuration Type
	DomainConfigurationType pulumi.StringPtrOutput `pulumi:"domainConfigurationType"`
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// DomainSecurity Settings
	DomainSecuritySettings DomainSecuritySettingsResponsePtrOutput `pulumi:"domainSecuritySettings"`
	// Resource etag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync pulumi.StringPtrOutput `pulumi:"filteredSync"`
	// Secure LDAP Settings
	LdapsSettings LdapsSettingsResponsePtrOutput `pulumi:"ldapsSettings"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Migration Properties
	MigrationProperties MigrationPropertiesResponseOutput `pulumi:"migrationProperties"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Notification Settings
	NotificationSettings NotificationSettingsResponsePtrOutput `pulumi:"notificationSettings"`
	// the current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// List of ReplicaSets
	ReplicaSets ReplicaSetResponseArrayOutput `pulumi:"replicaSets"`
	// Resource Forest Settings
	ResourceForestSettings ResourceForestSettingsResponsePtrOutput `pulumi:"resourceForestSettings"`
	// Sku Type
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// SyncOwner ReplicaSet Id
	SyncOwner pulumi.StringOutput `pulumi:"syncOwner"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure Active Directory Tenant Id
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Data Model Version
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewDomainService registers a new resource with the given unique name, arguments, and options.
func NewDomainService(ctx *pulumi.Context,
	name string, args *DomainServiceArgs, opts ...pulumi.ResourceOption) (*DomainService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:aad/v20210301:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad:DomainService"),
		},
		{
			Type: pulumi.String("azure-nextgen:aad:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20170101:DomainService"),
		},
		{
			Type: pulumi.String("azure-nextgen:aad/v20170101:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20170601:DomainService"),
		},
		{
			Type: pulumi.String("azure-nextgen:aad/v20170601:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20200101:DomainService"),
		},
		{
			Type: pulumi.String("azure-nextgen:aad/v20200101:DomainService"),
		},
	})
	opts = append(opts, aliases)
	var resource DomainService
	err := ctx.RegisterResource("azure-native:aad/v20210301:DomainService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainService gets an existing DomainService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainServiceState, opts ...pulumi.ResourceOption) (*DomainService, error) {
	var resource DomainService
	err := ctx.ReadResource("azure-native:aad/v20210301:DomainService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainService resources.
type domainServiceState struct {
	// Deployment Id
	DeploymentId *string `pulumi:"deploymentId"`
	// Domain Configuration Type
	DomainConfigurationType *string `pulumi:"domainConfigurationType"`
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName *string `pulumi:"domainName"`
	// DomainSecurity Settings
	DomainSecuritySettings *DomainSecuritySettingsResponse `pulumi:"domainSecuritySettings"`
	// Resource etag
	Etag *string `pulumi:"etag"`
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync *string `pulumi:"filteredSync"`
	// Secure LDAP Settings
	LdapsSettings *LdapsSettingsResponse `pulumi:"ldapsSettings"`
	// Resource location
	Location *string `pulumi:"location"`
	// Migration Properties
	MigrationProperties *MigrationPropertiesResponse `pulumi:"migrationProperties"`
	// Resource name
	Name *string `pulumi:"name"`
	// Notification Settings
	NotificationSettings *NotificationSettingsResponse `pulumi:"notificationSettings"`
	// the current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// List of ReplicaSets
	ReplicaSets []ReplicaSetResponse `pulumi:"replicaSets"`
	// Resource Forest Settings
	ResourceForestSettings *ResourceForestSettingsResponse `pulumi:"resourceForestSettings"`
	// Sku Type
	Sku *string `pulumi:"sku"`
	// SyncOwner ReplicaSet Id
	SyncOwner *string `pulumi:"syncOwner"`
	// The system meta data relating to this resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure Active Directory Tenant Id
	TenantId *string `pulumi:"tenantId"`
	// Resource type
	Type *string `pulumi:"type"`
	// Data Model Version
	Version *int `pulumi:"version"`
}

type DomainServiceState struct {
	// Deployment Id
	DeploymentId pulumi.StringPtrInput
	// Domain Configuration Type
	DomainConfigurationType pulumi.StringPtrInput
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName pulumi.StringPtrInput
	// DomainSecurity Settings
	DomainSecuritySettings DomainSecuritySettingsResponsePtrInput
	// Resource etag
	Etag pulumi.StringPtrInput
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync pulumi.StringPtrInput
	// Secure LDAP Settings
	LdapsSettings LdapsSettingsResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Migration Properties
	MigrationProperties MigrationPropertiesResponsePtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Notification Settings
	NotificationSettings NotificationSettingsResponsePtrInput
	// the current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// List of ReplicaSets
	ReplicaSets ReplicaSetResponseArrayInput
	// Resource Forest Settings
	ResourceForestSettings ResourceForestSettingsResponsePtrInput
	// Sku Type
	Sku pulumi.StringPtrInput
	// SyncOwner ReplicaSet Id
	SyncOwner pulumi.StringPtrInput
	// The system meta data relating to this resource.
	SystemData SystemDataResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Azure Active Directory Tenant Id
	TenantId pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
	// Data Model Version
	Version pulumi.IntPtrInput
}

func (DomainServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainServiceState)(nil)).Elem()
}

type domainServiceArgs struct {
	// Domain Configuration Type
	DomainConfigurationType *string `pulumi:"domainConfigurationType"`
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName *string `pulumi:"domainName"`
	// DomainSecurity Settings
	DomainSecuritySettings *DomainSecuritySettings `pulumi:"domainSecuritySettings"`
	// The name of the domain service.
	DomainServiceName *string `pulumi:"domainServiceName"`
	// Resource etag
	Etag *string `pulumi:"etag"`
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync *string `pulumi:"filteredSync"`
	// Secure LDAP Settings
	LdapsSettings *LdapsSettings `pulumi:"ldapsSettings"`
	// Resource location
	Location *string `pulumi:"location"`
	// Notification Settings
	NotificationSettings *NotificationSettings `pulumi:"notificationSettings"`
	// List of ReplicaSets
	ReplicaSets []ReplicaSet `pulumi:"replicaSets"`
	// Resource Forest Settings
	ResourceForestSettings *ResourceForestSettings `pulumi:"resourceForestSettings"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sku Type
	Sku *string `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DomainService resource.
type DomainServiceArgs struct {
	// Domain Configuration Type
	DomainConfigurationType pulumi.StringPtrInput
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName pulumi.StringPtrInput
	// DomainSecurity Settings
	DomainSecuritySettings DomainSecuritySettingsPtrInput
	// The name of the domain service.
	DomainServiceName pulumi.StringPtrInput
	// Resource etag
	Etag pulumi.StringPtrInput
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync pulumi.StringPtrInput
	// Secure LDAP Settings
	LdapsSettings LdapsSettingsPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Notification Settings
	NotificationSettings NotificationSettingsPtrInput
	// List of ReplicaSets
	ReplicaSets ReplicaSetArrayInput
	// Resource Forest Settings
	ResourceForestSettings ResourceForestSettingsPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sku Type
	Sku pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (DomainServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainServiceArgs)(nil)).Elem()
}

type DomainServiceInput interface {
	pulumi.Input

	ToDomainServiceOutput() DomainServiceOutput
	ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput
}

func (*DomainService) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainService)(nil))
}

func (i *DomainService) ToDomainServiceOutput() DomainServiceOutput {
	return i.ToDomainServiceOutputWithContext(context.Background())
}

func (i *DomainService) ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainServiceOutput)
}

type DomainServiceOutput struct {
	*pulumi.OutputState
}

func (DomainServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainService)(nil))
}

func (o DomainServiceOutput) ToDomainServiceOutput() DomainServiceOutput {
	return o
}

func (o DomainServiceOutput) ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainServiceOutput{})
}
