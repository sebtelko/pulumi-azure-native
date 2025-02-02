// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The diagnostic setting resource.
 * API Version: 2017-05-01-preview.
 */
export class DiagnosticSetting extends pulumi.CustomResource {
    /**
     * Get an existing DiagnosticSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DiagnosticSetting {
        return new DiagnosticSetting(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:insights:DiagnosticSetting';

    /**
     * Returns true if the given object is an instance of DiagnosticSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DiagnosticSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DiagnosticSetting.__pulumiType;
    }

    /**
     * The resource Id for the event hub authorization rule.
     */
    public readonly eventHubAuthorizationRuleId!: pulumi.Output<string | undefined>;
    /**
     * The name of the event hub. If none is specified, the default event hub will be selected.
     */
    public readonly eventHubName!: pulumi.Output<string | undefined>;
    /**
     * A string indicating whether the export to Log Analytics should use the default destination type, i.e. AzureDiagnostics, or use a destination type constructed as follows: <normalized service identity>_<normalized category name>. Possible values are: Dedicated and null (null is default.)
     */
    public readonly logAnalyticsDestinationType!: pulumi.Output<string | undefined>;
    /**
     * The list of logs settings.
     */
    public readonly logs!: pulumi.Output<outputs.insights.LogSettingsResponse[] | undefined>;
    /**
     * The list of metric settings.
     */
    public readonly metrics!: pulumi.Output<outputs.insights.MetricSettingsResponse[] | undefined>;
    /**
     * Azure resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
     */
    public readonly serviceBusRuleId!: pulumi.Output<string | undefined>;
    /**
     * The resource ID of the storage account to which you would like to send Diagnostic Logs.
     */
    public readonly storageAccountId!: pulumi.Output<string | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
     */
    public readonly workspaceId!: pulumi.Output<string | undefined>;

    /**
     * Create a DiagnosticSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiagnosticSettingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceUri'");
            }
            inputs["eventHubAuthorizationRuleId"] = args ? args.eventHubAuthorizationRuleId : undefined;
            inputs["eventHubName"] = args ? args.eventHubName : undefined;
            inputs["logAnalyticsDestinationType"] = args ? args.logAnalyticsDestinationType : undefined;
            inputs["logs"] = args ? args.logs : undefined;
            inputs["metrics"] = args ? args.metrics : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceUri"] = args ? args.resourceUri : undefined;
            inputs["serviceBusRuleId"] = args ? args.serviceBusRuleId : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            inputs["workspaceId"] = args ? args.workspaceId : undefined;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["eventHubAuthorizationRuleId"] = undefined /*out*/;
            inputs["eventHubName"] = undefined /*out*/;
            inputs["logAnalyticsDestinationType"] = undefined /*out*/;
            inputs["logs"] = undefined /*out*/;
            inputs["metrics"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["serviceBusRuleId"] = undefined /*out*/;
            inputs["storageAccountId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["workspaceId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:insights:DiagnosticSetting" }, { type: "azure-native:insights/v20170501preview:DiagnosticSetting" }, { type: "azure-nextgen:insights/v20170501preview:DiagnosticSetting" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DiagnosticSetting.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DiagnosticSetting resource.
 */
export interface DiagnosticSettingArgs {
    /**
     * The resource Id for the event hub authorization rule.
     */
    readonly eventHubAuthorizationRuleId?: pulumi.Input<string>;
    /**
     * The name of the event hub. If none is specified, the default event hub will be selected.
     */
    readonly eventHubName?: pulumi.Input<string>;
    /**
     * A string indicating whether the export to Log Analytics should use the default destination type, i.e. AzureDiagnostics, or use a destination type constructed as follows: <normalized service identity>_<normalized category name>. Possible values are: Dedicated and null (null is default.)
     */
    readonly logAnalyticsDestinationType?: pulumi.Input<string>;
    /**
     * The list of logs settings.
     */
    readonly logs?: pulumi.Input<pulumi.Input<inputs.insights.LogSettingsArgs>[]>;
    /**
     * The list of metric settings.
     */
    readonly metrics?: pulumi.Input<pulumi.Input<inputs.insights.MetricSettingsArgs>[]>;
    /**
     * The name of the diagnostic setting.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The identifier of the resource.
     */
    readonly resourceUri: pulumi.Input<string>;
    /**
     * The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
     */
    readonly serviceBusRuleId?: pulumi.Input<string>;
    /**
     * The resource ID of the storage account to which you would like to send Diagnostic Logs.
     */
    readonly storageAccountId?: pulumi.Input<string>;
    /**
     * The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
     */
    readonly workspaceId?: pulumi.Input<string>;
}
