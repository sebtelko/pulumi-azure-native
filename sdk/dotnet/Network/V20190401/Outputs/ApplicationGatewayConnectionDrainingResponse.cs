// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20190401.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayConnectionDrainingResponse
    {
        /// <summary>
        /// The number of seconds connection draining is active. Acceptable values are from 1 second to 3600 seconds.
        /// </summary>
        public readonly int DrainTimeoutInSec;
        /// <summary>
        /// Whether connection draining is enabled or not.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private ApplicationGatewayConnectionDrainingResponse(
            int drainTimeoutInSec,

            bool enabled)
        {
            DrainTimeoutInSec = drainTimeoutInSec;
            Enabled = enabled;
        }
    }
}
