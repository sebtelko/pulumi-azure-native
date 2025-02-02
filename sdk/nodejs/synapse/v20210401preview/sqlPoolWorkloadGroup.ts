// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Workload group operations for a sql pool
 */
export class SqlPoolWorkloadGroup extends pulumi.CustomResource {
    /**
     * Get an existing SqlPoolWorkloadGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SqlPoolWorkloadGroup {
        return new SqlPoolWorkloadGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:synapse/v20210401preview:SqlPoolWorkloadGroup';

    /**
     * Returns true if the given object is an instance of SqlPoolWorkloadGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SqlPoolWorkloadGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SqlPoolWorkloadGroup.__pulumiType;
    }

    /**
     * The workload group importance level.
     */
    public readonly importance!: pulumi.Output<string | undefined>;
    /**
     * The workload group cap percentage resource.
     */
    public readonly maxResourcePercent!: pulumi.Output<number>;
    /**
     * The workload group request maximum grant percentage.
     */
    public readonly maxResourcePercentPerRequest!: pulumi.Output<number | undefined>;
    /**
     * The workload group minimum percentage resource.
     */
    public readonly minResourcePercent!: pulumi.Output<number>;
    /**
     * The workload group request minimum grant percentage.
     */
    public readonly minResourcePercentPerRequest!: pulumi.Output<number>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The workload group query execution timeout.
     */
    public readonly queryExecutionTimeout!: pulumi.Output<number | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a SqlPoolWorkloadGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SqlPoolWorkloadGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.maxResourcePercent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxResourcePercent'");
            }
            if ((!args || args.minResourcePercent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'minResourcePercent'");
            }
            if ((!args || args.minResourcePercentPerRequest === undefined) && !opts.urn) {
                throw new Error("Missing required property 'minResourcePercentPerRequest'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sqlPoolName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sqlPoolName'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["importance"] = args ? args.importance : undefined;
            inputs["maxResourcePercent"] = args ? args.maxResourcePercent : undefined;
            inputs["maxResourcePercentPerRequest"] = args ? args.maxResourcePercentPerRequest : undefined;
            inputs["minResourcePercent"] = args ? args.minResourcePercent : undefined;
            inputs["minResourcePercentPerRequest"] = args ? args.minResourcePercentPerRequest : undefined;
            inputs["queryExecutionTimeout"] = args ? args.queryExecutionTimeout : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sqlPoolName"] = args ? args.sqlPoolName : undefined;
            inputs["workloadGroupName"] = args ? args.workloadGroupName : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["importance"] = undefined /*out*/;
            inputs["maxResourcePercent"] = undefined /*out*/;
            inputs["maxResourcePercentPerRequest"] = undefined /*out*/;
            inputs["minResourcePercent"] = undefined /*out*/;
            inputs["minResourcePercentPerRequest"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["queryExecutionTimeout"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:synapse/v20210401preview:SqlPoolWorkloadGroup" }, { type: "azure-native:synapse:SqlPoolWorkloadGroup" }, { type: "azure-nextgen:synapse:SqlPoolWorkloadGroup" }, { type: "azure-native:synapse/v20190601preview:SqlPoolWorkloadGroup" }, { type: "azure-nextgen:synapse/v20190601preview:SqlPoolWorkloadGroup" }, { type: "azure-native:synapse/v20201201:SqlPoolWorkloadGroup" }, { type: "azure-nextgen:synapse/v20201201:SqlPoolWorkloadGroup" }, { type: "azure-native:synapse/v20210301:SqlPoolWorkloadGroup" }, { type: "azure-nextgen:synapse/v20210301:SqlPoolWorkloadGroup" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(SqlPoolWorkloadGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SqlPoolWorkloadGroup resource.
 */
export interface SqlPoolWorkloadGroupArgs {
    /**
     * The workload group importance level.
     */
    readonly importance?: pulumi.Input<string>;
    /**
     * The workload group cap percentage resource.
     */
    readonly maxResourcePercent: pulumi.Input<number>;
    /**
     * The workload group request maximum grant percentage.
     */
    readonly maxResourcePercentPerRequest?: pulumi.Input<number>;
    /**
     * The workload group minimum percentage resource.
     */
    readonly minResourcePercent: pulumi.Input<number>;
    /**
     * The workload group request minimum grant percentage.
     */
    readonly minResourcePercentPerRequest: pulumi.Input<number>;
    /**
     * The workload group query execution timeout.
     */
    readonly queryExecutionTimeout?: pulumi.Input<number>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * SQL pool name
     */
    readonly sqlPoolName: pulumi.Input<string>;
    /**
     * The name of the workload group.
     */
    readonly workloadGroupName?: pulumi.Input<string>;
    /**
     * The name of the workspace
     */
    readonly workspaceName: pulumi.Input<string>;
}
