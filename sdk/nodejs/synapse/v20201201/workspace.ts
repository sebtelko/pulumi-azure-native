// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * A workspace
 */
export class Workspace extends pulumi.CustomResource {
    /**
     * Get an existing Workspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workspace {
        return new Workspace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:synapse/v20201201:Workspace';

    /**
     * Returns true if the given object is an instance of Workspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workspace.__pulumiType;
    }

    /**
     * The ADLA resource ID.
     */
    public /*out*/ readonly adlaResourceId!: pulumi.Output<string>;
    /**
     * Connectivity endpoints
     */
    public readonly connectivityEndpoints!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Workspace default data lake storage account details
     */
    public readonly defaultDataLakeStorage!: pulumi.Output<outputs.synapse.v20201201.DataLakeStorageAccountDetailsResponse | undefined>;
    /**
     * The encryption details of the workspace
     */
    public readonly encryption!: pulumi.Output<outputs.synapse.v20201201.EncryptionDetailsResponse | undefined>;
    /**
     * Workspace level configs and feature flags
     */
    public /*out*/ readonly extraProperties!: pulumi.Output<{[key: string]: any}>;
    /**
     * Identity of the workspace
     */
    public readonly identity!: pulumi.Output<outputs.synapse.v20201201.ManagedIdentityResponse | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Workspace managed resource group. The resource group name uniquely identifies the resource group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
     */
    public readonly managedResourceGroupName!: pulumi.Output<string | undefined>;
    /**
     * Setting this to 'default' will ensure that all compute for this workspace is in a virtual network managed on behalf of the user.
     */
    public readonly managedVirtualNetwork!: pulumi.Output<string | undefined>;
    /**
     * Managed Virtual Network Settings
     */
    public readonly managedVirtualNetworkSettings!: pulumi.Output<outputs.synapse.v20201201.ManagedVirtualNetworkSettingsResponse | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Private endpoint connections to the workspace
     */
    public readonly privateEndpointConnections!: pulumi.Output<outputs.synapse.v20201201.PrivateEndpointConnectionResponse[] | undefined>;
    /**
     * Resource provisioning state
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Purview Configuration
     */
    public readonly purviewConfiguration!: pulumi.Output<outputs.synapse.v20201201.PurviewConfigurationResponse | undefined>;
    /**
     * Login for workspace SQL active directory administrator
     */
    public readonly sqlAdministratorLogin!: pulumi.Output<string | undefined>;
    /**
     * SQL administrator login password
     */
    public readonly sqlAdministratorLoginPassword!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Virtual Network profile
     */
    public readonly virtualNetworkProfile!: pulumi.Output<outputs.synapse.v20201201.VirtualNetworkProfileResponse | undefined>;
    /**
     * Git integration settings
     */
    public readonly workspaceRepositoryConfiguration!: pulumi.Output<outputs.synapse.v20201201.WorkspaceRepositoryConfigurationResponse | undefined>;
    /**
     * The workspace unique identifier
     */
    public /*out*/ readonly workspaceUID!: pulumi.Output<string>;

    /**
     * Create a Workspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["connectivityEndpoints"] = args ? args.connectivityEndpoints : undefined;
            inputs["defaultDataLakeStorage"] = args ? args.defaultDataLakeStorage : undefined;
            inputs["encryption"] = args ? args.encryption : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedResourceGroupName"] = args ? args.managedResourceGroupName : undefined;
            inputs["managedVirtualNetwork"] = args ? args.managedVirtualNetwork : undefined;
            inputs["managedVirtualNetworkSettings"] = args ? args.managedVirtualNetworkSettings : undefined;
            inputs["privateEndpointConnections"] = args ? args.privateEndpointConnections : undefined;
            inputs["purviewConfiguration"] = args ? args.purviewConfiguration : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sqlAdministratorLogin"] = args ? args.sqlAdministratorLogin : undefined;
            inputs["sqlAdministratorLoginPassword"] = args ? args.sqlAdministratorLoginPassword : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualNetworkProfile"] = args ? args.virtualNetworkProfile : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["workspaceRepositoryConfiguration"] = args ? args.workspaceRepositoryConfiguration : undefined;
            inputs["adlaResourceId"] = undefined /*out*/;
            inputs["extraProperties"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["workspaceUID"] = undefined /*out*/;
        } else {
            inputs["adlaResourceId"] = undefined /*out*/;
            inputs["connectivityEndpoints"] = undefined /*out*/;
            inputs["defaultDataLakeStorage"] = undefined /*out*/;
            inputs["encryption"] = undefined /*out*/;
            inputs["extraProperties"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["managedResourceGroupName"] = undefined /*out*/;
            inputs["managedVirtualNetwork"] = undefined /*out*/;
            inputs["managedVirtualNetworkSettings"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["purviewConfiguration"] = undefined /*out*/;
            inputs["sqlAdministratorLogin"] = undefined /*out*/;
            inputs["sqlAdministratorLoginPassword"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetworkProfile"] = undefined /*out*/;
            inputs["workspaceRepositoryConfiguration"] = undefined /*out*/;
            inputs["workspaceUID"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:synapse/v20201201:Workspace" }, { type: "azure-native:synapse:Workspace" }, { type: "azure-nextgen:synapse:Workspace" }, { type: "azure-native:synapse/v20190601preview:Workspace" }, { type: "azure-nextgen:synapse/v20190601preview:Workspace" }, { type: "azure-native:synapse/v20210301:Workspace" }, { type: "azure-nextgen:synapse/v20210301:Workspace" }, { type: "azure-native:synapse/v20210401preview:Workspace" }, { type: "azure-nextgen:synapse/v20210401preview:Workspace" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Workspace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * Connectivity endpoints
     */
    readonly connectivityEndpoints?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Workspace default data lake storage account details
     */
    readonly defaultDataLakeStorage?: pulumi.Input<inputs.synapse.v20201201.DataLakeStorageAccountDetailsArgs>;
    /**
     * The encryption details of the workspace
     */
    readonly encryption?: pulumi.Input<inputs.synapse.v20201201.EncryptionDetailsArgs>;
    /**
     * Identity of the workspace
     */
    readonly identity?: pulumi.Input<inputs.synapse.v20201201.ManagedIdentityArgs>;
    /**
     * The geo-location where the resource lives
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Workspace managed resource group. The resource group name uniquely identifies the resource group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
     */
    readonly managedResourceGroupName?: pulumi.Input<string>;
    /**
     * Setting this to 'default' will ensure that all compute for this workspace is in a virtual network managed on behalf of the user.
     */
    readonly managedVirtualNetwork?: pulumi.Input<string>;
    /**
     * Managed Virtual Network Settings
     */
    readonly managedVirtualNetworkSettings?: pulumi.Input<inputs.synapse.v20201201.ManagedVirtualNetworkSettingsArgs>;
    /**
     * Private endpoint connections to the workspace
     */
    readonly privateEndpointConnections?: pulumi.Input<pulumi.Input<inputs.synapse.v20201201.PrivateEndpointConnectionArgs>[]>;
    /**
     * Purview Configuration
     */
    readonly purviewConfiguration?: pulumi.Input<inputs.synapse.v20201201.PurviewConfigurationArgs>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Login for workspace SQL active directory administrator
     */
    readonly sqlAdministratorLogin?: pulumi.Input<string>;
    /**
     * SQL administrator login password
     */
    readonly sqlAdministratorLoginPassword?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Virtual Network profile
     */
    readonly virtualNetworkProfile?: pulumi.Input<inputs.synapse.v20201201.VirtualNetworkProfileArgs>;
    /**
     * The name of the workspace
     */
    readonly workspaceName?: pulumi.Input<string>;
    /**
     * Git integration settings
     */
    readonly workspaceRepositoryConfiguration?: pulumi.Input<inputs.synapse.v20201201.WorkspaceRepositoryConfigurationArgs>;
}
