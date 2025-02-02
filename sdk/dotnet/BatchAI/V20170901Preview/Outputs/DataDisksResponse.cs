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
    public sealed class DataDisksResponse
    {
        public readonly string CachingType;
        public readonly int DiskCount;
        public readonly int DiskSizeInGB;
        public readonly string StorageAccountType;

        [OutputConstructor]
        private DataDisksResponse(
            string cachingType,

            int diskCount,

            int diskSizeInGB,

            string storageAccountType)
        {
            CachingType = cachingType;
            DiskCount = diskCount;
            DiskSizeInGB = diskSizeInGB;
            StorageAccountType = storageAccountType;
        }
    }
}
