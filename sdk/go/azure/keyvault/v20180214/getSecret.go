// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180214

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource information with extended details.
func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	var rv LookupSecretResult
	err := ctx.Invoke("azure-native:keyvault/v20180214:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecretArgs struct {
	// The name of the Resource Group to which the vault belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the secret.
	SecretName string `pulumi:"secretName"`
	// The name of the vault.
	VaultName string `pulumi:"vaultName"`
}

// Resource information with extended details.
type LookupSecretResult struct {
	// Fully qualified identifier of the key vault resource.
	Id string `pulumi:"id"`
	// Azure location of the key vault resource.
	Location string `pulumi:"location"`
	// Name of the key vault resource.
	Name string `pulumi:"name"`
	// Properties of the secret
	Properties SecretPropertiesResponse `pulumi:"properties"`
	// Tags assigned to the key vault resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type of the key vault resource.
	Type string `pulumi:"type"`
}
