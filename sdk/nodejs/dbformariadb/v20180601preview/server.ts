// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Represents a server.
 */
export class Server extends pulumi.CustomResource {
    /**
     * Get an existing Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Server {
        return new Server(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:dbformariadb/v20180601preview:Server';

    /**
     * Returns true if the given object is an instance of Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Server.__pulumiType;
    }

    /**
     * The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
     */
    public /*out*/ readonly administratorLogin!: pulumi.Output<string | undefined>;
    /**
     * Earliest restore point creation time (ISO8601 format)
     */
    public /*out*/ readonly earliestRestoreDate!: pulumi.Output<string | undefined>;
    /**
     * The fully qualified domain name of a server.
     */
    public /*out*/ readonly fullyQualifiedDomainName!: pulumi.Output<string | undefined>;
    /**
     * The Azure Active Directory identity of the server.
     */
    public /*out*/ readonly identity!: pulumi.Output<outputs.dbformariadb.v20180601preview.ResourceIdentityResponse | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The master server id of a replica server.
     */
    public /*out*/ readonly masterServerId!: pulumi.Output<string | undefined>;
    /**
     * Enforce a minimal Tls version for the server.
     */
    public /*out*/ readonly minimalTlsVersion!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The maximum number of replicas that a master server can have.
     */
    public /*out*/ readonly replicaCapacity!: pulumi.Output<number | undefined>;
    /**
     * The replication role of the server.
     */
    public /*out*/ readonly replicationRole!: pulumi.Output<string | undefined>;
    /**
     * The SKU (pricing tier) of the server.
     */
    public readonly sku!: pulumi.Output<outputs.dbformariadb.v20180601preview.SkuResponse | undefined>;
    /**
     * Enable ssl enforcement or not when connect to server.
     */
    public /*out*/ readonly sslEnforcement!: pulumi.Output<string | undefined>;
    /**
     * Storage profile of a server.
     */
    public /*out*/ readonly storageProfile!: pulumi.Output<outputs.dbformariadb.v20180601preview.StorageProfileResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A state of a server that is visible to user.
     */
    public /*out*/ readonly userVisibleState!: pulumi.Output<string | undefined>;
    /**
     * Server version.
     */
    public /*out*/ readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerArgs, opts?: pulumi.CustomResourceOptions) {
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
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["administratorLogin"] = undefined /*out*/;
            inputs["earliestRestoreDate"] = undefined /*out*/;
            inputs["fullyQualifiedDomainName"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["masterServerId"] = undefined /*out*/;
            inputs["minimalTlsVersion"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["replicaCapacity"] = undefined /*out*/;
            inputs["replicationRole"] = undefined /*out*/;
            inputs["sslEnforcement"] = undefined /*out*/;
            inputs["storageProfile"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["userVisibleState"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        } else {
            inputs["administratorLogin"] = undefined /*out*/;
            inputs["earliestRestoreDate"] = undefined /*out*/;
            inputs["fullyQualifiedDomainName"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["masterServerId"] = undefined /*out*/;
            inputs["minimalTlsVersion"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["replicaCapacity"] = undefined /*out*/;
            inputs["replicationRole"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["sslEnforcement"] = undefined /*out*/;
            inputs["storageProfile"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["userVisibleState"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:dbformariadb/v20180601preview:Server" }, { type: "azure-native:dbformariadb:Server" }, { type: "azure-nextgen:dbformariadb:Server" }, { type: "azure-native:dbformariadb/v20180601:Server" }, { type: "azure-nextgen:dbformariadb/v20180601:Server" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Server.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * The location the resource resides in.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Properties of the server.
     */
    readonly properties: pulumi.Input<inputs.dbformariadb.v20180601preview.ServerPropertiesForDefaultCreateArgs | inputs.dbformariadb.v20180601preview.ServerPropertiesForGeoRestoreArgs | inputs.dbformariadb.v20180601preview.ServerPropertiesForReplicaArgs | inputs.dbformariadb.v20180601preview.ServerPropertiesForRestoreArgs>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName?: pulumi.Input<string>;
    /**
     * The SKU (pricing tier) of the server.
     */
    readonly sku?: pulumi.Input<inputs.dbformariadb.v20180601preview.SkuArgs>;
    /**
     * Application-specific metadata in the form of key-value pairs.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
