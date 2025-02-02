// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.V20180301.Outputs
{

    [OutputType]
    public sealed class JobPropertiesResponseExecutionInfo
    {
        /// <summary>
        /// This property is only returned if the job is in completed state.
        /// </summary>
        public readonly string? EndTime;
        public readonly ImmutableArray<Outputs.BatchAIErrorResponse> Errors;
        /// <summary>
        /// This property is only returned if the job is in completed state.
        /// </summary>
        public readonly int? ExitCode;
        /// <summary>
        /// 'Running' corresponds to the running state. If the job has been restarted or retried, this is the most recent time at which the job started running. This property is present only for job that are in the running or completed state.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private JobPropertiesResponseExecutionInfo(
            string? endTime,

            ImmutableArray<Outputs.BatchAIErrorResponse> errors,

            int? exitCode,

            string startTime)
        {
            EndTime = endTime;
            Errors = errors;
            ExitCode = exitCode;
            StartTime = startTime;
        }
    }
}
