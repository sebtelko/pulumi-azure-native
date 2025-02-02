// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Information about managed application.
 * API Version: 2019-07-01.
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:solutions:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * The fully qualified path of managed application definition Id.
     */
    public readonly applicationDefinitionId!: pulumi.Output<string | undefined>;
    /**
     * The collection of managed application artifacts.
     */
    public /*out*/ readonly artifacts!: pulumi.Output<outputs.solutions.ApplicationArtifactResponse[]>;
    /**
     * The  read-only authorizations property that is retrieved from the application package.
     */
    public /*out*/ readonly authorizations!: pulumi.Output<outputs.solutions.ApplicationAuthorizationResponse[]>;
    /**
     * The managed application billing details.
     */
    public /*out*/ readonly billingDetails!: pulumi.Output<outputs.solutions.ApplicationBillingDetailsDefinitionResponse>;
    /**
     * The client entity that created the JIT request.
     */
    public /*out*/ readonly createdBy!: pulumi.Output<outputs.solutions.ApplicationClientDetailsResponse>;
    /**
     * The read-only customer support property that is retrieved from the application package.
     */
    public /*out*/ readonly customerSupport!: pulumi.Output<outputs.solutions.ApplicationPackageContactResponse>;
    /**
     * The identity of the resource.
     */
    public readonly identity!: pulumi.Output<outputs.solutions.IdentityResponse | undefined>;
    /**
     * The managed application Jit access policy.
     */
    public readonly jitAccessPolicy!: pulumi.Output<outputs.solutions.ApplicationJitAccessPolicyResponse | undefined>;
    /**
     * The kind of the managed application. Allowed values are MarketPlace and ServiceCatalog.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * ID of the resource that manages this resource.
     */
    public readonly managedBy!: pulumi.Output<string | undefined>;
    /**
     * The managed resource group Id.
     */
    public readonly managedResourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The managed application management mode.
     */
    public /*out*/ readonly managementMode!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Name and value pairs that define the managed application outputs.
     */
    public /*out*/ readonly outputs!: pulumi.Output<any>;
    /**
     * Name and value pairs that define the managed application parameters. It can be a JObject or a well formed JSON string.
     */
    public readonly parameters!: pulumi.Output<any | undefined>;
    /**
     * The plan information.
     */
    public readonly plan!: pulumi.Output<outputs.solutions.PlanResponse | undefined>;
    /**
     * The managed application provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The publisher tenant Id.
     */
    public /*out*/ readonly publisherTenantId!: pulumi.Output<string>;
    /**
     * The SKU of the resource.
     */
    public readonly sku!: pulumi.Output<outputs.solutions.SkuResponse | undefined>;
    /**
     * The read-only support URLs property that is retrieved from the application package.
     */
    public /*out*/ readonly supportUrls!: pulumi.Output<outputs.solutions.ApplicationPackageSupportUrlsResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The client entity that last updated the JIT request.
     */
    public /*out*/ readonly updatedBy!: pulumi.Output<outputs.solutions.ApplicationClientDetailsResponse>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["applicationDefinitionId"] = args ? args.applicationDefinitionId : undefined;
            inputs["applicationName"] = args ? args.applicationName : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["jitAccessPolicy"] = args ? args.jitAccessPolicy : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedBy"] = args ? args.managedBy : undefined;
            inputs["managedResourceGroupId"] = args ? args.managedResourceGroupId : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["artifacts"] = undefined /*out*/;
            inputs["authorizations"] = undefined /*out*/;
            inputs["billingDetails"] = undefined /*out*/;
            inputs["createdBy"] = undefined /*out*/;
            inputs["customerSupport"] = undefined /*out*/;
            inputs["managementMode"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["outputs"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publisherTenantId"] = undefined /*out*/;
            inputs["supportUrls"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedBy"] = undefined /*out*/;
        } else {
            inputs["applicationDefinitionId"] = undefined /*out*/;
            inputs["artifacts"] = undefined /*out*/;
            inputs["authorizations"] = undefined /*out*/;
            inputs["billingDetails"] = undefined /*out*/;
            inputs["createdBy"] = undefined /*out*/;
            inputs["customerSupport"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["jitAccessPolicy"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["managedBy"] = undefined /*out*/;
            inputs["managedResourceGroupId"] = undefined /*out*/;
            inputs["managementMode"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["outputs"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["plan"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publisherTenantId"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["supportUrls"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedBy"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:solutions:Application" }, { type: "azure-native:solutions/v20160901preview:Application" }, { type: "azure-nextgen:solutions/v20160901preview:Application" }, { type: "azure-native:solutions/v20170901:Application" }, { type: "azure-nextgen:solutions/v20170901:Application" }, { type: "azure-native:solutions/v20180601:Application" }, { type: "azure-nextgen:solutions/v20180601:Application" }, { type: "azure-native:solutions/v20190701:Application" }, { type: "azure-nextgen:solutions/v20190701:Application" }, { type: "azure-native:solutions/v20200821preview:Application" }, { type: "azure-nextgen:solutions/v20200821preview:Application" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Application.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * The fully qualified path of managed application definition Id.
     */
    readonly applicationDefinitionId?: pulumi.Input<string>;
    /**
     * The name of the managed application.
     */
    readonly applicationName?: pulumi.Input<string>;
    /**
     * The identity of the resource.
     */
    readonly identity?: pulumi.Input<inputs.solutions.IdentityArgs>;
    /**
     * The managed application Jit access policy.
     */
    readonly jitAccessPolicy?: pulumi.Input<inputs.solutions.ApplicationJitAccessPolicyArgs>;
    /**
     * The kind of the managed application. Allowed values are MarketPlace and ServiceCatalog.
     */
    readonly kind: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * ID of the resource that manages this resource.
     */
    readonly managedBy?: pulumi.Input<string>;
    /**
     * The managed resource group Id.
     */
    readonly managedResourceGroupId?: pulumi.Input<string>;
    /**
     * Name and value pairs that define the managed application parameters. It can be a JObject or a well formed JSON string.
     */
    readonly parameters?: any;
    /**
     * The plan information.
     */
    readonly plan?: pulumi.Input<inputs.solutions.PlanArgs>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU of the resource.
     */
    readonly sku?: pulumi.Input<inputs.solutions.SkuArgs>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
