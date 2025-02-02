// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * An application type version resource for the specified application type name resource.
 */
export function getApplicationTypeVersion(args: GetApplicationTypeVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationTypeVersionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:servicefabric/v20170701preview:getApplicationTypeVersion", {
        "applicationTypeName": args.applicationTypeName,
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
        "version": args.version,
    }, opts);
}

export interface GetApplicationTypeVersionArgs {
    /**
     * The name of the application type name resource.
     */
    readonly applicationTypeName: string;
    /**
     * The name of the cluster resource.
     */
    readonly clusterName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The application type version.
     */
    readonly version: string;
}

/**
 * An application type version resource for the specified application type name resource.
 */
export interface GetApplicationTypeVersionResult {
    /**
     * The URL to the application package
     */
    readonly appPackageUrl: string;
    /**
     * List of application type parameters that can be overridden when creating or updating the application.
     */
    readonly defaultParameterList: {[key: string]: string};
    /**
     * Azure resource identifier.
     */
    readonly id: string;
    /**
     * Azure resource location.
     */
    readonly location?: string;
    /**
     * Azure resource name.
     */
    readonly name: string;
    /**
     * The current deployment or provisioning state, which only appears in the response
     */
    readonly provisioningState: string;
    /**
     * Azure resource type.
     */
    readonly type: string;
}
