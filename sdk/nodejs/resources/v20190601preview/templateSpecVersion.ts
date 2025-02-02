// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Template Spec Version object.
 */
export class TemplateSpecVersion extends pulumi.CustomResource {
    /**
     * Get an existing TemplateSpecVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TemplateSpecVersion {
        return new TemplateSpecVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:resources/v20190601preview:TemplateSpecVersion';

    /**
     * Returns true if the given object is an instance of TemplateSpecVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TemplateSpecVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TemplateSpecVersion.__pulumiType;
    }

    /**
     * An array of Template Spec artifacts.
     */
    public readonly artifacts!: pulumi.Output<outputs.resources.v20190601preview.TemplateSpecTemplateArtifactResponse[] | undefined>;
    /**
     * Template Spec version description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The location of the Template Spec Version. It must match the location of the parent Template Spec.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of this resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.resources.v20190601preview.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Azure Resource Manager template content.
     */
    public readonly template!: pulumi.Output<any | undefined>;
    /**
     * Type of this resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a TemplateSpecVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemplateSpecVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.templateSpecName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateSpecName'");
            }
            inputs["artifacts"] = args ? args.artifacts : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["template"] = args ? args.template : undefined;
            inputs["templateSpecName"] = args ? args.templateSpecName : undefined;
            inputs["templateSpecVersion"] = args ? args.templateSpecVersion : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["artifacts"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["template"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:resources/v20190601preview:TemplateSpecVersion" }, { type: "azure-native:resources:TemplateSpecVersion" }, { type: "azure-nextgen:resources:TemplateSpecVersion" }, { type: "azure-native:resources/v20210301preview:TemplateSpecVersion" }, { type: "azure-nextgen:resources/v20210301preview:TemplateSpecVersion" }, { type: "azure-native:resources/v20210501:TemplateSpecVersion" }, { type: "azure-nextgen:resources/v20210501:TemplateSpecVersion" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(TemplateSpecVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TemplateSpecVersion resource.
 */
export interface TemplateSpecVersionArgs {
    /**
     * An array of Template Spec artifacts.
     */
    readonly artifacts?: pulumi.Input<pulumi.Input<inputs.resources.v20190601preview.TemplateSpecTemplateArtifactArgs>[]>;
    /**
     * Template Spec version description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The location of the Template Spec Version. It must match the location of the parent Template Spec.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Azure Resource Manager template content.
     */
    readonly template?: any;
    /**
     * Name of the Template Spec.
     */
    readonly templateSpecName: pulumi.Input<string>;
    /**
     * The version of the Template Spec.
     */
    readonly templateSpecVersion?: pulumi.Input<string>;
}
