// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The scheduled query rule resource.
type ScheduledQueryRule struct {
	pulumi.CustomResourceState

	// Actions to invoke when the alert fires.
	Actions ActionResponseArrayOutput `pulumi:"actions"`
	// The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.
	AutoMitigate pulumi.BoolPtrOutput `pulumi:"autoMitigate"`
	// The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
	CheckWorkspaceAlertsStorageConfigured pulumi.BoolPtrOutput `pulumi:"checkWorkspaceAlertsStorageConfigured"`
	// The api-version used when creating this alert rule
	CreatedWithApiVersion pulumi.StringOutput `pulumi:"createdWithApiVersion"`
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria ScheduledQueryRuleCriteriaResponseOutput `pulumi:"criteria"`
	// The description of the scheduled query rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the alert rule
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.
	EvaluationFrequency pulumi.StringPtrOutput `pulumi:"evaluationFrequency"`
	// True if alert rule is legacy Log Analytic rule
	IsLegacyLogAnalyticsRule pulumi.BoolOutput `pulumi:"isLegacyLogAnalyticsRule"`
	// The flag which indicates whether this scheduled query rule has been configured to be stored in the customer's storage. The default is false.
	IsWorkspaceAlertsStorageConfigured pulumi.BoolOutput `pulumi:"isWorkspaceAlertsStorageConfigured"`
	// Indicates the type of scheduled query rule. The default is LogAlert.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.
	MuteActionsDuration pulumi.StringPtrOutput `pulumi:"muteActionsDuration"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
	OverrideQueryTimeRange pulumi.StringPtrOutput `pulumi:"overrideQueryTimeRange"`
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.
	Severity pulumi.Float64PtrOutput `pulumi:"severity"`
	// The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.
	SkipQueryValidation pulumi.BoolPtrOutput `pulumi:"skipQueryValidation"`
	// SystemData of ScheduledQueryRule.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert
	TargetResourceTypes pulumi.StringArrayOutput `pulumi:"targetResourceTypes"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.
	WindowSize pulumi.StringPtrOutput `pulumi:"windowSize"`
}

// NewScheduledQueryRule registers a new resource with the given unique name, arguments, and options.
func NewScheduledQueryRule(ctx *pulumi.Context,
	name string, args *ScheduledQueryRuleArgs, opts ...pulumi.ResourceOption) (*ScheduledQueryRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights/v20210201preview:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180416:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20180416:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20200501preview:ScheduledQueryRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20200501preview:ScheduledQueryRule"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledQueryRule
	err := ctx.RegisterResource("azure-native:insights/v20210201preview:ScheduledQueryRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledQueryRule gets an existing ScheduledQueryRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledQueryRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledQueryRuleState, opts ...pulumi.ResourceOption) (*ScheduledQueryRule, error) {
	var resource ScheduledQueryRule
	err := ctx.ReadResource("azure-native:insights/v20210201preview:ScheduledQueryRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledQueryRule resources.
type scheduledQueryRuleState struct {
	// Actions to invoke when the alert fires.
	Actions []ActionResponse `pulumi:"actions"`
	// The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.
	AutoMitigate *bool `pulumi:"autoMitigate"`
	// The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
	CheckWorkspaceAlertsStorageConfigured *bool `pulumi:"checkWorkspaceAlertsStorageConfigured"`
	// The api-version used when creating this alert rule
	CreatedWithApiVersion *string `pulumi:"createdWithApiVersion"`
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria *ScheduledQueryRuleCriteriaResponse `pulumi:"criteria"`
	// The description of the scheduled query rule.
	Description *string `pulumi:"description"`
	// The display name of the alert rule
	DisplayName *string `pulumi:"displayName"`
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled *bool `pulumi:"enabled"`
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag *string `pulumi:"etag"`
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.
	EvaluationFrequency *string `pulumi:"evaluationFrequency"`
	// True if alert rule is legacy Log Analytic rule
	IsLegacyLogAnalyticsRule *bool `pulumi:"isLegacyLogAnalyticsRule"`
	// The flag which indicates whether this scheduled query rule has been configured to be stored in the customer's storage. The default is false.
	IsWorkspaceAlertsStorageConfigured *bool `pulumi:"isWorkspaceAlertsStorageConfigured"`
	// Indicates the type of scheduled query rule. The default is LogAlert.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.
	MuteActionsDuration *string `pulumi:"muteActionsDuration"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
	OverrideQueryTimeRange *string `pulumi:"overrideQueryTimeRange"`
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes []string `pulumi:"scopes"`
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.
	Severity *float64 `pulumi:"severity"`
	// The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.
	SkipQueryValidation *bool `pulumi:"skipQueryValidation"`
	// SystemData of ScheduledQueryRule.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert
	TargetResourceTypes []string `pulumi:"targetResourceTypes"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.
	WindowSize *string `pulumi:"windowSize"`
}

type ScheduledQueryRuleState struct {
	// Actions to invoke when the alert fires.
	Actions ActionResponseArrayInput
	// The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.
	AutoMitigate pulumi.BoolPtrInput
	// The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
	CheckWorkspaceAlertsStorageConfigured pulumi.BoolPtrInput
	// The api-version used when creating this alert rule
	CreatedWithApiVersion pulumi.StringPtrInput
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria ScheduledQueryRuleCriteriaResponsePtrInput
	// The description of the scheduled query rule.
	Description pulumi.StringPtrInput
	// The display name of the alert rule
	DisplayName pulumi.StringPtrInput
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled pulumi.BoolPtrInput
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag pulumi.StringPtrInput
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.
	EvaluationFrequency pulumi.StringPtrInput
	// True if alert rule is legacy Log Analytic rule
	IsLegacyLogAnalyticsRule pulumi.BoolPtrInput
	// The flag which indicates whether this scheduled query rule has been configured to be stored in the customer's storage. The default is false.
	IsWorkspaceAlertsStorageConfigured pulumi.BoolPtrInput
	// Indicates the type of scheduled query rule. The default is LogAlert.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.
	MuteActionsDuration pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
	OverrideQueryTimeRange pulumi.StringPtrInput
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes pulumi.StringArrayInput
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.
	Severity pulumi.Float64PtrInput
	// The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.
	SkipQueryValidation pulumi.BoolPtrInput
	// SystemData of ScheduledQueryRule.
	SystemData SystemDataResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert
	TargetResourceTypes pulumi.StringArrayInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.
	WindowSize pulumi.StringPtrInput
}

func (ScheduledQueryRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledQueryRuleState)(nil)).Elem()
}

type scheduledQueryRuleArgs struct {
	// Actions to invoke when the alert fires.
	Actions []Action `pulumi:"actions"`
	// The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.
	AutoMitigate *bool `pulumi:"autoMitigate"`
	// The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
	CheckWorkspaceAlertsStorageConfigured *bool `pulumi:"checkWorkspaceAlertsStorageConfigured"`
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria ScheduledQueryRuleCriteria `pulumi:"criteria"`
	// The description of the scheduled query rule.
	Description *string `pulumi:"description"`
	// The display name of the alert rule
	DisplayName *string `pulumi:"displayName"`
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled bool `pulumi:"enabled"`
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.
	EvaluationFrequency *string `pulumi:"evaluationFrequency"`
	// Indicates the type of scheduled query rule. The default is LogAlert.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.
	MuteActionsDuration *string `pulumi:"muteActionsDuration"`
	// If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
	OverrideQueryTimeRange *string `pulumi:"overrideQueryTimeRange"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the rule.
	RuleName *string `pulumi:"ruleName"`
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes []string `pulumi:"scopes"`
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.
	Severity *float64 `pulumi:"severity"`
	// The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.
	SkipQueryValidation *bool `pulumi:"skipQueryValidation"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert
	TargetResourceTypes []string `pulumi:"targetResourceTypes"`
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.
	WindowSize *string `pulumi:"windowSize"`
}

// The set of arguments for constructing a ScheduledQueryRule resource.
type ScheduledQueryRuleArgs struct {
	// Actions to invoke when the alert fires.
	Actions ActionArrayInput
	// The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.
	AutoMitigate pulumi.BoolPtrInput
	// The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
	CheckWorkspaceAlertsStorageConfigured pulumi.BoolPtrInput
	// The rule criteria that defines the conditions of the scheduled query rule.
	Criteria ScheduledQueryRuleCriteriaInput
	// The description of the scheduled query rule.
	Description pulumi.StringPtrInput
	// The display name of the alert rule
	DisplayName pulumi.StringPtrInput
	// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled pulumi.BoolInput
	// How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.
	EvaluationFrequency pulumi.StringPtrInput
	// Indicates the type of scheduled query rule. The default is LogAlert.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.
	MuteActionsDuration pulumi.StringPtrInput
	// If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
	OverrideQueryTimeRange pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the rule.
	RuleName pulumi.StringPtrInput
	// The list of resource id's that this scheduled query rule is scoped to.
	Scopes pulumi.StringArrayInput
	// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.
	Severity pulumi.Float64PtrInput
	// The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.
	SkipQueryValidation pulumi.BoolPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert
	TargetResourceTypes pulumi.StringArrayInput
	// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.
	WindowSize pulumi.StringPtrInput
}

func (ScheduledQueryRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledQueryRuleArgs)(nil)).Elem()
}

type ScheduledQueryRuleInput interface {
	pulumi.Input

	ToScheduledQueryRuleOutput() ScheduledQueryRuleOutput
	ToScheduledQueryRuleOutputWithContext(ctx context.Context) ScheduledQueryRuleOutput
}

func (*ScheduledQueryRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRule)(nil))
}

func (i *ScheduledQueryRule) ToScheduledQueryRuleOutput() ScheduledQueryRuleOutput {
	return i.ToScheduledQueryRuleOutputWithContext(context.Background())
}

func (i *ScheduledQueryRule) ToScheduledQueryRuleOutputWithContext(ctx context.Context) ScheduledQueryRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRuleOutput)
}

type ScheduledQueryRuleOutput struct {
	*pulumi.OutputState
}

func (ScheduledQueryRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRule)(nil))
}

func (o ScheduledQueryRuleOutput) ToScheduledQueryRuleOutput() ScheduledQueryRuleOutput {
	return o
}

func (o ScheduledQueryRuleOutput) ToScheduledQueryRuleOutputWithContext(ctx context.Context) ScheduledQueryRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScheduledQueryRuleOutput{})
}
