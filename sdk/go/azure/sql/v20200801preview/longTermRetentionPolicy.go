// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A long term retention policy.
type LongTermRetentionPolicy struct {
	pulumi.CustomResourceState

	// The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention pulumi.StringPtrOutput `pulumi:"monthlyRetention"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear pulumi.IntPtrOutput `pulumi:"weekOfYear"`
	// The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention pulumi.StringPtrOutput `pulumi:"weeklyRetention"`
	// The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention pulumi.StringPtrOutput `pulumi:"yearlyRetention"`
}

// NewLongTermRetentionPolicy registers a new resource with the given unique name, arguments, and options.
func NewLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, args *LongTermRetentionPolicyArgs, opts ...pulumi.ResourceOption) (*LongTermRetentionPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20170301preview:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:LongTermRetentionPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource LongTermRetentionPolicy
	err := ctx.RegisterResource("azure-native:sql/v20200801preview:LongTermRetentionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLongTermRetentionPolicy gets an existing LongTermRetentionPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LongTermRetentionPolicyState, opts ...pulumi.ResourceOption) (*LongTermRetentionPolicy, error) {
	var resource LongTermRetentionPolicy
	err := ctx.ReadResource("azure-native:sql/v20200801preview:LongTermRetentionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LongTermRetentionPolicy resources.
type longTermRetentionPolicyState struct {
	// The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention *string `pulumi:"monthlyRetention"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear *int `pulumi:"weekOfYear"`
	// The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention *string `pulumi:"weeklyRetention"`
	// The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention *string `pulumi:"yearlyRetention"`
}

type LongTermRetentionPolicyState struct {
	// The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear pulumi.IntPtrInput
	// The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention pulumi.StringPtrInput
	// The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention pulumi.StringPtrInput
}

func (LongTermRetentionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*longTermRetentionPolicyState)(nil)).Elem()
}

type longTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention *string `pulumi:"monthlyRetention"`
	// The policy name. Should always be Default.
	PolicyName *string `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear *int `pulumi:"weekOfYear"`
	// The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention *string `pulumi:"weeklyRetention"`
	// The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention *string `pulumi:"yearlyRetention"`
}

// The set of arguments for constructing a LongTermRetentionPolicy resource.
type LongTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention pulumi.StringPtrInput
	// The policy name. Should always be Default.
	PolicyName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear pulumi.IntPtrInput
	// The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention pulumi.StringPtrInput
	// The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention pulumi.StringPtrInput
}

func (LongTermRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*longTermRetentionPolicyArgs)(nil)).Elem()
}

type LongTermRetentionPolicyInput interface {
	pulumi.Input

	ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput
	ToLongTermRetentionPolicyOutputWithContext(ctx context.Context) LongTermRetentionPolicyOutput
}

func (*LongTermRetentionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermRetentionPolicy)(nil))
}

func (i *LongTermRetentionPolicy) ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput {
	return i.ToLongTermRetentionPolicyOutputWithContext(context.Background())
}

func (i *LongTermRetentionPolicy) ToLongTermRetentionPolicyOutputWithContext(ctx context.Context) LongTermRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LongTermRetentionPolicyOutput)
}

type LongTermRetentionPolicyOutput struct {
	*pulumi.OutputState
}

func (LongTermRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermRetentionPolicy)(nil))
}

func (o LongTermRetentionPolicyOutput) ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput {
	return o
}

func (o LongTermRetentionPolicyOutput) ToLongTermRetentionPolicyOutputWithContext(ctx context.Context) LongTermRetentionPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LongTermRetentionPolicyOutput{})
}
