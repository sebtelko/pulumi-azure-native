// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export function getResourceGuardProxy(args: GetResourceGuardProxyArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceGuardProxyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:recoveryservices/v20210201preview:getResourceGuardProxy", {
        "resourceGroupName": args.resourceGroupName,
        "resourceGuardProxyName": args.resourceGuardProxyName,
        "vaultName": args.vaultName,
    }, opts);
}

export interface GetResourceGuardProxyArgs {
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: string;
    readonly resourceGuardProxyName: string;
    /**
     * The name of the recovery services vault.
     */
    readonly vaultName: string;
}

export interface GetResourceGuardProxyResult {
    /**
     * Optional ETag.
     */
    readonly eTag?: string;
    /**
     * Resource Id represents the complete path to the resource.
     */
    readonly id: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name associated with the resource.
     */
    readonly name: string;
    /**
     * ResourceGuardProxyBaseResource properties
     */
    readonly properties: outputs.recoveryservices.v20210201preview.ResourceGuardProxyBaseResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
     */
    readonly type: string;
}
