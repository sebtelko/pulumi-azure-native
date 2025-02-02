// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource information with extended details.
 * API Version: 2019-09-01.
 */
export class Vault extends pulumi.CustomResource {
    /**
     * Get an existing Vault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Vault {
        return new Vault(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:keyvault:Vault';

    /**
     * Returns true if the given object is an instance of Vault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vault.__pulumiType;
    }

    /**
     * Azure location of the key vault resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Name of the key vault resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties of the vault
     */
    public readonly properties!: pulumi.Output<outputs.keyvault.VaultPropertiesResponse>;
    /**
     * Tags assigned to the key vault resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type of the key vault resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Vault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VaultArgs, opts?: pulumi.CustomResourceOptions) {
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
            inputs["vaultName"] = args ? args.vaultName : undefined;
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
        const aliasOpts = { aliases: [{ type: "azure-nextgen:keyvault:Vault" }, { type: "azure-native:keyvault/v20150601:Vault" }, { type: "azure-nextgen:keyvault/v20150601:Vault" }, { type: "azure-native:keyvault/v20161001:Vault" }, { type: "azure-nextgen:keyvault/v20161001:Vault" }, { type: "azure-native:keyvault/v20180214:Vault" }, { type: "azure-nextgen:keyvault/v20180214:Vault" }, { type: "azure-native:keyvault/v20180214preview:Vault" }, { type: "azure-nextgen:keyvault/v20180214preview:Vault" }, { type: "azure-native:keyvault/v20190901:Vault" }, { type: "azure-nextgen:keyvault/v20190901:Vault" }, { type: "azure-native:keyvault/v20200401preview:Vault" }, { type: "azure-nextgen:keyvault/v20200401preview:Vault" }, { type: "azure-native:keyvault/v20210401preview:Vault" }, { type: "azure-nextgen:keyvault/v20210401preview:Vault" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Vault.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Vault resource.
 */
export interface VaultArgs {
    /**
     * The supported Azure location where the key vault should be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Properties of the vault
     */
    readonly properties: pulumi.Input<inputs.keyvault.VaultPropertiesArgs>;
    /**
     * The name of the Resource Group to which the server belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The tags that will be assigned to the key vault.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the vault
     */
    readonly vaultName?: pulumi.Input<string>;
}
