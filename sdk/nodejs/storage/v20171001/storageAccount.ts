// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The storage account.
 */
export class StorageAccount extends pulumi.CustomResource {
    /**
     * Get an existing StorageAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StorageAccount {
        return new StorageAccount(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:storage/v20171001:StorageAccount';

    /**
     * Returns true if the given object is an instance of StorageAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageAccount.__pulumiType;
    }

    /**
     * Required for storage accounts where kind = BlobStorage. The access tier used for billing.
     */
    public readonly accessTier!: pulumi.Output<string>;
    /**
     * Gets the creation date and time of the storage account in UTC.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Gets the custom domain the user assigned to this storage account.
     */
    public readonly customDomain!: pulumi.Output<outputs.storage.v20171001.CustomDomainResponse>;
    /**
     * Allows https traffic only to storage service if sets to true.
     */
    public readonly enableHttpsTrafficOnly!: pulumi.Output<boolean | undefined>;
    /**
     * Gets the encryption settings on the account. If unspecified, the account is unencrypted.
     */
    public readonly encryption!: pulumi.Output<outputs.storage.v20171001.EncryptionResponse>;
    /**
     * The identity of the resource.
     */
    public readonly identity!: pulumi.Output<outputs.storage.v20171001.IdentityResponse | undefined>;
    /**
     * Gets the Kind.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
     */
    public /*out*/ readonly lastGeoFailoverTime!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Network rule set
     */
    public readonly networkRuleSet!: pulumi.Output<outputs.storage.v20171001.NetworkRuleSetResponse>;
    /**
     * Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
     */
    public /*out*/ readonly primaryEndpoints!: pulumi.Output<outputs.storage.v20171001.EndpointsResponse>;
    /**
     * Gets the location of the primary data center for the storage account.
     */
    public /*out*/ readonly primaryLocation!: pulumi.Output<string>;
    /**
     * Gets the status of the storage account at the time the operation was called.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
     */
    public /*out*/ readonly secondaryEndpoints!: pulumi.Output<outputs.storage.v20171001.EndpointsResponse>;
    /**
     * Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
     */
    public /*out*/ readonly secondaryLocation!: pulumi.Output<string>;
    /**
     * Gets the SKU.
     */
    public readonly sku!: pulumi.Output<outputs.storage.v20171001.SkuResponse>;
    /**
     * Gets the status indicating whether the primary location of the storage account is available or unavailable.
     */
    public /*out*/ readonly statusOfPrimary!: pulumi.Output<string>;
    /**
     * Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
     */
    public /*out*/ readonly statusOfSecondary!: pulumi.Output<string>;
    /**
     * Tags assigned to a resource; can be used for viewing and grouping a resource (across resource groups).
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a StorageAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageAccountArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sku === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["accessTier"] = args ? args.accessTier : undefined;
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["customDomain"] = args ? args.customDomain : undefined;
            inputs["enableHttpsTrafficOnly"] = (args ? args.enableHttpsTrafficOnly : undefined) ?? false;
            inputs["encryption"] = args ? args.encryption : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkRuleSet"] = args ? args.networkRuleSet : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["lastGeoFailoverTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["primaryEndpoints"] = undefined /*out*/;
            inputs["primaryLocation"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["secondaryEndpoints"] = undefined /*out*/;
            inputs["secondaryLocation"] = undefined /*out*/;
            inputs["statusOfPrimary"] = undefined /*out*/;
            inputs["statusOfSecondary"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["accessTier"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["customDomain"] = undefined /*out*/;
            inputs["enableHttpsTrafficOnly"] = undefined /*out*/;
            inputs["encryption"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["lastGeoFailoverTime"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkRuleSet"] = undefined /*out*/;
            inputs["primaryEndpoints"] = undefined /*out*/;
            inputs["primaryLocation"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["secondaryEndpoints"] = undefined /*out*/;
            inputs["secondaryLocation"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["statusOfPrimary"] = undefined /*out*/;
            inputs["statusOfSecondary"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:storage/v20171001:StorageAccount" }, { type: "azure-native:storage:StorageAccount" }, { type: "azure-nextgen:storage:StorageAccount" }, { type: "azure-native:storage/v20150501preview:StorageAccount" }, { type: "azure-nextgen:storage/v20150501preview:StorageAccount" }, { type: "azure-native:storage/v20150615:StorageAccount" }, { type: "azure-nextgen:storage/v20150615:StorageAccount" }, { type: "azure-native:storage/v20160101:StorageAccount" }, { type: "azure-nextgen:storage/v20160101:StorageAccount" }, { type: "azure-native:storage/v20160501:StorageAccount" }, { type: "azure-nextgen:storage/v20160501:StorageAccount" }, { type: "azure-native:storage/v20161201:StorageAccount" }, { type: "azure-nextgen:storage/v20161201:StorageAccount" }, { type: "azure-native:storage/v20170601:StorageAccount" }, { type: "azure-nextgen:storage/v20170601:StorageAccount" }, { type: "azure-native:storage/v20180201:StorageAccount" }, { type: "azure-nextgen:storage/v20180201:StorageAccount" }, { type: "azure-native:storage/v20180301preview:StorageAccount" }, { type: "azure-nextgen:storage/v20180301preview:StorageAccount" }, { type: "azure-native:storage/v20180701:StorageAccount" }, { type: "azure-nextgen:storage/v20180701:StorageAccount" }, { type: "azure-native:storage/v20181101:StorageAccount" }, { type: "azure-nextgen:storage/v20181101:StorageAccount" }, { type: "azure-native:storage/v20190401:StorageAccount" }, { type: "azure-nextgen:storage/v20190401:StorageAccount" }, { type: "azure-native:storage/v20190601:StorageAccount" }, { type: "azure-nextgen:storage/v20190601:StorageAccount" }, { type: "azure-native:storage/v20200801preview:StorageAccount" }, { type: "azure-nextgen:storage/v20200801preview:StorageAccount" }, { type: "azure-native:storage/v20210101:StorageAccount" }, { type: "azure-nextgen:storage/v20210101:StorageAccount" }, { type: "azure-native:storage/v20210201:StorageAccount" }, { type: "azure-nextgen:storage/v20210201:StorageAccount" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(StorageAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StorageAccount resource.
 */
export interface StorageAccountArgs {
    /**
     * Required for storage accounts where kind = BlobStorage. The access tier used for billing.
     */
    readonly accessTier?: pulumi.Input<enums.storage.v20171001.AccessTier>;
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    readonly accountName?: pulumi.Input<string>;
    /**
     * User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name property.
     */
    readonly customDomain?: pulumi.Input<inputs.storage.v20171001.CustomDomainArgs>;
    /**
     * Allows https traffic only to storage service if sets to true.
     */
    readonly enableHttpsTrafficOnly?: pulumi.Input<boolean>;
    /**
     * Provides the encryption settings on the account. If left unspecified the account encryption settings will remain the same. The default setting is unencrypted.
     */
    readonly encryption?: pulumi.Input<inputs.storage.v20171001.EncryptionArgs>;
    /**
     * The identity of the resource.
     */
    readonly identity?: pulumi.Input<inputs.storage.v20171001.IdentityArgs>;
    /**
     * Required. Indicates the type of storage account.
     */
    readonly kind: pulumi.Input<enums.storage.v20171001.Kind>;
    /**
     * Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Network rule set
     */
    readonly networkRuleSet?: pulumi.Input<inputs.storage.v20171001.NetworkRuleSetArgs>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Required. Gets or sets the sku name.
     */
    readonly sku: pulumi.Input<inputs.storage.v20171001.SkuArgs>;
    /**
     * Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
