// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure SQL Database sync member.
func LookupSyncMember(ctx *pulumi.Context, args *LookupSyncMemberArgs, opts ...pulumi.InvokeOption) (*LookupSyncMemberResult, error) {
	var rv LookupSyncMemberResult
	err := ctx.Invoke("azure-native:sql/v20201101preview:getSyncMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncMemberArgs struct {
	// The name of the database on which the sync group is hosted.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the sync group on which the sync member is hosted.
	SyncGroupName string `pulumi:"syncGroupName"`
	// The name of the sync member.
	SyncMemberName string `pulumi:"syncMemberName"`
}

// An Azure SQL Database sync member.
type LookupSyncMemberResult struct {
	// Database name of the member database in the sync member.
	DatabaseName *string `pulumi:"databaseName"`
	// Database type of the sync member.
	DatabaseType *string `pulumi:"databaseType"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Private endpoint name of the sync member if use private link connection is enabled, for sync members in Azure.
	PrivateEndpointName string `pulumi:"privateEndpointName"`
	// Server name of the member database in the sync member
	ServerName *string `pulumi:"serverName"`
	// SQL Server database id of the sync member.
	SqlServerDatabaseId *string `pulumi:"sqlServerDatabaseId"`
	// ARM resource id of the sync agent in the sync member.
	SyncAgentId *string `pulumi:"syncAgentId"`
	// Sync direction of the sync member.
	SyncDirection *string `pulumi:"syncDirection"`
	// ARM resource id of the sync member logical database, for sync members in Azure.
	SyncMemberAzureDatabaseResourceId *string `pulumi:"syncMemberAzureDatabaseResourceId"`
	// Sync state of the sync member.
	SyncState string `pulumi:"syncState"`
	// Resource type.
	Type string `pulumi:"type"`
	// Whether to use private link connection.
	UsePrivateLinkConnection *bool `pulumi:"usePrivateLinkConnection"`
	// User name of the member database in the sync member.
	UserName *string `pulumi:"userName"`
}
