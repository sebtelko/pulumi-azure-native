// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The key.
 */
export function listManagerActivationKey(args: ListManagerActivationKeyArgs, opts?: pulumi.InvokeOptions): Promise<ListManagerActivationKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:storsimple/v20170601:listManagerActivationKey", {
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListManagerActivationKeyArgs {
    /**
     * The manager name
     */
    readonly managerName: string;
    /**
     * The resource group name
     */
    readonly resourceGroupName: string;
}

/**
 * The key.
 */
export interface ListManagerActivationKeyResult {
    /**
     * The activation key for the device.
     */
    readonly activationKey: string;
}
