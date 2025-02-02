// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The live event.
 * API Version: 2020-05-01.
 */
export class LiveEvent extends pulumi.CustomResource {
    /**
     * Get an existing LiveEvent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LiveEvent {
        return new LiveEvent(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:media:LiveEvent';

    /**
     * Returns true if the given object is an instance of LiveEvent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LiveEvent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LiveEvent.__pulumiType;
    }

    /**
     * The creation time for the live event
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * Live event cross site access policies.
     */
    public readonly crossSiteAccessPolicies!: pulumi.Output<outputs.media.CrossSiteAccessPoliciesResponse | undefined>;
    /**
     * A description for the live event.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
     */
    public readonly encoding!: pulumi.Output<outputs.media.LiveEventEncodingResponse | undefined>;
    /**
     * When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
     */
    public readonly hostnamePrefix!: pulumi.Output<string | undefined>;
    /**
     * Live event input settings. It defines how the live event receives input from a contribution encoder.
     */
    public readonly input!: pulumi.Output<outputs.media.LiveEventInputResponse>;
    /**
     * The last modified time of the live event.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
     */
    public readonly preview!: pulumi.Output<outputs.media.LiveEventPreviewResponse | undefined>;
    /**
     * The provisioning state of the live event.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource state of the live event. See https://go.microsoft.com/fwlink/?linkid=2139012 for more information.
     */
    public /*out*/ readonly resourceState!: pulumi.Output<string>;
    /**
     * The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
     */
    public readonly streamOptions!: pulumi.Output<string[] | undefined>;
    /**
     * The system metadata relating to this resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.media.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
     */
    public readonly transcriptions!: pulumi.Output<outputs.media.LiveEventTranscriptionResponse[] | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
     */
    public readonly useStaticHostname!: pulumi.Output<boolean | undefined>;

    /**
     * Create a LiveEvent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LiveEventArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.input === undefined) && !opts.urn) {
                throw new Error("Missing required property 'input'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["autoStart"] = args ? args.autoStart : undefined;
            inputs["crossSiteAccessPolicies"] = args ? args.crossSiteAccessPolicies : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encoding"] = args ? args.encoding : undefined;
            inputs["hostnamePrefix"] = args ? args.hostnamePrefix : undefined;
            inputs["input"] = args ? args.input : undefined;
            inputs["liveEventName"] = args ? args.liveEventName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["preview"] = args ? args.preview : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["streamOptions"] = args ? args.streamOptions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["transcriptions"] = args ? args.transcriptions : undefined;
            inputs["useStaticHostname"] = args ? args.useStaticHostname : undefined;
            inputs["created"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["created"] = undefined /*out*/;
            inputs["crossSiteAccessPolicies"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["encoding"] = undefined /*out*/;
            inputs["hostnamePrefix"] = undefined /*out*/;
            inputs["input"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["preview"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["streamOptions"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["transcriptions"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["useStaticHostname"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:media:LiveEvent" }, { type: "azure-native:media/v20180330preview:LiveEvent" }, { type: "azure-nextgen:media/v20180330preview:LiveEvent" }, { type: "azure-native:media/v20180601preview:LiveEvent" }, { type: "azure-nextgen:media/v20180601preview:LiveEvent" }, { type: "azure-native:media/v20180701:LiveEvent" }, { type: "azure-nextgen:media/v20180701:LiveEvent" }, { type: "azure-native:media/v20190501preview:LiveEvent" }, { type: "azure-nextgen:media/v20190501preview:LiveEvent" }, { type: "azure-native:media/v20200501:LiveEvent" }, { type: "azure-nextgen:media/v20200501:LiveEvent" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(LiveEvent.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LiveEvent resource.
 */
export interface LiveEventArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The flag indicates if the resource should be automatically started on creation.
     */
    readonly autoStart?: pulumi.Input<boolean>;
    /**
     * Live event cross site access policies.
     */
    readonly crossSiteAccessPolicies?: pulumi.Input<inputs.media.CrossSiteAccessPoliciesArgs>;
    /**
     * A description for the live event.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
     */
    readonly encoding?: pulumi.Input<inputs.media.LiveEventEncodingArgs>;
    /**
     * When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
     */
    readonly hostnamePrefix?: pulumi.Input<string>;
    /**
     * Live event input settings. It defines how the live event receives input from a contribution encoder.
     */
    readonly input: pulumi.Input<inputs.media.LiveEventInputArgs>;
    /**
     * The name of the live event, maximum length is 32.
     */
    readonly liveEventName?: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
     */
    readonly preview?: pulumi.Input<inputs.media.LiveEventPreviewArgs>;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
     */
    readonly streamOptions?: pulumi.Input<pulumi.Input<string | enums.media.StreamOptionsFlag>[]>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
     */
    readonly transcriptions?: pulumi.Input<pulumi.Input<inputs.media.LiveEventTranscriptionArgs>[]>;
    /**
     * Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
     */
    readonly useStaticHostname?: pulumi.Input<boolean>;
}
