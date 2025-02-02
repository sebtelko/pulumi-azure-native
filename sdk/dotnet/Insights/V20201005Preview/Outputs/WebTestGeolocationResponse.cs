// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Insights.V20201005Preview.Outputs
{

    [OutputType]
    public sealed class WebTestGeolocationResponse
    {
        /// <summary>
        /// Location ID for the WebTest to run from.
        /// </summary>
        public readonly string? Location;

        [OutputConstructor]
        private WebTestGeolocationResponse(string? location)
        {
            Location = location;
        }
    }
}
