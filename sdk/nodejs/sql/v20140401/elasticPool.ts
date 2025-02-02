// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Represents a database elastic pool.
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
    public static readonly __pulumiType = 'azure-native:sql/v20140401:ElasticPool';

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
     * The maximum DTU any one database can consume.
     */
    public readonly databaseDtuMax!: pulumi.Output<number | undefined>;
    /**
     * The minimum DTU all databases are guaranteed.
     */
    public readonly databaseDtuMin!: pulumi.Output<number | undefined>;
    /**
     * The total shared DTU for the database elastic pool.
     */
    public readonly dtu!: pulumi.Output<number | undefined>;
    /**
     * The edition of the elastic pool.
     */
    public readonly edition!: pulumi.Output<string | undefined>;
    /**
     * Kind of elastic pool.  This is metadata used for the Azure portal experience.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The state of the elastic pool.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Gets storage limit for the database elastic pool in MB.
     */
    public readonly storageMB!: pulumi.Output<number | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Whether or not this database elastic pool is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
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
            inputs["databaseDtuMax"] = args ? args.databaseDtuMax : undefined;
            inputs["databaseDtuMin"] = args ? args.databaseDtuMin : undefined;
            inputs["dtu"] = args ? args.dtu : undefined;
            inputs["edition"] = args ? args.edition : undefined;
            inputs["elasticPoolName"] = args ? args.elasticPoolName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["storageMB"] = args ? args.storageMB : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zoneRedundant"] = args ? args.zoneRedundant : undefined;
            inputs["creationDate"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["creationDate"] = undefined /*out*/;
            inputs["databaseDtuMax"] = undefined /*out*/;
            inputs["databaseDtuMin"] = undefined /*out*/;
            inputs["dtu"] = undefined /*out*/;
            inputs["edition"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["storageMB"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["zoneRedundant"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:sql/v20140401:ElasticPool" }, { type: "azure-native:sql:ElasticPool" }, { type: "azure-nextgen:sql:ElasticPool" }, { type: "azure-native:sql/v20171001preview:ElasticPool" }, { type: "azure-nextgen:sql/v20171001preview:ElasticPool" }, { type: "azure-native:sql/v20200202preview:ElasticPool" }, { type: "azure-nextgen:sql/v20200202preview:ElasticPool" }, { type: "azure-native:sql/v20200801preview:ElasticPool" }, { type: "azure-nextgen:sql/v20200801preview:ElasticPool" }, { type: "azure-native:sql/v20201101preview:ElasticPool" }, { type: "azure-nextgen:sql/v20201101preview:ElasticPool" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ElasticPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ElasticPool resource.
 */
export interface ElasticPoolArgs {
    /**
     * The maximum DTU any one database can consume.
     */
    readonly databaseDtuMax?: pulumi.Input<number>;
    /**
     * The minimum DTU all databases are guaranteed.
     */
    readonly databaseDtuMin?: pulumi.Input<number>;
    /**
     * The total shared DTU for the database elastic pool.
     */
    readonly dtu?: pulumi.Input<number>;
    /**
     * The edition of the elastic pool.
     */
    readonly edition?: pulumi.Input<string | enums.sql.v20140401.ElasticPoolEdition>;
    /**
     * The name of the elastic pool to be operated on (updated or created).
     */
    readonly elasticPoolName?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * Gets storage limit for the database elastic pool in MB.
     */
    readonly storageMB?: pulumi.Input<number>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether or not this database elastic pool is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
     */
    readonly zoneRedundant?: pulumi.Input<boolean>;
}
