// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * IoT sensor model
 */
export function getSensor(args: GetSensorArgs, opts?: pulumi.InvokeOptions): Promise<GetSensorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:iotsecurity/v20210201preview:getSensor", {
        "scope": args.scope,
        "sensorName": args.sensorName,
    }, opts);
}

export interface GetSensorArgs {
    /**
     * Scope of the query (IoT Hub, /providers/Microsoft.Devices/iotHubs/myHub)
     */
    readonly scope: string;
    /**
     * Name of the IoT sensor
     */
    readonly sensorName: string;
}

/**
 * IoT sensor model
 */
export interface GetSensorResult {
    /**
     * Last connectivity time of the IoT sensor
     */
    readonly connectivityTime: string;
    /**
     * Dynamic mode status of the IoT sensor
     */
    readonly dynamicLearning: boolean;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * Learning mode status of the IoT sensor
     */
    readonly learningMode: boolean;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Status of the IoT sensor
     */
    readonly sensorStatus: string;
    /**
     * Type of sensor
     */
    readonly sensorType?: string;
    /**
     * Version of the IoT sensor
     */
    readonly sensorVersion: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.iotsecurity.v20210201preview.SystemDataResponse;
    /**
     * TI Automatic mode status of the IoT sensor
     */
    readonly tiAutomaticUpdates?: boolean;
    /**
     * TI Status of the IoT sensor
     */
    readonly tiStatus: string;
    /**
     * TI Version of the IoT sensor
     */
    readonly tiVersion: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Zone of the IoT sensor
     */
    readonly zone?: string;
}
