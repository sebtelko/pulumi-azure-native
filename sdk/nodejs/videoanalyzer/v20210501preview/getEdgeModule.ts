// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The representation of an edge module.
 */
export function getEdgeModule(args: GetEdgeModuleArgs, opts?: pulumi.InvokeOptions): Promise<GetEdgeModuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:videoanalyzer/v20210501preview:getEdgeModule", {
        "accountName": args.accountName,
        "edgeModuleName": args.edgeModuleName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetEdgeModuleArgs {
    /**
     * The Azure Video Analyzer account name.
     */
    readonly accountName: string;
    /**
     * The name of the edge module to retrieve.
     */
    readonly edgeModuleName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * The representation of an edge module.
 */
export interface GetEdgeModuleResult {
    /**
     * Internal ID generated for the instance of the Video Analyzer edge module.
     */
    readonly edgeModuleId: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The system metadata relating to this resource.
     */
    readonly systemData: outputs.videoanalyzer.v20210501preview.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
