// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * String dictionary resource.
 */
export function listWebAppConnectionStrings(args: ListWebAppConnectionStringsArgs, opts?: pulumi.InvokeOptions): Promise<ListWebAppConnectionStringsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:web/v20190801:listWebAppConnectionStrings", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListWebAppConnectionStringsArgs {
    /**
     * Name of the app.
     */
    readonly name: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * String dictionary resource.
 */
export interface ListWebAppConnectionStringsResult {
    /**
     * Resource Id.
     */
    readonly id: string;
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * Connection strings.
     */
    readonly properties: {[key: string]: outputs.web.v20190801.ConnStringValueTypePairResponse};
    /**
     * Resource type.
     */
    readonly type: string;
}
