// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Migration item.
 */
export function getReplicationMigrationItem(args: GetReplicationMigrationItemArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationMigrationItemResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:recoveryservices/v20210210:getReplicationMigrationItem", {
        "fabricName": args.fabricName,
        "migrationItemName": args.migrationItemName,
        "protectionContainerName": args.protectionContainerName,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetReplicationMigrationItemArgs {
    /**
     * Fabric unique name.
     */
    readonly fabricName: string;
    /**
     * Migration item name.
     */
    readonly migrationItemName: string;
    /**
     * Protection container name.
     */
    readonly protectionContainerName: string;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the recovery services vault.
     */
    readonly resourceName: string;
}

/**
 * Migration item.
 */
export interface GetReplicationMigrationItemResult {
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Resource Location
     */
    readonly location?: string;
    /**
     * Resource Name
     */
    readonly name: string;
    /**
     * The migration item properties.
     */
    readonly properties: outputs.recoveryservices.v20210210.MigrationItemPropertiesResponse;
    /**
     * Resource Type
     */
    readonly type: string;
}
