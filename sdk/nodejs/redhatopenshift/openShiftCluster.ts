// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
 * API Version: 2020-04-30.
 */
export class OpenShiftCluster extends pulumi.CustomResource {
    /**
     * Get an existing OpenShiftCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OpenShiftCluster {
        return new OpenShiftCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:redhatopenshift:OpenShiftCluster';

    /**
     * Returns true if the given object is an instance of OpenShiftCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OpenShiftCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OpenShiftCluster.__pulumiType;
    }

    /**
     * The cluster API server profile.
     */
    public readonly apiserverProfile!: pulumi.Output<outputs.redhatopenshift.APIServerProfileResponse | undefined>;
    /**
     * The cluster profile.
     */
    public readonly clusterProfile!: pulumi.Output<outputs.redhatopenshift.ClusterProfileResponse | undefined>;
    /**
     * The console profile.
     */
    public readonly consoleProfile!: pulumi.Output<outputs.redhatopenshift.ConsoleProfileResponse | undefined>;
    /**
     * The cluster ingress profiles.
     */
    public readonly ingressProfiles!: pulumi.Output<outputs.redhatopenshift.IngressProfileResponse[] | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The cluster master profile.
     */
    public readonly masterProfile!: pulumi.Output<outputs.redhatopenshift.MasterProfileResponse | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The cluster network profile.
     */
    public readonly networkProfile!: pulumi.Output<outputs.redhatopenshift.NetworkProfileResponse | undefined>;
    /**
     * The cluster provisioning state (immutable).
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The cluster service principal profile.
     */
    public readonly servicePrincipalProfile!: pulumi.Output<outputs.redhatopenshift.ServicePrincipalProfileResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The cluster worker profiles.
     */
    public readonly workerProfiles!: pulumi.Output<outputs.redhatopenshift.WorkerProfileResponse[] | undefined>;

    /**
     * Create a OpenShiftCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OpenShiftClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["apiserverProfile"] = args ? args.apiserverProfile : undefined;
            inputs["clusterProfile"] = args ? args.clusterProfile : undefined;
            inputs["consoleProfile"] = args ? args.consoleProfile : undefined;
            inputs["ingressProfiles"] = args ? args.ingressProfiles : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["masterProfile"] = args ? args.masterProfile : undefined;
            inputs["networkProfile"] = args ? args.networkProfile : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["servicePrincipalProfile"] = args ? args.servicePrincipalProfile : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workerProfiles"] = args ? args.workerProfiles : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["apiserverProfile"] = undefined /*out*/;
            inputs["clusterProfile"] = undefined /*out*/;
            inputs["consoleProfile"] = undefined /*out*/;
            inputs["ingressProfiles"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["masterProfile"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkProfile"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["servicePrincipalProfile"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["workerProfiles"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:redhatopenshift:OpenShiftCluster" }, { type: "azure-native:redhatopenshift/v20200430:OpenShiftCluster" }, { type: "azure-nextgen:redhatopenshift/v20200430:OpenShiftCluster" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(OpenShiftCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OpenShiftCluster resource.
 */
export interface OpenShiftClusterArgs {
    /**
     * The cluster API server profile.
     */
    readonly apiserverProfile?: pulumi.Input<inputs.redhatopenshift.APIServerProfileArgs>;
    /**
     * The cluster profile.
     */
    readonly clusterProfile?: pulumi.Input<inputs.redhatopenshift.ClusterProfileArgs>;
    /**
     * The console profile.
     */
    readonly consoleProfile?: pulumi.Input<inputs.redhatopenshift.ConsoleProfileArgs>;
    /**
     * The cluster ingress profiles.
     */
    readonly ingressProfiles?: pulumi.Input<pulumi.Input<inputs.redhatopenshift.IngressProfileArgs>[]>;
    /**
     * The geo-location where the resource lives
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The cluster master profile.
     */
    readonly masterProfile?: pulumi.Input<inputs.redhatopenshift.MasterProfileArgs>;
    /**
     * The cluster network profile.
     */
    readonly networkProfile?: pulumi.Input<inputs.redhatopenshift.NetworkProfileArgs>;
    /**
     * The cluster provisioning state (immutable).
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the OpenShift cluster resource.
     */
    readonly resourceName?: pulumi.Input<string>;
    /**
     * The cluster service principal profile.
     */
    readonly servicePrincipalProfile?: pulumi.Input<inputs.redhatopenshift.ServicePrincipalProfileArgs>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The cluster worker profiles.
     */
    readonly workerProfiles?: pulumi.Input<pulumi.Input<inputs.redhatopenshift.WorkerProfileArgs>[]>;
}
