// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An action group resource.
type ActionGroup struct {
	pulumi.CustomResourceState

	// The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
	ArmRoleReceivers ArmRoleReceiverResponseArrayOutput `pulumi:"armRoleReceivers"`
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers AutomationRunbookReceiverResponseArrayOutput `pulumi:"automationRunbookReceivers"`
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers AzureAppPushReceiverResponseArrayOutput `pulumi:"azureAppPushReceivers"`
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers AzureFunctionReceiverResponseArrayOutput `pulumi:"azureFunctionReceivers"`
	// The list of email receivers that are part of this action group.
	EmailReceivers EmailReceiverResponseArrayOutput `pulumi:"emailReceivers"`
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName pulumi.StringOutput `pulumi:"groupShortName"`
	// Azure resource identity
	Identity pulumi.StringOutput `pulumi:"identity"`
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers ItsmReceiverResponseArrayOutput `pulumi:"itsmReceivers"`
	// Azure resource kind
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers LogicAppReceiverResponseArrayOutput `pulumi:"logicAppReceivers"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of SMS receivers that are part of this action group.
	SmsReceivers SmsReceiverResponseArrayOutput `pulumi:"smsReceivers"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of voice receivers that are part of this action group.
	VoiceReceivers VoiceReceiverResponseArrayOutput `pulumi:"voiceReceivers"`
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers WebhookReceiverResponseArrayOutput `pulumi:"webhookReceivers"`
}

// NewActionGroup registers a new resource with the given unique name, arguments, and options.
func NewActionGroup(ctx *pulumi.Context,
	name string, args *ActionGroupArgs, opts ...pulumi.ResourceOption) (*ActionGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.GroupShortName == nil {
		return nil, errors.New("invalid value for required argument 'GroupShortName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Enabled == nil {
		args.Enabled = pulumi.Bool(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights/v20190301:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20170401:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20170401:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180301:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20180301:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180901:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20180901:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20190601:ActionGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20190601:ActionGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ActionGroup
	err := ctx.RegisterResource("azure-native:insights/v20190301:ActionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionGroup gets an existing ActionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionGroupState, opts ...pulumi.ResourceOption) (*ActionGroup, error) {
	var resource ActionGroup
	err := ctx.ReadResource("azure-native:insights/v20190301:ActionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionGroup resources.
type actionGroupState struct {
	// The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
	ArmRoleReceivers []ArmRoleReceiverResponse `pulumi:"armRoleReceivers"`
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers []AutomationRunbookReceiverResponse `pulumi:"automationRunbookReceivers"`
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers []AzureAppPushReceiverResponse `pulumi:"azureAppPushReceivers"`
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers []AzureFunctionReceiverResponse `pulumi:"azureFunctionReceivers"`
	// The list of email receivers that are part of this action group.
	EmailReceivers []EmailReceiverResponse `pulumi:"emailReceivers"`
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled *bool `pulumi:"enabled"`
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName *string `pulumi:"groupShortName"`
	// Azure resource identity
	Identity *string `pulumi:"identity"`
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers []ItsmReceiverResponse `pulumi:"itsmReceivers"`
	// Azure resource kind
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers []LogicAppReceiverResponse `pulumi:"logicAppReceivers"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// The list of SMS receivers that are part of this action group.
	SmsReceivers []SmsReceiverResponse `pulumi:"smsReceivers"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type *string `pulumi:"type"`
	// The list of voice receivers that are part of this action group.
	VoiceReceivers []VoiceReceiverResponse `pulumi:"voiceReceivers"`
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers []WebhookReceiverResponse `pulumi:"webhookReceivers"`
}

type ActionGroupState struct {
	// The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
	ArmRoleReceivers ArmRoleReceiverResponseArrayInput
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers AutomationRunbookReceiverResponseArrayInput
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers AzureAppPushReceiverResponseArrayInput
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers AzureFunctionReceiverResponseArrayInput
	// The list of email receivers that are part of this action group.
	EmailReceivers EmailReceiverResponseArrayInput
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled pulumi.BoolPtrInput
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName pulumi.StringPtrInput
	// Azure resource identity
	Identity pulumi.StringPtrInput
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers ItsmReceiverResponseArrayInput
	// Azure resource kind
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers LogicAppReceiverResponseArrayInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// The list of SMS receivers that are part of this action group.
	SmsReceivers SmsReceiverResponseArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Azure resource type
	Type pulumi.StringPtrInput
	// The list of voice receivers that are part of this action group.
	VoiceReceivers VoiceReceiverResponseArrayInput
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers WebhookReceiverResponseArrayInput
}

func (ActionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionGroupState)(nil)).Elem()
}

type actionGroupArgs struct {
	// The name of the action group.
	ActionGroupName *string `pulumi:"actionGroupName"`
	// The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
	ArmRoleReceivers []ArmRoleReceiver `pulumi:"armRoleReceivers"`
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers []AutomationRunbookReceiver `pulumi:"automationRunbookReceivers"`
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers []AzureAppPushReceiver `pulumi:"azureAppPushReceivers"`
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers []AzureFunctionReceiver `pulumi:"azureFunctionReceivers"`
	// The list of email receivers that are part of this action group.
	EmailReceivers []EmailReceiver `pulumi:"emailReceivers"`
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled bool `pulumi:"enabled"`
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName string `pulumi:"groupShortName"`
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers []ItsmReceiver `pulumi:"itsmReceivers"`
	// Resource location
	Location *string `pulumi:"location"`
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers []LogicAppReceiver `pulumi:"logicAppReceivers"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of SMS receivers that are part of this action group.
	SmsReceivers []SmsReceiver `pulumi:"smsReceivers"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The list of voice receivers that are part of this action group.
	VoiceReceivers []VoiceReceiver `pulumi:"voiceReceivers"`
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers []WebhookReceiver `pulumi:"webhookReceivers"`
}

// The set of arguments for constructing a ActionGroup resource.
type ActionGroupArgs struct {
	// The name of the action group.
	ActionGroupName pulumi.StringPtrInput
	// The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
	ArmRoleReceivers ArmRoleReceiverArrayInput
	// The list of AutomationRunbook receivers that are part of this action group.
	AutomationRunbookReceivers AutomationRunbookReceiverArrayInput
	// The list of AzureAppPush receivers that are part of this action group.
	AzureAppPushReceivers AzureAppPushReceiverArrayInput
	// The list of azure function receivers that are part of this action group.
	AzureFunctionReceivers AzureFunctionReceiverArrayInput
	// The list of email receivers that are part of this action group.
	EmailReceivers EmailReceiverArrayInput
	// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
	Enabled pulumi.BoolInput
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName pulumi.StringInput
	// The list of ITSM receivers that are part of this action group.
	ItsmReceivers ItsmReceiverArrayInput
	// Resource location
	Location pulumi.StringPtrInput
	// The list of logic app receivers that are part of this action group.
	LogicAppReceivers LogicAppReceiverArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The list of SMS receivers that are part of this action group.
	SmsReceivers SmsReceiverArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The list of voice receivers that are part of this action group.
	VoiceReceivers VoiceReceiverArrayInput
	// The list of webhook receivers that are part of this action group.
	WebhookReceivers WebhookReceiverArrayInput
}

func (ActionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionGroupArgs)(nil)).Elem()
}

type ActionGroupInput interface {
	pulumi.Input

	ToActionGroupOutput() ActionGroupOutput
	ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput
}

func (*ActionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroup)(nil))
}

func (i *ActionGroup) ToActionGroupOutput() ActionGroupOutput {
	return i.ToActionGroupOutputWithContext(context.Background())
}

func (i *ActionGroup) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupOutput)
}

type ActionGroupOutput struct {
	*pulumi.OutputState
}

func (ActionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroup)(nil))
}

func (o ActionGroupOutput) ToActionGroupOutput() ActionGroupOutput {
	return o
}

func (o ActionGroupOutput) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ActionGroupOutput{})
}
