// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The log profile resource.
 */
export class LogProfile extends pulumi.CustomResource {
    /**
     * Get an existing LogProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LogProfile {
        return new LogProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:insights/v20160301:LogProfile';

    /**
     * Returns true if the given object is an instance of LogProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogProfile.__pulumiType;
    }

    /**
     * the categories of the logs. These categories are created as is convenient to the user. Some values are: 'Write', 'Delete', and/or 'Action.'
     */
    public readonly categories!: pulumi.Output<string[]>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * List of regions for which Activity Log events should be stored or streamed. It is a comma separated list of valid ARM locations including the 'global' location.
     */
    public readonly locations!: pulumi.Output<string[]>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * the retention policy for the events in the log.
     */
    public readonly retentionPolicy!: pulumi.Output<outputs.insights.v20160301.RetentionPolicyResponse>;
    /**
     * The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log. The rule ID is of the format: '{service bus resource ID}/authorizationrules/{key name}'.
     */
    public readonly serviceBusRuleId!: pulumi.Output<string | undefined>;
    /**
     * the resource id of the storage account to which you would like to send the Activity Log.
     */
    public readonly storageAccountId!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a LogProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.categories === undefined) && !opts.urn) {
                throw new Error("Missing required property 'categories'");
            }
            if ((!args || args.locations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locations'");
            }
            if ((!args || args.retentionPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'retentionPolicy'");
            }
            inputs["categories"] = args ? args.categories : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["locations"] = args ? args.locations : undefined;
            inputs["logProfileName"] = args ? args.logProfileName : undefined;
            inputs["retentionPolicy"] = args ? args.retentionPolicy : undefined;
            inputs["serviceBusRuleId"] = args ? args.serviceBusRuleId : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["categories"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["locations"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["retentionPolicy"] = undefined /*out*/;
            inputs["serviceBusRuleId"] = undefined /*out*/;
            inputs["storageAccountId"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:insights/v20160301:LogProfile" }, { type: "azure-native:insights:LogProfile" }, { type: "azure-nextgen:insights:LogProfile" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(LogProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LogProfile resource.
 */
export interface LogProfileArgs {
    /**
     * the categories of the logs. These categories are created as is convenient to the user. Some values are: 'Write', 'Delete', and/or 'Action.'
     */
    readonly categories: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * List of regions for which Activity Log events should be stored or streamed. It is a comma separated list of valid ARM locations including the 'global' location.
     */
    readonly locations: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the log profile.
     */
    readonly logProfileName?: pulumi.Input<string>;
    /**
     * the retention policy for the events in the log.
     */
    readonly retentionPolicy: pulumi.Input<inputs.insights.v20160301.RetentionPolicyArgs>;
    /**
     * The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log. The rule ID is of the format: '{service bus resource ID}/authorizationrules/{key name}'.
     */
    readonly serviceBusRuleId?: pulumi.Input<string>;
    /**
     * the resource id of the storage account to which you would like to send the Activity Log.
     */
    readonly storageAccountId?: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
