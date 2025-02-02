// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * An Application Insights component linked storage accounts
 */
export function getComponentLinkedStorageAccount(args: GetComponentLinkedStorageAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetComponentLinkedStorageAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:insights/v20200301preview:getComponentLinkedStorageAccount", {
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
        "storageType": args.storageType,
    }, opts);
}

export interface GetComponentLinkedStorageAccountArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the Application Insights component resource.
     */
    readonly resourceName: string;
    /**
     * The type of the Application Insights component data source for the linked storage account.
     */
    readonly storageType: string;
}

/**
 * An Application Insights component linked storage accounts
 */
export interface GetComponentLinkedStorageAccountResult {
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * Linked storage account resource ID
     */
    readonly linkedStorageAccount?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
