// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The integration account certificate.
// API Version: 2019-05-01.
func LookupIntegrationAccountCertificate(ctx *pulumi.Context, args *LookupIntegrationAccountCertificateArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountCertificateResult, error) {
	var rv LookupIntegrationAccountCertificateResult
	err := ctx.Invoke("azure-native:logic:getIntegrationAccountCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountCertificateArgs struct {
	// The integration account certificate name.
	CertificateName string `pulumi:"certificateName"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The integration account certificate.
type LookupIntegrationAccountCertificateResult struct {
	// The changed time.
	ChangedTime string `pulumi:"changedTime"`
	// The created time.
	CreatedTime string `pulumi:"createdTime"`
	// The resource id.
	Id string `pulumi:"id"`
	// The key details in the key vault.
	Key *KeyVaultKeyReferenceResponse `pulumi:"key"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// Gets the resource name.
	Name string `pulumi:"name"`
	// The public certificate.
	PublicCertificate *string `pulumi:"publicCertificate"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type string `pulumi:"type"`
}
