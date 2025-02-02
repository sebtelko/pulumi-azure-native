// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The extended Info of the Data Box Edge/Gateway device.
func GetDeviceExtendedInformation(ctx *pulumi.Context, args *GetDeviceExtendedInformationArgs, opts ...pulumi.InvokeOption) (*GetDeviceExtendedInformationResult, error) {
	var rv GetDeviceExtendedInformationResult
	err := ctx.Invoke("azure-native:databoxedge/v20190801:getDeviceExtendedInformation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDeviceExtendedInformationArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The extended Info of the Data Box Edge/Gateway device.
type GetDeviceExtendedInformationResult struct {
	// The public part of the encryption certificate. Client uses this to encrypt any secret.
	EncryptionKey *string `pulumi:"encryptionKey"`
	// The digital signature of encrypted certificate.
	EncryptionKeyThumbprint *string `pulumi:"encryptionKeyThumbprint"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// The object name.
	Name string `pulumi:"name"`
	// The Resource ID of the Resource.
	ResourceKey string `pulumi:"resourceKey"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
}
