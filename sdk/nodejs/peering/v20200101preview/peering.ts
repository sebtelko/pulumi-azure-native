// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Peering is a logical representation of a set of connections to the Microsoft Cloud Edge at a location.
 */
export class Peering extends pulumi.CustomResource {
    /**
     * Get an existing Peering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Peering {
        return new Peering(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:peering/v20200101preview:Peering';

    /**
     * Returns true if the given object is an instance of Peering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Peering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Peering.__pulumiType;
    }

    /**
     * The properties that define a direct peering.
     */
    public readonly direct!: pulumi.Output<outputs.peering.v20200101preview.PeeringPropertiesDirectResponse | undefined>;
    /**
     * The properties that define an exchange peering.
     */
    public readonly exchange!: pulumi.Output<outputs.peering.v20200101preview.PeeringPropertiesExchangeResponse | undefined>;
    /**
     * The kind of the peering.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The location of the peering.
     */
    public readonly peeringLocation!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The SKU that defines the tier and kind of the peering.
     */
    public readonly sku!: pulumi.Output<outputs.peering.v20200101preview.PeeringSkuResponse>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Peering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PeeringArgs, opts?: pulumi.CustomResourceOptions) {
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
            inputs["direct"] = args ? args.direct : undefined;
            inputs["exchange"] = args ? args.exchange : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["peeringLocation"] = args ? args.peeringLocation : undefined;
            inputs["peeringName"] = args ? args.peeringName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["direct"] = undefined /*out*/;
            inputs["exchange"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["peeringLocation"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:peering/v20200101preview:Peering" }, { type: "azure-native:peering:Peering" }, { type: "azure-nextgen:peering:Peering" }, { type: "azure-native:peering/v20190801preview:Peering" }, { type: "azure-nextgen:peering/v20190801preview:Peering" }, { type: "azure-native:peering/v20190901preview:Peering" }, { type: "azure-nextgen:peering/v20190901preview:Peering" }, { type: "azure-native:peering/v20200401:Peering" }, { type: "azure-nextgen:peering/v20200401:Peering" }, { type: "azure-native:peering/v20201001:Peering" }, { type: "azure-nextgen:peering/v20201001:Peering" }, { type: "azure-native:peering/v20210101:Peering" }, { type: "azure-nextgen:peering/v20210101:Peering" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Peering.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Peering resource.
 */
export interface PeeringArgs {
    /**
     * The properties that define a direct peering.
     */
    readonly direct?: pulumi.Input<inputs.peering.v20200101preview.PeeringPropertiesDirectArgs>;
    /**
     * The properties that define an exchange peering.
     */
    readonly exchange?: pulumi.Input<inputs.peering.v20200101preview.PeeringPropertiesExchangeArgs>;
    /**
     * The kind of the peering.
     */
    readonly kind: pulumi.Input<string | enums.peering.v20200101preview.Kind>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The location of the peering.
     */
    readonly peeringLocation?: pulumi.Input<string>;
    /**
     * The name of the peering.
     */
    readonly peeringName?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU that defines the tier and kind of the peering.
     */
    readonly sku: pulumi.Input<inputs.peering.v20200101preview.PeeringSkuArgs>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
