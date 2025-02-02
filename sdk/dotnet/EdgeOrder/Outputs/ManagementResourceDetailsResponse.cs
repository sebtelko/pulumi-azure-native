// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EdgeOrder.Outputs
{

    [OutputType]
    public sealed class ManagementResourceDetailsResponse
    {
        /// <summary>
        /// Management resource ARM ID
        /// </summary>
        public readonly string ManagementResourceArmId;

        [OutputConstructor]
        private ManagementResourceDetailsResponse(string managementResourceArmId)
        {
            ManagementResourceArmId = managementResourceArmId;
        }
    }
}
