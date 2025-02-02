// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170426

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The image definition.
func GetImageUploadUrlForData(ctx *pulumi.Context, args *GetImageUploadUrlForDataArgs, opts ...pulumi.InvokeOption) (*GetImageUploadUrlForDataResult, error) {
	var rv GetImageUploadUrlForDataResult
	err := ctx.Invoke("azure-native:customerinsights/v20170426:getImageUploadUrlForData", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetImageUploadUrlForDataArgs struct {
	// Type of entity. Can be Profile or Interaction.
	EntityType *string `pulumi:"entityType"`
	// Name of the entity type.
	EntityTypeName *string `pulumi:"entityTypeName"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// Relative path of the image.
	RelativePath *string `pulumi:"relativePath"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The image definition.
type GetImageUploadUrlForDataResult struct {
	// Content URL for the image blob.
	ContentUrl *string `pulumi:"contentUrl"`
	// Whether image exists already.
	ImageExists *bool `pulumi:"imageExists"`
	// Relative path of the image.
	RelativePath *string `pulumi:"relativePath"`
}
