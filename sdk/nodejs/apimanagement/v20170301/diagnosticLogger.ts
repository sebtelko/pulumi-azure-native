// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Logger details.
 */
export class DiagnosticLogger extends pulumi.CustomResource {
    /**
     * Get an existing DiagnosticLogger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DiagnosticLogger {
        return new DiagnosticLogger(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:apimanagement/v20170301:DiagnosticLogger';

    /**
     * Returns true if the given object is an instance of DiagnosticLogger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DiagnosticLogger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DiagnosticLogger.__pulumiType;
    }

    /**
     * The name and SendRule connection string of the event hub for azureEventHub logger.
     * Instrumentation key for applicationInsights logger.
     */
    public /*out*/ readonly credentials!: pulumi.Output<{[key: string]: string}>;
    /**
     * Logger description.
     */
    public /*out*/ readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether records are buffered in the logger before publishing. Default is assumed to be true.
     */
    public /*out*/ readonly isBuffered!: pulumi.Output<boolean | undefined>;
    /**
     * Logger type.
     */
    public /*out*/ readonly loggerType!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Sampling settings for an ApplicationInsights logger.
     */
    public /*out*/ readonly sampling!: pulumi.Output<outputs.apimanagement.v20170301.LoggerSamplingContractResponse | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DiagnosticLogger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiagnosticLoggerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.diagnosticId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diagnosticId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["diagnosticId"] = args ? args.diagnosticId : undefined;
            inputs["loggerid"] = args ? args.loggerid : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["credentials"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["isBuffered"] = undefined /*out*/;
            inputs["loggerType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sampling"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["credentials"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["isBuffered"] = undefined /*out*/;
            inputs["loggerType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sampling"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:apimanagement/v20170301:DiagnosticLogger" }, { type: "azure-native:apimanagement:DiagnosticLogger" }, { type: "azure-nextgen:apimanagement:DiagnosticLogger" }, { type: "azure-native:apimanagement/v20180101:DiagnosticLogger" }, { type: "azure-nextgen:apimanagement/v20180101:DiagnosticLogger" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DiagnosticLogger.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DiagnosticLogger resource.
 */
export interface DiagnosticLoggerArgs {
    /**
     * Diagnostic identifier. Must be unique in the current API Management service instance.
     */
    readonly diagnosticId: pulumi.Input<string>;
    /**
     * Logger identifier. Must be unique in the API Management service instance.
     */
    readonly loggerid?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
}
