// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * An elastic pool.
 */
export class ElasticPool extends pulumi.CustomResource {
    /**
     * Get an existing ElasticPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ElasticPool {
        return new ElasticPool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:sql/v20171001preview:ElasticPool';

    /**
     * Returns true if the given object is an instance of ElasticPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ElasticPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ElasticPool.__pulumiType;
    }

    /**
     * The creation date of the elastic pool (ISO8601 format).
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * Kind of elastic pool. This is metadata used for the Azure portal experience.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * The license type to apply for this elastic pool.
     */
    public readonly licenseType!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The storage limit for the database elastic pool in bytes.
     */
    public readonly maxSizeBytes!: pulumi.Output<number | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The per database settings for the elastic pool.
     */
    public readonly perDatabaseSettings!: pulumi.Output<outputs.sql.v20171001preview.ElasticPoolPerDatabaseSettingsResponse | undefined>;
    /**
     * The elastic pool SKU.
     * 
     * The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
     * 
     * ```azurecli
     * az sql elastic-pool list-editions -l <location> -o table
     * ````
     */
    public readonly sku!: pulumi.Output<outputs.sql.v20171001preview.SkuResponse | undefined>;
    /**
     * The state of the elastic pool.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
     */
    public readonly zoneRedundant!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ElasticPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ElasticPoolArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serverName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverName'");
            }
            inputs["elasticPoolName"] = args ? args.elasticPoolName : undefined;
            inputs["licenseType"] = args ? args.licenseType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maxSizeBytes"] = args ? args.maxSizeBytes : undefined;
            inputs["perDatabaseSettings"] = args ? args.perDatabaseSettings : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zoneRedundant"] = args ? args.zoneRedundant : undefined;
            inputs["creationDate"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["creationDate"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["licenseType"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["maxSizeBytes"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["perDatabaseSettings"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["zoneRedundant"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:sql/v20171001preview:ElasticPool" }, { type: "azure-native:sql:ElasticPool" }, { type: "azure-nextgen:sql:ElasticPool" }, { type: "azure-native:sql/v20140401:ElasticPool" }, { type: "azure-nextgen:sql/v20140401:ElasticPool" }, { type: "azure-native:sql/v20200202preview:ElasticPool" }, { type: "azure-nextgen:sql/v20200202preview:ElasticPool" }, { type: "azure-native:sql/v20200801preview:ElasticPool" }, { type: "azure-nextgen:sql/v20200801preview:ElasticPool" }, { type: "azure-native:sql/v20201101preview:ElasticPool" }, { type: "azure-nextgen:sql/v20201101preview:ElasticPool" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ElasticPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ElasticPool resource.
 */
export interface ElasticPoolArgs {
    /**
     * The name of the elastic pool.
     */
    readonly elasticPoolName?: pulumi.Input<string>;
    /**
     * The license type to apply for this elastic pool.
     */
    readonly licenseType?: pulumi.Input<string | enums.sql.v20171001preview.ElasticPoolLicenseType>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The storage limit for the database elastic pool in bytes.
     */
    readonly maxSizeBytes?: pulumi.Input<number>;
    /**
     * The per database settings for the elastic pool.
     */
    readonly perDatabaseSettings?: pulumi.Input<inputs.sql.v20171001preview.ElasticPoolPerDatabaseSettingsArgs>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * The elastic pool SKU.
     * 
     * The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or the following command:
     * 
     * ```azurecli
     * az sql elastic-pool list-editions -l <location> -o table
     * ````
     */
    readonly sku?: pulumi.Input<inputs.sql.v20171001preview.SkuArgs>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether or not this elastic pool is zone redundant, which means the replicas of this elastic pool will be spread across multiple availability zones.
     */
    readonly zoneRedundant?: pulumi.Input<boolean>;
}
