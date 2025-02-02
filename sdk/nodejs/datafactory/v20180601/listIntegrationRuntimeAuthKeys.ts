// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The integration runtime authentication keys.
 */
export function listIntegrationRuntimeAuthKeys(args: ListIntegrationRuntimeAuthKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListIntegrationRuntimeAuthKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:datafactory/v20180601:listIntegrationRuntimeAuthKeys", {
        "factoryName": args.factoryName,
        "integrationRuntimeName": args.integrationRuntimeName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListIntegrationRuntimeAuthKeysArgs {
    /**
     * The factory name.
     */
    readonly factoryName: string;
    /**
     * The integration runtime name.
     */
    readonly integrationRuntimeName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * The integration runtime authentication keys.
 */
export interface ListIntegrationRuntimeAuthKeysResult {
    /**
     * The primary integration runtime authentication key.
     */
    readonly authKey1?: string;
    /**
     * The secondary integration runtime authentication key.
     */
    readonly authKey2?: string;
}
