// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.V20180501.Outputs
{

    [OutputType]
    public sealed class AzureStorageCredentialsInfoResponse
    {
        /// <summary>
        /// Storage account key. One of accountKey or accountKeySecretReference must be specified.
        /// </summary>
        public readonly string? AccountKey;
        /// <summary>
        /// Information about KeyVault secret storing the storage account key. One of accountKey or accountKeySecretReference must be specified.
        /// </summary>
        public readonly Outputs.KeyVaultSecretReferenceResponse? AccountKeySecretReference;

        [OutputConstructor]
        private AzureStorageCredentialsInfoResponse(
            string? accountKey,

            Outputs.KeyVaultSecretReferenceResponse? accountKeySecretReference)
        {
            AccountKey = accountKey;
            AccountKeySecretReference = accountKeySecretReference;
        }
    }
}
