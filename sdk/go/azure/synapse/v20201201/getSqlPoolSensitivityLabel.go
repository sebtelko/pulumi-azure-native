// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A sensitivity label.
func LookupSqlPoolSensitivityLabel(ctx *pulumi.Context, args *LookupSqlPoolSensitivityLabelArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolSensitivityLabelResult, error) {
	var rv LookupSqlPoolSensitivityLabelResult
	err := ctx.Invoke("azure-native:synapse/v20201201:getSqlPoolSensitivityLabel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolSensitivityLabelArgs struct {
	// The name of the column.
	ColumnName string `pulumi:"columnName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the schema.
	SchemaName string `pulumi:"schemaName"`
	// The source of the sensitivity label.
	SensitivityLabelSource string `pulumi:"sensitivityLabelSource"`
	// SQL pool name
	SqlPoolName string `pulumi:"sqlPoolName"`
	// The name of the table.
	TableName string `pulumi:"tableName"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// A sensitivity label.
type LookupSqlPoolSensitivityLabelResult struct {
	// The column name.
	ColumnName string `pulumi:"columnName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The information type.
	InformationType *string `pulumi:"informationType"`
	// The information type ID.
	InformationTypeId *string `pulumi:"informationTypeId"`
	// Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
	IsDisabled bool `pulumi:"isDisabled"`
	// The label ID.
	LabelId *string `pulumi:"labelId"`
	// The label name.
	LabelName *string `pulumi:"labelName"`
	// managed by
	ManagedBy string `pulumi:"managedBy"`
	// The name of the resource
	Name string  `pulumi:"name"`
	Rank *string `pulumi:"rank"`
	// The schema name.
	SchemaName string `pulumi:"schemaName"`
	// The table name.
	TableName string `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
