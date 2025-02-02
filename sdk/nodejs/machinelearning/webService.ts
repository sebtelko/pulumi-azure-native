// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Instance of an Azure ML web service resource.
 * API Version: 2017-01-01.
 */
export class WebService extends pulumi.CustomResource {
    /**
     * Get an existing WebService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebService {
        return new WebService(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:machinelearning:WebService';

    /**
     * Returns true if the given object is an instance of WebService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebService.__pulumiType;
    }

    /**
     * Specifies the location of the resource.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Contains the property payload that describes the web service.
     */
    public readonly properties!: pulumi.Output<outputs.machinelearning.WebServicePropertiesForGraphResponse>;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WebService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.properties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'properties'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["webServiceName"] = args ? args.webServiceName : undefined;
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
        const aliasOpts = { aliases: [{ type: "azure-nextgen:machinelearning:WebService" }, { type: "azure-native:machinelearning/v20160501preview:WebService" }, { type: "azure-nextgen:machinelearning/v20160501preview:WebService" }, { type: "azure-native:machinelearning/v20170101:WebService" }, { type: "azure-nextgen:machinelearning/v20170101:WebService" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(WebService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebService resource.
 */
export interface WebServiceArgs {
    /**
     * Specifies the location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Contains the property payload that describes the web service.
     */
    readonly properties: pulumi.Input<inputs.machinelearning.WebServicePropertiesForGraphArgs>;
    /**
     * Name of the resource group in which the web service is located.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the web service.
     */
    readonly webServiceName?: pulumi.Input<string>;
}
