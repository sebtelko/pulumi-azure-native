// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Dedicated cloud node model
 */
export class DedicatedCloudNode extends pulumi.CustomResource {
    /**
     * Get an existing DedicatedCloudNode resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DedicatedCloudNode {
        return new DedicatedCloudNode(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:vmwarecloudsimple/v20190401:DedicatedCloudNode';

    /**
     * Returns true if the given object is an instance of DedicatedCloudNode.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DedicatedCloudNode {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DedicatedCloudNode.__pulumiType;
    }

    /**
     * Availability Zone id, e.g. "az1"
     */
    public readonly availabilityZoneId!: pulumi.Output<string>;
    /**
     * Availability Zone name, e.g. "Availability Zone 1"
     */
    public /*out*/ readonly availabilityZoneName!: pulumi.Output<string>;
    /**
     * VMWare Cloud Rack Name
     */
    public /*out*/ readonly cloudRackName!: pulumi.Output<string>;
    /**
     * date time the resource was created
     */
    public /*out*/ readonly created!: pulumi.Output<any>;
    /**
     * Azure region
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * SKU's name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * count of nodes to create
     */
    public readonly nodesCount!: pulumi.Output<number>;
    /**
     * Placement Group id, e.g. "n1"
     */
    public readonly placementGroupId!: pulumi.Output<string>;
    /**
     * Placement Name, e.g. "Placement Group 1"
     */
    public /*out*/ readonly placementGroupName!: pulumi.Output<string>;
    /**
     * Private Cloud Id
     */
    public /*out*/ readonly privateCloudId!: pulumi.Output<string>;
    /**
     * Resource Pool Name
     */
    public /*out*/ readonly privateCloudName!: pulumi.Output<string>;
    /**
     * The provisioning status of the resource
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * purchase id
     */
    public readonly purchaseId!: pulumi.Output<string>;
    /**
     * Dedicated Cloud Nodes SKU
     */
    public readonly sku!: pulumi.Output<outputs.vmwarecloudsimple.v20190401.SkuResponse | undefined>;
    /**
     * Node status, indicates is private cloud set up on this node or not
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Dedicated Cloud Nodes tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * {resourceProviderNamespace}/{resourceType}
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * VMWare Cluster Name
     */
    public /*out*/ readonly vmwareClusterName!: pulumi.Output<string>;

    /**
     * Create a DedicatedCloudNode resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DedicatedCloudNodeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.availabilityZoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availabilityZoneId'");
            }
            if ((!args || args.id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'id'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.nodesCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodesCount'");
            }
            if ((!args || args.placementGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'placementGroupId'");
            }
            if ((!args || args.purchaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'purchaseId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["availabilityZoneId"] = args ? args.availabilityZoneId : undefined;
            inputs["dedicatedCloudNodeName"] = args ? args.dedicatedCloudNodeName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodesCount"] = args ? args.nodesCount : undefined;
            inputs["placementGroupId"] = args ? args.placementGroupId : undefined;
            inputs["purchaseId"] = args ? args.purchaseId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["availabilityZoneName"] = undefined /*out*/;
            inputs["cloudRackName"] = undefined /*out*/;
            inputs["created"] = undefined /*out*/;
            inputs["placementGroupName"] = undefined /*out*/;
            inputs["privateCloudId"] = undefined /*out*/;
            inputs["privateCloudName"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vmwareClusterName"] = undefined /*out*/;
        } else {
            inputs["availabilityZoneId"] = undefined /*out*/;
            inputs["availabilityZoneName"] = undefined /*out*/;
            inputs["cloudRackName"] = undefined /*out*/;
            inputs["created"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nodesCount"] = undefined /*out*/;
            inputs["placementGroupId"] = undefined /*out*/;
            inputs["placementGroupName"] = undefined /*out*/;
            inputs["privateCloudId"] = undefined /*out*/;
            inputs["privateCloudName"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["purchaseId"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vmwareClusterName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:vmwarecloudsimple/v20190401:DedicatedCloudNode" }, { type: "azure-native:vmwarecloudsimple:DedicatedCloudNode" }, { type: "azure-nextgen:vmwarecloudsimple:DedicatedCloudNode" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DedicatedCloudNode.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DedicatedCloudNode resource.
 */
export interface DedicatedCloudNodeArgs {
    /**
     * Availability Zone id, e.g. "az1"
     */
    readonly availabilityZoneId: pulumi.Input<string>;
    /**
     * dedicated cloud node name
     */
    readonly dedicatedCloudNodeName?: pulumi.Input<string>;
    /**
     * SKU's id
     */
    readonly id: pulumi.Input<string>;
    /**
     * Azure region
     */
    readonly location?: pulumi.Input<string>;
    /**
     * SKU's name
     */
    readonly name: pulumi.Input<string>;
    /**
     * count of nodes to create
     */
    readonly nodesCount: pulumi.Input<number>;
    /**
     * Placement Group id, e.g. "n1"
     */
    readonly placementGroupId: pulumi.Input<string>;
    /**
     * purchase id
     */
    readonly purchaseId: pulumi.Input<string>;
    /**
     * The name of the resource group
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Dedicated Cloud Nodes SKU
     */
    readonly sku?: pulumi.Input<inputs.vmwarecloudsimple.v20190401.SkuArgs>;
    /**
     * Dedicated Cloud Nodes tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
