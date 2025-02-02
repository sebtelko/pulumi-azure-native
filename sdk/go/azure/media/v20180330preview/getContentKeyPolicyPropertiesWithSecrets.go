// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180330preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The properties of the Content Key Policy.
func GetContentKeyPolicyPropertiesWithSecrets(ctx *pulumi.Context, args *GetContentKeyPolicyPropertiesWithSecretsArgs, opts ...pulumi.InvokeOption) (*GetContentKeyPolicyPropertiesWithSecretsResult, error) {
	var rv GetContentKeyPolicyPropertiesWithSecretsResult
	err := ctx.Invoke("azure-native:media/v20180330preview:getContentKeyPolicyPropertiesWithSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetContentKeyPolicyPropertiesWithSecretsArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Content Key Policy name.
	ContentKeyPolicyName string `pulumi:"contentKeyPolicyName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The properties of the Content Key Policy.
type GetContentKeyPolicyPropertiesWithSecretsResult struct {
	// The creation date of the Policy
	Created string `pulumi:"created"`
	// A description for the Policy.
	Description *string `pulumi:"description"`
	// The last modified date of the Policy
	LastModified string `pulumi:"lastModified"`
	// The Key Policy options.
	Options []ContentKeyPolicyOptionResponse `pulumi:"options"`
	// The legacy Policy ID.
	PolicyId string `pulumi:"policyId"`
}
