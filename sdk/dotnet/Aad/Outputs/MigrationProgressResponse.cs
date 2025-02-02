// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Aad.Outputs
{

    [OutputType]
    public sealed class MigrationProgressResponse
    {
        /// <summary>
        /// Completion Percentage
        /// </summary>
        public readonly double? CompletionPercentage;
        /// <summary>
        /// Progress Message
        /// </summary>
        public readonly string? ProgressMessage;

        [OutputConstructor]
        private MigrationProgressResponse(
            double? completionPercentage,

            string? progressMessage)
        {
            CompletionPercentage = completionPercentage;
            ProgressMessage = progressMessage;
        }
    }
}
