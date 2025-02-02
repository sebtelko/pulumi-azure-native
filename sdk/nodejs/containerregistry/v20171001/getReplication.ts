// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * An object that represents a replication for a container registry.
 */
export function getReplication(args: GetReplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:containerregistry/v20171001:getReplication", {
        "registryName": args.registryName,
        "replicationName": args.replicationName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetReplicationArgs {
    /**
     * The name of the container registry.
     */
    readonly registryName: string;
    /**
     * The name of the replication.
     */
    readonly replicationName: string;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * An object that represents a replication for a container registry.
 */
export interface GetReplicationResult {
    /**
     * The resource ID.
     */
    readonly id: string;
    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    readonly location: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The provisioning state of the replication at the time the operation was called.
     */
    readonly provisioningState: string;
    /**
     * The status of the replication at the time the operation was called.
     */
    readonly status: outputs.containerregistry.v20171001.StatusResponse;
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}
