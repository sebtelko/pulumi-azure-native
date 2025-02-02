// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Custom domain resource payload.
 */
export class CustomDomain extends pulumi.CustomResource {
    /**
     * Get an existing CustomDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomDomain {
        return new CustomDomain(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:appplatform/v20190501preview:CustomDomain';

    /**
     * Returns true if the given object is an instance of CustomDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomDomain.__pulumiType;
    }

    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties of the custom domain resource.
     */
    public readonly properties!: pulumi.Output<outputs.appplatform.v20190501preview.CustomDomainPropertiesResponse>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a CustomDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomDomainArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.appName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["appName"] = args ? args.appName : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:appplatform/v20190501preview:CustomDomain" }, { type: "azure-native:appplatform:CustomDomain" }, { type: "azure-nextgen:appplatform:CustomDomain" }, { type: "azure-native:appplatform/v20200701:CustomDomain" }, { type: "azure-nextgen:appplatform/v20200701:CustomDomain" }, { type: "azure-native:appplatform/v20201101preview:CustomDomain" }, { type: "azure-nextgen:appplatform/v20201101preview:CustomDomain" }, { type: "azure-native:appplatform/v20210303preview:CustomDomain" }, { type: "azure-nextgen:appplatform/v20210303preview:CustomDomain" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(CustomDomain.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomDomain resource.
 */
export interface CustomDomainArgs {
    /**
     * The name of the App resource.
     */
    readonly appName: pulumi.Input<string>;
    /**
     * The name of the custom domain resource.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * Properties of the custom domain resource.
     */
    readonly properties?: pulumi.Input<inputs.appplatform.v20190501preview.CustomDomainPropertiesArgs>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Service resource.
     */
    readonly serviceName: pulumi.Input<string>;
}
