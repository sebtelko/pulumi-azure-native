// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * maintenance configuration.
 */
export function getMaintenanceConfiguration(args: GetMaintenanceConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetMaintenanceConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:containerservice/v20210301:getMaintenanceConfiguration", {
        "configName": args.configName,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetMaintenanceConfigurationArgs {
    /**
     * The name of the maintenance configuration.
     */
    readonly configName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the managed cluster resource.
     */
    readonly resourceName: string;
}

/**
 * maintenance configuration.
 */
export interface GetMaintenanceConfigurationResult {
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name: string;
    /**
     * Time slots on which upgrade is not allowed.
     */
    readonly notAllowedTime?: outputs.containerservice.v20210301.TimeSpanResponse[];
    /**
     * The system meta data relating to this resource.
     */
    readonly systemData: outputs.containerservice.v20210301.SystemDataResponse;
    /**
     * Weekday time slots allowed to upgrade.
     */
    readonly timeInWeek?: outputs.containerservice.v20210301.TimeInWeekResponse[];
    /**
     * Resource type
     */
    readonly type: string;
}
