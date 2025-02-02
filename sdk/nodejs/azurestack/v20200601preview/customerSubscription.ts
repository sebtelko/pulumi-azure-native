// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Customer subscription.
 */
export class CustomerSubscription extends pulumi.CustomResource {
    /**
     * Get an existing CustomerSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomerSubscription {
        return new CustomerSubscription(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:azurestack/v20200601preview:CustomerSubscription';

    /**
     * Returns true if the given object is an instance of CustomerSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomerSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomerSubscription.__pulumiType;
    }

    /**
     * The entity tag used for optimistic concurrency when modifying the resource.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.azurestack.v20200601preview.SystemDataResponse>;
    /**
     * Tenant Id.
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * Type of Resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a CustomerSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomerSubscriptionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.registrationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registrationName'");
            }
            if ((!args || args.resourceGroup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroup'");
            }
            inputs["customerSubscriptionName"] = args ? args.customerSubscriptionName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["registrationName"] = args ? args.registrationName : undefined;
            inputs["resourceGroup"] = args ? args.resourceGroup : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:azurestack/v20200601preview:CustomerSubscription" }, { type: "azure-native:azurestack:CustomerSubscription" }, { type: "azure-nextgen:azurestack:CustomerSubscription" }, { type: "azure-native:azurestack/v20170601:CustomerSubscription" }, { type: "azure-nextgen:azurestack/v20170601:CustomerSubscription" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(CustomerSubscription.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomerSubscription resource.
 */
export interface CustomerSubscriptionArgs {
    /**
     * Name of the product.
     */
    readonly customerSubscriptionName?: pulumi.Input<string>;
    /**
     * The entity tag used for optimistic concurrency when modifying the resource.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Name of the Azure Stack registration.
     */
    readonly registrationName: pulumi.Input<string>;
    /**
     * Name of the resource group.
     */
    readonly resourceGroup: pulumi.Input<string>;
    /**
     * Tenant Id.
     */
    readonly tenantId?: pulumi.Input<string>;
}
