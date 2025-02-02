// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Shared access keys of the Domain.
 */
export function listDomainSharedAccessKeys(args: ListDomainSharedAccessKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListDomainSharedAccessKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:eventgrid/v20190601:listDomainSharedAccessKeys", {
        "domainName": args.domainName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListDomainSharedAccessKeysArgs {
    /**
     * Name of the domain.
     */
    readonly domainName: string;
    /**
     * The name of the resource group within the user's subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * Shared access keys of the Domain.
 */
export interface ListDomainSharedAccessKeysResult {
    /**
     * Shared access key1 for the domain.
     */
    readonly key1?: string;
    /**
     * Shared access key2 for the domain.
     */
    readonly key2?: string;
}
