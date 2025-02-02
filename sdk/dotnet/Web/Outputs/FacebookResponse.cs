// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Outputs
{

    [OutputType]
    public sealed class FacebookResponse
    {
        /// <summary>
        /// &lt;code&gt;false&lt;/code&gt; if the Facebook provider should not be enabled despite the set registration; otherwise, &lt;code&gt;true&lt;/code&gt;.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The version of the Facebook api to be used while logging in.
        /// </summary>
        public readonly string? GraphApiVersion;
        /// <summary>
        /// Resource Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The configuration settings of the login flow.
        /// </summary>
        public readonly Outputs.LoginScopesResponse? Login;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The configuration settings of the app registration for the Facebook provider.
        /// </summary>
        public readonly Outputs.AppRegistrationResponse? Registration;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private FacebookResponse(
            bool? enabled,

            string? graphApiVersion,

            string id,

            string? kind,

            Outputs.LoginScopesResponse? login,

            string name,

            Outputs.AppRegistrationResponse? registration,

            string type)
        {
            Enabled = enabled;
            GraphApiVersion = graphApiVersion;
            Id = id;
            Kind = kind;
            Login = login;
            Name = name;
            Registration = registration;
            Type = type;
        }
    }
}
