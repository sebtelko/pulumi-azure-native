// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Redis cache access keys.
type RedisAccessKeysResponse struct {
	// The current primary key that clients can use to authenticate with Redis cache.
	PrimaryKey string `pulumi:"primaryKey"`
	// The current secondary key that clients can use to authenticate with Redis cache.
	SecondaryKey string `pulumi:"secondaryKey"`
}

// RedisAccessKeysResponseInput is an input type that accepts RedisAccessKeysResponseArgs and RedisAccessKeysResponseOutput values.
// You can construct a concrete instance of `RedisAccessKeysResponseInput` via:
//
//          RedisAccessKeysResponseArgs{...}
type RedisAccessKeysResponseInput interface {
	pulumi.Input

	ToRedisAccessKeysResponseOutput() RedisAccessKeysResponseOutput
	ToRedisAccessKeysResponseOutputWithContext(context.Context) RedisAccessKeysResponseOutput
}

// Redis cache access keys.
type RedisAccessKeysResponseArgs struct {
	// The current primary key that clients can use to authenticate with Redis cache.
	PrimaryKey pulumi.StringInput `pulumi:"primaryKey"`
	// The current secondary key that clients can use to authenticate with Redis cache.
	SecondaryKey pulumi.StringInput `pulumi:"secondaryKey"`
}

func (RedisAccessKeysResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisAccessKeysResponse)(nil)).Elem()
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponseOutput() RedisAccessKeysResponseOutput {
	return i.ToRedisAccessKeysResponseOutputWithContext(context.Background())
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponseOutputWithContext(ctx context.Context) RedisAccessKeysResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisAccessKeysResponseOutput)
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return i.ToRedisAccessKeysResponsePtrOutputWithContext(context.Background())
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisAccessKeysResponseOutput).ToRedisAccessKeysResponsePtrOutputWithContext(ctx)
}

// RedisAccessKeysResponsePtrInput is an input type that accepts RedisAccessKeysResponseArgs, RedisAccessKeysResponsePtr and RedisAccessKeysResponsePtrOutput values.
// You can construct a concrete instance of `RedisAccessKeysResponsePtrInput` via:
//
//          RedisAccessKeysResponseArgs{...}
//
//  or:
//
//          nil
type RedisAccessKeysResponsePtrInput interface {
	pulumi.Input

	ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput
	ToRedisAccessKeysResponsePtrOutputWithContext(context.Context) RedisAccessKeysResponsePtrOutput
}

type redisAccessKeysResponsePtrType RedisAccessKeysResponseArgs

func RedisAccessKeysResponsePtr(v *RedisAccessKeysResponseArgs) RedisAccessKeysResponsePtrInput {
	return (*redisAccessKeysResponsePtrType)(v)
}

func (*redisAccessKeysResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisAccessKeysResponse)(nil)).Elem()
}

func (i *redisAccessKeysResponsePtrType) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return i.ToRedisAccessKeysResponsePtrOutputWithContext(context.Background())
}

func (i *redisAccessKeysResponsePtrType) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisAccessKeysResponsePtrOutput)
}

// Redis cache access keys.
type RedisAccessKeysResponseOutput struct{ *pulumi.OutputState }

func (RedisAccessKeysResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisAccessKeysResponse)(nil)).Elem()
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponseOutput() RedisAccessKeysResponseOutput {
	return o
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponseOutputWithContext(ctx context.Context) RedisAccessKeysResponseOutput {
	return o
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return o.ToRedisAccessKeysResponsePtrOutputWithContext(context.Background())
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return o.ApplyT(func(v RedisAccessKeysResponse) *RedisAccessKeysResponse {
		return &v
	}).(RedisAccessKeysResponsePtrOutput)
}

// The current primary key that clients can use to authenticate with Redis cache.
func (o RedisAccessKeysResponseOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v RedisAccessKeysResponse) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// The current secondary key that clients can use to authenticate with Redis cache.
func (o RedisAccessKeysResponseOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v RedisAccessKeysResponse) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

type RedisAccessKeysResponsePtrOutput struct{ *pulumi.OutputState }

func (RedisAccessKeysResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisAccessKeysResponse)(nil)).Elem()
}

func (o RedisAccessKeysResponsePtrOutput) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return o
}

func (o RedisAccessKeysResponsePtrOutput) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return o
}

func (o RedisAccessKeysResponsePtrOutput) Elem() RedisAccessKeysResponseOutput {
	return o.ApplyT(func(v *RedisAccessKeysResponse) RedisAccessKeysResponse { return *v }).(RedisAccessKeysResponseOutput)
}

// The current primary key that clients can use to authenticate with Redis cache.
func (o RedisAccessKeysResponsePtrOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisAccessKeysResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrimaryKey
	}).(pulumi.StringPtrOutput)
}

// The current secondary key that clients can use to authenticate with Redis cache.
func (o RedisAccessKeysResponsePtrOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisAccessKeysResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecondaryKey
	}).(pulumi.StringPtrOutput)
}

// Details of single instance of redis.
type RedisInstanceDetailsResponse struct {
	// Specifies whether the instance is a master node.
	IsMaster bool `pulumi:"isMaster"`
	// If enableNonSslPort is true, provides Redis instance Non-SSL port.
	NonSslPort int `pulumi:"nonSslPort"`
	// If clustering is enabled, the Shard ID of Redis Instance
	ShardId int `pulumi:"shardId"`
	// Redis instance SSL port.
	SslPort int `pulumi:"sslPort"`
	// If the Cache uses availability zones, specifies availability zone where this instance is located.
	Zone string `pulumi:"zone"`
}

// RedisInstanceDetailsResponseInput is an input type that accepts RedisInstanceDetailsResponseArgs and RedisInstanceDetailsResponseOutput values.
// You can construct a concrete instance of `RedisInstanceDetailsResponseInput` via:
//
//          RedisInstanceDetailsResponseArgs{...}
type RedisInstanceDetailsResponseInput interface {
	pulumi.Input

	ToRedisInstanceDetailsResponseOutput() RedisInstanceDetailsResponseOutput
	ToRedisInstanceDetailsResponseOutputWithContext(context.Context) RedisInstanceDetailsResponseOutput
}

// Details of single instance of redis.
type RedisInstanceDetailsResponseArgs struct {
	// Specifies whether the instance is a master node.
	IsMaster pulumi.BoolInput `pulumi:"isMaster"`
	// If enableNonSslPort is true, provides Redis instance Non-SSL port.
	NonSslPort pulumi.IntInput `pulumi:"nonSslPort"`
	// If clustering is enabled, the Shard ID of Redis Instance
	ShardId pulumi.IntInput `pulumi:"shardId"`
	// Redis instance SSL port.
	SslPort pulumi.IntInput `pulumi:"sslPort"`
	// If the Cache uses availability zones, specifies availability zone where this instance is located.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (RedisInstanceDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisInstanceDetailsResponse)(nil)).Elem()
}

func (i RedisInstanceDetailsResponseArgs) ToRedisInstanceDetailsResponseOutput() RedisInstanceDetailsResponseOutput {
	return i.ToRedisInstanceDetailsResponseOutputWithContext(context.Background())
}

func (i RedisInstanceDetailsResponseArgs) ToRedisInstanceDetailsResponseOutputWithContext(ctx context.Context) RedisInstanceDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisInstanceDetailsResponseOutput)
}

// RedisInstanceDetailsResponseArrayInput is an input type that accepts RedisInstanceDetailsResponseArray and RedisInstanceDetailsResponseArrayOutput values.
// You can construct a concrete instance of `RedisInstanceDetailsResponseArrayInput` via:
//
//          RedisInstanceDetailsResponseArray{ RedisInstanceDetailsResponseArgs{...} }
type RedisInstanceDetailsResponseArrayInput interface {
	pulumi.Input

	ToRedisInstanceDetailsResponseArrayOutput() RedisInstanceDetailsResponseArrayOutput
	ToRedisInstanceDetailsResponseArrayOutputWithContext(context.Context) RedisInstanceDetailsResponseArrayOutput
}

type RedisInstanceDetailsResponseArray []RedisInstanceDetailsResponseInput

func (RedisInstanceDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisInstanceDetailsResponse)(nil)).Elem()
}

func (i RedisInstanceDetailsResponseArray) ToRedisInstanceDetailsResponseArrayOutput() RedisInstanceDetailsResponseArrayOutput {
	return i.ToRedisInstanceDetailsResponseArrayOutputWithContext(context.Background())
}

func (i RedisInstanceDetailsResponseArray) ToRedisInstanceDetailsResponseArrayOutputWithContext(ctx context.Context) RedisInstanceDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisInstanceDetailsResponseArrayOutput)
}

// Details of single instance of redis.
type RedisInstanceDetailsResponseOutput struct{ *pulumi.OutputState }

func (RedisInstanceDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisInstanceDetailsResponse)(nil)).Elem()
}

func (o RedisInstanceDetailsResponseOutput) ToRedisInstanceDetailsResponseOutput() RedisInstanceDetailsResponseOutput {
	return o
}

func (o RedisInstanceDetailsResponseOutput) ToRedisInstanceDetailsResponseOutputWithContext(ctx context.Context) RedisInstanceDetailsResponseOutput {
	return o
}

// Specifies whether the instance is a master node.
func (o RedisInstanceDetailsResponseOutput) IsMaster() pulumi.BoolOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) bool { return v.IsMaster }).(pulumi.BoolOutput)
}

// If enableNonSslPort is true, provides Redis instance Non-SSL port.
func (o RedisInstanceDetailsResponseOutput) NonSslPort() pulumi.IntOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) int { return v.NonSslPort }).(pulumi.IntOutput)
}

// If clustering is enabled, the Shard ID of Redis Instance
func (o RedisInstanceDetailsResponseOutput) ShardId() pulumi.IntOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) int { return v.ShardId }).(pulumi.IntOutput)
}

// Redis instance SSL port.
func (o RedisInstanceDetailsResponseOutput) SslPort() pulumi.IntOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) int { return v.SslPort }).(pulumi.IntOutput)
}

// If the Cache uses availability zones, specifies availability zone where this instance is located.
func (o RedisInstanceDetailsResponseOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) string { return v.Zone }).(pulumi.StringOutput)
}

type RedisInstanceDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (RedisInstanceDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisInstanceDetailsResponse)(nil)).Elem()
}

func (o RedisInstanceDetailsResponseArrayOutput) ToRedisInstanceDetailsResponseArrayOutput() RedisInstanceDetailsResponseArrayOutput {
	return o
}

func (o RedisInstanceDetailsResponseArrayOutput) ToRedisInstanceDetailsResponseArrayOutputWithContext(ctx context.Context) RedisInstanceDetailsResponseArrayOutput {
	return o
}

func (o RedisInstanceDetailsResponseArrayOutput) Index(i pulumi.IntInput) RedisInstanceDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RedisInstanceDetailsResponse {
		return vs[0].([]RedisInstanceDetailsResponse)[vs[1].(int)]
	}).(RedisInstanceDetailsResponseOutput)
}

// Linked server Id
type RedisLinkedServerResponse struct {
	// Linked server Id.
	Id string `pulumi:"id"`
}

// RedisLinkedServerResponseInput is an input type that accepts RedisLinkedServerResponseArgs and RedisLinkedServerResponseOutput values.
// You can construct a concrete instance of `RedisLinkedServerResponseInput` via:
//
//          RedisLinkedServerResponseArgs{...}
type RedisLinkedServerResponseInput interface {
	pulumi.Input

	ToRedisLinkedServerResponseOutput() RedisLinkedServerResponseOutput
	ToRedisLinkedServerResponseOutputWithContext(context.Context) RedisLinkedServerResponseOutput
}

// Linked server Id
type RedisLinkedServerResponseArgs struct {
	// Linked server Id.
	Id pulumi.StringInput `pulumi:"id"`
}

func (RedisLinkedServerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisLinkedServerResponse)(nil)).Elem()
}

func (i RedisLinkedServerResponseArgs) ToRedisLinkedServerResponseOutput() RedisLinkedServerResponseOutput {
	return i.ToRedisLinkedServerResponseOutputWithContext(context.Background())
}

func (i RedisLinkedServerResponseArgs) ToRedisLinkedServerResponseOutputWithContext(ctx context.Context) RedisLinkedServerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisLinkedServerResponseOutput)
}

// RedisLinkedServerResponseArrayInput is an input type that accepts RedisLinkedServerResponseArray and RedisLinkedServerResponseArrayOutput values.
// You can construct a concrete instance of `RedisLinkedServerResponseArrayInput` via:
//
//          RedisLinkedServerResponseArray{ RedisLinkedServerResponseArgs{...} }
type RedisLinkedServerResponseArrayInput interface {
	pulumi.Input

	ToRedisLinkedServerResponseArrayOutput() RedisLinkedServerResponseArrayOutput
	ToRedisLinkedServerResponseArrayOutputWithContext(context.Context) RedisLinkedServerResponseArrayOutput
}

type RedisLinkedServerResponseArray []RedisLinkedServerResponseInput

func (RedisLinkedServerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisLinkedServerResponse)(nil)).Elem()
}

func (i RedisLinkedServerResponseArray) ToRedisLinkedServerResponseArrayOutput() RedisLinkedServerResponseArrayOutput {
	return i.ToRedisLinkedServerResponseArrayOutputWithContext(context.Background())
}

func (i RedisLinkedServerResponseArray) ToRedisLinkedServerResponseArrayOutputWithContext(ctx context.Context) RedisLinkedServerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisLinkedServerResponseArrayOutput)
}

// Linked server Id
type RedisLinkedServerResponseOutput struct{ *pulumi.OutputState }

func (RedisLinkedServerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisLinkedServerResponse)(nil)).Elem()
}

func (o RedisLinkedServerResponseOutput) ToRedisLinkedServerResponseOutput() RedisLinkedServerResponseOutput {
	return o
}

func (o RedisLinkedServerResponseOutput) ToRedisLinkedServerResponseOutputWithContext(ctx context.Context) RedisLinkedServerResponseOutput {
	return o
}

// Linked server Id.
func (o RedisLinkedServerResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RedisLinkedServerResponse) string { return v.Id }).(pulumi.StringOutput)
}

type RedisLinkedServerResponseArrayOutput struct{ *pulumi.OutputState }

func (RedisLinkedServerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisLinkedServerResponse)(nil)).Elem()
}

func (o RedisLinkedServerResponseArrayOutput) ToRedisLinkedServerResponseArrayOutput() RedisLinkedServerResponseArrayOutput {
	return o
}

func (o RedisLinkedServerResponseArrayOutput) ToRedisLinkedServerResponseArrayOutputWithContext(ctx context.Context) RedisLinkedServerResponseArrayOutput {
	return o
}

func (o RedisLinkedServerResponseArrayOutput) Index(i pulumi.IntInput) RedisLinkedServerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RedisLinkedServerResponse {
		return vs[0].([]RedisLinkedServerResponse)[vs[1].(int)]
	}).(RedisLinkedServerResponseOutput)
}

// Patch schedule entry for a Premium Redis Cache.
type ScheduleEntry struct {
	// Day of the week when a cache can be patched.
	DayOfWeek string `pulumi:"dayOfWeek"`
	// ISO8601 timespan specifying how much time cache patching can take.
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	// Start hour after which cache patching can start.
	StartHourUtc int `pulumi:"startHourUtc"`
}

// ScheduleEntryInput is an input type that accepts ScheduleEntryArgs and ScheduleEntryOutput values.
// You can construct a concrete instance of `ScheduleEntryInput` via:
//
//          ScheduleEntryArgs{...}
type ScheduleEntryInput interface {
	pulumi.Input

	ToScheduleEntryOutput() ScheduleEntryOutput
	ToScheduleEntryOutputWithContext(context.Context) ScheduleEntryOutput
}

// Patch schedule entry for a Premium Redis Cache.
type ScheduleEntryArgs struct {
	// Day of the week when a cache can be patched.
	DayOfWeek DayOfWeek `pulumi:"dayOfWeek"`
	// ISO8601 timespan specifying how much time cache patching can take.
	MaintenanceWindow pulumi.StringPtrInput `pulumi:"maintenanceWindow"`
	// Start hour after which cache patching can start.
	StartHourUtc pulumi.IntInput `pulumi:"startHourUtc"`
}

func (ScheduleEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntry)(nil)).Elem()
}

func (i ScheduleEntryArgs) ToScheduleEntryOutput() ScheduleEntryOutput {
	return i.ToScheduleEntryOutputWithContext(context.Background())
}

func (i ScheduleEntryArgs) ToScheduleEntryOutputWithContext(ctx context.Context) ScheduleEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryOutput)
}

// ScheduleEntryArrayInput is an input type that accepts ScheduleEntryArray and ScheduleEntryArrayOutput values.
// You can construct a concrete instance of `ScheduleEntryArrayInput` via:
//
//          ScheduleEntryArray{ ScheduleEntryArgs{...} }
type ScheduleEntryArrayInput interface {
	pulumi.Input

	ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput
	ToScheduleEntryArrayOutputWithContext(context.Context) ScheduleEntryArrayOutput
}

type ScheduleEntryArray []ScheduleEntryInput

func (ScheduleEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntry)(nil)).Elem()
}

func (i ScheduleEntryArray) ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput {
	return i.ToScheduleEntryArrayOutputWithContext(context.Background())
}

func (i ScheduleEntryArray) ToScheduleEntryArrayOutputWithContext(ctx context.Context) ScheduleEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryArrayOutput)
}

// Patch schedule entry for a Premium Redis Cache.
type ScheduleEntryOutput struct{ *pulumi.OutputState }

func (ScheduleEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntry)(nil)).Elem()
}

func (o ScheduleEntryOutput) ToScheduleEntryOutput() ScheduleEntryOutput {
	return o
}

func (o ScheduleEntryOutput) ToScheduleEntryOutputWithContext(ctx context.Context) ScheduleEntryOutput {
	return o
}

// Day of the week when a cache can be patched.
func (o ScheduleEntryOutput) DayOfWeek() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleEntry) string { return v.DayOfWeek }).(pulumi.StringOutput)
}

// ISO8601 timespan specifying how much time cache patching can take.
func (o ScheduleEntryOutput) MaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleEntry) *string { return v.MaintenanceWindow }).(pulumi.StringPtrOutput)
}

// Start hour after which cache patching can start.
func (o ScheduleEntryOutput) StartHourUtc() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleEntry) int { return v.StartHourUtc }).(pulumi.IntOutput)
}

type ScheduleEntryArrayOutput struct{ *pulumi.OutputState }

func (ScheduleEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntry)(nil)).Elem()
}

func (o ScheduleEntryArrayOutput) ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput {
	return o
}

func (o ScheduleEntryArrayOutput) ToScheduleEntryArrayOutputWithContext(ctx context.Context) ScheduleEntryArrayOutput {
	return o
}

func (o ScheduleEntryArrayOutput) Index(i pulumi.IntInput) ScheduleEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleEntry {
		return vs[0].([]ScheduleEntry)[vs[1].(int)]
	}).(ScheduleEntryOutput)
}

// Patch schedule entry for a Premium Redis Cache.
type ScheduleEntryResponse struct {
	// Day of the week when a cache can be patched.
	DayOfWeek string `pulumi:"dayOfWeek"`
	// ISO8601 timespan specifying how much time cache patching can take.
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	// Start hour after which cache patching can start.
	StartHourUtc int `pulumi:"startHourUtc"`
}

// ScheduleEntryResponseInput is an input type that accepts ScheduleEntryResponseArgs and ScheduleEntryResponseOutput values.
// You can construct a concrete instance of `ScheduleEntryResponseInput` via:
//
//          ScheduleEntryResponseArgs{...}
type ScheduleEntryResponseInput interface {
	pulumi.Input

	ToScheduleEntryResponseOutput() ScheduleEntryResponseOutput
	ToScheduleEntryResponseOutputWithContext(context.Context) ScheduleEntryResponseOutput
}

// Patch schedule entry for a Premium Redis Cache.
type ScheduleEntryResponseArgs struct {
	// Day of the week when a cache can be patched.
	DayOfWeek pulumi.StringInput `pulumi:"dayOfWeek"`
	// ISO8601 timespan specifying how much time cache patching can take.
	MaintenanceWindow pulumi.StringPtrInput `pulumi:"maintenanceWindow"`
	// Start hour after which cache patching can start.
	StartHourUtc pulumi.IntInput `pulumi:"startHourUtc"`
}

func (ScheduleEntryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntryResponse)(nil)).Elem()
}

func (i ScheduleEntryResponseArgs) ToScheduleEntryResponseOutput() ScheduleEntryResponseOutput {
	return i.ToScheduleEntryResponseOutputWithContext(context.Background())
}

func (i ScheduleEntryResponseArgs) ToScheduleEntryResponseOutputWithContext(ctx context.Context) ScheduleEntryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryResponseOutput)
}

// ScheduleEntryResponseArrayInput is an input type that accepts ScheduleEntryResponseArray and ScheduleEntryResponseArrayOutput values.
// You can construct a concrete instance of `ScheduleEntryResponseArrayInput` via:
//
//          ScheduleEntryResponseArray{ ScheduleEntryResponseArgs{...} }
type ScheduleEntryResponseArrayInput interface {
	pulumi.Input

	ToScheduleEntryResponseArrayOutput() ScheduleEntryResponseArrayOutput
	ToScheduleEntryResponseArrayOutputWithContext(context.Context) ScheduleEntryResponseArrayOutput
}

type ScheduleEntryResponseArray []ScheduleEntryResponseInput

func (ScheduleEntryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntryResponse)(nil)).Elem()
}

func (i ScheduleEntryResponseArray) ToScheduleEntryResponseArrayOutput() ScheduleEntryResponseArrayOutput {
	return i.ToScheduleEntryResponseArrayOutputWithContext(context.Background())
}

func (i ScheduleEntryResponseArray) ToScheduleEntryResponseArrayOutputWithContext(ctx context.Context) ScheduleEntryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryResponseArrayOutput)
}

// Patch schedule entry for a Premium Redis Cache.
type ScheduleEntryResponseOutput struct{ *pulumi.OutputState }

func (ScheduleEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntryResponse)(nil)).Elem()
}

func (o ScheduleEntryResponseOutput) ToScheduleEntryResponseOutput() ScheduleEntryResponseOutput {
	return o
}

func (o ScheduleEntryResponseOutput) ToScheduleEntryResponseOutputWithContext(ctx context.Context) ScheduleEntryResponseOutput {
	return o
}

// Day of the week when a cache can be patched.
func (o ScheduleEntryResponseOutput) DayOfWeek() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) string { return v.DayOfWeek }).(pulumi.StringOutput)
}

// ISO8601 timespan specifying how much time cache patching can take.
func (o ScheduleEntryResponseOutput) MaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) *string { return v.MaintenanceWindow }).(pulumi.StringPtrOutput)
}

// Start hour after which cache patching can start.
func (o ScheduleEntryResponseOutput) StartHourUtc() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) int { return v.StartHourUtc }).(pulumi.IntOutput)
}

type ScheduleEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (ScheduleEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntryResponse)(nil)).Elem()
}

func (o ScheduleEntryResponseArrayOutput) ToScheduleEntryResponseArrayOutput() ScheduleEntryResponseArrayOutput {
	return o
}

func (o ScheduleEntryResponseArrayOutput) ToScheduleEntryResponseArrayOutputWithContext(ctx context.Context) ScheduleEntryResponseArrayOutput {
	return o
}

func (o ScheduleEntryResponseArrayOutput) Index(i pulumi.IntInput) ScheduleEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleEntryResponse {
		return vs[0].([]ScheduleEntryResponse)[vs[1].(int)]
	}).(ScheduleEntryResponseOutput)
}

// SKU parameters supplied to the create Redis operation.
type Sku struct {
	// The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4, 5).
	Capacity int `pulumi:"capacity"`
	// The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family string `pulumi:"family"`
	// The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
	Name string `pulumi:"name"`
}

// SkuInput is an input type that accepts SkuArgs and SkuOutput values.
// You can construct a concrete instance of `SkuInput` via:
//
//          SkuArgs{...}
type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

// SKU parameters supplied to the create Redis operation.
type SkuArgs struct {
	// The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4, 5).
	Capacity pulumi.IntInput `pulumi:"capacity"`
	// The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family pulumi.StringInput `pulumi:"family"`
	// The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}

// SkuPtrInput is an input type that accepts SkuArgs, SkuPtr and SkuPtrOutput values.
// You can construct a concrete instance of `SkuPtrInput` via:
//
//          SkuArgs{...}
//
//  or:
//
//          nil
type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

// SKU parameters supplied to the create Redis operation.
type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyT(func(v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

// The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4, 5).
func (o SkuOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v Sku) int { return v.Capacity }).(pulumi.IntOutput)
}

// The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
func (o SkuOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Family }).(pulumi.StringOutput)
}

// The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku { return *v }).(SkuOutput)
}

// The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4, 5).
func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

// The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

// The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// SKU parameters supplied to the create Redis operation.
type SkuResponse struct {
	// The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4, 5).
	Capacity int `pulumi:"capacity"`
	// The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family string `pulumi:"family"`
	// The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
	Name string `pulumi:"name"`
}

// SkuResponseInput is an input type that accepts SkuResponseArgs and SkuResponseOutput values.
// You can construct a concrete instance of `SkuResponseInput` via:
//
//          SkuResponseArgs{...}
type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

// SKU parameters supplied to the create Redis operation.
type SkuResponseArgs struct {
	// The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4, 5).
	Capacity pulumi.IntInput `pulumi:"capacity"`
	// The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family pulumi.StringInput `pulumi:"family"`
	// The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}

// SkuResponsePtrInput is an input type that accepts SkuResponseArgs, SkuResponsePtr and SkuResponsePtrOutput values.
// You can construct a concrete instance of `SkuResponsePtrInput` via:
//
//          SkuResponseArgs{...}
//
//  or:
//
//          nil
type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

// SKU parameters supplied to the create Redis operation.
type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyT(func(v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

// The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4, 5).
func (o SkuResponseOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v SkuResponse) int { return v.Capacity }).(pulumi.IntOutput)
}

// The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
func (o SkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Family }).(pulumi.StringOutput)
}

// The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse { return *v }).(SkuResponseOutput)
}

// The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4, 5).
func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

// The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

// The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RedisAccessKeysResponseOutput{})
	pulumi.RegisterOutputType(RedisAccessKeysResponsePtrOutput{})
	pulumi.RegisterOutputType(RedisInstanceDetailsResponseOutput{})
	pulumi.RegisterOutputType(RedisInstanceDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(RedisLinkedServerResponseOutput{})
	pulumi.RegisterOutputType(RedisLinkedServerResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduleEntryOutput{})
	pulumi.RegisterOutputType(ScheduleEntryArrayOutput{})
	pulumi.RegisterOutputType(ScheduleEntryResponseOutput{})
	pulumi.RegisterOutputType(ScheduleEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
