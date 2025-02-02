// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Base class for backup ProtectionIntent.
 */
export function getProtectionIntent(args: GetProtectionIntentArgs, opts?: pulumi.InvokeOptions): Promise<GetProtectionIntentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:recoveryservices/v20210210:getProtectionIntent", {
        "fabricName": args.fabricName,
        "intentObjectName": args.intentObjectName,
        "resourceGroupName": args.resourceGroupName,
        "vaultName": args.vaultName,
    }, opts);
}

export interface GetProtectionIntentArgs {
    /**
     * Fabric name associated with the backed up item.
     */
    readonly fabricName: string;
    /**
     * Backed up item name whose details are to be fetched.
     */
    readonly intentObjectName: string;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the recovery services vault.
     */
    readonly vaultName: string;
}

/**
 * Base class for backup ProtectionIntent.
 */
export interface GetProtectionIntentResult {
    /**
     * Optional ETag.
     */
    readonly eTag?: string;
    /**
     * Resource Id represents the complete path to the resource.
     */
    readonly id: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name associated with the resource.
     */
    readonly name: string;
    /**
     * ProtectionIntentResource properties
     */
    readonly properties: outputs.recoveryservices.v20210210.AzureRecoveryServiceVaultProtectionIntentResponse | outputs.recoveryservices.v20210210.AzureResourceProtectionIntentResponse | outputs.recoveryservices.v20210210.AzureWorkloadAutoProtectionIntentResponse | outputs.recoveryservices.v20210210.AzureWorkloadSQLAutoProtectionIntentResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
     */
    readonly type: string;
}
