// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The kind of the alert rule
type AlertRuleKind pulumi.String

const (
	AlertRuleKindScheduled                         = AlertRuleKind("Scheduled")
	AlertRuleKindMicrosoftSecurityIncidentCreation = AlertRuleKind("MicrosoftSecurityIncidentCreation")
	AlertRuleKindFusion                            = AlertRuleKind("Fusion")
	AlertRuleKindMLBehaviorAnalytics               = AlertRuleKind("MLBehaviorAnalytics")
	AlertRuleKindThreatIntelligence                = AlertRuleKind("ThreatIntelligence")
)

func (AlertRuleKind) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AlertRuleKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlertRuleKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlertRuleKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AlertRuleKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The type of the automation rule action
type AutomationRuleActionType pulumi.String

const (
	// Modify an object's properties
	AutomationRuleActionTypeModifyProperties = AutomationRuleActionType("ModifyProperties")
	// Run a playbook on an object
	AutomationRuleActionTypeRunPlaybook = AutomationRuleActionType("RunPlaybook")
)

func (AutomationRuleActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AutomationRuleActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRuleActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRuleActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutomationRuleActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The type of the automation rule condition
type AutomationRuleConditionType pulumi.String

const (
	// Evaluate an object property value
	AutomationRuleConditionTypeProperty = AutomationRuleConditionType("Property")
)

func (AutomationRuleConditionType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AutomationRuleConditionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRuleConditionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRuleConditionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutomationRuleConditionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The operator to use for evaluation the condition
type AutomationRulePropertyConditionSupportedOperator pulumi.String

const (
	// Evaluates if the property equals at least one of the condition values
	AutomationRulePropertyConditionSupportedOperatorEquals = AutomationRulePropertyConditionSupportedOperator("Equals")
	// Evaluates if the property does not equal any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotEquals = AutomationRulePropertyConditionSupportedOperator("NotEquals")
	// Evaluates if the property contains at least one of the condition values
	AutomationRulePropertyConditionSupportedOperatorContains = AutomationRulePropertyConditionSupportedOperator("Contains")
	// Evaluates if the property does not contain any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotContains = AutomationRulePropertyConditionSupportedOperator("NotContains")
	// Evaluates if the property starts with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorStartsWith = AutomationRulePropertyConditionSupportedOperator("StartsWith")
	// Evaluates if the property does not start with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotStartsWith = AutomationRulePropertyConditionSupportedOperator("NotStartsWith")
	// Evaluates if the property ends with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorEndsWith = AutomationRulePropertyConditionSupportedOperator("EndsWith")
	// Evaluates if the property does not end with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotEndsWith = AutomationRulePropertyConditionSupportedOperator("NotEndsWith")
)

func (AutomationRulePropertyConditionSupportedOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AutomationRulePropertyConditionSupportedOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRulePropertyConditionSupportedOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRulePropertyConditionSupportedOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutomationRulePropertyConditionSupportedOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The property to evaluate
type AutomationRulePropertyConditionSupportedProperty pulumi.String

const (
	// The title of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentTitle = AutomationRulePropertyConditionSupportedProperty("IncidentTitle")
	// The description of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentDescription = AutomationRulePropertyConditionSupportedProperty("IncidentDescription")
	// The severity of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentSeverity = AutomationRulePropertyConditionSupportedProperty("IncidentSeverity")
	// The status of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentStatus = AutomationRulePropertyConditionSupportedProperty("IncidentStatus")
	// The tactics of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentTactics = AutomationRulePropertyConditionSupportedProperty("IncidentTactics")
	// The related Analytic rule ids of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentRelatedAnalyticRuleIds = AutomationRulePropertyConditionSupportedProperty("IncidentRelatedAnalyticRuleIds")
	// The provider name of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentProviderName = AutomationRulePropertyConditionSupportedProperty("IncidentProviderName")
	// The account Azure Active Directory tenant id
	AutomationRulePropertyConditionSupportedPropertyAccountAadTenantId = AutomationRulePropertyConditionSupportedProperty("AccountAadTenantId")
	// The account Azure Active Directory user id.
	AutomationRulePropertyConditionSupportedPropertyAccountAadUserId = AutomationRulePropertyConditionSupportedProperty("AccountAadUserId")
	// The account name
	AutomationRulePropertyConditionSupportedPropertyAccountName = AutomationRulePropertyConditionSupportedProperty("AccountName")
	// The account NetBIOS domain name
	AutomationRulePropertyConditionSupportedPropertyAccountNTDomain = AutomationRulePropertyConditionSupportedProperty("AccountNTDomain")
	// The account Azure Active Directory Passport User ID
	AutomationRulePropertyConditionSupportedPropertyAccountPUID = AutomationRulePropertyConditionSupportedProperty("AccountPUID")
	// The account security identifier
	AutomationRulePropertyConditionSupportedPropertyAccountSid = AutomationRulePropertyConditionSupportedProperty("AccountSid")
	// The account unique identifier
	AutomationRulePropertyConditionSupportedPropertyAccountObjectGuid = AutomationRulePropertyConditionSupportedProperty("AccountObjectGuid")
	// The account user principal name suffix
	AutomationRulePropertyConditionSupportedPropertyAccountUPNSuffix = AutomationRulePropertyConditionSupportedProperty("AccountUPNSuffix")
	// The Azure resource id
	AutomationRulePropertyConditionSupportedPropertyAzureResourceResourceId = AutomationRulePropertyConditionSupportedProperty("AzureResourceResourceId")
	// The Azure resource subscription id
	AutomationRulePropertyConditionSupportedPropertyAzureResourceSubscriptionId = AutomationRulePropertyConditionSupportedProperty("AzureResourceSubscriptionId")
	// The cloud application identifier
	AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppId = AutomationRulePropertyConditionSupportedProperty("CloudApplicationAppId")
	// The cloud application name
	AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppName = AutomationRulePropertyConditionSupportedProperty("CloudApplicationAppName")
	// The dns record domain name
	AutomationRulePropertyConditionSupportedPropertyDNSDomainName = AutomationRulePropertyConditionSupportedProperty("DNSDomainName")
	// The file directory full path
	AutomationRulePropertyConditionSupportedPropertyFileDirectory = AutomationRulePropertyConditionSupportedProperty("FileDirectory")
	// The file name without path
	AutomationRulePropertyConditionSupportedPropertyFileName = AutomationRulePropertyConditionSupportedProperty("FileName")
	// The file hash value
	AutomationRulePropertyConditionSupportedPropertyFileHashValue = AutomationRulePropertyConditionSupportedProperty("FileHashValue")
	// The host Azure resource id
	AutomationRulePropertyConditionSupportedPropertyHostAzureID = AutomationRulePropertyConditionSupportedProperty("HostAzureID")
	// The host name without domain
	AutomationRulePropertyConditionSupportedPropertyHostName = AutomationRulePropertyConditionSupportedProperty("HostName")
	// The host NetBIOS name
	AutomationRulePropertyConditionSupportedPropertyHostNetBiosName = AutomationRulePropertyConditionSupportedProperty("HostNetBiosName")
	// The host NT domain
	AutomationRulePropertyConditionSupportedPropertyHostNTDomain = AutomationRulePropertyConditionSupportedProperty("HostNTDomain")
	// The host operating system
	AutomationRulePropertyConditionSupportedPropertyHostOSVersion = AutomationRulePropertyConditionSupportedProperty("HostOSVersion")
	// The IoT device id
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceId = AutomationRulePropertyConditionSupportedProperty("IoTDeviceId")
	// The IoT device name
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceName = AutomationRulePropertyConditionSupportedProperty("IoTDeviceName")
	// The IoT device type
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceType = AutomationRulePropertyConditionSupportedProperty("IoTDeviceType")
	// The IoT device vendor
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceVendor = AutomationRulePropertyConditionSupportedProperty("IoTDeviceVendor")
	// The IoT device model
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceModel = AutomationRulePropertyConditionSupportedProperty("IoTDeviceModel")
	// The IoT device operating system
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceOperatingSystem = AutomationRulePropertyConditionSupportedProperty("IoTDeviceOperatingSystem")
	// The IP address
	AutomationRulePropertyConditionSupportedPropertyIPAddress = AutomationRulePropertyConditionSupportedProperty("IPAddress")
	// The mailbox display name
	AutomationRulePropertyConditionSupportedPropertyMailboxDisplayName = AutomationRulePropertyConditionSupportedProperty("MailboxDisplayName")
	// The mailbox primary address
	AutomationRulePropertyConditionSupportedPropertyMailboxPrimaryAddress = AutomationRulePropertyConditionSupportedProperty("MailboxPrimaryAddress")
	// The mailbox user principal name
	AutomationRulePropertyConditionSupportedPropertyMailboxUPN = AutomationRulePropertyConditionSupportedProperty("MailboxUPN")
	// The mail message delivery action
	AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryAction = AutomationRulePropertyConditionSupportedProperty("MailMessageDeliveryAction")
	// The mail message delivery location
	AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryLocation = AutomationRulePropertyConditionSupportedProperty("MailMessageDeliveryLocation")
	// The mail message recipient
	AutomationRulePropertyConditionSupportedPropertyMailMessageRecipient = AutomationRulePropertyConditionSupportedProperty("MailMessageRecipient")
	// The mail message sender IP address
	AutomationRulePropertyConditionSupportedPropertyMailMessageSenderIP = AutomationRulePropertyConditionSupportedProperty("MailMessageSenderIP")
	// The mail message subject
	AutomationRulePropertyConditionSupportedPropertyMailMessageSubject = AutomationRulePropertyConditionSupportedProperty("MailMessageSubject")
	// The mail message P1 sender
	AutomationRulePropertyConditionSupportedPropertyMailMessageP1Sender = AutomationRulePropertyConditionSupportedProperty("MailMessageP1Sender")
	// The mail message P2 sender
	AutomationRulePropertyConditionSupportedPropertyMailMessageP2Sender = AutomationRulePropertyConditionSupportedProperty("MailMessageP2Sender")
	// The malware category
	AutomationRulePropertyConditionSupportedPropertyMalwareCategory = AutomationRulePropertyConditionSupportedProperty("MalwareCategory")
	// The malware name
	AutomationRulePropertyConditionSupportedPropertyMalwareName = AutomationRulePropertyConditionSupportedProperty("MalwareName")
	// The process execution command line
	AutomationRulePropertyConditionSupportedPropertyProcessCommandLine = AutomationRulePropertyConditionSupportedProperty("ProcessCommandLine")
	// The process id
	AutomationRulePropertyConditionSupportedPropertyProcessId = AutomationRulePropertyConditionSupportedProperty("ProcessId")
	// The registry key path
	AutomationRulePropertyConditionSupportedPropertyRegistryKey = AutomationRulePropertyConditionSupportedProperty("RegistryKey")
	// The registry key value in string formatted representation
	AutomationRulePropertyConditionSupportedPropertyRegistryValueData = AutomationRulePropertyConditionSupportedProperty("RegistryValueData")
	// The url
	AutomationRulePropertyConditionSupportedPropertyUrl = AutomationRulePropertyConditionSupportedProperty("Url")
)

func (AutomationRulePropertyConditionSupportedProperty) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AutomationRulePropertyConditionSupportedProperty) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRulePropertyConditionSupportedProperty) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRulePropertyConditionSupportedProperty) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutomationRulePropertyConditionSupportedProperty) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The severity of the incident
type CaseSeverity pulumi.String

const (
	// Critical severity
	CaseSeverityCritical = CaseSeverity("Critical")
	// High severity
	CaseSeverityHigh = CaseSeverity("High")
	// Medium severity
	CaseSeverityMedium = CaseSeverity("Medium")
	// Low severity
	CaseSeverityLow = CaseSeverity("Low")
	// Informational severity
	CaseSeverityInformational = CaseSeverity("Informational")
)

func (CaseSeverity) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e CaseSeverity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CaseSeverity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CaseSeverity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CaseSeverity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The kind of the data connector
type DataConnectorKind pulumi.String

const (
	DataConnectorKindAzureActiveDirectory                      = DataConnectorKind("AzureActiveDirectory")
	DataConnectorKindAzureSecurityCenter                       = DataConnectorKind("AzureSecurityCenter")
	DataConnectorKindMicrosoftCloudAppSecurity                 = DataConnectorKind("MicrosoftCloudAppSecurity")
	DataConnectorKindThreatIntelligence                        = DataConnectorKind("ThreatIntelligence")
	DataConnectorKindThreatIntelligenceTaxii                   = DataConnectorKind("ThreatIntelligenceTaxii")
	DataConnectorKindOffice365                                 = DataConnectorKind("Office365")
	DataConnectorKindOfficeATP                                 = DataConnectorKind("OfficeATP")
	DataConnectorKindAmazonWebServicesCloudTrail               = DataConnectorKind("AmazonWebServicesCloudTrail")
	DataConnectorKindAzureAdvancedThreatProtection             = DataConnectorKind("AzureAdvancedThreatProtection")
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection = DataConnectorKind("MicrosoftDefenderAdvancedThreatProtection")
	DataConnectorKindDynamics365                               = DataConnectorKind("Dynamics365")
	DataConnectorKindMicrosoftThreatProtection                 = DataConnectorKind("MicrosoftThreatProtection")
	DataConnectorKindMicrosoftThreatIntelligence               = DataConnectorKind("MicrosoftThreatIntelligence")
)

func (DataConnectorKind) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e DataConnectorKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataConnectorKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataConnectorKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataConnectorKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The entity query kind
type EntityTimelineKind pulumi.String

const (
	// activity
	EntityTimelineKindActivity = EntityTimelineKind("Activity")
	// bookmarks
	EntityTimelineKindBookmark = EntityTimelineKind("Bookmark")
	// security alerts
	EntityTimelineKindSecurityAlert = EntityTimelineKind("SecurityAlert")
)

func (EntityTimelineKind) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e EntityTimelineKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityTimelineKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityTimelineKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EntityTimelineKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The reason the incident was closed
type IncidentClassification pulumi.String

const (
	// Incident classification was undetermined
	IncidentClassificationUndetermined = IncidentClassification("Undetermined")
	// Incident was true positive
	IncidentClassificationTruePositive = IncidentClassification("TruePositive")
	// Incident was benign positive
	IncidentClassificationBenignPositive = IncidentClassification("BenignPositive")
	// Incident was false positive
	IncidentClassificationFalsePositive = IncidentClassification("FalsePositive")
)

func (IncidentClassification) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e IncidentClassification) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassification) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassification) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentClassification) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The classification reason the incident was closed with
type IncidentClassificationReason pulumi.String

const (
	// Classification reason was suspicious activity
	IncidentClassificationReasonSuspiciousActivity = IncidentClassificationReason("SuspiciousActivity")
	// Classification reason was suspicious but expected
	IncidentClassificationReasonSuspiciousButExpected = IncidentClassificationReason("SuspiciousButExpected")
	// Classification reason was incorrect alert logic
	IncidentClassificationReasonIncorrectAlertLogic = IncidentClassificationReason("IncorrectAlertLogic")
	// Classification reason was inaccurate data
	IncidentClassificationReasonInaccurateData = IncidentClassificationReason("InaccurateData")
)

func (IncidentClassificationReason) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e IncidentClassificationReason) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassificationReason) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassificationReason) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentClassificationReason) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The severity of the incident
type IncidentSeverity pulumi.String

const (
	// High severity
	IncidentSeverityHigh = IncidentSeverity("High")
	// Medium severity
	IncidentSeverityMedium = IncidentSeverity("Medium")
	// Low severity
	IncidentSeverityLow = IncidentSeverity("Low")
	// Informational severity
	IncidentSeverityInformational = IncidentSeverity("Informational")
)

func (IncidentSeverity) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e IncidentSeverity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentSeverity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentSeverity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentSeverity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The status of the incident
type IncidentStatus pulumi.String

const (
	// An active incident which isn't being handled currently
	IncidentStatusNew = IncidentStatus("New")
	// An active incident which is being handled
	IncidentStatusActive = IncidentStatus("Active")
	// A non-active incident
	IncidentStatusClosed = IncidentStatus("Closed")
)

func (IncidentStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e IncidentStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The kind of the setting
type SettingKind pulumi.String

const (
	SettingKindEyesOn          = SettingKind("EyesOn")
	SettingKindEntityAnalytics = SettingKind("EntityAnalytics")
	SettingKindUeba            = SettingKind("Ueba")
)

func (SettingKind) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e SettingKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SettingKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SettingKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SettingKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The source of the watchlist
type Source pulumi.String

const (
	Source_Local_file     = Source("Local file")
	Source_Remote_storage = Source("Remote storage")
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

// The kind of the entity.
type ThreatIntelligenceResourceKind pulumi.String

const (
	// Entity represents threat intelligence indicator in the system.
	ThreatIntelligenceResourceKindIndicator = ThreatIntelligenceResourceKind("indicator")
)

func (ThreatIntelligenceResourceKind) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ThreatIntelligenceResourceKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ThreatIntelligenceResourceKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ThreatIntelligenceResourceKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ThreatIntelligenceResourceKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The type of object the automation rule triggers on
type TriggersOn pulumi.String

const (
	// Trigger on Incidents
	TriggersOnIncidents = TriggersOn("Incidents")
)

func (TriggersOn) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e TriggersOn) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggersOn) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggersOn) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggersOn) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The type of event the automation rule triggers on
type TriggersWhen pulumi.String

const (
	// Trigger on created objects
	TriggersWhenCreated = TriggersWhen("Created")
)

func (TriggersWhen) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e TriggersWhen) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggersWhen) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggersWhen) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggersWhen) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
