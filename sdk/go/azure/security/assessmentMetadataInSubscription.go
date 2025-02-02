// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Security assessment metadata
// API Version: 2020-01-01.
type AssessmentMetadataInSubscription struct {
	pulumi.CustomResourceState

	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType pulumi.StringOutput      `pulumi:"assessmentType"`
	Categories     pulumi.StringArrayOutput `pulumi:"categories"`
	// Human readable description of the assessment
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User friendly display name of the assessment
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The implementation effort required to remediate this assessment
	ImplementationEffort pulumi.StringPtrOutput `pulumi:"implementationEffort"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes the partner that created the assessment
	PartnerData SecurityAssessmentMetadataPartnerDataResponsePtrOutput `pulumi:"partnerData"`
	// Azure resource ID of the policy definition that turns this assessment calculation on
	PolicyDefinitionId pulumi.StringOutput `pulumi:"policyDefinitionId"`
	// True if this assessment is in preview release status
	Preview pulumi.BoolPtrOutput `pulumi:"preview"`
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription pulumi.StringPtrOutput `pulumi:"remediationDescription"`
	// The severity level of the assessment
	Severity pulumi.StringOutput      `pulumi:"severity"`
	Threats  pulumi.StringArrayOutput `pulumi:"threats"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The user impact of the assessment
	UserImpact pulumi.StringPtrOutput `pulumi:"userImpact"`
}

// NewAssessmentMetadataInSubscription registers a new resource with the given unique name, arguments, and options.
func NewAssessmentMetadataInSubscription(ctx *pulumi.Context,
	name string, args *AssessmentMetadataInSubscriptionArgs, opts ...pulumi.ResourceOption) (*AssessmentMetadataInSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssessmentType == nil {
		return nil, errors.New("invalid value for required argument 'AssessmentType'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:security:AssessmentMetadataInSubscription"),
		},
		{
			Type: pulumi.String("azure-native:security/v20190101preview:AssessmentMetadataInSubscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:security/v20190101preview:AssessmentMetadataInSubscription"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101:AssessmentMetadataInSubscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:security/v20200101:AssessmentMetadataInSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource AssessmentMetadataInSubscription
	err := ctx.RegisterResource("azure-native:security:AssessmentMetadataInSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssessmentMetadataInSubscription gets an existing AssessmentMetadataInSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssessmentMetadataInSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentMetadataInSubscriptionState, opts ...pulumi.ResourceOption) (*AssessmentMetadataInSubscription, error) {
	var resource AssessmentMetadataInSubscription
	err := ctx.ReadResource("azure-native:security:AssessmentMetadataInSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssessmentMetadataInSubscription resources.
type assessmentMetadataInSubscriptionState struct {
	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType *string  `pulumi:"assessmentType"`
	Categories     []string `pulumi:"categories"`
	// Human readable description of the assessment
	Description *string `pulumi:"description"`
	// User friendly display name of the assessment
	DisplayName *string `pulumi:"displayName"`
	// The implementation effort required to remediate this assessment
	ImplementationEffort *string `pulumi:"implementationEffort"`
	// Resource name
	Name *string `pulumi:"name"`
	// Describes the partner that created the assessment
	PartnerData *SecurityAssessmentMetadataPartnerDataResponse `pulumi:"partnerData"`
	// Azure resource ID of the policy definition that turns this assessment calculation on
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
	// True if this assessment is in preview release status
	Preview *bool `pulumi:"preview"`
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription *string `pulumi:"remediationDescription"`
	// The severity level of the assessment
	Severity *string  `pulumi:"severity"`
	Threats  []string `pulumi:"threats"`
	// Resource type
	Type *string `pulumi:"type"`
	// The user impact of the assessment
	UserImpact *string `pulumi:"userImpact"`
}

type AssessmentMetadataInSubscriptionState struct {
	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType pulumi.StringPtrInput
	Categories     pulumi.StringArrayInput
	// Human readable description of the assessment
	Description pulumi.StringPtrInput
	// User friendly display name of the assessment
	DisplayName pulumi.StringPtrInput
	// The implementation effort required to remediate this assessment
	ImplementationEffort pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Describes the partner that created the assessment
	PartnerData SecurityAssessmentMetadataPartnerDataResponsePtrInput
	// Azure resource ID of the policy definition that turns this assessment calculation on
	PolicyDefinitionId pulumi.StringPtrInput
	// True if this assessment is in preview release status
	Preview pulumi.BoolPtrInput
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription pulumi.StringPtrInput
	// The severity level of the assessment
	Severity pulumi.StringPtrInput
	Threats  pulumi.StringArrayInput
	// Resource type
	Type pulumi.StringPtrInput
	// The user impact of the assessment
	UserImpact pulumi.StringPtrInput
}

func (AssessmentMetadataInSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentMetadataInSubscriptionState)(nil)).Elem()
}

type assessmentMetadataInSubscriptionArgs struct {
	// The Assessment Key - Unique key for the assessment type
	AssessmentMetadataName *string `pulumi:"assessmentMetadataName"`
	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType string   `pulumi:"assessmentType"`
	Categories     []string `pulumi:"categories"`
	// Human readable description of the assessment
	Description *string `pulumi:"description"`
	// User friendly display name of the assessment
	DisplayName string `pulumi:"displayName"`
	// The implementation effort required to remediate this assessment
	ImplementationEffort *string `pulumi:"implementationEffort"`
	// Describes the partner that created the assessment
	PartnerData *SecurityAssessmentMetadataPartnerData `pulumi:"partnerData"`
	// True if this assessment is in preview release status
	Preview *bool `pulumi:"preview"`
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription *string `pulumi:"remediationDescription"`
	// The severity level of the assessment
	Severity string   `pulumi:"severity"`
	Threats  []string `pulumi:"threats"`
	// The user impact of the assessment
	UserImpact *string `pulumi:"userImpact"`
}

// The set of arguments for constructing a AssessmentMetadataInSubscription resource.
type AssessmentMetadataInSubscriptionArgs struct {
	// The Assessment Key - Unique key for the assessment type
	AssessmentMetadataName pulumi.StringPtrInput
	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType pulumi.StringInput
	Categories     pulumi.StringArrayInput
	// Human readable description of the assessment
	Description pulumi.StringPtrInput
	// User friendly display name of the assessment
	DisplayName pulumi.StringInput
	// The implementation effort required to remediate this assessment
	ImplementationEffort pulumi.StringPtrInput
	// Describes the partner that created the assessment
	PartnerData SecurityAssessmentMetadataPartnerDataPtrInput
	// True if this assessment is in preview release status
	Preview pulumi.BoolPtrInput
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription pulumi.StringPtrInput
	// The severity level of the assessment
	Severity pulumi.StringInput
	Threats  pulumi.StringArrayInput
	// The user impact of the assessment
	UserImpact pulumi.StringPtrInput
}

func (AssessmentMetadataInSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentMetadataInSubscriptionArgs)(nil)).Elem()
}

type AssessmentMetadataInSubscriptionInput interface {
	pulumi.Input

	ToAssessmentMetadataInSubscriptionOutput() AssessmentMetadataInSubscriptionOutput
	ToAssessmentMetadataInSubscriptionOutputWithContext(ctx context.Context) AssessmentMetadataInSubscriptionOutput
}

func (*AssessmentMetadataInSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentMetadataInSubscription)(nil))
}

func (i *AssessmentMetadataInSubscription) ToAssessmentMetadataInSubscriptionOutput() AssessmentMetadataInSubscriptionOutput {
	return i.ToAssessmentMetadataInSubscriptionOutputWithContext(context.Background())
}

func (i *AssessmentMetadataInSubscription) ToAssessmentMetadataInSubscriptionOutputWithContext(ctx context.Context) AssessmentMetadataInSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentMetadataInSubscriptionOutput)
}

type AssessmentMetadataInSubscriptionOutput struct {
	*pulumi.OutputState
}

func (AssessmentMetadataInSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentMetadataInSubscription)(nil))
}

func (o AssessmentMetadataInSubscriptionOutput) ToAssessmentMetadataInSubscriptionOutput() AssessmentMetadataInSubscriptionOutput {
	return o
}

func (o AssessmentMetadataInSubscriptionOutput) ToAssessmentMetadataInSubscriptionOutputWithContext(ctx context.Context) AssessmentMetadataInSubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AssessmentMetadataInSubscriptionOutput{})
}
