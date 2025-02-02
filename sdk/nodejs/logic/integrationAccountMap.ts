// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The integration account map.
 * API Version: 2019-05-01.
 */
export class IntegrationAccountMap extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationAccountMap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IntegrationAccountMap {
        return new IntegrationAccountMap(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:logic:IntegrationAccountMap';

    /**
     * Returns true if the given object is an instance of IntegrationAccountMap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationAccountMap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationAccountMap.__pulumiType;
    }

    /**
     * The changed time.
     */
    public /*out*/ readonly changedTime!: pulumi.Output<string>;
    /**
     * The content.
     */
    public readonly content!: pulumi.Output<string | undefined>;
    /**
     * The content link.
     */
    public /*out*/ readonly contentLink!: pulumi.Output<outputs.logic.ContentLinkResponse>;
    /**
     * The content type.
     */
    public readonly contentType!: pulumi.Output<string | undefined>;
    /**
     * The created time.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The map type.
     */
    public readonly mapType!: pulumi.Output<string>;
    /**
     * The metadata.
     */
    public readonly metadata!: pulumi.Output<any | undefined>;
    /**
     * Gets the resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parameters schema of integration account map.
     */
    public readonly parametersSchema!: pulumi.Output<outputs.logic.IntegrationAccountMapPropertiesResponseParametersSchema | undefined>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Gets the resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IntegrationAccountMap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationAccountMapArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.integrationAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationAccountName'");
            }
            if ((!args || args.mapType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mapType'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["content"] = args ? args.content : undefined;
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["integrationAccountName"] = args ? args.integrationAccountName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["mapName"] = args ? args.mapName : undefined;
            inputs["mapType"] = args ? args.mapType : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["parametersSchema"] = args ? args.parametersSchema : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["changedTime"] = undefined /*out*/;
            inputs["contentLink"] = undefined /*out*/;
            inputs["createdTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["changedTime"] = undefined /*out*/;
            inputs["content"] = undefined /*out*/;
            inputs["contentLink"] = undefined /*out*/;
            inputs["contentType"] = undefined /*out*/;
            inputs["createdTime"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["mapType"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parametersSchema"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:logic:IntegrationAccountMap" }, { type: "azure-native:logic/v20150801preview:IntegrationAccountMap" }, { type: "azure-nextgen:logic/v20150801preview:IntegrationAccountMap" }, { type: "azure-native:logic/v20160601:IntegrationAccountMap" }, { type: "azure-nextgen:logic/v20160601:IntegrationAccountMap" }, { type: "azure-native:logic/v20180701preview:IntegrationAccountMap" }, { type: "azure-nextgen:logic/v20180701preview:IntegrationAccountMap" }, { type: "azure-native:logic/v20190501:IntegrationAccountMap" }, { type: "azure-nextgen:logic/v20190501:IntegrationAccountMap" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(IntegrationAccountMap.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IntegrationAccountMap resource.
 */
export interface IntegrationAccountMapArgs {
    /**
     * The content.
     */
    readonly content?: pulumi.Input<string>;
    /**
     * The content type.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * The integration account name.
     */
    readonly integrationAccountName: pulumi.Input<string>;
    /**
     * The resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The integration account map name.
     */
    readonly mapName?: pulumi.Input<string>;
    /**
     * The map type.
     */
    readonly mapType: pulumi.Input<string | enums.logic.MapType>;
    /**
     * The metadata.
     */
    readonly metadata?: any;
    /**
     * The parameters schema of integration account map.
     */
    readonly parametersSchema?: pulumi.Input<inputs.logic.IntegrationAccountMapPropertiesParametersSchemaArgs>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
