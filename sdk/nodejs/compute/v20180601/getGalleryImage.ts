// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Specifies information about the gallery Image Definition that you want to create or update.
 */
export function getGalleryImage(args: GetGalleryImageArgs, opts?: pulumi.InvokeOptions): Promise<GetGalleryImageResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:compute/v20180601:getGalleryImage", {
        "galleryImageName": args.galleryImageName,
        "galleryName": args.galleryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetGalleryImageArgs {
    /**
     * The name of the gallery Image Definition to be retrieved.
     */
    readonly galleryImageName: string;
    /**
     * The name of the Shared Image Gallery from which the Image Definitions are to be retrieved.
     */
    readonly galleryName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Specifies information about the gallery Image Definition that you want to create or update.
 */
export interface GetGalleryImageResult {
    /**
     * The description of this gallery Image Definition resource. This property is updatable.
     */
    readonly description?: string;
    /**
     * Describes the disallowed disk types.
     */
    readonly disallowed?: outputs.compute.v20180601.DisallowedResponse;
    /**
     * The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
     */
    readonly endOfLifeDate?: string;
    /**
     * The Eula agreement for the gallery Image Definition.
     */
    readonly eula?: string;
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * This is the gallery Image Definition identifier.
     */
    readonly identifier: outputs.compute.v20180601.GalleryImageIdentifierResponse;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The allowed values for OS State are 'Generalized'.
     */
    readonly osState: string;
    /**
     * This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
     */
    readonly osType: string;
    /**
     * The privacy statement uri.
     */
    readonly privacyStatementUri?: string;
    /**
     * The provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * Describes the gallery Image Definition purchase plan. This is used by marketplace images.
     */
    readonly purchasePlan?: outputs.compute.v20180601.ImagePurchasePlanResponse;
    /**
     * The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
     */
    readonly recommended?: outputs.compute.v20180601.RecommendedMachineConfigurationResponse;
    /**
     * The release note uri.
     */
    readonly releaseNoteUri?: string;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
