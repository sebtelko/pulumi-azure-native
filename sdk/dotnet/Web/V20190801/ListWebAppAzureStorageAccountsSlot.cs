// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20190801
{
    public static class ListWebAppAzureStorageAccountsSlot
    {
        /// <summary>
        /// AzureStorageInfo dictionary resource.
        /// </summary>
        public static Task<ListWebAppAzureStorageAccountsSlotResult> InvokeAsync(ListWebAppAzureStorageAccountsSlotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListWebAppAzureStorageAccountsSlotResult>("azure-native:web/v20190801:listWebAppAzureStorageAccountsSlot", args ?? new ListWebAppAzureStorageAccountsSlotArgs(), options.WithVersion());
    }


    public sealed class ListWebAppAzureStorageAccountsSlotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API will update the Azure storage account configurations for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public ListWebAppAzureStorageAccountsSlotArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListWebAppAzureStorageAccountsSlotResult
    {
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
        /// Azure storage accounts.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.AzureStorageInfoValueResponse> Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListWebAppAzureStorageAccountsSlotResult(
            string id,

            string? kind,

            string name,

            ImmutableDictionary<string, Outputs.AzureStorageInfoValueResponse> properties,

            string type)
        {
            Id = id;
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
