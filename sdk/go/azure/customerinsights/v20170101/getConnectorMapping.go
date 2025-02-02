// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The connector mapping resource format.
func LookupConnectorMapping(ctx *pulumi.Context, args *LookupConnectorMappingArgs, opts ...pulumi.InvokeOption) (*LookupConnectorMappingResult, error) {
	var rv LookupConnectorMappingResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getConnectorMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorMappingArgs struct {
	// The name of the connector.
	ConnectorName string `pulumi:"connectorName"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the connector mapping.
	MappingName string `pulumi:"mappingName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The connector mapping resource format.
type LookupConnectorMappingResult struct {
	// The connector mapping name
	ConnectorMappingName string `pulumi:"connectorMappingName"`
	// The connector name.
	ConnectorName string `pulumi:"connectorName"`
	// Type of connector.
	ConnectorType *string `pulumi:"connectorType"`
	// The created time.
	Created string `pulumi:"created"`
	// The DataFormat ID.
	DataFormatId string `pulumi:"dataFormatId"`
	// The description of the connector mapping.
	Description *string `pulumi:"description"`
	// Display name for the connector mapping.
	DisplayName *string `pulumi:"displayName"`
	// Defines which entity type the file should map to.
	EntityType string `pulumi:"entityType"`
	// The mapping entity name.
	EntityTypeName string `pulumi:"entityTypeName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The last modified time.
	LastModified string `pulumi:"lastModified"`
	// The properties of the mapping.
	MappingProperties ConnectorMappingPropertiesResponse `pulumi:"mappingProperties"`
	// Resource name.
	Name string `pulumi:"name"`
	// The next run time based on customer's settings.
	NextRunTime string `pulumi:"nextRunTime"`
	// The RunId.
	RunId string `pulumi:"runId"`
	// State of connector mapping.
	State string `pulumi:"state"`
	// The hub name.
	TenantId string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
}
