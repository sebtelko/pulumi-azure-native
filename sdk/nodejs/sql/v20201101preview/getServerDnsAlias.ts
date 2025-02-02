// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * A server DNS alias.
 */
export function getServerDnsAlias(args: GetServerDnsAliasArgs, opts?: pulumi.InvokeOptions): Promise<GetServerDnsAliasResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:sql/v20201101preview:getServerDnsAlias", {
        "dnsAliasName": args.dnsAliasName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetServerDnsAliasArgs {
    /**
     * The name of the server dns alias.
     */
    readonly dnsAliasName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server that the alias is pointing to.
     */
    readonly serverName: string;
}

/**
 * A server DNS alias.
 */
export interface GetServerDnsAliasResult {
    /**
     * The fully qualified DNS record for alias
     */
    readonly azureDnsRecord: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
