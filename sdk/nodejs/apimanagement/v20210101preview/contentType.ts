// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Content type contract details.
 */
export class ContentType extends pulumi.CustomResource {
    /**
     * Get an existing ContentType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ContentType {
        return new ContentType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:apimanagement/v20210101preview:ContentType';

    /**
     * Returns true if the given object is an instance of ContentType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContentType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContentType.__pulumiType;
    }

    /**
     * Content type description.
     */
    public /*out*/ readonly description!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Content type schema.
     */
    public /*out*/ readonly schema!: pulumi.Output<any | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Content type version.
     */
    public /*out*/ readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a ContentType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContentTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["contentTypeId"] = args ? args.contentTypeId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["description"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["schema"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["schema"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:apimanagement/v20210101preview:ContentType" }, { type: "azure-native:apimanagement:ContentType" }, { type: "azure-nextgen:apimanagement:ContentType" }, { type: "azure-native:apimanagement/v20191201:ContentType" }, { type: "azure-nextgen:apimanagement/v20191201:ContentType" }, { type: "azure-native:apimanagement/v20200601preview:ContentType" }, { type: "azure-nextgen:apimanagement/v20200601preview:ContentType" }, { type: "azure-native:apimanagement/v20201201:ContentType" }, { type: "azure-nextgen:apimanagement/v20201201:ContentType" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ContentType.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ContentType resource.
 */
export interface ContentTypeArgs {
    /**
     * Content type identifier.
     */
    readonly contentTypeId?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
}
