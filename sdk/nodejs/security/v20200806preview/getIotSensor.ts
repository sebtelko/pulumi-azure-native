// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * IoT sensor model
 */
export function getIotSensor(args: GetIotSensorArgs, opts?: pulumi.InvokeOptions): Promise<GetIotSensorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:security/v20200806preview:getIotSensor", {
        "iotSensorName": args.iotSensorName,
        "scope": args.scope,
    }, opts);
}

export interface GetIotSensorArgs {
    /**
     * Name of the IoT sensor
     */
    readonly iotSensorName: string;
    /**
     * Scope of the query (IoT Hub, /providers/Microsoft.Devices/iotHubs/myHub)
     */
    readonly scope: string;
}

/**
 * IoT sensor model
 */
export interface GetIotSensorResult {
    /**
     * Last connectivity time of the IoT sensor
     */
    readonly connectivityTime: string;
    /**
     * Creation time of the IoT sensor
     */
    readonly creationTime: string;
    /**
     * Dynamic mode status of the IoT sensor
     */
    readonly dynamicLearning: boolean;
    /**
     * Resource Id
     */
    readonly id: string;
    /**
     * Learning mode status of the IoT sensor
     */
    readonly learningMode: boolean;
    /**
     * Resource name
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
     * Resource type
     */
    readonly type: string;
    /**
     * Zone of the IoT sensor
     */
    readonly zone?: string;
}
