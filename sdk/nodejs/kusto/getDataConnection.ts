// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Class representing an data connection.
 * API Version: 2021-01-01.
 */
export function getDataConnection(args: GetDataConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetDataConnectionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:kusto:getDataConnection", {
        "clusterName": args.clusterName,
        "dataConnectionName": args.dataConnectionName,
        "databaseName": args.databaseName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDataConnectionArgs {
    /**
     * The name of the Kusto cluster.
     */
    readonly clusterName: string;
    /**
     * The name of the data connection.
     */
    readonly dataConnectionName: string;
    /**
     * The name of the database in the Kusto cluster.
     */
    readonly databaseName: string;
    /**
     * The name of the resource group containing the Kusto cluster.
     */
    readonly resourceGroupName: string;
}

/**
 * Class representing an data connection.
 */
export interface GetDataConnectionResult {
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * Kind of the endpoint for the data connection
     */
    readonly kind: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
