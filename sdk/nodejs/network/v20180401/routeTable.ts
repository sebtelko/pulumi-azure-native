// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Route table resource.
 */
export class RouteTable extends pulumi.CustomResource {
    /**
     * Get an existing RouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RouteTable {
        return new RouteTable(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:network/v20180401:RouteTable';

    /**
     * Returns true if the given object is an instance of RouteTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteTable.__pulumiType;
    }

    /**
     * Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
     */
    public readonly disableBgpRoutePropagation!: pulumi.Output<boolean | undefined>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Collection of routes contained within a route table.
     */
    public readonly routes!: pulumi.Output<outputs.network.v20180401.RouteResponse[] | undefined>;
    /**
     * A collection of references to subnets.
     */
    public /*out*/ readonly subnets!: pulumi.Output<outputs.network.v20180401.SubnetResponse[]>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteTableArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["disableBgpRoutePropagation"] = args ? args.disableBgpRoutePropagation : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routeTableName"] = args ? args.routeTableName : undefined;
            inputs["routes"] = args ? args.routes : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["subnets"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["disableBgpRoutePropagation"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["routes"] = undefined /*out*/;
            inputs["subnets"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:network/v20180401:RouteTable" }, { type: "azure-native:network:RouteTable" }, { type: "azure-nextgen:network:RouteTable" }, { type: "azure-native:network/v20150501preview:RouteTable" }, { type: "azure-nextgen:network/v20150501preview:RouteTable" }, { type: "azure-native:network/v20150615:RouteTable" }, { type: "azure-nextgen:network/v20150615:RouteTable" }, { type: "azure-native:network/v20160330:RouteTable" }, { type: "azure-nextgen:network/v20160330:RouteTable" }, { type: "azure-native:network/v20160601:RouteTable" }, { type: "azure-nextgen:network/v20160601:RouteTable" }, { type: "azure-native:network/v20160901:RouteTable" }, { type: "azure-nextgen:network/v20160901:RouteTable" }, { type: "azure-native:network/v20161201:RouteTable" }, { type: "azure-nextgen:network/v20161201:RouteTable" }, { type: "azure-native:network/v20170301:RouteTable" }, { type: "azure-nextgen:network/v20170301:RouteTable" }, { type: "azure-native:network/v20170601:RouteTable" }, { type: "azure-nextgen:network/v20170601:RouteTable" }, { type: "azure-native:network/v20170801:RouteTable" }, { type: "azure-nextgen:network/v20170801:RouteTable" }, { type: "azure-native:network/v20170901:RouteTable" }, { type: "azure-nextgen:network/v20170901:RouteTable" }, { type: "azure-native:network/v20171001:RouteTable" }, { type: "azure-nextgen:network/v20171001:RouteTable" }, { type: "azure-native:network/v20171101:RouteTable" }, { type: "azure-nextgen:network/v20171101:RouteTable" }, { type: "azure-native:network/v20180101:RouteTable" }, { type: "azure-nextgen:network/v20180101:RouteTable" }, { type: "azure-native:network/v20180201:RouteTable" }, { type: "azure-nextgen:network/v20180201:RouteTable" }, { type: "azure-native:network/v20180601:RouteTable" }, { type: "azure-nextgen:network/v20180601:RouteTable" }, { type: "azure-native:network/v20180701:RouteTable" }, { type: "azure-nextgen:network/v20180701:RouteTable" }, { type: "azure-native:network/v20180801:RouteTable" }, { type: "azure-nextgen:network/v20180801:RouteTable" }, { type: "azure-native:network/v20181001:RouteTable" }, { type: "azure-nextgen:network/v20181001:RouteTable" }, { type: "azure-native:network/v20181101:RouteTable" }, { type: "azure-nextgen:network/v20181101:RouteTable" }, { type: "azure-native:network/v20181201:RouteTable" }, { type: "azure-nextgen:network/v20181201:RouteTable" }, { type: "azure-native:network/v20190201:RouteTable" }, { type: "azure-nextgen:network/v20190201:RouteTable" }, { type: "azure-native:network/v20190401:RouteTable" }, { type: "azure-nextgen:network/v20190401:RouteTable" }, { type: "azure-native:network/v20190601:RouteTable" }, { type: "azure-nextgen:network/v20190601:RouteTable" }, { type: "azure-native:network/v20190701:RouteTable" }, { type: "azure-nextgen:network/v20190701:RouteTable" }, { type: "azure-native:network/v20190801:RouteTable" }, { type: "azure-nextgen:network/v20190801:RouteTable" }, { type: "azure-native:network/v20190901:RouteTable" }, { type: "azure-nextgen:network/v20190901:RouteTable" }, { type: "azure-native:network/v20191101:RouteTable" }, { type: "azure-nextgen:network/v20191101:RouteTable" }, { type: "azure-native:network/v20191201:RouteTable" }, { type: "azure-nextgen:network/v20191201:RouteTable" }, { type: "azure-native:network/v20200301:RouteTable" }, { type: "azure-nextgen:network/v20200301:RouteTable" }, { type: "azure-native:network/v20200401:RouteTable" }, { type: "azure-nextgen:network/v20200401:RouteTable" }, { type: "azure-native:network/v20200501:RouteTable" }, { type: "azure-nextgen:network/v20200501:RouteTable" }, { type: "azure-native:network/v20200601:RouteTable" }, { type: "azure-nextgen:network/v20200601:RouteTable" }, { type: "azure-native:network/v20200701:RouteTable" }, { type: "azure-nextgen:network/v20200701:RouteTable" }, { type: "azure-native:network/v20200801:RouteTable" }, { type: "azure-nextgen:network/v20200801:RouteTable" }, { type: "azure-native:network/v20201101:RouteTable" }, { type: "azure-nextgen:network/v20201101:RouteTable" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(RouteTable.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RouteTable resource.
 */
export interface RouteTableArgs {
    /**
     * Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
     */
    readonly disableBgpRoutePropagation?: pulumi.Input<boolean>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the route table.
     */
    readonly routeTableName?: pulumi.Input<string>;
    /**
     * Collection of routes contained within a route table.
     */
    readonly routes?: pulumi.Input<pulumi.Input<inputs.network.v20180401.RouteArgs>[]>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
