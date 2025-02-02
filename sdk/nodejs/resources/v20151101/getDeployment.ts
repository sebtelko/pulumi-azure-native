// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Deployment information.
 */
export function getDeployment(args: GetDeploymentArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:resources/v20151101:getDeployment", {
        "deploymentName": args.deploymentName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDeploymentArgs {
    /**
     * The name of the deployment.
     */
    readonly deploymentName: string;
    /**
     * The name of the resource group to get. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Deployment information.
 */
export interface GetDeploymentResult {
    /**
     * Gets or sets the ID of the deployment.
     */
    readonly id: string;
    /**
     * Gets or sets the name of the deployment.
     */
    readonly name: string;
    /**
     * Gets or sets deployment properties.
     */
    readonly properties: outputs.resources.v20151101.DeploymentPropertiesExtendedResponse;
}
