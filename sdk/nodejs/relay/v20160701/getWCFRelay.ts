// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Description of WcfRelays Resource.
 */
export function getWCFRelay(args: GetWCFRelayArgs, opts?: pulumi.InvokeOptions): Promise<GetWCFRelayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:relay/v20160701:getWCFRelay", {
        "namespaceName": args.namespaceName,
        "relayName": args.relayName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetWCFRelayArgs {
    /**
     * The Namespace Name
     */
    readonly namespaceName: string;
    /**
     * The relay name
     */
    readonly relayName: string;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * Description of WcfRelays Resource.
 */
export interface GetWCFRelayResult {
    /**
     * The time the WCFRelay was created.
     */
    readonly createdAt: string;
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * true if the relay is dynamic; otherwise, false.
     */
    readonly isDynamic: boolean;
    /**
     * The number of listeners for this relay. min : 1 and max:25 supported
     */
    readonly listenerCount: number;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * WCFRelay Type.
     */
    readonly relayType?: string;
    /**
     * true if client authorization is needed for this relay; otherwise, false.
     */
    readonly requiresClientAuthorization?: boolean;
    /**
     * true if transport security is needed for this relay; otherwise, false.
     */
    readonly requiresTransportSecurity?: boolean;
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * The time the namespace was updated.
     */
    readonly updatedAt: string;
    /**
     * usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
     */
    readonly userMetadata?: string;
}
