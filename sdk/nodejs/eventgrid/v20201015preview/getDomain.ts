// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * EventGrid Domain.
 */
export function getDomain(args: GetDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:eventgrid/v20201015preview:getDomain", {
        "domainName": args.domainName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDomainArgs {
    /**
     * Name of the domain.
     */
    readonly domainName: string;
    /**
     * The name of the resource group within the user's subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * EventGrid Domain.
 */
export interface GetDomainResult {
    /**
     * Endpoint for the domain.
     */
    readonly endpoint: string;
    /**
     * Fully qualified identifier of the resource.
     */
    readonly id: string;
    /**
     * Identity information for the resource.
     */
    readonly identity?: outputs.eventgrid.v20201015preview.IdentityInfoResponse;
    /**
     * This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
     */
    readonly inboundIpRules?: outputs.eventgrid.v20201015preview.InboundIpRuleResponse[];
    /**
     * This determines the format that Event Grid should expect for incoming events published to the domain.
     */
    readonly inputSchema?: string;
    /**
     * Information about the InputSchemaMapping which specified the info about mapping event payload.
     */
    readonly inputSchemaMapping?: outputs.eventgrid.v20201015preview.JsonInputSchemaMappingResponse;
    /**
     * Location of the resource.
     */
    readonly location: string;
    /**
     * Metric resource id for the domain.
     */
    readonly metricResourceId: string;
    /**
     * Name of the resource.
     */
    readonly name: string;
    /**
     * List of private endpoint connections.
     */
    readonly privateEndpointConnections: outputs.eventgrid.v20201015preview.PrivateEndpointConnectionResponse[];
    /**
     * Provisioning state of the domain.
     */
    readonly provisioningState: string;
    /**
     * This determines if traffic is allowed over public network. By default it is enabled. 
     * You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
     */
    readonly publicNetworkAccess?: string;
    /**
     * The Sku pricing tier for the domain.
     */
    readonly sku?: outputs.eventgrid.v20201015preview.ResourceSkuResponse;
    /**
     * The system metadata relating to Domain resource.
     */
    readonly systemData: outputs.eventgrid.v20201015preview.SystemDataResponse;
    /**
     * Tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Type of the resource.
     */
    readonly type: string;
}
