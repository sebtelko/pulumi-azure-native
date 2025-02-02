// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The properties of the EventHubConsumerGroupInfo object.
 */
export function getIotHubResourceEventHubConsumerGroup(args: GetIotHubResourceEventHubConsumerGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetIotHubResourceEventHubConsumerGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:devices/v20190701preview:getIotHubResourceEventHubConsumerGroup", {
        "eventHubEndpointName": args.eventHubEndpointName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetIotHubResourceEventHubConsumerGroupArgs {
    /**
     * The name of the Event Hub-compatible endpoint in the IoT hub.
     */
    readonly eventHubEndpointName: string;
    /**
     * The name of the consumer group to retrieve.
     */
    readonly name: string;
    /**
     * The name of the resource group that contains the IoT hub.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the IoT hub.
     */
    readonly resourceName: string;
}

/**
 * The properties of the EventHubConsumerGroupInfo object.
 */
export interface GetIotHubResourceEventHubConsumerGroupResult {
    /**
     * The etag.
     */
    readonly etag: string;
    /**
     * The Event Hub-compatible consumer group identifier.
     */
    readonly id: string;
    /**
     * The Event Hub-compatible consumer group name.
     */
    readonly name: string;
    /**
     * The tags.
     */
    readonly properties: {[key: string]: string};
    /**
     * the resource type.
     */
    readonly type: string;
}
