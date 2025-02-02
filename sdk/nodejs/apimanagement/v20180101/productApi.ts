// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * API details.
 */
export class ProductApi extends pulumi.CustomResource {
    /**
     * Get an existing ProductApi resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ProductApi {
        return new ProductApi(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:apimanagement/v20180101:ProductApi';

    /**
     * Returns true if the given object is an instance of ProductApi.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProductApi {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProductApi.__pulumiType;
    }

    /**
     * Describes the Revision of the Api. If no value is provided, default revision 1 is created
     */
    public /*out*/ readonly apiRevision!: pulumi.Output<string | undefined>;
    /**
     * Description of the Api Revision.
     */
    public /*out*/ readonly apiRevisionDescription!: pulumi.Output<string | undefined>;
    /**
     * Type of API.
     */
    public /*out*/ readonly apiType!: pulumi.Output<string | undefined>;
    /**
     * Indicates the Version identifier of the API if the API is versioned
     */
    public /*out*/ readonly apiVersion!: pulumi.Output<string | undefined>;
    /**
     * Description of the Api Version.
     */
    public /*out*/ readonly apiVersionDescription!: pulumi.Output<string | undefined>;
    /**
     * An API Version Set contains the common configuration for a set of API Versions relating 
     */
    public /*out*/ readonly apiVersionSet!: pulumi.Output<outputs.apimanagement.v20180101.ApiVersionSetContractDetailsResponse | undefined>;
    /**
     * A resource identifier for the related ApiVersionSet.
     */
    public /*out*/ readonly apiVersionSetId!: pulumi.Output<string | undefined>;
    /**
     * Collection of authentication settings included into this API.
     */
    public /*out*/ readonly authenticationSettings!: pulumi.Output<outputs.apimanagement.v20180101.AuthenticationSettingsContractResponse | undefined>;
    /**
     * Description of the API. May include HTML formatting tags.
     */
    public /*out*/ readonly description!: pulumi.Output<string | undefined>;
    /**
     * API name.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Indicates if API revision is current api revision.
     */
    public /*out*/ readonly isCurrent!: pulumi.Output<boolean>;
    /**
     * Indicates if API revision is accessible via the gateway.
     */
    public /*out*/ readonly isOnline!: pulumi.Output<boolean>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
     */
    public /*out*/ readonly path!: pulumi.Output<string>;
    /**
     * Describes on which protocols the operations in this API can be invoked.
     */
    public /*out*/ readonly protocols!: pulumi.Output<string[] | undefined>;
    /**
     * Absolute URL of the backend service implementing this API.
     */
    public /*out*/ readonly serviceUrl!: pulumi.Output<string | undefined>;
    /**
     * Protocols over which API is made available.
     */
    public /*out*/ readonly subscriptionKeyParameterNames!: pulumi.Output<outputs.apimanagement.v20180101.SubscriptionKeyParameterNamesContractResponse | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ProductApi resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductApiArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.productId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["productId"] = args ? args.productId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["apiRevision"] = undefined /*out*/;
            inputs["apiRevisionDescription"] = undefined /*out*/;
            inputs["apiType"] = undefined /*out*/;
            inputs["apiVersion"] = undefined /*out*/;
            inputs["apiVersionDescription"] = undefined /*out*/;
            inputs["apiVersionSet"] = undefined /*out*/;
            inputs["apiVersionSetId"] = undefined /*out*/;
            inputs["authenticationSettings"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["isCurrent"] = undefined /*out*/;
            inputs["isOnline"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["path"] = undefined /*out*/;
            inputs["protocols"] = undefined /*out*/;
            inputs["serviceUrl"] = undefined /*out*/;
            inputs["subscriptionKeyParameterNames"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["apiRevision"] = undefined /*out*/;
            inputs["apiRevisionDescription"] = undefined /*out*/;
            inputs["apiType"] = undefined /*out*/;
            inputs["apiVersion"] = undefined /*out*/;
            inputs["apiVersionDescription"] = undefined /*out*/;
            inputs["apiVersionSet"] = undefined /*out*/;
            inputs["apiVersionSetId"] = undefined /*out*/;
            inputs["authenticationSettings"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["isCurrent"] = undefined /*out*/;
            inputs["isOnline"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["path"] = undefined /*out*/;
            inputs["protocols"] = undefined /*out*/;
            inputs["serviceUrl"] = undefined /*out*/;
            inputs["subscriptionKeyParameterNames"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:apimanagement/v20180101:ProductApi" }, { type: "azure-native:apimanagement:ProductApi" }, { type: "azure-nextgen:apimanagement:ProductApi" }, { type: "azure-native:apimanagement/v20170301:ProductApi" }, { type: "azure-nextgen:apimanagement/v20170301:ProductApi" }, { type: "azure-native:apimanagement/v20180601preview:ProductApi" }, { type: "azure-nextgen:apimanagement/v20180601preview:ProductApi" }, { type: "azure-native:apimanagement/v20190101:ProductApi" }, { type: "azure-nextgen:apimanagement/v20190101:ProductApi" }, { type: "azure-native:apimanagement/v20191201:ProductApi" }, { type: "azure-nextgen:apimanagement/v20191201:ProductApi" }, { type: "azure-native:apimanagement/v20191201preview:ProductApi" }, { type: "azure-nextgen:apimanagement/v20191201preview:ProductApi" }, { type: "azure-native:apimanagement/v20200601preview:ProductApi" }, { type: "azure-nextgen:apimanagement/v20200601preview:ProductApi" }, { type: "azure-native:apimanagement/v20201201:ProductApi" }, { type: "azure-nextgen:apimanagement/v20201201:ProductApi" }, { type: "azure-native:apimanagement/v20210101preview:ProductApi" }, { type: "azure-nextgen:apimanagement/v20210101preview:ProductApi" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ProductApi.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ProductApi resource.
 */
export interface ProductApiArgs {
    /**
     * API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
     */
    readonly apiId?: pulumi.Input<string>;
    /**
     * Product identifier. Must be unique in the current API Management service instance.
     */
    readonly productId: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
}
