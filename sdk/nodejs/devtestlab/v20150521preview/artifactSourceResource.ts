// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Properties of an artifact source.
 */
export class ArtifactSourceResource extends pulumi.CustomResource {
    /**
     * Get an existing ArtifactSourceResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ArtifactSourceResource {
        return new ArtifactSourceResource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:devtestlab/v20150521preview:ArtifactSourceResource';

    /**
     * Returns true if the given object is an instance of ArtifactSourceResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ArtifactSourceResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ArtifactSourceResource.__pulumiType;
    }

    /**
     * The branch reference of the artifact source.
     */
    public readonly branchRef!: pulumi.Output<string | undefined>;
    /**
     * The display name of the artifact source.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The folder path of the artifact source.
     */
    public readonly folderPath!: pulumi.Output<string | undefined>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The provisioning status of the resource.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The security token of the artifact source.
     */
    public readonly securityToken!: pulumi.Output<string | undefined>;
    /**
     * The type of the artifact source.
     */
    public readonly sourceType!: pulumi.Output<string | undefined>;
    /**
     * The status of the artifact source.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The URI of the artifact source.
     */
    public readonly uri!: pulumi.Output<string | undefined>;

    /**
     * Create a ArtifactSourceResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ArtifactSourceResourceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.labName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'labName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["branchRef"] = args ? args.branchRef : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["folderPath"] = args ? args.folderPath : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["securityToken"] = args ? args.securityToken : undefined;
            inputs["sourceType"] = args ? args.sourceType : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["uri"] = args ? args.uri : undefined;
        } else {
            inputs["branchRef"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["folderPath"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["securityToken"] = undefined /*out*/;
            inputs["sourceType"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uri"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:devtestlab/v20150521preview:ArtifactSourceResource" }, { type: "azure-native:devtestlab:ArtifactSourceResource" }, { type: "azure-nextgen:devtestlab:ArtifactSourceResource" }, { type: "azure-native:devtestlab/v20160515:ArtifactSourceResource" }, { type: "azure-nextgen:devtestlab/v20160515:ArtifactSourceResource" }, { type: "azure-native:devtestlab/v20180915:ArtifactSourceResource" }, { type: "azure-nextgen:devtestlab/v20180915:ArtifactSourceResource" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ArtifactSourceResource.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ArtifactSourceResource resource.
 */
export interface ArtifactSourceResourceArgs {
    /**
     * The branch reference of the artifact source.
     */
    readonly branchRef?: pulumi.Input<string>;
    /**
     * The display name of the artifact source.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The folder path of the artifact source.
     */
    readonly folderPath?: pulumi.Input<string>;
    /**
     * The identifier of the resource.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the lab.
     */
    readonly labName: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The provisioning status of the resource.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The security token of the artifact source.
     */
    readonly securityToken?: pulumi.Input<string>;
    /**
     * The type of the artifact source.
     */
    readonly sourceType?: pulumi.Input<string | enums.devtestlab.v20150521preview.SourceControlType>;
    /**
     * The status of the artifact source.
     */
    readonly status?: pulumi.Input<string | enums.devtestlab.v20150521preview.EnableStatus>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the resource.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The URI of the artifact source.
     */
    readonly uri?: pulumi.Input<string>;
}
