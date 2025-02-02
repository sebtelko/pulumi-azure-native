// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Cluster details.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:azurestackhci/v20200301preview:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * App id of cluster AAD identity.
     */
    public readonly aadClientId!: pulumi.Output<string>;
    /**
     * Tenant id of cluster AAD identity.
     */
    public readonly aadTenantId!: pulumi.Output<string>;
    /**
     * Type of billing applied to the resource.
     */
    public /*out*/ readonly billingModel!: pulumi.Output<string>;
    /**
     * Unique, immutable resource id.
     */
    public /*out*/ readonly cloudId!: pulumi.Output<string>;
    /**
     * Most recent billing meter timestamp.
     */
    public /*out*/ readonly lastBillingTimestamp!: pulumi.Output<string>;
    /**
     * Most recent cluster sync timestamp.
     */
    public /*out*/ readonly lastSyncTimestamp!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * First cluster sync timestamp.
     */
    public /*out*/ readonly registrationTimestamp!: pulumi.Output<string>;
    /**
     * Properties reported by cluster agent.
     */
    public /*out*/ readonly reportedProperties!: pulumi.Output<outputs.azurestackhci.v20200301preview.ClusterReportedPropertiesResponse | undefined>;
    /**
     * Status of the cluster agent.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Number of days remaining in the trial period.
     */
    public /*out*/ readonly trialDaysRemaining!: pulumi.Output<number>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.aadClientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aadClientId'");
            }
            if ((!args || args.aadTenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aadTenantId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["aadClientId"] = args ? args.aadClientId : undefined;
            inputs["aadTenantId"] = args ? args.aadTenantId : undefined;
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["billingModel"] = undefined /*out*/;
            inputs["cloudId"] = undefined /*out*/;
            inputs["lastBillingTimestamp"] = undefined /*out*/;
            inputs["lastSyncTimestamp"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["registrationTimestamp"] = undefined /*out*/;
            inputs["reportedProperties"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["trialDaysRemaining"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["aadClientId"] = undefined /*out*/;
            inputs["aadTenantId"] = undefined /*out*/;
            inputs["billingModel"] = undefined /*out*/;
            inputs["cloudId"] = undefined /*out*/;
            inputs["lastBillingTimestamp"] = undefined /*out*/;
            inputs["lastSyncTimestamp"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["registrationTimestamp"] = undefined /*out*/;
            inputs["reportedProperties"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["trialDaysRemaining"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:azurestackhci/v20200301preview:Cluster" }, { type: "azure-native:azurestackhci:Cluster" }, { type: "azure-nextgen:azurestackhci:Cluster" }, { type: "azure-native:azurestackhci/v20201001:Cluster" }, { type: "azure-nextgen:azurestackhci/v20201001:Cluster" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * App id of cluster AAD identity.
     */
    readonly aadClientId: pulumi.Input<string>;
    /**
     * Tenant id of cluster AAD identity.
     */
    readonly aadTenantId: pulumi.Input<string>;
    /**
     * The name of the cluster.
     */
    readonly clusterName?: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
