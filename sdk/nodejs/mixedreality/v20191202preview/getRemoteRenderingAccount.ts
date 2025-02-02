// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * RemoteRenderingAccount Response.
 */
export function getRemoteRenderingAccount(args: GetRemoteRenderingAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetRemoteRenderingAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:mixedreality/v20191202preview:getRemoteRenderingAccount", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRemoteRenderingAccountArgs {
    /**
     * Name of an Mixed Reality Account.
     */
    readonly accountName: string;
    /**
     * Name of an Azure resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * RemoteRenderingAccount Response.
 */
export interface GetRemoteRenderingAccountResult {
    /**
     * Correspond domain name of certain Spatial Anchors Account
     */
    readonly accountDomain: string;
    /**
     * unique id of certain account.
     */
    readonly accountId: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The identity associated with this account
     */
    readonly identity?: outputs.mixedreality.v20191202preview.IdentityResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
