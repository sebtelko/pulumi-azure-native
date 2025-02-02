// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200202preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A server key.
func LookupServerKey(ctx *pulumi.Context, args *LookupServerKeyArgs, opts ...pulumi.InvokeOption) (*LookupServerKeyResult, error) {
	var rv LookupServerKeyResult
	err := ctx.Invoke("azure-native:sql/v20200202preview:getServerKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerKeyArgs struct {
	// The name of the server key to be retrieved.
	KeyName string `pulumi:"keyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A server key.
type LookupServerKeyResult struct {
	// The server key creation date.
	CreationDate string `pulumi:"creationDate"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Kind of encryption protector. This is metadata used for the Azure portal experience.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Subregion of the server key.
	Subregion string `pulumi:"subregion"`
	// Thumbprint of the server key.
	Thumbprint string `pulumi:"thumbprint"`
	// Resource type.
	Type string `pulumi:"type"`
}
