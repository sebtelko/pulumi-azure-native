// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// External OAuth authorization server settings.
func LookupAuthorizationServer(ctx *pulumi.Context, args *LookupAuthorizationServerArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationServerResult, error) {
	var rv LookupAuthorizationServerResult
	err := ctx.Invoke("azure-native:apimanagement/v20201201:getAuthorizationServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationServerArgs struct {
	// Identifier of the authorization server.
	Authsid string `pulumi:"authsid"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// External OAuth authorization server settings.
type LookupAuthorizationServerResult struct {
	// OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
	AuthorizationEndpoint string `pulumi:"authorizationEndpoint"`
	// HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
	AuthorizationMethods []string `pulumi:"authorizationMethods"`
	// Specifies the mechanism by which access token is passed to the API.
	BearerTokenSendingMethods []string `pulumi:"bearerTokenSendingMethods"`
	// Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
	ClientAuthenticationMethod []string `pulumi:"clientAuthenticationMethod"`
	// Client or app id registered with this authorization server.
	ClientId string `pulumi:"clientId"`
	// Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
	ClientRegistrationEndpoint string `pulumi:"clientRegistrationEndpoint"`
	// Client or app secret registered with this authorization server. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	ClientSecret *string `pulumi:"clientSecret"`
	// Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
	DefaultScope *string `pulumi:"defaultScope"`
	// Description of the authorization server. Can contain HTML formatting tags.
	Description *string `pulumi:"description"`
	// User-friendly authorization server name.
	DisplayName string `pulumi:"displayName"`
	// Form of an authorization grant, which the client uses to request the access token.
	GrantTypes []string `pulumi:"grantTypes"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
	ResourceOwnerPassword *string `pulumi:"resourceOwnerPassword"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
	ResourceOwnerUsername *string `pulumi:"resourceOwnerUsername"`
	// If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
	SupportState *bool `pulumi:"supportState"`
	// Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
	TokenBodyParameters []TokenBodyParameterContractResponse `pulumi:"tokenBodyParameters"`
	// OAuth token endpoint. Contains absolute URI to entity being referenced.
	TokenEndpoint *string `pulumi:"tokenEndpoint"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}
