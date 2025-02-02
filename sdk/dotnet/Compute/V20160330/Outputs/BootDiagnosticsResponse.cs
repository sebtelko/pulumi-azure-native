// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.V20160330.Outputs
{

    [OutputType]
    public sealed class BootDiagnosticsResponse
    {
        /// <summary>
        /// Whether boot diagnostics should be enabled on the Virtual Machine.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Uri of the storage account to use for placing the console output and screenshot.
        /// </summary>
        public readonly string? StorageUri;

        [OutputConstructor]
        private BootDiagnosticsResponse(
            bool? enabled,

            string? storageUri)
        {
            Enabled = enabled;
            StorageUri = storageUri;
        }
    }
}
