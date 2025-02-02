// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Role Assignments
 */
export class RoleAssignment extends pulumi.CustomResource {
    /**
     * Get an existing RoleAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RoleAssignment {
        return new RoleAssignment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:authorization/v20180101preview:RoleAssignment';

    /**
     * Returns true if the given object is an instance of RoleAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoleAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoleAssignment.__pulumiType;
    }

    /**
     * The Delegation flag for the role assignment
     */
    public readonly canDelegate!: pulumi.Output<boolean | undefined>;
    /**
     * The role assignment name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The principal ID.
     */
    public readonly principalId!: pulumi.Output<string | undefined>;
    /**
     * The role definition ID.
     */
    public readonly roleDefinitionId!: pulumi.Output<string | undefined>;
    /**
     * The role assignment scope.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * The role assignment type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RoleAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleAssignmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.principalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principalId'");
            }
            if ((!args || args.roleDefinitionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleDefinitionId'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["canDelegate"] = args ? args.canDelegate : undefined;
            inputs["principalId"] = args ? args.principalId : undefined;
            inputs["roleAssignmentName"] = args ? args.roleAssignmentName : undefined;
            inputs["roleDefinitionId"] = args ? args.roleDefinitionId : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["canDelegate"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["principalId"] = undefined /*out*/;
            inputs["roleDefinitionId"] = undefined /*out*/;
            inputs["scope"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:authorization/v20180101preview:RoleAssignment" }, { type: "azure-native:authorization:RoleAssignment" }, { type: "azure-nextgen:authorization:RoleAssignment" }, { type: "azure-native:authorization/v20150701:RoleAssignment" }, { type: "azure-nextgen:authorization/v20150701:RoleAssignment" }, { type: "azure-native:authorization/v20171001preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20171001preview:RoleAssignment" }, { type: "azure-native:authorization/v20180901preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20180901preview:RoleAssignment" }, { type: "azure-native:authorization/v20200301preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20200301preview:RoleAssignment" }, { type: "azure-native:authorization/v20200401preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20200401preview:RoleAssignment" }, { type: "azure-native:authorization/v20200801preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20200801preview:RoleAssignment" }, { type: "azure-native:authorization/v20201001preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20201001preview:RoleAssignment" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(RoleAssignment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RoleAssignment resource.
 */
export interface RoleAssignmentArgs {
    /**
     * The delegation flag used for creating a role assignment
     */
    readonly canDelegate?: pulumi.Input<boolean>;
    /**
     * The principal ID assigned to the role. This maps to the ID inside the Active Directory. It can point to a user, service principal, or security group.
     */
    readonly principalId: pulumi.Input<string>;
    /**
     * The name of the role assignment to create. It can be any valid GUID.
     */
    readonly roleAssignmentName?: pulumi.Input<string>;
    /**
     * The role definition ID used in the role assignment.
     */
    readonly roleDefinitionId: pulumi.Input<string>;
    /**
     * The scope of the role assignment to create. The scope can be any REST resource instance. For example, use '/subscriptions/{subscription-id}/' for a subscription, '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}' for a resource.
     */
    readonly scope: pulumi.Input<string>;
}
