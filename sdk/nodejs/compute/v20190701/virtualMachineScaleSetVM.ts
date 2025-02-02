// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Describes a virtual machine scale set virtual machine.
 */
export class VirtualMachineScaleSetVM extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachineScaleSetVM resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualMachineScaleSetVM {
        return new VirtualMachineScaleSetVM(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:compute/v20190701:VirtualMachineScaleSetVM';

    /**
     * Returns true if the given object is an instance of VirtualMachineScaleSetVM.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualMachineScaleSetVM {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMachineScaleSetVM.__pulumiType;
    }

    /**
     * Specifies additional capabilities enabled or disabled on the virtual machine in the scale set. For instance: whether the virtual machine has the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
     */
    public readonly additionalCapabilities!: pulumi.Output<outputs.compute.v20190701.AdditionalCapabilitiesResponse | undefined>;
    /**
     * Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
     */
    public readonly availabilitySet!: pulumi.Output<outputs.compute.v20190701.SubResourceResponse | undefined>;
    /**
     * Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
     */
    public readonly diagnosticsProfile!: pulumi.Output<outputs.compute.v20190701.DiagnosticsProfileResponse | undefined>;
    /**
     * Specifies the hardware settings for the virtual machine.
     */
    public readonly hardwareProfile!: pulumi.Output<outputs.compute.v20190701.HardwareProfileResponse | undefined>;
    /**
     * The virtual machine instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The virtual machine instance view.
     */
    public /*out*/ readonly instanceView!: pulumi.Output<outputs.compute.v20190701.VirtualMachineScaleSetVMInstanceViewResponse>;
    /**
     * Specifies whether the latest model has been applied to the virtual machine.
     */
    public /*out*/ readonly latestModelApplied!: pulumi.Output<boolean>;
    /**
     * Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
     */
    public readonly licenseType!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies whether the model applied to the virtual machine is the model of the virtual machine scale set or the customized model for the virtual machine.
     */
    public /*out*/ readonly modelDefinitionApplied!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specifies the network interfaces of the virtual machine.
     */
    public readonly networkProfile!: pulumi.Output<outputs.compute.v20190701.NetworkProfileResponse | undefined>;
    /**
     * Specifies the network profile configuration of the virtual machine.
     */
    public readonly networkProfileConfiguration!: pulumi.Output<outputs.compute.v20190701.VirtualMachineScaleSetVMNetworkProfileConfigurationResponse | undefined>;
    /**
     * Specifies the operating system settings for the virtual machine.
     */
    public readonly osProfile!: pulumi.Output<outputs.compute.v20190701.OSProfileResponse | undefined>;
    /**
     * Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
     */
    public readonly plan!: pulumi.Output<outputs.compute.v20190701.PlanResponse | undefined>;
    /**
     * Specifies the protection policy of the virtual machine.
     */
    public readonly protectionPolicy!: pulumi.Output<outputs.compute.v20190701.VirtualMachineScaleSetVMProtectionPolicyResponse | undefined>;
    /**
     * The provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The virtual machine child extension resources.
     */
    public /*out*/ readonly resources!: pulumi.Output<outputs.compute.v20190701.VirtualMachineExtensionResponse[]>;
    /**
     * The virtual machine SKU.
     */
    public /*out*/ readonly sku!: pulumi.Output<outputs.compute.v20190701.SkuResponse>;
    /**
     * Specifies the storage settings for the virtual machine disks.
     */
    public readonly storageProfile!: pulumi.Output<outputs.compute.v20190701.StorageProfileResponse | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Azure VM unique ID.
     */
    public /*out*/ readonly vmId!: pulumi.Output<string>;
    /**
     * The virtual machine zones.
     */
    public /*out*/ readonly zones!: pulumi.Output<string[]>;

    /**
     * Create a VirtualMachineScaleSetVM resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineScaleSetVMArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.vmScaleSetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmScaleSetName'");
            }
            inputs["additionalCapabilities"] = args ? args.additionalCapabilities : undefined;
            inputs["availabilitySet"] = args ? args.availabilitySet : undefined;
            inputs["diagnosticsProfile"] = args ? args.diagnosticsProfile : undefined;
            inputs["hardwareProfile"] = args ? args.hardwareProfile : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["licenseType"] = args ? args.licenseType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkProfile"] = args ? args.networkProfile : undefined;
            inputs["networkProfileConfiguration"] = args ? args.networkProfileConfiguration : undefined;
            inputs["osProfile"] = args ? args.osProfile : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["protectionPolicy"] = args ? args.protectionPolicy : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageProfile"] = args ? args.storageProfile : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vmScaleSetName"] = args ? args.vmScaleSetName : undefined;
            inputs["instanceView"] = undefined /*out*/;
            inputs["latestModelApplied"] = undefined /*out*/;
            inputs["modelDefinitionApplied"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resources"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vmId"] = undefined /*out*/;
            inputs["zones"] = undefined /*out*/;
        } else {
            inputs["additionalCapabilities"] = undefined /*out*/;
            inputs["availabilitySet"] = undefined /*out*/;
            inputs["diagnosticsProfile"] = undefined /*out*/;
            inputs["hardwareProfile"] = undefined /*out*/;
            inputs["instanceId"] = undefined /*out*/;
            inputs["instanceView"] = undefined /*out*/;
            inputs["latestModelApplied"] = undefined /*out*/;
            inputs["licenseType"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["modelDefinitionApplied"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkProfile"] = undefined /*out*/;
            inputs["networkProfileConfiguration"] = undefined /*out*/;
            inputs["osProfile"] = undefined /*out*/;
            inputs["plan"] = undefined /*out*/;
            inputs["protectionPolicy"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resources"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["storageProfile"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vmId"] = undefined /*out*/;
            inputs["zones"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:compute/v20190701:VirtualMachineScaleSetVM" }, { type: "azure-native:compute:VirtualMachineScaleSetVM" }, { type: "azure-nextgen:compute:VirtualMachineScaleSetVM" }, { type: "azure-native:compute/v20171201:VirtualMachineScaleSetVM" }, { type: "azure-nextgen:compute/v20171201:VirtualMachineScaleSetVM" }, { type: "azure-native:compute/v20180401:VirtualMachineScaleSetVM" }, { type: "azure-nextgen:compute/v20180401:VirtualMachineScaleSetVM" }, { type: "azure-native:compute/v20180601:VirtualMachineScaleSetVM" }, { type: "azure-nextgen:compute/v20180601:VirtualMachineScaleSetVM" }, { type: "azure-native:compute/v20181001:VirtualMachineScaleSetVM" }, { type: "azure-nextgen:compute/v20181001:VirtualMachineScaleSetVM" }, { type: "azure-native:compute/v20190301:VirtualMachineScaleSetVM" }, { type: "azure-nextgen:compute/v20190301:VirtualMachineScaleSetVM" }, { type: "azure-native:compute/v20191201:VirtualMachineScaleSetVM" }, { type: "azure-nextgen:compute/v20191201:VirtualMachineScaleSetVM" }, { type: "azure-native:compute/v20200601:VirtualMachineScaleSetVM" }, { type: "azure-nextgen:compute/v20200601:VirtualMachineScaleSetVM" }, { type: "azure-native:compute/v20201201:VirtualMachineScaleSetVM" }, { type: "azure-nextgen:compute/v20201201:VirtualMachineScaleSetVM" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(VirtualMachineScaleSetVM.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualMachineScaleSetVM resource.
 */
export interface VirtualMachineScaleSetVMArgs {
    /**
     * Specifies additional capabilities enabled or disabled on the virtual machine in the scale set. For instance: whether the virtual machine has the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
     */
    readonly additionalCapabilities?: pulumi.Input<inputs.compute.v20190701.AdditionalCapabilitiesArgs>;
    /**
     * Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
     */
    readonly availabilitySet?: pulumi.Input<inputs.compute.v20190701.SubResourceArgs>;
    /**
     * Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
     */
    readonly diagnosticsProfile?: pulumi.Input<inputs.compute.v20190701.DiagnosticsProfileArgs>;
    /**
     * Specifies the hardware settings for the virtual machine.
     */
    readonly hardwareProfile?: pulumi.Input<inputs.compute.v20190701.HardwareProfileArgs>;
    /**
     * The instance ID of the virtual machine.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
     */
    readonly licenseType?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the network interfaces of the virtual machine.
     */
    readonly networkProfile?: pulumi.Input<inputs.compute.v20190701.NetworkProfileArgs>;
    /**
     * Specifies the network profile configuration of the virtual machine.
     */
    readonly networkProfileConfiguration?: pulumi.Input<inputs.compute.v20190701.VirtualMachineScaleSetVMNetworkProfileConfigurationArgs>;
    /**
     * Specifies the operating system settings for the virtual machine.
     */
    readonly osProfile?: pulumi.Input<inputs.compute.v20190701.OSProfileArgs>;
    /**
     * Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
     */
    readonly plan?: pulumi.Input<inputs.compute.v20190701.PlanArgs>;
    /**
     * Specifies the protection policy of the virtual machine.
     */
    readonly protectionPolicy?: pulumi.Input<inputs.compute.v20190701.VirtualMachineScaleSetVMProtectionPolicyArgs>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the storage settings for the virtual machine disks.
     */
    readonly storageProfile?: pulumi.Input<inputs.compute.v20190701.StorageProfileArgs>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the VM scale set where the extension should be create or updated.
     */
    readonly vmScaleSetName: pulumi.Input<string>;
}
