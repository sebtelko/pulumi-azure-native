// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Private Endpoint Connection ARM resource.
 * API Version: 2018-06-01.
 */
export function getPrivateEndpointConnection(args: GetPrivateEndpointConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateEndpointConnectionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:datafactory:getPrivateEndpointConnection", {
        "factoryName": args.factoryName,
        "privateEndpointConnectionName": args.privateEndpointConnectionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPrivateEndpointConnectionArgs {
    /**
     * The factory name.
     */
    readonly factoryName: string;
    /**
     * The private endpoint connection name.
     */
    readonly privateEndpointConnectionName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * Private Endpoint Connection ARM resource.
 */
export interface GetPrivateEndpointConnectionResult {
    /**
     * Etag identifies change in the resource.
     */
    readonly etag: string;
    /**
     * The resource identifier.
     */
    readonly id: string;
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * Core resource properties
     */
    readonly properties: outputs.datafactory.RemotePrivateEndpointConnectionResponse;
    /**
     * The resource type.
     */
    readonly type: string;
}
