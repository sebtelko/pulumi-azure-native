// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The customer's prefix that is registered by the peering service provider.
 */
export function getRegisteredPrefix(args: GetRegisteredPrefixArgs, opts?: pulumi.InvokeOptions): Promise<GetRegisteredPrefixResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:peering/v20210101:getRegisteredPrefix", {
        "peeringName": args.peeringName,
        "registeredPrefixName": args.registeredPrefixName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRegisteredPrefixArgs {
    /**
     * The name of the peering.
     */
    readonly peeringName: string;
    /**
     * The name of the registered prefix.
     */
    readonly registeredPrefixName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * The customer's prefix that is registered by the peering service provider.
 */
export interface GetRegisteredPrefixResult {
    /**
     * The error message associated with the validation state, if any.
     */
    readonly errorMessage: string;
    /**
     * The ID of the resource.
     */
    readonly id: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The peering service prefix key that is to be shared with the customer.
     */
    readonly peeringServicePrefixKey: string;
    /**
     * The customer's prefix from which traffic originates.
     */
    readonly prefix?: string;
    /**
     * The prefix validation state.
     */
    readonly prefixValidationState: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
