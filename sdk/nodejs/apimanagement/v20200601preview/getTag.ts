// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Tag Contract details.
 */
export function getTag(args: GetTagArgs, opts?: pulumi.InvokeOptions): Promise<GetTagResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:apimanagement/v20200601preview:getTag", {
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
        "tagId": args.tagId,
    }, opts);
}

export interface GetTagArgs {
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: string;
    /**
     * Tag identifier. Must be unique in the current API Management service instance.
     */
    readonly tagId: string;
}

/**
 * Tag Contract details.
 */
export interface GetTagResult {
    /**
     * Tag name.
     */
    readonly displayName: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Resource type for API Management resource.
     */
    readonly type: string;
}
