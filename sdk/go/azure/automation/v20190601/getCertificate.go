// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the certificate.
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:automation/v20190601:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The name of certificate.
	CertificateName string `pulumi:"certificateName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the certificate.
type LookupCertificateResult struct {
	// Gets the creation time.
	CreationTime string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets the expiry time of the certificate.
	ExpiryTime string `pulumi:"expiryTime"`
	// Fully qualified resource Id for the resource
	Id string `pulumi:"id"`
	// Gets the is exportable flag of the certificate.
	IsExportable bool `pulumi:"isExportable"`
	// Gets the last modified time.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the thumbprint of the certificate.
	Thumbprint string `pulumi:"thumbprint"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
