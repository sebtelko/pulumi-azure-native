// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * ExpressRoute gateway resource.
 */
export function getExpressRouteGateway(args: GetExpressRouteGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetExpressRouteGatewayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:network/v20181101:getExpressRouteGateway", {
        "expressRouteGatewayName": args.expressRouteGatewayName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetExpressRouteGatewayArgs {
    /**
     * The name of the ExpressRoute gateway.
     */
    readonly expressRouteGatewayName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * ExpressRoute gateway resource.
 */
export interface GetExpressRouteGatewayResult {
    /**
     * Configuration for auto scaling.
     */
    readonly autoScaleConfiguration?: outputs.network.v20181101.ExpressRouteGatewayPropertiesResponseAutoScaleConfiguration;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * List of ExpressRoute connections to the ExpressRoute gateway.
     */
    readonly expressRouteConnections: outputs.network.v20181101.ExpressRouteConnectionResponse[];
    /**
     * Resource ID.
     */
    readonly id?: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The Virtual Hub where the ExpressRoute gateway is or will be deployed.
     */
    readonly virtualHub: outputs.network.v20181101.VirtualHubIdResponse;
}
