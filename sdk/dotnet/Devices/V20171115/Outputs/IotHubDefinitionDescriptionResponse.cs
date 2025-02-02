// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Devices.V20171115.Outputs
{

    [OutputType]
    public sealed class IotHubDefinitionDescriptionResponse
    {
        /// <summary>
        /// Weight to apply for a given IoT hub.
        /// </summary>
        public readonly int? AllocationWeight;
        /// <summary>
        /// Flag for applying allocationPolicy or not for a given IoT hub.
        /// </summary>
        public readonly bool? ApplyAllocationPolicy;
        /// <summary>
        /// Connection string of the IoT hub.
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// ARM region of the IoT hub.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Host name of the IoT hub.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private IotHubDefinitionDescriptionResponse(
            int? allocationWeight,

            bool? applyAllocationPolicy,

            string connectionString,

            string location,

            string name)
        {
            AllocationWeight = allocationWeight;
            ApplyAllocationPolicy = applyAllocationPolicy;
            ConnectionString = connectionString;
            Location = location;
            Name = name;
        }
    }
}
