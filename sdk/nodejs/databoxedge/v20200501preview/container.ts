// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Represents a container on the  Data Box Edge/Gateway device.
 */
export class Container extends pulumi.CustomResource {
    /**
     * Get an existing Container resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Container {
        return new Container(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:databoxedge/v20200501preview:Container';

    /**
     * Returns true if the given object is an instance of Container.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Container {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Container.__pulumiType;
    }

    /**
     * Current status of the container.
     */
    public /*out*/ readonly containerStatus!: pulumi.Output<string>;
    /**
     * The UTC time when container got created.
     */
    public /*out*/ readonly createdDateTime!: pulumi.Output<string>;
    /**
     * DataFormat for Container
     */
    public readonly dataFormat!: pulumi.Output<string>;
    /**
     * The object name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Details of the refresh job on this container.
     */
    public /*out*/ readonly refreshDetails!: pulumi.Output<outputs.databoxedge.v20200501preview.RefreshDetailsResponse>;
    /**
     * The hierarchical type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Container resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dataFormat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataFormat'");
            }
            if ((!args || args.deviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.storageAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageAccountName'");
            }
            inputs["containerName"] = args ? args.containerName : undefined;
            inputs["dataFormat"] = args ? args.dataFormat : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageAccountName"] = args ? args.storageAccountName : undefined;
            inputs["containerStatus"] = undefined /*out*/;
            inputs["createdDateTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["refreshDetails"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["containerStatus"] = undefined /*out*/;
            inputs["createdDateTime"] = undefined /*out*/;
            inputs["dataFormat"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["refreshDetails"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:databoxedge/v20200501preview:Container" }, { type: "azure-native:databoxedge:Container" }, { type: "azure-nextgen:databoxedge:Container" }, { type: "azure-native:databoxedge/v20190801:Container" }, { type: "azure-nextgen:databoxedge/v20190801:Container" }, { type: "azure-native:databoxedge/v20200901:Container" }, { type: "azure-nextgen:databoxedge/v20200901:Container" }, { type: "azure-native:databoxedge/v20200901preview:Container" }, { type: "azure-nextgen:databoxedge/v20200901preview:Container" }, { type: "azure-native:databoxedge/v20201201:Container" }, { type: "azure-nextgen:databoxedge/v20201201:Container" }, { type: "azure-native:databoxedge/v20210201preview:Container" }, { type: "azure-nextgen:databoxedge/v20210201preview:Container" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Container.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Container resource.
 */
export interface ContainerArgs {
    /**
     * The container name.
     */
    readonly containerName?: pulumi.Input<string>;
    /**
     * DataFormat for Container
     */
    readonly dataFormat: pulumi.Input<string | enums.databoxedge.v20200501preview.AzureContainerDataFormat>;
    /**
     * The device name.
     */
    readonly deviceName: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The Storage Account Name
     */
    readonly storageAccountName: pulumi.Input<string>;
}
