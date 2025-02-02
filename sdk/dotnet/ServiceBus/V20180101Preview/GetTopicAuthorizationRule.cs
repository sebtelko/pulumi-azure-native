// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceBus.V20180101Preview
{
    public static class GetTopicAuthorizationRule
    {
        /// <summary>
        /// Description of a namespace authorization rule.
        /// </summary>
        public static Task<GetTopicAuthorizationRuleResult> InvokeAsync(GetTopicAuthorizationRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicAuthorizationRuleResult>("azure-native:servicebus/v20180101preview:getTopicAuthorizationRule", args ?? new GetTopicAuthorizationRuleArgs(), options.WithVersion());
    }


    public sealed class GetTopicAuthorizationRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The authorization rule name.
        /// </summary>
        [Input("authorizationRuleName", required: true)]
        public string AuthorizationRuleName { get; set; } = null!;

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The topic name.
        /// </summary>
        [Input("topicName", required: true)]
        public string TopicName { get; set; } = null!;

        public GetTopicAuthorizationRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTopicAuthorizationRuleResult
    {
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The rights associated with the rule.
        /// </summary>
        public readonly ImmutableArray<string> Rights;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTopicAuthorizationRuleResult(
            string id,

            string name,

            ImmutableArray<string> rights,

            string type)
        {
            Id = id;
            Name = name;
            Rights = rights;
            Type = type;
        }
    }
}
