// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Description of Rule Resource.
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:servicebus/v20210101preview:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
     */
    public readonly action!: pulumi.Output<outputs.servicebus.v20210101preview.ActionResponse | undefined>;
    /**
     * Properties of correlationFilter
     */
    public readonly correlationFilter!: pulumi.Output<outputs.servicebus.v20210101preview.CorrelationFilterResponse | undefined>;
    /**
     * Filter type that is evaluated against a BrokeredMessage.
     */
    public readonly filterType!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties of sqlFilter
     */
    public readonly sqlFilter!: pulumi.Output<outputs.servicebus.v20210101preview.SqlFilterResponse | undefined>;
    /**
     * The system meta data relating to this resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.servicebus.v20210101preview.SystemDataResponse>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.namespaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.subscriptionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscriptionName'");
            }
            if ((!args || args.topicName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicName'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["correlationFilter"] = args ? args.correlationFilter : undefined;
            inputs["filterType"] = args ? args.filterType : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["ruleName"] = args ? args.ruleName : undefined;
            inputs["sqlFilter"] = args ? args.sqlFilter : undefined;
            inputs["subscriptionName"] = args ? args.subscriptionName : undefined;
            inputs["topicName"] = args ? args.topicName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["action"] = undefined /*out*/;
            inputs["correlationFilter"] = undefined /*out*/;
            inputs["filterType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sqlFilter"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:servicebus/v20210101preview:Rule" }, { type: "azure-native:servicebus:Rule" }, { type: "azure-nextgen:servicebus:Rule" }, { type: "azure-native:servicebus/v20170401:Rule" }, { type: "azure-nextgen:servicebus/v20170401:Rule" }, { type: "azure-native:servicebus/v20180101preview:Rule" }, { type: "azure-nextgen:servicebus/v20180101preview:Rule" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Rule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
     */
    readonly action?: pulumi.Input<inputs.servicebus.v20210101preview.ActionArgs>;
    /**
     * Properties of correlationFilter
     */
    readonly correlationFilter?: pulumi.Input<inputs.servicebus.v20210101preview.CorrelationFilterArgs>;
    /**
     * Filter type that is evaluated against a BrokeredMessage.
     */
    readonly filterType?: pulumi.Input<enums.servicebus.v20210101preview.FilterType>;
    /**
     * The namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The rule name.
     */
    readonly ruleName?: pulumi.Input<string>;
    /**
     * Properties of sqlFilter
     */
    readonly sqlFilter?: pulumi.Input<inputs.servicebus.v20210101preview.SqlFilterArgs>;
    /**
     * The subscription name.
     */
    readonly subscriptionName: pulumi.Input<string>;
    /**
     * The topic name.
     */
    readonly topicName: pulumi.Input<string>;
}
