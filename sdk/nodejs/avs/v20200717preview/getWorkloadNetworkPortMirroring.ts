// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * NSX Port Mirroring
 */
export function getWorkloadNetworkPortMirroring(args: GetWorkloadNetworkPortMirroringArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkloadNetworkPortMirroringResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:avs/v20200717preview:getWorkloadNetworkPortMirroring", {
        "portMirroringId": args.portMirroringId,
        "privateCloudName": args.privateCloudName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetWorkloadNetworkPortMirroringArgs {
    /**
     * NSX Port Mirroring identifier. Generally the same as the Port Mirroring display name
     */
    readonly portMirroringId: string;
    /**
     * Name of the private cloud
     */
    readonly privateCloudName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * NSX Port Mirroring
 */
export interface GetWorkloadNetworkPortMirroringResult {
    /**
     * Destination VM Group.
     */
    readonly destination?: string;
    /**
     * Direction of port mirroring profile.
     */
    readonly direction?: string;
    /**
     * Display name of the port mirroring profile.
     */
    readonly displayName?: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state
     */
    readonly provisioningState: string;
    /**
     * NSX revision number.
     */
    readonly revision?: number;
    /**
     * Source VM Group.
     */
    readonly source?: string;
    /**
     * Port Mirroring Status.
     */
    readonly status: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
