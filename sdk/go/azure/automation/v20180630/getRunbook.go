// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180630

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the runbook type.
func LookupRunbook(ctx *pulumi.Context, args *LookupRunbookArgs, opts ...pulumi.InvokeOption) (*LookupRunbookResult, error) {
	var rv LookupRunbookResult
	err := ctx.Invoke("azure-native:automation/v20180630:getRunbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRunbookArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The runbook name.
	RunbookName string `pulumi:"runbookName"`
}

// Definition of the runbook type.
type LookupRunbookResult struct {
	// Gets or sets the creation time.
	CreationTime *string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets or sets the draft runbook properties.
	Draft *RunbookDraftResponse `pulumi:"draft"`
	// Gets or sets the etag of the resource.
	Etag *string `pulumi:"etag"`
	// Fully qualified resource Id for the resource
	Id string `pulumi:"id"`
	// Gets or sets the job count of the runbook.
	JobCount *int `pulumi:"jobCount"`
	// Gets or sets the last modified by.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// Gets or sets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// Gets or sets the option to log activity trace of the runbook.
	LogActivityTrace *int `pulumi:"logActivityTrace"`
	// Gets or sets progress log option.
	LogProgress *bool `pulumi:"logProgress"`
	// Gets or sets verbose log option.
	LogVerbose *bool `pulumi:"logVerbose"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets or sets the runbook output types.
	OutputTypes []string `pulumi:"outputTypes"`
	// Gets or sets the runbook parameters.
	Parameters map[string]RunbookParameterResponse `pulumi:"parameters"`
	// Gets or sets the provisioning state of the runbook.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets or sets the published runbook content link.
	PublishContentLink *ContentLinkResponse `pulumi:"publishContentLink"`
	// Gets or sets the type of the runbook.
	RunbookType *string `pulumi:"runbookType"`
	// Gets or sets the state of the runbook.
	State *string `pulumi:"state"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}
