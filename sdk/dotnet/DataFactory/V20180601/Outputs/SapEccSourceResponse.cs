// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.V20180601.Outputs
{

    [OutputType]
    public sealed class SapEccSourceResponse
    {
        /// <summary>
        /// Specifies the additional columns to be added to source data. Type: array of objects (or Expression with resultType array of objects).
        /// </summary>
        public readonly ImmutableArray<Outputs.AdditionalColumnsResponse> AdditionalColumns;
        /// <summary>
        /// The timeout (TimeSpan) to get an HTTP response. It is the timeout to get a response, not the timeout to read response data. Default value: 00:05:00. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        public readonly object? HttpRequestTimeout;
        /// <summary>
        /// The maximum concurrent connection count for the source data store. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? MaxConcurrentConnections;
        /// <summary>
        /// SAP ECC OData query. For example, "$top=1". Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Query;
        /// <summary>
        /// Query timeout. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        public readonly object? QueryTimeout;
        /// <summary>
        /// Source retry count. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? SourceRetryCount;
        /// <summary>
        /// Source retry wait. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        public readonly object? SourceRetryWait;
        /// <summary>
        /// Copy source type.
        /// Expected value is 'SapEccSource'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SapEccSourceResponse(
            ImmutableArray<Outputs.AdditionalColumnsResponse> additionalColumns,

            object? httpRequestTimeout,

            object? maxConcurrentConnections,

            object? query,

            object? queryTimeout,

            object? sourceRetryCount,

            object? sourceRetryWait,

            string type)
        {
            AdditionalColumns = additionalColumns;
            HttpRequestTimeout = httpRequestTimeout;
            MaxConcurrentConnections = maxConcurrentConnections;
            Query = query;
            QueryTimeout = queryTimeout;
            SourceRetryCount = sourceRetryCount;
            SourceRetryWait = sourceRetryWait;
            Type = type;
        }
    }
}
