// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logz.Outputs
{

    [OutputType]
    public sealed class LogzOrganizationPropertiesResponse
    {
        /// <summary>
        /// Name of the Logz organization.
        /// </summary>
        public readonly string? CompanyName;
        /// <summary>
        /// The Id of the Enterprise App used for Single sign on.
        /// </summary>
        public readonly string? EnterpriseAppId;
        /// <summary>
        /// Id of the Logz organization.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private LogzOrganizationPropertiesResponse(
            string? companyName,

            string? enterpriseAppId,

            string id)
        {
            CompanyName = companyName;
            EnterpriseAppId = enterpriseAppId;
            Id = id;
        }
    }
}
