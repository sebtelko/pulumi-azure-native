// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The task that has the ARM resource and task properties.
// The task will have all information to schedule a run against it.
func LookupTask(ctx *pulumi.Context, args *LookupTaskArgs, opts ...pulumi.InvokeOption) (*LookupTaskResult, error) {
	var rv LookupTaskResult
	err := ctx.Invoke("azure-native:containerregistry/v20180901:getTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTaskArgs struct {
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the container registry task.
	TaskName string `pulumi:"taskName"`
}

// The task that has the ARM resource and task properties.
// The task will have all information to schedule a run against it.
type LookupTaskResult struct {
	// The machine configuration of the run agent.
	AgentConfiguration *AgentPropertiesResponse `pulumi:"agentConfiguration"`
	// The creation date of task.
	CreationDate string `pulumi:"creationDate"`
	// The properties that describes a set of credentials that will be used when this run is invoked.
	Credentials *CredentialsResponse `pulumi:"credentials"`
	// The resource ID.
	Id string `pulumi:"id"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The platform properties against which the run has to happen.
	Platform PlatformPropertiesResponse `pulumi:"platform"`
	// The provisioning state of the task.
	ProvisioningState string `pulumi:"provisioningState"`
	// The current status of task.
	Status *string `pulumi:"status"`
	// The properties of a task step.
	Step interface{} `pulumi:"step"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Run timeout in seconds.
	Timeout *int `pulumi:"timeout"`
	// The properties that describe all triggers for the task.
	Trigger *TriggerPropertiesResponse `pulumi:"trigger"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
