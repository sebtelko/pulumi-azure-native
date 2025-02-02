// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * A report resource.
 * API Version: 2018-08-01-preview.
 */
export class ReportByDepartment extends pulumi.CustomResource {
    /**
     * Get an existing ReportByDepartment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReportByDepartment {
        return new ReportByDepartment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:costmanagement:ReportByDepartment';

    /**
     * Returns true if the given object is an instance of ReportByDepartment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReportByDepartment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReportByDepartment.__pulumiType;
    }

    /**
     * Has definition for the report.
     */
    public readonly definition!: pulumi.Output<outputs.costmanagement.ReportDefinitionResponse>;
    /**
     * Has delivery information for the report.
     */
    public readonly deliveryInfo!: pulumi.Output<outputs.costmanagement.ReportDeliveryInfoResponse>;
    /**
     * The format of the report being delivered.
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Has schedule information for the report.
     */
    public readonly schedule!: pulumi.Output<outputs.costmanagement.ReportScheduleResponse | undefined>;
    /**
     * Resource tags.
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ReportByDepartment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReportByDepartmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.definition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'definition'");
            }
            if ((!args || args.deliveryInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deliveryInfo'");
            }
            if ((!args || args.departmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'departmentId'");
            }
            inputs["definition"] = args ? args.definition : undefined;
            inputs["deliveryInfo"] = args ? args.deliveryInfo : undefined;
            inputs["departmentId"] = args ? args.departmentId : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["reportName"] = args ? args.reportName : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["definition"] = undefined /*out*/;
            inputs["deliveryInfo"] = undefined /*out*/;
            inputs["format"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["schedule"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:costmanagement:ReportByDepartment" }, { type: "azure-native:costmanagement/v20180801preview:ReportByDepartment" }, { type: "azure-nextgen:costmanagement/v20180801preview:ReportByDepartment" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ReportByDepartment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReportByDepartment resource.
 */
export interface ReportByDepartmentArgs {
    /**
     * Has definition for the report.
     */
    readonly definition: pulumi.Input<inputs.costmanagement.ReportDefinitionArgs>;
    /**
     * Has delivery information for the report.
     */
    readonly deliveryInfo: pulumi.Input<inputs.costmanagement.ReportDeliveryInfoArgs>;
    /**
     * Department ID
     */
    readonly departmentId: pulumi.Input<string>;
    /**
     * The format of the report being delivered.
     */
    readonly format?: pulumi.Input<string | enums.costmanagement.FormatType>;
    /**
     * Report Name.
     */
    readonly reportName?: pulumi.Input<string>;
    /**
     * Has schedule information for the report.
     */
    readonly schedule?: pulumi.Input<inputs.costmanagement.ReportScheduleArgs>;
}
