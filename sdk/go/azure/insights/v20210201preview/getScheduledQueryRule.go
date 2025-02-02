// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The scheduled query rule resource.
func LookupScheduledQueryRule(ctx *pulumi.Context, args *LookupScheduledQueryRuleArgs, opts ...pulumi.InvokeOption) (*LookupScheduledQueryRuleResult, error) {
	var rv LookupScheduledQueryRuleResult
	err := ctx.Invoke("azure-native:insights/v20210201preview:getScheduledQueryRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledQueryRuleArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the rule.
	RuleName string `pulumi:"ruleName"`
}

// The scheduled query rule resource.
type LookupScheduledQueryRuleResult struct {
	// Actions to invoke when the alert fires.
	Actions []ActionResponse `pulumi:"actions"`
	// The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.
	AutoMitigate *bool `pulumi:"autoMitigate"`
	// The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
	CheckWorkspaceAlertsStorageConfigured *bool `pulumi:"checkWorkspaceAlertsStorageConfigured"`
	// The api-version used when creating this alert rule
	CreatedWithApiVersion string `pulumi:"createdWithApiVersion"`
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria ScheduledQueryRuleCriteriaResponse `pulumi:"criteria"`
	// The description of the scheduled query rule.
	Description *string `pulumi:"description"`
	// The display name of the alert rule
	DisplayName *string `pulumi:"displayName"`
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled bool `pulumi:"enabled"`
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag string `pulumi:"etag"`
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.
	EvaluationFrequency *string `pulumi:"evaluationFrequency"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// True if alert rule is legacy Log Analytic rule
	IsLegacyLogAnalyticsRule bool `pulumi:"isLegacyLogAnalyticsRule"`
	// The flag which indicates whether this scheduled query rule has been configured to be stored in the customer's storage. The default is false.
	IsWorkspaceAlertsStorageConfigured bool `pulumi:"isWorkspaceAlertsStorageConfigured"`
	// Indicates the type of scheduled query rule. The default is LogAlert.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.
	MuteActionsDuration *string `pulumi:"muteActionsDuration"`
	// The name of the resource
	Name string `pulumi:"name"`
	// If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
	OverrideQueryTimeRange *string `pulumi:"overrideQueryTimeRange"`
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes []string `pulumi:"scopes"`
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.
	Severity *float64 `pulumi:"severity"`
	// The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.
	SkipQueryValidation *bool `pulumi:"skipQueryValidation"`
	// SystemData of ScheduledQueryRule.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert
	TargetResourceTypes []string `pulumi:"targetResourceTypes"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.
	WindowSize *string `pulumi:"windowSize"`
}
