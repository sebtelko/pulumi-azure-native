// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * API Version: 2019-01-01-preview.
 */
export class B2CTenant extends pulumi.CustomResource {
    /**
     * Get an existing B2CTenant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): B2CTenant {
        return new B2CTenant(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:azureactivedirectory:B2CTenant';

    /**
     * Returns true if the given object is an instance of B2CTenant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is B2CTenant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === B2CTenant.__pulumiType;
    }

    /**
     * The billing configuration for the tenant.
     */
    public /*out*/ readonly billingConfig!: pulumi.Output<outputs.azureactivedirectory.B2CTenantResourcePropertiesResponseBillingConfig | undefined>;
    /**
     * The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia' (preview). Refer to [this documentation](https://aka.ms/B2CDataResidency) for more information.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the B2C tenant resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * SKU properties of the Azure AD B2C tenant. Learn more about Azure AD B2C billing at [aka.ms/b2cBilling](https://aka.ms/b2cBilling).
     */
    public readonly sku!: pulumi.Output<outputs.azureactivedirectory.B2CResourceSKUResponse>;
    /**
     * Resource Tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * An identifier of the B2C tenant.
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * The type of the B2C tenant resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a B2CTenant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: B2CTenantArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.properties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'properties'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sku === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["billingConfig"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["billingConfig"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:azureactivedirectory:B2CTenant" }, { type: "azure-native:azureactivedirectory/v20190101preview:B2CTenant" }, { type: "azure-nextgen:azureactivedirectory/v20190101preview:B2CTenant" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(B2CTenant.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a B2CTenant resource.
 */
export interface B2CTenantArgs {
    /**
     * The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia' (preview). Refer to [this documentation](https://aka.ms/B2CDataResidency) for more information.
     */
    readonly location?: pulumi.Input<string>;
    readonly properties: pulumi.Input<inputs.azureactivedirectory.CreateTenantRequestBodyPropertiesArgs>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The initial domain name of the B2C tenant.
     */
    readonly resourceName?: pulumi.Input<string>;
    /**
     * SKU properties of the Azure AD B2C tenant. Learn more about Azure AD B2C billing at [aka.ms/b2cBilling](https://aka.ms/b2cBilling).
     */
    readonly sku: pulumi.Input<inputs.azureactivedirectory.B2CResourceSKUArgs>;
    /**
     * Resource Tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
