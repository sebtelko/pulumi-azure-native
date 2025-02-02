// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gateway authentication keys.
 */
export function listGatewayKeys(args: ListGatewayKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListGatewayKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:apimanagement/v20210101preview:listGatewayKeys", {
        "gatewayId": args.gatewayId,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface ListGatewayKeysArgs {
    /**
     * Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
     */
    readonly gatewayId: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: string;
}

/**
 * Gateway authentication keys.
 */
export interface ListGatewayKeysResult {
    /**
     * Primary gateway key.
     */
    readonly primary?: string;
    /**
     * Secondary gateway key.
     */
    readonly secondary?: string;
}
