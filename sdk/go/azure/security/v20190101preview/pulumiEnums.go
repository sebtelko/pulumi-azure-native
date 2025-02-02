// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of the action that will be triggered by the Automation
type ActionType pulumi.String

const (
	ActionTypeLogicApp  = ActionType("LogicApp")
	ActionTypeEventHub  = ActionType("EventHub")
	ActionTypeWorkspace = ActionType("Workspace")
)

func (ActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Programmatic code for the status of the assessment
type AssessmentStatusCode pulumi.String

const (
	// The resource is healthy
	AssessmentStatusCodeHealthy = AssessmentStatusCode("Healthy")
	// The resource has a security issue that needs to be addressed
	AssessmentStatusCodeUnhealthy = AssessmentStatusCode("Unhealthy")
	// Assessment for this resource did not happen
	AssessmentStatusCodeNotApplicable = AssessmentStatusCode("NotApplicable")
)

func (AssessmentStatusCode) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AssessmentStatusCode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentStatusCode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentStatusCode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentStatusCode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
type AssessmentType pulumi.String

const (
	// Azure Security Center managed assessments
	AssessmentTypeBuiltIn = AssessmentType("BuiltIn")
	// User defined policies that are automatically ingested from Azure Policy to Azure Security Center
	AssessmentTypeCustomPolicy = AssessmentType("CustomPolicy")
	// User assessments pushed directly by the user or other third party to Azure Security Center
	AssessmentTypeCustomerManaged = AssessmentType("CustomerManaged")
)

func (AssessmentType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AssessmentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The categories of resource that is at risk when the assessment is unhealthy
type Categories pulumi.String

const (
	CategoriesCompute           = Categories("Compute")
	CategoriesNetworking        = Categories("Networking")
	CategoriesData              = Categories("Data")
	CategoriesIdentityAndAccess = Categories("IdentityAndAccess")
	CategoriesIoT               = Categories("IoT")
)

func (Categories) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Categories) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Categories) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Categories) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Categories) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// A valid event source type.
type EventSource pulumi.String

const (
	EventSourceAssessments                            = EventSource("Assessments")
	EventSourceSubAssessments                         = EventSource("SubAssessments")
	EventSourceAlerts                                 = EventSource("Alerts")
	EventSourceSecureScores                           = EventSource("SecureScores")
	EventSourceSecureScoresSnapshot                   = EventSource("SecureScoresSnapshot")
	EventSourceSecureScoreControls                    = EventSource("SecureScoreControls")
	EventSourceSecureScoreControlsSnapshot            = EventSource("SecureScoreControlsSnapshot")
	EventSourceRegulatoryComplianceAssessment         = EventSource("RegulatoryComplianceAssessment")
	EventSourceRegulatoryComplianceAssessmentSnapshot = EventSource("RegulatoryComplianceAssessmentSnapshot")
)

func (EventSource) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e EventSource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventSource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The implementation effort required to remediate this assessment
type ImplementationEffort pulumi.String

const (
	ImplementationEffortLow      = ImplementationEffort("Low")
	ImplementationEffortModerate = ImplementationEffort("Moderate")
	ImplementationEffortHigh     = ImplementationEffort("High")
)

func (ImplementationEffort) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ImplementationEffort) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImplementationEffort) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImplementationEffort) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ImplementationEffort) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// A valid comparer operator to use. A case-insensitive comparison will be applied for String PropertyType.
type Operator pulumi.String

const (
	// Applies for decimal and non-decimal operands
	OperatorEquals = Operator("Equals")
	// Applies only for decimal operands
	OperatorGreaterThan = Operator("GreaterThan")
	// Applies only for decimal operands
	OperatorGreaterThanOrEqualTo = Operator("GreaterThanOrEqualTo")
	// Applies only for decimal operands
	OperatorLesserThan = Operator("LesserThan")
	// Applies only for decimal operands
	OperatorLesserThanOrEqualTo = Operator("LesserThanOrEqualTo")
	// Applies  for decimal and non-decimal operands
	OperatorNotEquals = Operator("NotEquals")
	// Applies only for non-decimal operands
	OperatorContains = Operator("Contains")
	// Applies only for non-decimal operands
	OperatorStartsWith = Operator("StartsWith")
	// Applies only for non-decimal operands
	OperatorEndsWith = Operator("EndsWith")
)

func (Operator) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Operator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Operator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The data type of the compared operands (string, integer, floating point number or a boolean [true/false]]
type PropertyType pulumi.String

const (
	PropertyTypeString  = PropertyType("String")
	PropertyTypeInteger = PropertyType("Integer")
	PropertyTypeNumber  = PropertyType("Number")
	PropertyTypeBoolean = PropertyType("Boolean")
)

func (PropertyType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e PropertyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PropertyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PropertyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PropertyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Possible states of the rule
type RuleState pulumi.String

const (
	RuleStateEnabled  = RuleState("Enabled")
	RuleStateDisabled = RuleState("Disabled")
	RuleStateExpired  = RuleState("Expired")
)

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e RuleState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RuleState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RuleState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RuleState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The severity level of the assessment
type Severity pulumi.String

const (
	SeverityLow    = Severity("Low")
	SeverityMedium = Severity("Medium")
	SeverityHigh   = Severity("High")
)

func (Severity) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Severity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Severity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Severity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Severity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The platform where the assessed resource resides
type Source pulumi.String

const (
	// Resource is in Azure
	SourceAzure = Source("Azure")
	// Resource in an on premise machine connected to Azure cloud
	SourceOnPremise = Source("OnPremise")
	// SQL Resource in an on premise machine connected to Azure cloud
	SourceOnPremiseSql = Source("OnPremiseSql")
)

func (Source) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Source) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Source) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Source) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Source) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Threats impact of the assessment
type Threats pulumi.String

const (
	ThreatsAccountBreach        = Threats("accountBreach")
	ThreatsDataExfiltration     = Threats("dataExfiltration")
	ThreatsDataSpillage         = Threats("dataSpillage")
	ThreatsMaliciousInsider     = Threats("maliciousInsider")
	ThreatsElevationOfPrivilege = Threats("elevationOfPrivilege")
	ThreatsThreatResistance     = Threats("threatResistance")
	ThreatsMissingCoverage      = Threats("missingCoverage")
	ThreatsDenialOfService      = Threats("denialOfService")
)

func (Threats) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Threats) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Threats) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Threats) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Threats) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The user impact of the assessment
type UserImpact pulumi.String

const (
	UserImpactLow      = UserImpact("Low")
	UserImpactModerate = UserImpact("Moderate")
	UserImpactHigh     = UserImpact("High")
)

func (UserImpact) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e UserImpact) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserImpact) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserImpact) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UserImpact) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
