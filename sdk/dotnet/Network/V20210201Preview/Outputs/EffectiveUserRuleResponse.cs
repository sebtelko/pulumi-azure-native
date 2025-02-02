// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20210201Preview.Outputs
{

    [OutputType]
    public sealed class EffectiveUserRuleResponse
    {
        /// <summary>
        /// A description for this rule.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The destination port ranges.
        /// </summary>
        public readonly ImmutableArray<string> DestinationPortRanges;
        /// <summary>
        /// The destination address prefixes. CIDR or destination IP ranges.
        /// </summary>
        public readonly ImmutableArray<Outputs.AddressPrefixItemResponse> Destinations;
        /// <summary>
        /// Indicates if the traffic matched against the rule in inbound or outbound.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// A friendly name for the rule.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Whether the rule collection is custom or default.
        /// Expected value is 'Custom'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Network protocol this rule applies to.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The provisioning state of the security configuration user rule resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The source port ranges.
        /// </summary>
        public readonly ImmutableArray<string> SourcePortRanges;
        /// <summary>
        /// The CIDR or source IP ranges.
        /// </summary>
        public readonly ImmutableArray<Outputs.AddressPrefixItemResponse> Sources;

        [OutputConstructor]
        private EffectiveUserRuleResponse(
            string? description,

            ImmutableArray<string> destinationPortRanges,

            ImmutableArray<Outputs.AddressPrefixItemResponse> destinations,

            string direction,

            string? displayName,

            string? id,

            string kind,

            string protocol,

            string provisioningState,

            ImmutableArray<string> sourcePortRanges,

            ImmutableArray<Outputs.AddressPrefixItemResponse> sources)
        {
            Description = description;
            DestinationPortRanges = destinationPortRanges;
            Destinations = destinations;
            Direction = direction;
            DisplayName = displayName;
            Id = id;
            Kind = kind;
            Protocol = protocol;
            ProvisioningState = provisioningState;
            SourcePortRanges = sourcePortRanges;
            Sources = sources;
        }
    }
}
