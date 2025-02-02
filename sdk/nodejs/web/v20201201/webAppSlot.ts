// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * A web app, a mobile app backend, or an API app.
 */
export class WebAppSlot extends pulumi.CustomResource {
    /**
     * Get an existing WebAppSlot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebAppSlot {
        return new WebAppSlot(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:web/v20201201:WebAppSlot';

    /**
     * Returns true if the given object is an instance of WebAppSlot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAppSlot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAppSlot.__pulumiType;
    }

    /**
     * Management information availability state for the app.
     */
    public /*out*/ readonly availabilityState!: pulumi.Output<string>;
    /**
     * <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
     */
    public readonly clientAffinityEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
     */
    public readonly clientCertEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * client certificate authentication comma-separated exclusion paths
     */
    public readonly clientCertExclusionPaths!: pulumi.Output<string | undefined>;
    /**
     * This composes with ClientCertEnabled setting.
     * - ClientCertEnabled: false means ClientCert is ignored.
     * - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required.
     * - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted.
     */
    public readonly clientCertMode!: pulumi.Output<string | undefined>;
    /**
     * Size of the function container.
     */
    public readonly containerSize!: pulumi.Output<number | undefined>;
    /**
     * Unique identifier that verifies the custom domains assigned to the app. Customer will add this id to a txt record for verification.
     */
    public readonly customDomainVerificationId!: pulumi.Output<string | undefined>;
    /**
     * Maximum allowed daily memory-time quota (applicable on dynamic apps only).
     */
    public readonly dailyMemoryTimeQuota!: pulumi.Output<number | undefined>;
    /**
     * Default hostname of the app. Read-only.
     */
    public /*out*/ readonly defaultHostName!: pulumi.Output<string>;
    /**
     * <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Enabled hostnames for the app.Hostnames need to be assigned (see HostNames) AND enabled. Otherwise,
     * the app is not served on those hostnames.
     */
    public /*out*/ readonly enabledHostNames!: pulumi.Output<string[]>;
    /**
     * Hostname SSL states are used to manage the SSL bindings for app's hostnames.
     */
    public readonly hostNameSslStates!: pulumi.Output<outputs.web.v20201201.HostNameSslStateResponse[] | undefined>;
    /**
     * Hostnames associated with the app.
     */
    public /*out*/ readonly hostNames!: pulumi.Output<string[]>;
    /**
     * <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
     *  If <code>true</code>, the app is only accessible via API management process.
     */
    public readonly hostNamesDisabled!: pulumi.Output<boolean | undefined>;
    /**
     * App Service Environment to use for the app.
     */
    public readonly hostingEnvironmentProfile!: pulumi.Output<outputs.web.v20201201.HostingEnvironmentProfileResponse | undefined>;
    /**
     * HttpsOnly: configures a web site to accept only https requests. Issues redirect for
     * http requests
     */
    public readonly httpsOnly!: pulumi.Output<boolean | undefined>;
    /**
     * Hyper-V sandbox.
     */
    public readonly hyperV!: pulumi.Output<boolean | undefined>;
    /**
     * Managed service identity.
     */
    public readonly identity!: pulumi.Output<outputs.web.v20201201.ManagedServiceIdentityResponse | undefined>;
    /**
     * Specifies an operation id if this site has a pending operation.
     */
    public /*out*/ readonly inProgressOperationId!: pulumi.Output<string>;
    /**
     * <code>true</code> if the app is a default container; otherwise, <code>false</code>.
     */
    public /*out*/ readonly isDefaultContainer!: pulumi.Output<boolean>;
    /**
     * Obsolete: Hyper-V sandbox.
     */
    public readonly isXenon!: pulumi.Output<boolean | undefined>;
    /**
     * Identity to use for Key Vault Reference authentication.
     */
    public readonly keyVaultReferenceIdentity!: pulumi.Output<string | undefined>;
    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Last time the app was modified, in UTC. Read-only.
     */
    public /*out*/ readonly lastModifiedTimeUtc!: pulumi.Output<string>;
    /**
     * Resource Location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Maximum number of workers.
     * This only applies to Functions container.
     */
    public /*out*/ readonly maxNumberOfWorkers!: pulumi.Output<number>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from tenants that site can be hosted with current settings. Read-only.
     */
    public /*out*/ readonly outboundIpAddresses!: pulumi.Output<string>;
    /**
     * List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from all tenants except dataComponent. Read-only.
     */
    public /*out*/ readonly possibleOutboundIpAddresses!: pulumi.Output<string>;
    /**
     * Site redundancy mode
     */
    public readonly redundancyMode!: pulumi.Output<string | undefined>;
    /**
     * Name of the repository site.
     */
    public /*out*/ readonly repositorySiteName!: pulumi.Output<string>;
    /**
     * <code>true</code> if reserved; otherwise, <code>false</code>.
     */
    public readonly reserved!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the resource group the app belongs to. Read-only.
     */
    public /*out*/ readonly resourceGroup!: pulumi.Output<string>;
    /**
     * <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
     */
    public readonly scmSiteAlsoStopped!: pulumi.Output<boolean | undefined>;
    /**
     * Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
     */
    public readonly serverFarmId!: pulumi.Output<string | undefined>;
    /**
     * Configuration of the app.
     */
    public readonly siteConfig!: pulumi.Output<outputs.web.v20201201.SiteConfigResponse | undefined>;
    /**
     * Status of the last deployment slot swap operation.
     */
    public /*out*/ readonly slotSwapStatus!: pulumi.Output<outputs.web.v20201201.SlotSwapStatusResponse>;
    /**
     * Current state of the app.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Checks if Customer provided storage account is required
     */
    public readonly storageAccountRequired!: pulumi.Output<boolean | undefined>;
    /**
     * App suspended till in case memory-time quota is exceeded.
     */
    public /*out*/ readonly suspendedTill!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies which deployment slot this app will swap into. Read-only.
     */
    public /*out*/ readonly targetSwapSlot!: pulumi.Output<string>;
    /**
     * Azure Traffic Manager hostnames associated with the app. Read-only.
     */
    public /*out*/ readonly trafficManagerHostNames!: pulumi.Output<string[]>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * State indicating whether the app has exceeded its quota usage. Read-only.
     */
    public /*out*/ readonly usageState!: pulumi.Output<string>;
    /**
     * Azure Resource Manager ID of the Virtual network and subnet to be joined by Regional VNET Integration.
     * This must be of the form /subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}
     */
    public readonly virtualNetworkSubnetId!: pulumi.Output<string | undefined>;

    /**
     * Create a WebAppSlot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppSlotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["clientAffinityEnabled"] = args ? args.clientAffinityEnabled : undefined;
            inputs["clientCertEnabled"] = args ? args.clientCertEnabled : undefined;
            inputs["clientCertExclusionPaths"] = args ? args.clientCertExclusionPaths : undefined;
            inputs["clientCertMode"] = args ? args.clientCertMode : undefined;
            inputs["cloningInfo"] = args ? args.cloningInfo : undefined;
            inputs["containerSize"] = args ? args.containerSize : undefined;
            inputs["customDomainVerificationId"] = args ? args.customDomainVerificationId : undefined;
            inputs["dailyMemoryTimeQuota"] = args ? args.dailyMemoryTimeQuota : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["hostNameSslStates"] = args ? args.hostNameSslStates : undefined;
            inputs["hostNamesDisabled"] = args ? args.hostNamesDisabled : undefined;
            inputs["hostingEnvironmentProfile"] = args ? args.hostingEnvironmentProfile : undefined;
            inputs["httpsOnly"] = args ? args.httpsOnly : undefined;
            inputs["hyperV"] = (args ? args.hyperV : undefined) ?? false;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["isXenon"] = (args ? args.isXenon : undefined) ?? false;
            inputs["keyVaultReferenceIdentity"] = args ? args.keyVaultReferenceIdentity : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["redundancyMode"] = args ? args.redundancyMode : undefined;
            inputs["reserved"] = (args ? args.reserved : undefined) ?? false;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scmSiteAlsoStopped"] = (args ? args.scmSiteAlsoStopped : undefined) ?? false;
            inputs["serverFarmId"] = args ? args.serverFarmId : undefined;
            inputs["siteConfig"] = args ? args.siteConfig : undefined;
            inputs["slot"] = args ? args.slot : undefined;
            inputs["storageAccountRequired"] = args ? args.storageAccountRequired : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualNetworkSubnetId"] = args ? args.virtualNetworkSubnetId : undefined;
            inputs["availabilityState"] = undefined /*out*/;
            inputs["defaultHostName"] = undefined /*out*/;
            inputs["enabledHostNames"] = undefined /*out*/;
            inputs["hostNames"] = undefined /*out*/;
            inputs["inProgressOperationId"] = undefined /*out*/;
            inputs["isDefaultContainer"] = undefined /*out*/;
            inputs["lastModifiedTimeUtc"] = undefined /*out*/;
            inputs["maxNumberOfWorkers"] = undefined /*out*/;
            inputs["outboundIpAddresses"] = undefined /*out*/;
            inputs["possibleOutboundIpAddresses"] = undefined /*out*/;
            inputs["repositorySiteName"] = undefined /*out*/;
            inputs["resourceGroup"] = undefined /*out*/;
            inputs["slotSwapStatus"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["suspendedTill"] = undefined /*out*/;
            inputs["targetSwapSlot"] = undefined /*out*/;
            inputs["trafficManagerHostNames"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["usageState"] = undefined /*out*/;
        } else {
            inputs["availabilityState"] = undefined /*out*/;
            inputs["clientAffinityEnabled"] = undefined /*out*/;
            inputs["clientCertEnabled"] = undefined /*out*/;
            inputs["clientCertExclusionPaths"] = undefined /*out*/;
            inputs["clientCertMode"] = undefined /*out*/;
            inputs["containerSize"] = undefined /*out*/;
            inputs["customDomainVerificationId"] = undefined /*out*/;
            inputs["dailyMemoryTimeQuota"] = undefined /*out*/;
            inputs["defaultHostName"] = undefined /*out*/;
            inputs["enabled"] = undefined /*out*/;
            inputs["enabledHostNames"] = undefined /*out*/;
            inputs["hostNameSslStates"] = undefined /*out*/;
            inputs["hostNames"] = undefined /*out*/;
            inputs["hostNamesDisabled"] = undefined /*out*/;
            inputs["hostingEnvironmentProfile"] = undefined /*out*/;
            inputs["httpsOnly"] = undefined /*out*/;
            inputs["hyperV"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["inProgressOperationId"] = undefined /*out*/;
            inputs["isDefaultContainer"] = undefined /*out*/;
            inputs["isXenon"] = undefined /*out*/;
            inputs["keyVaultReferenceIdentity"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["lastModifiedTimeUtc"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["maxNumberOfWorkers"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["outboundIpAddresses"] = undefined /*out*/;
            inputs["possibleOutboundIpAddresses"] = undefined /*out*/;
            inputs["redundancyMode"] = undefined /*out*/;
            inputs["repositorySiteName"] = undefined /*out*/;
            inputs["reserved"] = undefined /*out*/;
            inputs["resourceGroup"] = undefined /*out*/;
            inputs["scmSiteAlsoStopped"] = undefined /*out*/;
            inputs["serverFarmId"] = undefined /*out*/;
            inputs["siteConfig"] = undefined /*out*/;
            inputs["slotSwapStatus"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["storageAccountRequired"] = undefined /*out*/;
            inputs["suspendedTill"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["targetSwapSlot"] = undefined /*out*/;
            inputs["trafficManagerHostNames"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["usageState"] = undefined /*out*/;
            inputs["virtualNetworkSubnetId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:web/v20201201:WebAppSlot" }, { type: "azure-native:web:WebAppSlot" }, { type: "azure-nextgen:web:WebAppSlot" }, { type: "azure-native:web/v20150801:WebAppSlot" }, { type: "azure-nextgen:web/v20150801:WebAppSlot" }, { type: "azure-native:web/v20160801:WebAppSlot" }, { type: "azure-nextgen:web/v20160801:WebAppSlot" }, { type: "azure-native:web/v20180201:WebAppSlot" }, { type: "azure-nextgen:web/v20180201:WebAppSlot" }, { type: "azure-native:web/v20181101:WebAppSlot" }, { type: "azure-nextgen:web/v20181101:WebAppSlot" }, { type: "azure-native:web/v20190801:WebAppSlot" }, { type: "azure-nextgen:web/v20190801:WebAppSlot" }, { type: "azure-native:web/v20200601:WebAppSlot" }, { type: "azure-nextgen:web/v20200601:WebAppSlot" }, { type: "azure-native:web/v20200901:WebAppSlot" }, { type: "azure-nextgen:web/v20200901:WebAppSlot" }, { type: "azure-native:web/v20201001:WebAppSlot" }, { type: "azure-nextgen:web/v20201001:WebAppSlot" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(WebAppSlot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebAppSlot resource.
 */
export interface WebAppSlotArgs {
    /**
     * <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
     */
    readonly clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
     */
    readonly clientCertEnabled?: pulumi.Input<boolean>;
    /**
     * client certificate authentication comma-separated exclusion paths
     */
    readonly clientCertExclusionPaths?: pulumi.Input<string>;
    /**
     * This composes with ClientCertEnabled setting.
     * - ClientCertEnabled: false means ClientCert is ignored.
     * - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required.
     * - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted.
     */
    readonly clientCertMode?: pulumi.Input<enums.web.v20201201.ClientCertMode>;
    /**
     * If specified during app creation, the app is cloned from a source app.
     */
    readonly cloningInfo?: pulumi.Input<inputs.web.v20201201.CloningInfoArgs>;
    /**
     * Size of the function container.
     */
    readonly containerSize?: pulumi.Input<number>;
    /**
     * Unique identifier that verifies the custom domains assigned to the app. Customer will add this id to a txt record for verification.
     */
    readonly customDomainVerificationId?: pulumi.Input<string>;
    /**
     * Maximum allowed daily memory-time quota (applicable on dynamic apps only).
     */
    readonly dailyMemoryTimeQuota?: pulumi.Input<number>;
    /**
     * <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Hostname SSL states are used to manage the SSL bindings for app's hostnames.
     */
    readonly hostNameSslStates?: pulumi.Input<pulumi.Input<inputs.web.v20201201.HostNameSslStateArgs>[]>;
    /**
     * <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
     *  If <code>true</code>, the app is only accessible via API management process.
     */
    readonly hostNamesDisabled?: pulumi.Input<boolean>;
    /**
     * App Service Environment to use for the app.
     */
    readonly hostingEnvironmentProfile?: pulumi.Input<inputs.web.v20201201.HostingEnvironmentProfileArgs>;
    /**
     * HttpsOnly: configures a web site to accept only https requests. Issues redirect for
     * http requests
     */
    readonly httpsOnly?: pulumi.Input<boolean>;
    /**
     * Hyper-V sandbox.
     */
    readonly hyperV?: pulumi.Input<boolean>;
    /**
     * Managed service identity.
     */
    readonly identity?: pulumi.Input<inputs.web.v20201201.ManagedServiceIdentityArgs>;
    /**
     * Obsolete: Hyper-V sandbox.
     */
    readonly isXenon?: pulumi.Input<boolean>;
    /**
     * Identity to use for Key Vault Reference authentication.
     */
    readonly keyVaultReferenceIdentity?: pulumi.Input<string>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Unique name of the app to create or update. To create or update a deployment slot, use the {slot} parameter.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Site redundancy mode
     */
    readonly redundancyMode?: pulumi.Input<enums.web.v20201201.RedundancyMode>;
    /**
     * <code>true</code> if reserved; otherwise, <code>false</code>.
     */
    readonly reserved?: pulumi.Input<boolean>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
     */
    readonly scmSiteAlsoStopped?: pulumi.Input<boolean>;
    /**
     * Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
     */
    readonly serverFarmId?: pulumi.Input<string>;
    /**
     * Configuration of the app.
     */
    readonly siteConfig?: pulumi.Input<inputs.web.v20201201.SiteConfigArgs>;
    /**
     * Name of the deployment slot to create or update. By default, this API attempts to create or modify the production slot.
     */
    readonly slot?: pulumi.Input<string>;
    /**
     * Checks if Customer provided storage account is required
     */
    readonly storageAccountRequired?: pulumi.Input<boolean>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Azure Resource Manager ID of the Virtual network and subnet to be joined by Regional VNET Integration.
     * This must be of the form /subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}
     */
    readonly virtualNetworkSubnetId?: pulumi.Input<string>;
}
