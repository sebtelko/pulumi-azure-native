// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a connected cluster.
func LookupConnectedCluster(ctx *pulumi.Context, args *LookupConnectedClusterArgs, opts ...pulumi.InvokeOption) (*LookupConnectedClusterResult, error) {
	var rv LookupConnectedClusterResult
	err := ctx.Invoke("azure-native:kubernetes/v20200101preview:getConnectedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectedClusterArgs struct {
	// The name of the Kubernetes cluster on which get is called.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a connected cluster.
type LookupConnectedClusterResult struct {
	// AAD profile of the connected cluster.
	AadProfile ConnectedClusterAADProfileResponse `pulumi:"aadProfile"`
	// Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
	AgentPublicKeyCertificate string `pulumi:"agentPublicKeyCertificate"`
	// Version of the agent running on the connected cluster resource
	AgentVersion string `pulumi:"agentVersion"`
	// Represents the connectivity status of the connected cluster.
	ConnectivityStatus *string `pulumi:"connectivityStatus"`
	// The Kubernetes distribution running on this connected cluster.
	Distribution *string `pulumi:"distribution"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The identity of the connected cluster.
	Identity ConnectedClusterIdentityResponse `pulumi:"identity"`
	// The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
	Infrastructure *string `pulumi:"infrastructure"`
	// The Kubernetes version of the connected cluster resource
	KubernetesVersion string `pulumi:"kubernetesVersion"`
	// Time representing the last instance when heart beat was received from the cluster
	LastConnectivityTime string `pulumi:"lastConnectivityTime"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Expiration time of the managed identity certificate
	ManagedIdentityCertificateExpirationTime string `pulumi:"managedIdentityCertificateExpirationTime"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Connected cluster offering
	Offering string `pulumi:"offering"`
	// Provisioning state of the connected cluster resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Number of CPU cores present in the connected cluster resource
	TotalCoreCount int `pulumi:"totalCoreCount"`
	// Number of nodes present in the connected cluster resource
	TotalNodeCount int `pulumi:"totalNodeCount"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
