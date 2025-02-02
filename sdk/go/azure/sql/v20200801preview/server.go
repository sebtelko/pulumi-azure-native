// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure SQL Database server.
type Server struct {
	pulumi.CustomResourceState

	// Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin pulumi.StringPtrOutput `pulumi:"administratorLogin"`
	// The fully qualified domain name of the server.
	FullyQualifiedDomainName pulumi.StringOutput `pulumi:"fullyQualifiedDomainName"`
	// The Azure Active Directory identity of the server.
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// Kind of sql server. This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion pulumi.StringPtrOutput `pulumi:"minimalTlsVersion"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections on a server
	PrivateEndpointConnections ServerPrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The state of the server.
	State pulumi.StringOutput `pulumi:"state"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The version of the server.
	Version pulumi.StringPtrOutput `pulumi:"version"`
	// Whether or not existing server has a workspace created and if it allows connection from workspace
	WorkspaceFeature pulumi.StringOutput `pulumi:"workspaceFeature"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql:Server"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:Server"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20140401:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:Server"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20150501preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:Server"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20190601preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:Server"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:Server"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:Server"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:Server"),
		},
	})
	opts = append(opts, aliases)
	var resource Server
	err := ctx.RegisterResource("azure-native:sql/v20200801preview:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure-native:sql/v20200801preview:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
	// Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The fully qualified domain name of the server.
	FullyQualifiedDomainName *string `pulumi:"fullyQualifiedDomainName"`
	// The Azure Active Directory identity of the server.
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	// Kind of sql server. This is metadata used for the Azure portal experience.
	Kind *string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `pulumi:"minimalTlsVersion"`
	// Resource name.
	Name *string `pulumi:"name"`
	// List of private endpoint connections on a server
	PrivateEndpointConnections []ServerPrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The state of the server.
	State *string `pulumi:"state"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The version of the server.
	Version *string `pulumi:"version"`
	// Whether or not existing server has a workspace created and if it allows connection from workspace
	WorkspaceFeature *string `pulumi:"workspaceFeature"`
}

type ServerState struct {
	// Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin pulumi.StringPtrInput
	// The fully qualified domain name of the server.
	FullyQualifiedDomainName pulumi.StringPtrInput
	// The Azure Active Directory identity of the server.
	Identity ResourceIdentityResponsePtrInput
	// Kind of sql server. This is metadata used for the Azure portal experience.
	Kind pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// List of private endpoint connections on a server
	PrivateEndpointConnections ServerPrivateEndpointConnectionResponseArrayInput
	// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess pulumi.StringPtrInput
	// The state of the server.
	State pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The version of the server.
	Version pulumi.StringPtrInput
	// Whether or not existing server has a workspace created and if it allows connection from workspace
	WorkspaceFeature pulumi.StringPtrInput
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The administrator login password (required for server creation).
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// The Azure Active Directory identity of the server.
	Identity *ResourceIdentity `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `pulumi:"minimalTlsVersion"`
	// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName *string `pulumi:"serverName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The version of the server.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin pulumi.StringPtrInput
	// The administrator login password (required for server creation).
	AdministratorLoginPassword pulumi.StringPtrInput
	// The Azure Active Directory identity of the server.
	Identity ResourceIdentityPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion pulumi.StringPtrInput
	// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The version of the server.
	Version pulumi.StringPtrInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (*Server) ElementType() reflect.Type {
	return reflect.TypeOf((*Server)(nil))
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

type ServerOutput struct {
	*pulumi.OutputState
}

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Server)(nil))
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
