// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Pageable list of products.
 * API Version: 2017-06-01.
 */
export function getProducts(args: GetProductsArgs, opts?: pulumi.InvokeOptions): Promise<GetProductsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:azurestack:getProducts", {
        "productName": args.productName,
        "registrationName": args.registrationName,
        "resourceGroup": args.resourceGroup,
    }, opts);
}

export interface GetProductsArgs {
    /**
     * Name of the product.
     */
    readonly productName: string;
    /**
     * Name of the Azure Stack registration.
     */
    readonly registrationName: string;
    /**
     * Name of the resource group.
     */
    readonly resourceGroup: string;
}

/**
 * Pageable list of products.
 */
export interface GetProductsResult {
    /**
     * URI to the next page.
     */
    readonly nextLink?: string;
    /**
     * List of products.
     */
    readonly value?: outputs.azurestack.ProductResponse[];
}
