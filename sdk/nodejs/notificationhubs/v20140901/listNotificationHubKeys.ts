// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Namespace/NotificationHub Connection String
 */
export function listNotificationHubKeys(args: ListNotificationHubKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListNotificationHubKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:notificationhubs/v20140901:listNotificationHubKeys", {
        "authorizationRuleName": args.authorizationRuleName,
        "namespaceName": args.namespaceName,
        "notificationHubName": args.notificationHubName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListNotificationHubKeysArgs {
    /**
     * The connection string of the NotificationHub for the specified authorizationRule.
     */
    readonly authorizationRuleName: string;
    /**
     * The namespace name.
     */
    readonly namespaceName: string;
    /**
     * The notification hub name.
     */
    readonly notificationHubName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Namespace/NotificationHub Connection String
 */
export interface ListNotificationHubKeysResult {
    /**
     * Gets or sets the primaryConnectionString of the created Namespace AuthorizationRule.
     */
    readonly primaryConnectionString?: string;
    /**
     * Gets or sets the secondaryConnectionString of the created Namespace AuthorizationRule
     */
    readonly secondaryConnectionString?: string;
}
