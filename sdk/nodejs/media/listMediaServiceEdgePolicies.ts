// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * API Version: 2020-05-01.
 */
export function listMediaServiceEdgePolicies(args: ListMediaServiceEdgePoliciesArgs, opts?: pulumi.InvokeOptions): Promise<ListMediaServiceEdgePoliciesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:media:listMediaServiceEdgePolicies", {
        "accountName": args.accountName,
        "deviceId": args.deviceId,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListMediaServiceEdgePoliciesArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: string;
    /**
     * Unique identifier of the edge device.
     */
    readonly deviceId?: string;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

export interface ListMediaServiceEdgePoliciesResult {
    readonly usageDataCollectionPolicy?: outputs.media.EdgeUsageDataCollectionPolicyResponse;
}
