// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Description of a Namespace resource.
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:notificationhubs/v20170401:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * The time the namespace was created.
     */
    public readonly createdAt!: pulumi.Output<string | undefined>;
    /**
     * Whether or not the namespace is set as Critical.
     */
    public readonly critical!: pulumi.Output<boolean | undefined>;
    /**
     * Data center for the namespace
     */
    public readonly dataCenter!: pulumi.Output<string | undefined>;
    /**
     * Whether or not the namespace is currently enabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Identifier for Azure Insights metrics
     */
    public /*out*/ readonly metricId!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace type.
     */
    public readonly namespaceType!: pulumi.Output<string | undefined>;
    /**
     * Provisioning state of the Namespace.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Specifies the targeted region in which the namespace should be created. It can be any of the following values: Australia East, Australia Southeast, Central US, East US, East US 2, West US, North Central US, South Central US, East Asia, Southeast Asia, Brazil South, Japan East, Japan West, North Europe, West Europe
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * ScaleUnit where the namespace gets created
     */
    public readonly scaleUnit!: pulumi.Output<string | undefined>;
    /**
     * Endpoint you can use to perform NotificationHub operations.
     */
    public readonly serviceBusEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The sku of the created namespace
     */
    public readonly sku!: pulumi.Output<outputs.notificationhubs.v20170401.SkuResponse | undefined>;
    /**
     * Status of the namespace. It can be any of these values:1 = Created/Active2 = Creating3 = Suspended4 = Deleting
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The Id of the Azure subscription associated with the namespace.
     */
    public readonly subscriptionId!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The time the namespace was updated.
     */
    public readonly updatedAt!: pulumi.Output<string | undefined>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["createdAt"] = args ? args.createdAt : undefined;
            inputs["critical"] = args ? args.critical : undefined;
            inputs["dataCenter"] = args ? args.dataCenter : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["namespaceType"] = args ? args.namespaceType : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scaleUnit"] = args ? args.scaleUnit : undefined;
            inputs["serviceBusEndpoint"] = args ? args.serviceBusEndpoint : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["subscriptionId"] = args ? args.subscriptionId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["updatedAt"] = args ? args.updatedAt : undefined;
            inputs["metricId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["createdAt"] = undefined /*out*/;
            inputs["critical"] = undefined /*out*/;
            inputs["dataCenter"] = undefined /*out*/;
            inputs["enabled"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["metricId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["namespaceType"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["scaleUnit"] = undefined /*out*/;
            inputs["serviceBusEndpoint"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["subscriptionId"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:notificationhubs/v20170401:Namespace" }, { type: "azure-native:notificationhubs:Namespace" }, { type: "azure-nextgen:notificationhubs:Namespace" }, { type: "azure-native:notificationhubs/v20140901:Namespace" }, { type: "azure-nextgen:notificationhubs/v20140901:Namespace" }, { type: "azure-native:notificationhubs/v20160301:Namespace" }, { type: "azure-nextgen:notificationhubs/v20160301:Namespace" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Namespace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * The time the namespace was created.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * Whether or not the namespace is set as Critical.
     */
    readonly critical?: pulumi.Input<boolean>;
    /**
     * Data center for the namespace
     */
    readonly dataCenter?: pulumi.Input<string>;
    /**
     * Whether or not the namespace is currently enabled.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the namespace.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The namespace name.
     */
    readonly namespaceName?: pulumi.Input<string>;
    /**
     * The namespace type.
     */
    readonly namespaceType?: pulumi.Input<enums.notificationhubs.v20170401.NamespaceType>;
    /**
     * Provisioning state of the Namespace.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * Specifies the targeted region in which the namespace should be created. It can be any of the following values: Australia East, Australia Southeast, Central US, East US, East US 2, West US, North Central US, South Central US, East Asia, Southeast Asia, Brazil South, Japan East, Japan West, North Europe, West Europe
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * ScaleUnit where the namespace gets created
     */
    readonly scaleUnit?: pulumi.Input<string>;
    /**
     * Endpoint you can use to perform NotificationHub operations.
     */
    readonly serviceBusEndpoint?: pulumi.Input<string>;
    /**
     * The sku of the created namespace
     */
    readonly sku?: pulumi.Input<inputs.notificationhubs.v20170401.SkuArgs>;
    /**
     * Status of the namespace. It can be any of these values:1 = Created/Active2 = Creating3 = Suspended4 = Deleting
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The Id of the Azure subscription associated with the namespace.
     */
    readonly subscriptionId?: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The time the namespace was updated.
     */
    readonly updatedAt?: pulumi.Input<string>;
}
