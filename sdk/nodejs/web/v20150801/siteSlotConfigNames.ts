// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Slot Config names azure resource
 */
export class SiteSlotConfigNames extends pulumi.CustomResource {
    /**
     * Get an existing SiteSlotConfigNames resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SiteSlotConfigNames {
        return new SiteSlotConfigNames(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:web/v20150801:SiteSlotConfigNames';

    /**
     * Returns true if the given object is an instance of SiteSlotConfigNames.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SiteSlotConfigNames {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SiteSlotConfigNames.__pulumiType;
    }

    /**
     * List of application settings names
     */
    public readonly appSettingNames!: pulumi.Output<string[] | undefined>;
    /**
     * List of connection string names
     */
    public readonly connectionStringNames!: pulumi.Output<string[] | undefined>;
    /**
     * Kind of resource
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource Name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a SiteSlotConfigNames resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SiteSlotConfigNamesArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["appSettingNames"] = args ? args.appSettingNames : undefined;
            inputs["connectionStringNames"] = args ? args.connectionStringNames : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
        } else {
            inputs["appSettingNames"] = undefined /*out*/;
            inputs["connectionStringNames"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:web/v20150801:SiteSlotConfigNames" }, { type: "azure-native:web:SiteSlotConfigNames" }, { type: "azure-nextgen:web:SiteSlotConfigNames" }, { type: "azure-native:web/v20160801:SiteSlotConfigNames" }, { type: "azure-nextgen:web/v20160801:SiteSlotConfigNames" }, { type: "azure-native:web/v20180201:SiteSlotConfigNames" }, { type: "azure-nextgen:web/v20180201:SiteSlotConfigNames" }, { type: "azure-native:web/v20181101:SiteSlotConfigNames" }, { type: "azure-nextgen:web/v20181101:SiteSlotConfigNames" }, { type: "azure-native:web/v20190801:SiteSlotConfigNames" }, { type: "azure-nextgen:web/v20190801:SiteSlotConfigNames" }, { type: "azure-native:web/v20200601:SiteSlotConfigNames" }, { type: "azure-nextgen:web/v20200601:SiteSlotConfigNames" }, { type: "azure-native:web/v20200901:SiteSlotConfigNames" }, { type: "azure-nextgen:web/v20200901:SiteSlotConfigNames" }, { type: "azure-native:web/v20201001:SiteSlotConfigNames" }, { type: "azure-nextgen:web/v20201001:SiteSlotConfigNames" }, { type: "azure-native:web/v20201201:SiteSlotConfigNames" }, { type: "azure-nextgen:web/v20201201:SiteSlotConfigNames" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(SiteSlotConfigNames.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SiteSlotConfigNames resource.
 */
export interface SiteSlotConfigNamesArgs {
    /**
     * List of application settings names
     */
    readonly appSettingNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of connection string names
     */
    readonly connectionStringNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource Id
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Kind of resource
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Resource Name
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of resource group
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource type
     */
    readonly type?: pulumi.Input<string>;
}
