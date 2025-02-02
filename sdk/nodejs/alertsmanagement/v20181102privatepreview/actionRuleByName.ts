// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Action rule object containing target scope, conditions and suppression logic
 */
export class ActionRuleByName extends pulumi.CustomResource {
    /**
     * Get an existing ActionRuleByName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ActionRuleByName {
        return new ActionRuleByName(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:alertsmanagement/v20181102privatepreview:ActionRuleByName';

    /**
     * Returns true if the given object is an instance of ActionRuleByName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionRuleByName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionRuleByName.__pulumiType;
    }

    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Action rule properties defining scope, conditions, suppression logic for action rule
     */
    public readonly properties!: pulumi.Output<outputs.alertsmanagement.v20181102privatepreview.ActionRulePropertiesResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ActionRuleByName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionRuleByNameArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroup'");
            }
            inputs["actionRuleName"] = args ? args.actionRuleName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroup"] = args ? args.resourceGroup : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:alertsmanagement/v20181102privatepreview:ActionRuleByName" }, { type: "azure-native:alertsmanagement:ActionRuleByName" }, { type: "azure-nextgen:alertsmanagement:ActionRuleByName" }, { type: "azure-native:alertsmanagement/v20190505preview:ActionRuleByName" }, { type: "azure-nextgen:alertsmanagement/v20190505preview:ActionRuleByName" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ActionRuleByName.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ActionRuleByName resource.
 */
export interface ActionRuleByNameArgs {
    /**
     * The name of action rule that needs to be created/updated
     */
    readonly actionRuleName?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Action rule properties defining scope, conditions, suppression logic for action rule
     */
    readonly properties?: pulumi.Input<inputs.alertsmanagement.v20181102privatepreview.ActionRulePropertiesArgs>;
    /**
     * Resource group name where the resource is created.
     */
    readonly resourceGroup: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
