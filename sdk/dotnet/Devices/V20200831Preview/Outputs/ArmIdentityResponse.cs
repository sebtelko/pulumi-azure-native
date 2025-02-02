// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Devices.V20200831Preview.Outputs
{

    [OutputType]
    public sealed class ArmIdentityResponse
    {
        /// <summary>
        /// Identity type. Only allowed values are SystemAssigned and UserAssigned. Comma separated if both for ex: SystemAssigned,UserAssigned
        /// </summary>
        public readonly string? IdentityType;
        /// <summary>
        /// Principal Id
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// Tenant Id
        /// </summary>
        public readonly string TenantId;
        public readonly ImmutableDictionary<string, Outputs.ArmUserIdentityResponse>? UserAssignedIdentities;

        [OutputConstructor]
        private ArmIdentityResponse(
            string? identityType,

            string principalId,

            string tenantId,

            ImmutableDictionary<string, Outputs.ArmUserIdentityResponse>? userAssignedIdentities)
        {
            IdentityType = identityType;
            PrincipalId = principalId;
            TenantId = tenantId;
            UserAssignedIdentities = userAssignedIdentities;
        }
    }
}
