// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Base class for backup items.
 */
export class ProtectedItem extends pulumi.CustomResource {
    /**
     * Get an existing ProtectedItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ProtectedItem {
        return new ProtectedItem(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:recoveryservices/v20201001:ProtectedItem';

    /**
     * Returns true if the given object is an instance of ProtectedItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProtectedItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProtectedItem.__pulumiType;
    }

    /**
     * Optional ETag.
     */
    public readonly eTag!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name associated with the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * ProtectedItemResource properties
     */
    public readonly properties!: pulumi.Output<outputs.recoveryservices.v20201001.AzureFileshareProtectedItemResponse | outputs.recoveryservices.v20201001.AzureIaaSClassicComputeVMProtectedItemResponse | outputs.recoveryservices.v20201001.AzureIaaSComputeVMProtectedItemResponse | outputs.recoveryservices.v20201001.AzureIaaSVMProtectedItemResponse | outputs.recoveryservices.v20201001.AzureSqlProtectedItemResponse | outputs.recoveryservices.v20201001.AzureVmWorkloadProtectedItemResponse | outputs.recoveryservices.v20201001.AzureVmWorkloadSAPAseDatabaseProtectedItemResponse | outputs.recoveryservices.v20201001.AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse | outputs.recoveryservices.v20201001.AzureVmWorkloadSQLDatabaseProtectedItemResponse | outputs.recoveryservices.v20201001.DPMProtectedItemResponse | outputs.recoveryservices.v20201001.GenericProtectedItemResponse | outputs.recoveryservices.v20201001.MabFileFolderProtectedItemResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ProtectedItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProtectedItemArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.containerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'containerName'");
            }
            if ((!args || args.fabricName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fabricName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.vaultName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vaultName'");
            }
            inputs["containerName"] = args ? args.containerName : undefined;
            inputs["eTag"] = args ? args.eTag : undefined;
            inputs["fabricName"] = args ? args.fabricName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["protectedItemName"] = args ? args.protectedItemName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vaultName"] = args ? args.vaultName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["eTag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:recoveryservices/v20201001:ProtectedItem" }, { type: "azure-native:recoveryservices:ProtectedItem" }, { type: "azure-nextgen:recoveryservices:ProtectedItem" }, { type: "azure-native:recoveryservices/v20160601:ProtectedItem" }, { type: "azure-nextgen:recoveryservices/v20160601:ProtectedItem" }, { type: "azure-native:recoveryservices/v20190513:ProtectedItem" }, { type: "azure-nextgen:recoveryservices/v20190513:ProtectedItem" }, { type: "azure-native:recoveryservices/v20190615:ProtectedItem" }, { type: "azure-nextgen:recoveryservices/v20190615:ProtectedItem" }, { type: "azure-native:recoveryservices/v20201201:ProtectedItem" }, { type: "azure-nextgen:recoveryservices/v20201201:ProtectedItem" }, { type: "azure-native:recoveryservices/v20210101:ProtectedItem" }, { type: "azure-nextgen:recoveryservices/v20210101:ProtectedItem" }, { type: "azure-native:recoveryservices/v20210201:ProtectedItem" }, { type: "azure-nextgen:recoveryservices/v20210201:ProtectedItem" }, { type: "azure-native:recoveryservices/v20210201preview:ProtectedItem" }, { type: "azure-nextgen:recoveryservices/v20210201preview:ProtectedItem" }, { type: "azure-native:recoveryservices/v20210210:ProtectedItem" }, { type: "azure-nextgen:recoveryservices/v20210210:ProtectedItem" }, { type: "azure-native:recoveryservices/v20210301:ProtectedItem" }, { type: "azure-nextgen:recoveryservices/v20210301:ProtectedItem" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ProtectedItem.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ProtectedItem resource.
 */
export interface ProtectedItemArgs {
    /**
     * Container name associated with the backup item.
     */
    readonly containerName: pulumi.Input<string>;
    /**
     * Optional ETag.
     */
    readonly eTag?: pulumi.Input<string>;
    /**
     * Fabric name associated with the backup item.
     */
    readonly fabricName: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * ProtectedItemResource properties
     */
    readonly properties?: pulumi.Input<inputs.recoveryservices.v20201001.AzureFileshareProtectedItemArgs | inputs.recoveryservices.v20201001.AzureIaaSClassicComputeVMProtectedItemArgs | inputs.recoveryservices.v20201001.AzureIaaSComputeVMProtectedItemArgs | inputs.recoveryservices.v20201001.AzureIaaSVMProtectedItemArgs | inputs.recoveryservices.v20201001.AzureSqlProtectedItemArgs | inputs.recoveryservices.v20201001.AzureVmWorkloadProtectedItemArgs | inputs.recoveryservices.v20201001.AzureVmWorkloadSAPAseDatabaseProtectedItemArgs | inputs.recoveryservices.v20201001.AzureVmWorkloadSAPHanaDatabaseProtectedItemArgs | inputs.recoveryservices.v20201001.AzureVmWorkloadSQLDatabaseProtectedItemArgs | inputs.recoveryservices.v20201001.DPMProtectedItemArgs | inputs.recoveryservices.v20201001.GenericProtectedItemArgs | inputs.recoveryservices.v20201001.MabFileFolderProtectedItemArgs>;
    /**
     * Item name to be backed up.
     */
    readonly protectedItemName?: pulumi.Input<string>;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the recovery services vault.
     */
    readonly vaultName: pulumi.Input<string>;
}
