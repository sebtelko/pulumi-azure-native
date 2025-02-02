// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * A report resource.
 * API Version: 2018-08-01-preview.
 */
export function getReportByDepartment(args: GetReportByDepartmentArgs, opts?: pulumi.InvokeOptions): Promise<GetReportByDepartmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:costmanagement:getReportByDepartment", {
        "departmentId": args.departmentId,
        "reportName": args.reportName,
    }, opts);
}

export interface GetReportByDepartmentArgs {
    /**
     * Department ID
     */
    readonly departmentId: string;
    /**
     * Report Name.
     */
    readonly reportName: string;
}

/**
 * A report resource.
 */
export interface GetReportByDepartmentResult {
    /**
     * Has definition for the report.
     */
    readonly definition: outputs.costmanagement.ReportDefinitionResponse;
    /**
     * Has delivery information for the report.
     */
    readonly deliveryInfo: outputs.costmanagement.ReportDeliveryInfoResponse;
    /**
     * The format of the report being delivered.
     */
    readonly format?: string;
    /**
     * Resource Id.
     */
    readonly id: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Has schedule information for the report.
     */
    readonly schedule?: outputs.costmanagement.ReportScheduleResponse;
    /**
     * Resource tags.
     */
    readonly tags: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
