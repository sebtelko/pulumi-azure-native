// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Definition of the connection type.
 */
export class ConnectionType extends pulumi.CustomResource {
    /**
     * Get an existing ConnectionType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectionType {
        return new ConnectionType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:automation/v20190601:ConnectionType';

    /**
     * Returns true if the given object is an instance of ConnectionType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectionType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectionType.__pulumiType;
    }

    /**
     * Gets the creation time.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Gets or sets the description.
     */
    public /*out*/ readonly description!: pulumi.Output<string | undefined>;
    /**
     * Gets the field definitions of the connection type.
     */
    public readonly fieldDefinitions!: pulumi.Output<{[key: string]: outputs.automation.v20190601.FieldDefinitionResponse}>;
    /**
     * Gets or sets a Boolean value to indicate if the connection type is global.
     */
    public readonly isGlobal!: pulumi.Output<boolean | undefined>;
    /**
     * Gets or sets the last modified time.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string | undefined>;
    /**
     * Gets the name of the connection type.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ConnectionType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.automationAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if ((!args || args.fieldDefinitions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fieldDefinitions'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["connectionTypeName"] = args ? args.connectionTypeName : undefined;
            inputs["fieldDefinitions"] = args ? args.fieldDefinitions : undefined;
            inputs["isGlobal"] = args ? args.isGlobal : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["creationTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["fieldDefinitions"] = undefined /*out*/;
            inputs["isGlobal"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:automation/v20190601:ConnectionType" }, { type: "azure-native:automation:ConnectionType" }, { type: "azure-nextgen:automation:ConnectionType" }, { type: "azure-native:automation/v20151031:ConnectionType" }, { type: "azure-nextgen:automation/v20151031:ConnectionType" }, { type: "azure-native:automation/v20200113preview:ConnectionType" }, { type: "azure-nextgen:automation/v20200113preview:ConnectionType" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ConnectionType.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConnectionType resource.
 */
export interface ConnectionTypeArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * The parameters supplied to the create or update connection type operation.
     */
    readonly connectionTypeName?: pulumi.Input<string>;
    /**
     * Gets or sets the field definitions of the connection type.
     */
    readonly fieldDefinitions: pulumi.Input<{[key: string]: pulumi.Input<inputs.automation.v20190601.FieldDefinitionArgs>}>;
    /**
     * Gets or sets a Boolean value to indicate if the connection type is global.
     */
    readonly isGlobal?: pulumi.Input<boolean>;
    /**
     * Gets or sets the name of the connection type.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
