// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The container for solution.
 */
export function getManagementConfiguration(args: GetManagementConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetManagementConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:operationsmanagement/v20151101preview:getManagementConfiguration", {
        "managementConfigurationName": args.managementConfigurationName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagementConfigurationArgs {
    /**
     * User Management Configuration Name.
     */
    readonly managementConfigurationName: string;
    /**
     * The name of the resource group to get. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * The container for solution.
 */
export interface GetManagementConfigurationResult {
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Properties for ManagementConfiguration object supported by the OperationsManagement resource provider.
     */
    readonly properties: outputs.operationsmanagement.v20151101preview.ManagementConfigurationPropertiesResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}
