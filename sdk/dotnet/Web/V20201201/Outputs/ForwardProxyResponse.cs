// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20201201.Outputs
{

    [OutputType]
    public sealed class ForwardProxyResponse
    {
        /// <summary>
        /// The convention used to determine the url of the request made.
        /// </summary>
        public readonly string? Convention;
        /// <summary>
        /// The name of the header containing the host of the request.
        /// </summary>
        public readonly string? CustomHostHeaderName;
        /// <summary>
        /// The name of the header containing the scheme of the request.
        /// </summary>
        public readonly string? CustomProtoHeaderName;
        /// <summary>
        /// Resource Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ForwardProxyResponse(
            string? convention,

            string? customHostHeaderName,

            string? customProtoHeaderName,

            string id,

            string? kind,

            string name,

            string type)
        {
            Convention = convention;
            CustomHostHeaderName = customHostHeaderName;
            CustomProtoHeaderName = customProtoHeaderName;
            Id = id;
            Kind = kind;
            Name = name;
            Type = type;
        }
    }
}
