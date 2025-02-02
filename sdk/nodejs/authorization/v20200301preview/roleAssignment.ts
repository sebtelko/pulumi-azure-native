// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
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
    public static readonly __pulumiType = 'azure-native:authorization/v20200301preview:RoleAssignment';

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
     * The delegation flag used for creating a role assignment
     */
    public readonly canDelegate!: pulumi.Output<boolean | undefined>;
    /**
     * The conditions on the role assignment. This limits the resources it can be assigned to. e.g.: @Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'
     */
    public readonly condition!: pulumi.Output<string | undefined>;
    /**
     * Version of the condition. Currently accepted value is '2.0'
     */
    public readonly conditionVersion!: pulumi.Output<string | undefined>;
    /**
     * Id of the user who created the assignment
     */
    public /*out*/ readonly createdBy!: pulumi.Output<string>;
    /**
     * Time it was created
     */
    public /*out*/ readonly createdOn!: pulumi.Output<string>;
    /**
     * Id of the delegated managed identity resource
     */
    public readonly delegatedManagedIdentityResourceId!: pulumi.Output<string | undefined>;
    /**
     * The role assignment name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The principal ID.
     */
    public readonly principalId!: pulumi.Output<string>;
    /**
     * The principal type of the assigned principal ID.
     */
    public readonly principalType!: pulumi.Output<string | undefined>;
    /**
     * The role definition ID.
     */
    public readonly roleDefinitionId!: pulumi.Output<string>;
    /**
     * The role assignment scope.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * The role assignment type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Id of the user who updated the assignment
     */
    public /*out*/ readonly updatedBy!: pulumi.Output<string>;
    /**
     * Time it was updated
     */
    public /*out*/ readonly updatedOn!: pulumi.Output<string>;

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
            inputs["condition"] = args ? args.condition : undefined;
            inputs["conditionVersion"] = args ? args.conditionVersion : undefined;
            inputs["delegatedManagedIdentityResourceId"] = args ? args.delegatedManagedIdentityResourceId : undefined;
            inputs["principalId"] = args ? args.principalId : undefined;
            inputs["principalType"] = (args ? args.principalType : undefined) ?? "User";
            inputs["roleAssignmentName"] = args ? args.roleAssignmentName : undefined;
            inputs["roleDefinitionId"] = args ? args.roleDefinitionId : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["createdBy"] = undefined /*out*/;
            inputs["createdOn"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedBy"] = undefined /*out*/;
            inputs["updatedOn"] = undefined /*out*/;
        } else {
            inputs["canDelegate"] = undefined /*out*/;
            inputs["condition"] = undefined /*out*/;
            inputs["conditionVersion"] = undefined /*out*/;
            inputs["createdBy"] = undefined /*out*/;
            inputs["createdOn"] = undefined /*out*/;
            inputs["delegatedManagedIdentityResourceId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["principalId"] = undefined /*out*/;
            inputs["principalType"] = undefined /*out*/;
            inputs["roleDefinitionId"] = undefined /*out*/;
            inputs["scope"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedBy"] = undefined /*out*/;
            inputs["updatedOn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:authorization/v20200301preview:RoleAssignment" }, { type: "azure-native:authorization:RoleAssignment" }, { type: "azure-nextgen:authorization:RoleAssignment" }, { type: "azure-native:authorization/v20150701:RoleAssignment" }, { type: "azure-nextgen:authorization/v20150701:RoleAssignment" }, { type: "azure-native:authorization/v20171001preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20171001preview:RoleAssignment" }, { type: "azure-native:authorization/v20180101preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20180101preview:RoleAssignment" }, { type: "azure-native:authorization/v20180901preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20180901preview:RoleAssignment" }, { type: "azure-native:authorization/v20200401preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20200401preview:RoleAssignment" }, { type: "azure-native:authorization/v20200801preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20200801preview:RoleAssignment" }, { type: "azure-native:authorization/v20201001preview:RoleAssignment" }, { type: "azure-nextgen:authorization/v20201001preview:RoleAssignment" }] };
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
     * The conditions on the role assignment. This limits the resources it can be assigned to. e.g.: @Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'
     */
    readonly condition?: pulumi.Input<string>;
    /**
     * Version of the condition. Currently accepted value is '2.0'
     */
    readonly conditionVersion?: pulumi.Input<string>;
    /**
     * Id of the delegated managed identity resource
     */
    readonly delegatedManagedIdentityResourceId?: pulumi.Input<string>;
    /**
     * The principal ID.
     */
    readonly principalId: pulumi.Input<string>;
    /**
     * The principal type of the assigned principal ID.
     */
    readonly principalType?: pulumi.Input<string | enums.authorization.v20200301preview.PrincipalType>;
    /**
     * The name of the role assignment. It can be any valid GUID.
     */
    readonly roleAssignmentName?: pulumi.Input<string>;
    /**
     * The role definition ID.
     */
    readonly roleDefinitionId: pulumi.Input<string>;
    /**
     * The role assignment scope.
     */
    readonly scope: pulumi.Input<string>;
}
