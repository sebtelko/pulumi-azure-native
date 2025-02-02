// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * An object that represents a connected registry for a container registry.
 */
export class ConnectedRegistry extends pulumi.CustomResource {
    /**
     * Get an existing ConnectedRegistry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectedRegistry {
        return new ConnectedRegistry(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:containerregistry/v20201101preview:ConnectedRegistry';

    /**
     * Returns true if the given object is an instance of ConnectedRegistry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectedRegistry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectedRegistry.__pulumiType;
    }

    /**
     * The activation properties of the connected registry.
     */
    public /*out*/ readonly activation!: pulumi.Output<outputs.containerregistry.v20201101preview.ActivationPropertiesResponse>;
    /**
     * The list of the ACR token resource IDs used to authenticate clients to the connected registry.
     */
    public readonly clientTokenIds!: pulumi.Output<string[] | undefined>;
    /**
     * The current connection state of the connected registry.
     */
    public /*out*/ readonly connectionState!: pulumi.Output<string>;
    /**
     * The last activity time of the connected registry.
     */
    public /*out*/ readonly lastActivityTime!: pulumi.Output<string>;
    /**
     * The logging properties of the connected registry.
     */
    public readonly logging!: pulumi.Output<outputs.containerregistry.v20201101preview.LoggingPropertiesResponse | undefined>;
    /**
     * The login server properties of the connected registry.
     */
    public /*out*/ readonly loginServer!: pulumi.Output<outputs.containerregistry.v20201101preview.LoginServerPropertiesResponse | undefined>;
    /**
     * The mode of the connected registry resource that indicates the permissions of the registry.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parent of the connected registry.
     */
    public readonly parent!: pulumi.Output<outputs.containerregistry.v20201101preview.ParentPropertiesResponse>;
    /**
     * Provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The list of current statuses of the connected registry.
     */
    public /*out*/ readonly statusDetails!: pulumi.Output<outputs.containerregistry.v20201101preview.StatusDetailPropertiesResponse[]>;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.containerregistry.v20201101preview.SystemDataResponse>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The current version of ACR runtime on the connected registry.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a ConnectedRegistry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectedRegistryArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.mode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mode'");
            }
            if ((!args || args.parent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent'");
            }
            if ((!args || args.registryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["clientTokenIds"] = args ? args.clientTokenIds : undefined;
            inputs["connectedRegistryName"] = args ? args.connectedRegistryName : undefined;
            inputs["logging"] = args ? args.logging : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["registryName"] = args ? args.registryName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["activation"] = undefined /*out*/;
            inputs["connectionState"] = undefined /*out*/;
            inputs["lastActivityTime"] = undefined /*out*/;
            inputs["loginServer"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["statusDetails"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        } else {
            inputs["activation"] = undefined /*out*/;
            inputs["clientTokenIds"] = undefined /*out*/;
            inputs["connectionState"] = undefined /*out*/;
            inputs["lastActivityTime"] = undefined /*out*/;
            inputs["logging"] = undefined /*out*/;
            inputs["loginServer"] = undefined /*out*/;
            inputs["mode"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parent"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["statusDetails"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:containerregistry/v20201101preview:ConnectedRegistry" }, { type: "azure-native:containerregistry:ConnectedRegistry" }, { type: "azure-nextgen:containerregistry:ConnectedRegistry" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ConnectedRegistry.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConnectedRegistry resource.
 */
export interface ConnectedRegistryArgs {
    /**
     * The list of the ACR token resource IDs used to authenticate clients to the connected registry.
     */
    readonly clientTokenIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the connected registry.
     */
    readonly connectedRegistryName?: pulumi.Input<string>;
    /**
     * The logging properties of the connected registry.
     */
    readonly logging?: pulumi.Input<inputs.containerregistry.v20201101preview.LoggingPropertiesArgs>;
    /**
     * The mode of the connected registry resource that indicates the permissions of the registry.
     */
    readonly mode: pulumi.Input<string | enums.containerregistry.v20201101preview.ConnectedRegistryMode>;
    /**
     * The parent of the connected registry.
     */
    readonly parent: pulumi.Input<inputs.containerregistry.v20201101preview.ParentPropertiesArgs>;
    /**
     * The name of the container registry.
     */
    readonly registryName: pulumi.Input<string>;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
