// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Recipient User details.
 */
export class NotificationRecipientUser extends pulumi.CustomResource {
    /**
     * Get an existing NotificationRecipientUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NotificationRecipientUser {
        return new NotificationRecipientUser(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:apimanagement/v20180601preview:NotificationRecipientUser';

    /**
     * Returns true if the given object is an instance of NotificationRecipientUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationRecipientUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationRecipientUser.__pulumiType;
    }

    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * API Management UserId subscribed to notification.
     */
    public readonly userId!: pulumi.Output<string | undefined>;

    /**
     * Create a NotificationRecipientUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationRecipientUserArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.notificationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notificationName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["notificationName"] = args ? args.notificationName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["userId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:apimanagement/v20180601preview:NotificationRecipientUser" }, { type: "azure-native:apimanagement:NotificationRecipientUser" }, { type: "azure-nextgen:apimanagement:NotificationRecipientUser" }, { type: "azure-native:apimanagement/v20170301:NotificationRecipientUser" }, { type: "azure-nextgen:apimanagement/v20170301:NotificationRecipientUser" }, { type: "azure-native:apimanagement/v20180101:NotificationRecipientUser" }, { type: "azure-nextgen:apimanagement/v20180101:NotificationRecipientUser" }, { type: "azure-native:apimanagement/v20190101:NotificationRecipientUser" }, { type: "azure-nextgen:apimanagement/v20190101:NotificationRecipientUser" }, { type: "azure-native:apimanagement/v20191201:NotificationRecipientUser" }, { type: "azure-nextgen:apimanagement/v20191201:NotificationRecipientUser" }, { type: "azure-native:apimanagement/v20191201preview:NotificationRecipientUser" }, { type: "azure-nextgen:apimanagement/v20191201preview:NotificationRecipientUser" }, { type: "azure-native:apimanagement/v20200601preview:NotificationRecipientUser" }, { type: "azure-nextgen:apimanagement/v20200601preview:NotificationRecipientUser" }, { type: "azure-native:apimanagement/v20201201:NotificationRecipientUser" }, { type: "azure-nextgen:apimanagement/v20201201:NotificationRecipientUser" }, { type: "azure-native:apimanagement/v20210101preview:NotificationRecipientUser" }, { type: "azure-nextgen:apimanagement/v20210101preview:NotificationRecipientUser" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(NotificationRecipientUser.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NotificationRecipientUser resource.
 */
export interface NotificationRecipientUserArgs {
    /**
     * Notification Name Identifier.
     */
    readonly notificationName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * User identifier. Must be unique in the current API Management service instance.
     */
    readonly userId?: pulumi.Input<string>;
}
