// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security
{
    public static class GetIotDefenderSetting
    {
        /// <summary>
        /// IoT Defender settings
        /// API Version: 2020-08-06-preview.
        /// </summary>
        public static Task<GetIotDefenderSettingResult> InvokeAsync(GetIotDefenderSettingArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIotDefenderSettingResult>("azure-native:security:getIotDefenderSetting", args ?? new GetIotDefenderSettingArgs(), options.WithVersion());
    }


    public sealed class GetIotDefenderSettingArgs : Pulumi.InvokeArgs
    {
        public GetIotDefenderSettingArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIotDefenderSettingResult
    {
        /// <summary>
        /// Size of the device quota (as a opposed to a Pay as You Go billing model). Value is required to be in multiples of 1000.
        /// </summary>
        public readonly int DeviceQuota;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The kind of onboarding for the subscription
        /// </summary>
        public readonly string OnboardingKind;
        /// <summary>
        /// Sentinel Workspace Resource Ids
        /// </summary>
        public readonly ImmutableArray<string> SentinelWorkspaceResourceIds;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetIotDefenderSettingResult(
            int deviceQuota,

            string id,

            string name,

            string onboardingKind,

            ImmutableArray<string> sentinelWorkspaceResourceIds,

            string type)
        {
            DeviceQuota = deviceQuota;
            Id = id;
            Name = name;
            OnboardingKind = onboardingKind;
            SentinelWorkspaceResourceIds = sentinelWorkspaceResourceIds;
            Type = type;
        }
    }
}
