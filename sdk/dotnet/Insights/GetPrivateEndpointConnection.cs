// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Insights
{
    public static class GetPrivateEndpointConnection
    {
        /// <summary>
        /// A private endpoint connection
        /// API Version: 2019-10-17-preview.
        /// </summary>
        public static Task<GetPrivateEndpointConnectionResult> InvokeAsync(GetPrivateEndpointConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrivateEndpointConnectionResult>("azure-native:insights:getPrivateEndpointConnection", args ?? new GetPrivateEndpointConnectionArgs(), options.WithVersion());
    }


    public sealed class GetPrivateEndpointConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the private endpoint connection.
        /// </summary>
        [Input("privateEndpointConnectionName", required: true)]
        public string PrivateEndpointConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure Monitor PrivateLinkScope resource.
        /// </summary>
        [Input("scopeName", required: true)]
        public string ScopeName { get; set; } = null!;

        public GetPrivateEndpointConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrivateEndpointConnectionResult
    {
        /// <summary>
        /// Azure resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Azure resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Private endpoint which the connection belongs to.
        /// </summary>
        public readonly Outputs.PrivateEndpointPropertyResponse? PrivateEndpoint;
        /// <summary>
        /// Connection state of the private endpoint connection.
        /// </summary>
        public readonly Outputs.PrivateLinkServiceConnectionStatePropertyResponse? PrivateLinkServiceConnectionState;
        /// <summary>
        /// State of the private endpoint connection.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateEndpointConnectionResult(
            string id,

            string name,

            Outputs.PrivateEndpointPropertyResponse? privateEndpoint,

            Outputs.PrivateLinkServiceConnectionStatePropertyResponse? privateLinkServiceConnectionState,

            string provisioningState,

            string type)
        {
            Id = id;
            Name = name;
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
