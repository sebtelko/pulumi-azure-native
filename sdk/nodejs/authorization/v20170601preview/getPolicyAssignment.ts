// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The policy assignment.
 */
export function getPolicyAssignment(args: GetPolicyAssignmentArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyAssignmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:authorization/v20170601preview:getPolicyAssignment", {
        "policyAssignmentName": args.policyAssignmentName,
        "scope": args.scope,
    }, opts);
}

export interface GetPolicyAssignmentArgs {
    /**
     * The name of the policy assignment to get.
     */
    readonly policyAssignmentName: string;
    /**
     * The scope of the policy assignment.
     */
    readonly scope: string;
}

/**
 * The policy assignment.
 */
export interface GetPolicyAssignmentResult {
    /**
     * This message will be part of response in case of policy violation.
     */
    readonly description?: string;
    /**
     * The display name of the policy assignment.
     */
    readonly displayName?: string;
    /**
     * The ID of the policy assignment.
     */
    readonly id: string;
    /**
     * The policy assignment metadata.
     */
    readonly metadata?: any;
    /**
     * The name of the policy assignment.
     */
    readonly name: string;
    /**
     * The policy's excluded scopes.
     */
    readonly notScopes?: string[];
    /**
     * Required if a parameter is used in policy rule.
     */
    readonly parameters?: any;
    /**
     * The ID of the policy definition.
     */
    readonly policyDefinitionId?: string;
    /**
     * The scope for the policy assignment.
     */
    readonly scope?: string;
    /**
     * The policy sku.
     */
    readonly sku?: outputs.authorization.v20170601preview.PolicySkuResponse;
    /**
     * The type of the policy assignment.
     */
    readonly type: string;
}
