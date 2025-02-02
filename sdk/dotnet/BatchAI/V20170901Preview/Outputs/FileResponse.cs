// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.V20170901Preview.Outputs
{

    [OutputType]
    public sealed class FileResponse
    {
        /// <summary>
        /// The file size.
        /// </summary>
        public readonly double? ContentLength;
        /// <summary>
        /// This will be returned only if the model has been archived. During job run, this won't be returned and customers can use SSH tunneling to download. Users can use Get Remote Login Information API to get the IP address and port information of all the compute nodes running the job.
        /// </summary>
        public readonly string DownloadUrl;
        /// <summary>
        /// The time at which the file was last modified.
        /// </summary>
        public readonly string? LastModified;
        /// <summary>
        /// file name
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private FileResponse(
            double? contentLength,

            string downloadUrl,

            string? lastModified,

            string name)
        {
            ContentLength = contentLength;
            DownloadUrl = downloadUrl;
            LastModified = lastModified;
            Name = name;
        }
    }
}
