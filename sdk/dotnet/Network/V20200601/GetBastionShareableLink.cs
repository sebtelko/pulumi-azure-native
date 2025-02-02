// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20200601
{
    public static class GetBastionShareableLink
    {
        /// <summary>
        /// Response for all the Bastion Shareable Link endpoints.
        /// </summary>
        public static Task<GetBastionShareableLinkResult> InvokeAsync(GetBastionShareableLinkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBastionShareableLinkResult>("azure-native:network/v20200601:getBastionShareableLink", args ?? new GetBastionShareableLinkArgs(), options.WithVersion());
    }


    public sealed class GetBastionShareableLinkArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Bastion Host.
        /// </summary>
        [Input("bastionHostName", required: true)]
        public string BastionHostName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("vms")]
        private List<Inputs.BastionShareableLink>? _vms;

        /// <summary>
        /// List of VM references.
        /// </summary>
        public List<Inputs.BastionShareableLink> Vms
        {
            get => _vms ?? (_vms = new List<Inputs.BastionShareableLink>());
            set => _vms = value;
        }

        public GetBastionShareableLinkArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBastionShareableLinkResult
    {
        /// <summary>
        /// The URL to get the next set of results.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// List of Bastion Shareable Links for the request.
        /// </summary>
        public readonly ImmutableArray<Outputs.BastionShareableLinkResponse> Value;

        [OutputConstructor]
        private GetBastionShareableLinkResult(
            string? nextLink,

            ImmutableArray<Outputs.BastionShareableLinkResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
