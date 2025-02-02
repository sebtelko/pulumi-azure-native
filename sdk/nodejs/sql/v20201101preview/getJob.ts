// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * A job.
 */
export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:sql/v20201101preview:getJob", {
        "jobAgentName": args.jobAgentName,
        "jobName": args.jobName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetJobArgs {
    /**
     * The name of the job agent.
     */
    readonly jobAgentName: string;
    /**
     * The name of the job to get.
     */
    readonly jobName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server.
     */
    readonly serverName: string;
}

/**
 * A job.
 */
export interface GetJobResult {
    /**
     * User-defined description of the job.
     */
    readonly description?: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Schedule properties of the job.
     */
    readonly schedule?: outputs.sql.v20201101preview.JobScheduleResponse;
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The job version number.
     */
    readonly version: number;
}
