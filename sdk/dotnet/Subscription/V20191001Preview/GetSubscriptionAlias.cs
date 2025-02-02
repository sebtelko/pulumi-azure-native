// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Subscription.V20191001Preview
{
    public static class GetSubscriptionAlias
    {
        /// <summary>
        /// Subscription Information with the alias.
        /// </summary>
        public static Task<GetSubscriptionAliasResult> InvokeAsync(GetSubscriptionAliasArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionAliasResult>("azure-native:subscription/v20191001preview:getSubscriptionAlias", args ?? new GetSubscriptionAliasArgs(), options.WithVersion());
    }


    public sealed class GetSubscriptionAliasArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Alias Name
        /// </summary>
        [Input("aliasName", required: true)]
        public string AliasName { get; set; } = null!;

        public GetSubscriptionAliasArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSubscriptionAliasResult
    {
        /// <summary>
        /// Fully qualified ID for the alias resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Alias ID.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Put Alias response properties.
        /// </summary>
        public readonly Outputs.PutAliasResponsePropertiesResponse Properties;
        /// <summary>
        /// Resource type, Microsoft.Subscription/aliases.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSubscriptionAliasResult(
            string id,

            string name,

            Outputs.PutAliasResponsePropertiesResponse properties,

            string type)
        {
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
