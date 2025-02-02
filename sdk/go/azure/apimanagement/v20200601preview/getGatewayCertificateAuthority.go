// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gateway certificate authority details.
func LookupGatewayCertificateAuthority(ctx *pulumi.Context, args *LookupGatewayCertificateAuthorityArgs, opts ...pulumi.InvokeOption) (*LookupGatewayCertificateAuthorityResult, error) {
	var rv LookupGatewayCertificateAuthorityResult
	err := ctx.Invoke("azure-native:apimanagement/v20200601preview:getGatewayCertificateAuthority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayCertificateAuthorityArgs struct {
	// Identifier of the certificate entity. Must be unique in the current API Management service instance.
	CertificateId string `pulumi:"certificateId"`
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId string `pulumi:"gatewayId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Gateway certificate authority details.
type LookupGatewayCertificateAuthorityResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Determines whether certificate authority is trusted.
	IsTrusted *bool `pulumi:"isTrusted"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}
