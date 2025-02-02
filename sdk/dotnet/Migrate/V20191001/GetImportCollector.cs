// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.V20191001
{
    public static class GetImportCollector
    {
        public static Task<GetImportCollectorResult> InvokeAsync(GetImportCollectorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImportCollectorResult>("azure-native:migrate/v20191001:getImportCollector", args ?? new GetImportCollectorArgs(), options.WithVersion());
    }


    public sealed class GetImportCollectorArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique name of a Import collector within a project.
        /// </summary>
        [Input("importCollectorName", required: true)]
        public string ImportCollectorName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Migrate project.
        /// </summary>
        [Input("projectName", required: true)]
        public string ProjectName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Resource Group that project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetImportCollectorArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetImportCollectorResult
    {
        public readonly string? ETag;
        public readonly string Id;
        public readonly string Name;
        public readonly Outputs.ImportCollectorPropertiesResponse Properties;
        public readonly string Type;

        [OutputConstructor]
        private GetImportCollectorResult(
            string? eTag,

            string id,

            string name,

            Outputs.ImportCollectorPropertiesResponse properties,

            string type)
        {
            ETag = eTag;
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
