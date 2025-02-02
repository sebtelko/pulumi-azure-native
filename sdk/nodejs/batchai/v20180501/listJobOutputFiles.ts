// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Values returned by the List operation.
 */
export function listJobOutputFiles(args: ListJobOutputFilesArgs, opts?: pulumi.InvokeOptions): Promise<ListJobOutputFilesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:batchai/v20180501:listJobOutputFiles", {
        "directory": args.directory,
        "experimentName": args.experimentName,
        "jobName": args.jobName,
        "linkexpiryinminutes": args.linkexpiryinminutes,
        "maxResults": args.maxResults,
        "outputdirectoryid": args.outputdirectoryid,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface ListJobOutputFilesArgs {
    /**
     * The path to the directory.
     */
    readonly directory?: string;
    /**
     * The name of the experiment. Experiment names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly experimentName: string;
    /**
     * The name of the job within the specified resource group. Job names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly jobName: string;
    /**
     * The number of minutes after which the download link will expire.
     */
    readonly linkexpiryinminutes?: number;
    /**
     * The maximum number of items to return in the response. A maximum of 1000 files can be returned.
     */
    readonly maxResults?: number;
    /**
     * Id of the job output directory. This is the OutputDirectory-->id parameter that is given by the user during Create Job.
     */
    readonly outputdirectoryid: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly workspaceName: string;
}

/**
 * Values returned by the List operation.
 */
export interface ListJobOutputFilesResult {
    /**
     * The continuation token.
     */
    readonly nextLink: string;
    /**
     * The collection of returned job directories and files.
     */
    readonly value: outputs.batchai.v20180501.FileResponse[];
}
