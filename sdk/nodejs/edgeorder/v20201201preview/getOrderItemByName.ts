// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Represents order item contract
 */
export function getOrderItemByName(args: GetOrderItemByNameArgs, opts?: pulumi.InvokeOptions): Promise<GetOrderItemByNameResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:edgeorder/v20201201preview:getOrderItemByName", {
        "orderItemName": args.orderItemName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetOrderItemByNameArgs {
    /**
     * The name of the order item
     */
    readonly orderItemName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Represents order item contract
 */
export interface GetOrderItemByNameResult {
    /**
     * Represents shipping and return address for order
     */
    readonly addressDetails: outputs.edgeorder.v20201201preview.AddressDetailsResponse;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Id of the order to which order items belongs to
     */
    readonly orderId: string;
    /**
     * Represents order item details.
     */
    readonly orderItemDetails: outputs.edgeorder.v20201201preview.OrderItemDetailsResponse;
    /**
     * Start time of order item
     */
    readonly startTime: string;
    /**
     * Represents resource creation and update time
     */
    readonly systemData: outputs.edgeorder.v20201201preview.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
