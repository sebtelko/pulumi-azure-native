// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The cost allocation rule model definition
 * API Version: 2020-03-01-preview.
 */
export function getCostAllocationRule(args: GetCostAllocationRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetCostAllocationRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:costmanagement:getCostAllocationRule", {
        "billingAccountId": args.billingAccountId,
        "ruleName": args.ruleName,
    }, opts);
}

export interface GetCostAllocationRuleArgs {
    /**
     * BillingAccount ID
     */
    readonly billingAccountId: string;
    /**
     * Cost allocation rule name. The name cannot include spaces or any non alphanumeric characters other than '_' and '-'. The max length is 260 characters.
     */
    readonly ruleName: string;
}

/**
 * The cost allocation rule model definition
 */
export interface GetCostAllocationRuleResult {
    /**
     * Azure Resource Manager Id for the rule. This is a read ony value.
     */
    readonly id: string;
    /**
     * Name of the rule. This is a read only value.
     */
    readonly name: string;
    /**
     * Cost allocation rule properties
     */
    readonly properties: outputs.costmanagement.CostAllocationRulePropertiesResponse;
    /**
     * Resource type of the rule. This is a read only value of Microsoft.CostManagement/CostAllocationRule.
     */
    readonly type: string;
}
