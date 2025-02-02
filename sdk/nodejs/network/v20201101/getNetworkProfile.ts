// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Network profile resource.
 */
export function getNetworkProfile(args: GetNetworkProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkProfileResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:network/v20201101:getNetworkProfile", {
        "expand": args.expand,
        "networkProfileName": args.networkProfileName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNetworkProfileArgs {
    /**
     * Expands referenced resources.
     */
    readonly expand?: string;
    /**
     * The name of the public IP prefix.
     */
    readonly networkProfileName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Network profile resource.
 */
export interface GetNetworkProfileResult {
    /**
     * List of chid container network interface configurations.
     */
    readonly containerNetworkInterfaceConfigurations?: outputs.network.v20201101.ContainerNetworkInterfaceConfigurationResponse[];
    /**
     * List of child container network interfaces.
     */
    readonly containerNetworkInterfaces: outputs.network.v20201101.ContainerNetworkInterfaceResponse[];
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
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
     * The provisioning state of the network profile resource.
     */
    readonly provisioningState: string;
    /**
     * The resource GUID property of the network profile resource.
     */
    readonly resourceGuid: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
