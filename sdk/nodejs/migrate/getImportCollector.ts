// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * API Version: 2019-10-01.
 */
export function getImportCollector(args: GetImportCollectorArgs, opts?: pulumi.InvokeOptions): Promise<GetImportCollectorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:migrate:getImportCollector", {
        "importCollectorName": args.importCollectorName,
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetImportCollectorArgs {
    /**
     * Unique name of a Import collector within a project.
     */
    readonly importCollectorName: string;
    /**
     * Name of the Azure Migrate project.
     */
    readonly projectName: string;
    /**
     * Name of the Azure Resource Group that project is part of.
     */
    readonly resourceGroupName: string;
}

export interface GetImportCollectorResult {
    readonly eTag?: string;
    readonly id: string;
    readonly name: string;
    readonly properties: outputs.migrate.ImportCollectorPropertiesResponse;
    readonly type: string;
}
