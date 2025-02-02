// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Definition of the EnterprisePolicy.
 * API Version: 2020-10-30-preview.
 */
export function getEnterprisePolicy(args: GetEnterprisePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetEnterprisePolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:powerplatform:getEnterprisePolicy", {
        "enterprisePolicyName": args.enterprisePolicyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetEnterprisePolicyArgs {
    /**
     * The EnterprisePolicy name.
     */
    readonly enterprisePolicyName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Definition of the EnterprisePolicy.
 */
export interface GetEnterprisePolicyResult {
    /**
     * The encryption settings for a configuration store.
     */
    readonly encryption?: outputs.powerplatform.PropertiesResponseEncryption;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The identity of the EnterprisePolicy.
     */
    readonly identity?: outputs.powerplatform.EnterprisePolicyIdentityResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * Settings concerning lockbox.
     */
    readonly lockbox?: outputs.powerplatform.PropertiesResponseLockbox;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Settings concerning network injection.
     */
    readonly networkInjection?: outputs.powerplatform.PropertiesResponseNetworkInjection;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.powerplatform.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
