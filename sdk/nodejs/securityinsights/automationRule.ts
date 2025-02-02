// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Represents an automation rule.
 * API Version: 2019-01-01-preview.
 */
export class AutomationRule extends pulumi.CustomResource {
    /**
     * Get an existing AutomationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AutomationRule {
        return new AutomationRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:securityinsights:AutomationRule';

    /**
     * Returns true if the given object is an instance of AutomationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutomationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutomationRule.__pulumiType;
    }

    /**
     * The actions to execute when the automation rule is triggered
     */
    public readonly actions!: pulumi.Output<outputs.securityinsights.AutomationRuleModifyPropertiesActionResponse | outputs.securityinsights.AutomationRuleRunPlaybookActionResponse[]>;
    /**
     * Describes the client that created the automation rule
     */
    public /*out*/ readonly createdBy!: pulumi.Output<outputs.securityinsights.ClientInfoResponse>;
    /**
     * The time the automation rule was created
     */
    public /*out*/ readonly createdTimeUtc!: pulumi.Output<string>;
    /**
     * The display name of the automation  rule
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Etag of the azure resource
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Describes the client that last updated the automation rule
     */
    public /*out*/ readonly lastModifiedBy!: pulumi.Output<outputs.securityinsights.ClientInfoResponse>;
    /**
     * The last time the automation rule was updated
     */
    public /*out*/ readonly lastModifiedTimeUtc!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The order of execution of the automation rule
     */
    public readonly order!: pulumi.Output<number>;
    /**
     * The triggering logic of the automation rule
     */
    public readonly triggeringLogic!: pulumi.Output<outputs.securityinsights.AutomationRuleTriggeringLogicResponse>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AutomationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutomationRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.actions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actions'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.operationalInsightsResourceProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'operationalInsightsResourceProvider'");
            }
            if ((!args || args.order === undefined) && !opts.urn) {
                throw new Error("Missing required property 'order'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.triggeringLogic === undefined) && !opts.urn) {
                throw new Error("Missing required property 'triggeringLogic'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["automationRuleId"] = args ? args.automationRuleId : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["operationalInsightsResourceProvider"] = args ? args.operationalInsightsResourceProvider : undefined;
            inputs["order"] = args ? args.order : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["triggeringLogic"] = args ? args.triggeringLogic : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["createdBy"] = undefined /*out*/;
            inputs["createdTimeUtc"] = undefined /*out*/;
            inputs["lastModifiedBy"] = undefined /*out*/;
            inputs["lastModifiedTimeUtc"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["actions"] = undefined /*out*/;
            inputs["createdBy"] = undefined /*out*/;
            inputs["createdTimeUtc"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["lastModifiedBy"] = undefined /*out*/;
            inputs["lastModifiedTimeUtc"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["order"] = undefined /*out*/;
            inputs["triggeringLogic"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:securityinsights:AutomationRule" }, { type: "azure-native:securityinsights/v20190101preview:AutomationRule" }, { type: "azure-nextgen:securityinsights/v20190101preview:AutomationRule" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(AutomationRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AutomationRule resource.
 */
export interface AutomationRuleArgs {
    /**
     * The actions to execute when the automation rule is triggered
     */
    readonly actions: pulumi.Input<pulumi.Input<inputs.securityinsights.AutomationRuleModifyPropertiesActionArgs | inputs.securityinsights.AutomationRuleRunPlaybookActionArgs>[]>;
    /**
     * Automation rule ID
     */
    readonly automationRuleId?: pulumi.Input<string>;
    /**
     * The display name of the automation  rule
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Etag of the azure resource
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The namespace of workspaces resource provider- Microsoft.OperationalInsights.
     */
    readonly operationalInsightsResourceProvider: pulumi.Input<string>;
    /**
     * The order of execution of the automation rule
     */
    readonly order: pulumi.Input<number>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The triggering logic of the automation rule
     */
    readonly triggeringLogic: pulumi.Input<inputs.securityinsights.AutomationRuleTriggeringLogicArgs>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
