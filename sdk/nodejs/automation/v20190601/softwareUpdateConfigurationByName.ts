// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Software update configuration properties.
 */
export class SoftwareUpdateConfigurationByName extends pulumi.CustomResource {
    /**
     * Get an existing SoftwareUpdateConfigurationByName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SoftwareUpdateConfigurationByName {
        return new SoftwareUpdateConfigurationByName(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:automation/v20190601:SoftwareUpdateConfigurationByName';

    /**
     * Returns true if the given object is an instance of SoftwareUpdateConfigurationByName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SoftwareUpdateConfigurationByName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SoftwareUpdateConfigurationByName.__pulumiType;
    }

    /**
     * CreatedBy property, which only appears in the response.
     */
    public /*out*/ readonly createdBy!: pulumi.Output<string>;
    /**
     * Creation time of the resource, which only appears in the response.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Details of provisioning error
     */
    public readonly error!: pulumi.Output<outputs.automation.v20190601.ErrorResponseResponse | undefined>;
    /**
     * LastModifiedBy property, which only appears in the response.
     */
    public /*out*/ readonly lastModifiedBy!: pulumi.Output<string>;
    /**
     * Last time resource was modified, which only appears in the response.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state for the software update configuration, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Schedule information for the Software update configuration
     */
    public readonly scheduleInfo!: pulumi.Output<outputs.automation.v20190601.SUCSchedulePropertiesResponse>;
    /**
     * Tasks information for the Software update configuration.
     */
    public readonly tasks!: pulumi.Output<outputs.automation.v20190601.SoftwareUpdateConfigurationTasksResponse | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * update specific properties for the Software update configuration
     */
    public readonly updateConfiguration!: pulumi.Output<outputs.automation.v20190601.UpdateConfigurationResponse>;

    /**
     * Create a SoftwareUpdateConfigurationByName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SoftwareUpdateConfigurationByNameArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.automationAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.scheduleInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scheduleInfo'");
            }
            if ((!args || args.updateConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'updateConfiguration'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["error"] = args ? args.error : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scheduleInfo"] = args ? args.scheduleInfo : undefined;
            inputs["softwareUpdateConfigurationName"] = args ? args.softwareUpdateConfigurationName : undefined;
            inputs["tasks"] = args ? args.tasks : undefined;
            inputs["updateConfiguration"] = args ? args.updateConfiguration : undefined;
            inputs["createdBy"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["lastModifiedBy"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["createdBy"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["error"] = undefined /*out*/;
            inputs["lastModifiedBy"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["scheduleInfo"] = undefined /*out*/;
            inputs["tasks"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updateConfiguration"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:automation/v20190601:SoftwareUpdateConfigurationByName" }, { type: "azure-native:automation:SoftwareUpdateConfigurationByName" }, { type: "azure-nextgen:automation:SoftwareUpdateConfigurationByName" }, { type: "azure-native:automation/v20170515preview:SoftwareUpdateConfigurationByName" }, { type: "azure-nextgen:automation/v20170515preview:SoftwareUpdateConfigurationByName" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(SoftwareUpdateConfigurationByName.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SoftwareUpdateConfigurationByName resource.
 */
export interface SoftwareUpdateConfigurationByNameArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * Details of provisioning error
     */
    readonly error?: pulumi.Input<inputs.automation.v20190601.ErrorResponseArgs>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Schedule information for the Software update configuration
     */
    readonly scheduleInfo: pulumi.Input<inputs.automation.v20190601.SUCSchedulePropertiesArgs>;
    /**
     * The name of the software update configuration to be created.
     */
    readonly softwareUpdateConfigurationName?: pulumi.Input<string>;
    /**
     * Tasks information for the Software update configuration.
     */
    readonly tasks?: pulumi.Input<inputs.automation.v20190601.SoftwareUpdateConfigurationTasksArgs>;
    /**
     * update specific properties for the Software update configuration
     */
    readonly updateConfiguration: pulumi.Input<inputs.automation.v20190601.UpdateConfigurationArgs>;
}
