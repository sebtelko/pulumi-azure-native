// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Identity for the connected cluster.
type ConnectedClusterIdentity struct {
	// The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
	Type string `pulumi:"type"`
}

// ConnectedClusterIdentityInput is an input type that accepts ConnectedClusterIdentityArgs and ConnectedClusterIdentityOutput values.
// You can construct a concrete instance of `ConnectedClusterIdentityInput` via:
//
//          ConnectedClusterIdentityArgs{...}
type ConnectedClusterIdentityInput interface {
	pulumi.Input

	ToConnectedClusterIdentityOutput() ConnectedClusterIdentityOutput
	ToConnectedClusterIdentityOutputWithContext(context.Context) ConnectedClusterIdentityOutput
}

// Identity for the connected cluster.
type ConnectedClusterIdentityArgs struct {
	// The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
	Type ResourceIdentityType `pulumi:"type"`
}

func (ConnectedClusterIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentity)(nil)).Elem()
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityOutput() ConnectedClusterIdentityOutput {
	return i.ToConnectedClusterIdentityOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityOutputWithContext(ctx context.Context) ConnectedClusterIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityOutput)
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return i.ToConnectedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityOutput).ToConnectedClusterIdentityPtrOutputWithContext(ctx)
}

// ConnectedClusterIdentityPtrInput is an input type that accepts ConnectedClusterIdentityArgs, ConnectedClusterIdentityPtr and ConnectedClusterIdentityPtrOutput values.
// You can construct a concrete instance of `ConnectedClusterIdentityPtrInput` via:
//
//          ConnectedClusterIdentityArgs{...}
//
//  or:
//
//          nil
type ConnectedClusterIdentityPtrInput interface {
	pulumi.Input

	ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput
	ToConnectedClusterIdentityPtrOutputWithContext(context.Context) ConnectedClusterIdentityPtrOutput
}

type connectedClusterIdentityPtrType ConnectedClusterIdentityArgs

func ConnectedClusterIdentityPtr(v *ConnectedClusterIdentityArgs) ConnectedClusterIdentityPtrInput {
	return (*connectedClusterIdentityPtrType)(v)
}

func (*connectedClusterIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentity)(nil)).Elem()
}

func (i *connectedClusterIdentityPtrType) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return i.ToConnectedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i *connectedClusterIdentityPtrType) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityPtrOutput)
}

// Identity for the connected cluster.
type ConnectedClusterIdentityOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentity)(nil)).Elem()
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityOutput() ConnectedClusterIdentityOutput {
	return o
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityOutputWithContext(ctx context.Context) ConnectedClusterIdentityOutput {
	return o
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return o.ToConnectedClusterIdentityPtrOutputWithContext(context.Background())
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return o.ApplyT(func(v ConnectedClusterIdentity) *ConnectedClusterIdentity {
		return &v
	}).(ConnectedClusterIdentityPtrOutput)
}

// The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
func (o ConnectedClusterIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type ConnectedClusterIdentityPtrOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentity)(nil)).Elem()
}

func (o ConnectedClusterIdentityPtrOutput) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return o
}

func (o ConnectedClusterIdentityPtrOutput) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return o
}

func (o ConnectedClusterIdentityPtrOutput) Elem() ConnectedClusterIdentityOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentity) ConnectedClusterIdentity { return *v }).(ConnectedClusterIdentityOutput)
}

// The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
func (o ConnectedClusterIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

// Identity for the connected cluster.
type ConnectedClusterIdentityResponse struct {
	// The principal id of connected cluster identity. This property will only be provided for a system assigned identity.
	PrincipalId string `pulumi:"principalId"`
	// The tenant id associated with the connected cluster. This property will only be provided for a system assigned identity.
	TenantId string `pulumi:"tenantId"`
	// The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
	Type string `pulumi:"type"`
}

// ConnectedClusterIdentityResponseInput is an input type that accepts ConnectedClusterIdentityResponseArgs and ConnectedClusterIdentityResponseOutput values.
// You can construct a concrete instance of `ConnectedClusterIdentityResponseInput` via:
//
//          ConnectedClusterIdentityResponseArgs{...}
type ConnectedClusterIdentityResponseInput interface {
	pulumi.Input

	ToConnectedClusterIdentityResponseOutput() ConnectedClusterIdentityResponseOutput
	ToConnectedClusterIdentityResponseOutputWithContext(context.Context) ConnectedClusterIdentityResponseOutput
}

// Identity for the connected cluster.
type ConnectedClusterIdentityResponseArgs struct {
	// The principal id of connected cluster identity. This property will only be provided for a system assigned identity.
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	// The tenant id associated with the connected cluster. This property will only be provided for a system assigned identity.
	TenantId pulumi.StringInput `pulumi:"tenantId"`
	// The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ConnectedClusterIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponseOutput() ConnectedClusterIdentityResponseOutput {
	return i.ToConnectedClusterIdentityResponseOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponseOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityResponseOutput)
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return i.ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityResponseOutput).ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx)
}

// ConnectedClusterIdentityResponsePtrInput is an input type that accepts ConnectedClusterIdentityResponseArgs, ConnectedClusterIdentityResponsePtr and ConnectedClusterIdentityResponsePtrOutput values.
// You can construct a concrete instance of `ConnectedClusterIdentityResponsePtrInput` via:
//
//          ConnectedClusterIdentityResponseArgs{...}
//
//  or:
//
//          nil
type ConnectedClusterIdentityResponsePtrInput interface {
	pulumi.Input

	ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput
	ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Context) ConnectedClusterIdentityResponsePtrOutput
}

type connectedClusterIdentityResponsePtrType ConnectedClusterIdentityResponseArgs

func ConnectedClusterIdentityResponsePtr(v *ConnectedClusterIdentityResponseArgs) ConnectedClusterIdentityResponsePtrInput {
	return (*connectedClusterIdentityResponsePtrType)(v)
}

func (*connectedClusterIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (i *connectedClusterIdentityResponsePtrType) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return i.ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *connectedClusterIdentityResponsePtrType) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityResponsePtrOutput)
}

// Identity for the connected cluster.
type ConnectedClusterIdentityResponseOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponseOutput() ConnectedClusterIdentityResponseOutput {
	return o
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponseOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponseOutput {
	return o
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return o.ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) *ConnectedClusterIdentityResponse {
		return &v
	}).(ConnectedClusterIdentityResponsePtrOutput)
}

// The principal id of connected cluster identity. This property will only be provided for a system assigned identity.
func (o ConnectedClusterIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The tenant id associated with the connected cluster. This property will only be provided for a system assigned identity.
func (o ConnectedClusterIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
func (o ConnectedClusterIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ConnectedClusterIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (o ConnectedClusterIdentityResponsePtrOutput) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return o
}

func (o ConnectedClusterIdentityResponsePtrOutput) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return o
}

func (o ConnectedClusterIdentityResponsePtrOutput) Elem() ConnectedClusterIdentityResponseOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) ConnectedClusterIdentityResponse { return *v }).(ConnectedClusterIdentityResponseOutput)
}

// The principal id of connected cluster identity. This property will only be provided for a system assigned identity.
func (o ConnectedClusterIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The tenant id associated with the connected cluster. This property will only be provided for a system assigned identity.
func (o ConnectedClusterIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

// The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
func (o ConnectedClusterIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

// The credential result response.
type CredentialResultResponse struct {
	// The name of the credential.
	Name string `pulumi:"name"`
	// Base64-encoded Kubernetes configuration file.
	Value string `pulumi:"value"`
}

// CredentialResultResponseInput is an input type that accepts CredentialResultResponseArgs and CredentialResultResponseOutput values.
// You can construct a concrete instance of `CredentialResultResponseInput` via:
//
//          CredentialResultResponseArgs{...}
type CredentialResultResponseInput interface {
	pulumi.Input

	ToCredentialResultResponseOutput() CredentialResultResponseOutput
	ToCredentialResultResponseOutputWithContext(context.Context) CredentialResultResponseOutput
}

// The credential result response.
type CredentialResultResponseArgs struct {
	// The name of the credential.
	Name pulumi.StringInput `pulumi:"name"`
	// Base64-encoded Kubernetes configuration file.
	Value pulumi.StringInput `pulumi:"value"`
}

func (CredentialResultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialResultResponse)(nil)).Elem()
}

func (i CredentialResultResponseArgs) ToCredentialResultResponseOutput() CredentialResultResponseOutput {
	return i.ToCredentialResultResponseOutputWithContext(context.Background())
}

func (i CredentialResultResponseArgs) ToCredentialResultResponseOutputWithContext(ctx context.Context) CredentialResultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialResultResponseOutput)
}

// CredentialResultResponseArrayInput is an input type that accepts CredentialResultResponseArray and CredentialResultResponseArrayOutput values.
// You can construct a concrete instance of `CredentialResultResponseArrayInput` via:
//
//          CredentialResultResponseArray{ CredentialResultResponseArgs{...} }
type CredentialResultResponseArrayInput interface {
	pulumi.Input

	ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput
	ToCredentialResultResponseArrayOutputWithContext(context.Context) CredentialResultResponseArrayOutput
}

type CredentialResultResponseArray []CredentialResultResponseInput

func (CredentialResultResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CredentialResultResponse)(nil)).Elem()
}

func (i CredentialResultResponseArray) ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput {
	return i.ToCredentialResultResponseArrayOutputWithContext(context.Background())
}

func (i CredentialResultResponseArray) ToCredentialResultResponseArrayOutputWithContext(ctx context.Context) CredentialResultResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialResultResponseArrayOutput)
}

// The credential result response.
type CredentialResultResponseOutput struct{ *pulumi.OutputState }

func (CredentialResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialResultResponse)(nil)).Elem()
}

func (o CredentialResultResponseOutput) ToCredentialResultResponseOutput() CredentialResultResponseOutput {
	return o
}

func (o CredentialResultResponseOutput) ToCredentialResultResponseOutputWithContext(ctx context.Context) CredentialResultResponseOutput {
	return o
}

// The name of the credential.
func (o CredentialResultResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CredentialResultResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Base64-encoded Kubernetes configuration file.
func (o CredentialResultResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v CredentialResultResponse) string { return v.Value }).(pulumi.StringOutput)
}

type CredentialResultResponseArrayOutput struct{ *pulumi.OutputState }

func (CredentialResultResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CredentialResultResponse)(nil)).Elem()
}

func (o CredentialResultResponseArrayOutput) ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput {
	return o
}

func (o CredentialResultResponseArrayOutput) ToCredentialResultResponseArrayOutputWithContext(ctx context.Context) CredentialResultResponseArrayOutput {
	return o
}

func (o CredentialResultResponseArrayOutput) Index(i pulumi.IntInput) CredentialResultResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CredentialResultResponse {
		return vs[0].([]CredentialResultResponse)[vs[1].(int)]
	}).(CredentialResultResponseOutput)
}

// Contains the REP (rendezvous endpoint) and “Sender” access token.
type HybridConnectionConfigResponse struct {
	// Timestamp when this token will be expired.
	ExpirationTime float64 `pulumi:"expirationTime"`
	// Name of the connection
	HybridConnectionName string `pulumi:"hybridConnectionName"`
	// Name of the relay.
	Relay string `pulumi:"relay"`
	// Sender access token
	Token string `pulumi:"token"`
}

// HybridConnectionConfigResponseInput is an input type that accepts HybridConnectionConfigResponseArgs and HybridConnectionConfigResponseOutput values.
// You can construct a concrete instance of `HybridConnectionConfigResponseInput` via:
//
//          HybridConnectionConfigResponseArgs{...}
type HybridConnectionConfigResponseInput interface {
	pulumi.Input

	ToHybridConnectionConfigResponseOutput() HybridConnectionConfigResponseOutput
	ToHybridConnectionConfigResponseOutputWithContext(context.Context) HybridConnectionConfigResponseOutput
}

// Contains the REP (rendezvous endpoint) and “Sender” access token.
type HybridConnectionConfigResponseArgs struct {
	// Timestamp when this token will be expired.
	ExpirationTime pulumi.Float64Input `pulumi:"expirationTime"`
	// Name of the connection
	HybridConnectionName pulumi.StringInput `pulumi:"hybridConnectionName"`
	// Name of the relay.
	Relay pulumi.StringInput `pulumi:"relay"`
	// Sender access token
	Token pulumi.StringInput `pulumi:"token"`
}

func (HybridConnectionConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionConfigResponse)(nil)).Elem()
}

func (i HybridConnectionConfigResponseArgs) ToHybridConnectionConfigResponseOutput() HybridConnectionConfigResponseOutput {
	return i.ToHybridConnectionConfigResponseOutputWithContext(context.Background())
}

func (i HybridConnectionConfigResponseArgs) ToHybridConnectionConfigResponseOutputWithContext(ctx context.Context) HybridConnectionConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridConnectionConfigResponseOutput)
}

// Contains the REP (rendezvous endpoint) and “Sender” access token.
type HybridConnectionConfigResponseOutput struct{ *pulumi.OutputState }

func (HybridConnectionConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionConfigResponse)(nil)).Elem()
}

func (o HybridConnectionConfigResponseOutput) ToHybridConnectionConfigResponseOutput() HybridConnectionConfigResponseOutput {
	return o
}

func (o HybridConnectionConfigResponseOutput) ToHybridConnectionConfigResponseOutputWithContext(ctx context.Context) HybridConnectionConfigResponseOutput {
	return o
}

// Timestamp when this token will be expired.
func (o HybridConnectionConfigResponseOutput) ExpirationTime() pulumi.Float64Output {
	return o.ApplyT(func(v HybridConnectionConfigResponse) float64 { return v.ExpirationTime }).(pulumi.Float64Output)
}

// Name of the connection
func (o HybridConnectionConfigResponseOutput) HybridConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionConfigResponse) string { return v.HybridConnectionName }).(pulumi.StringOutput)
}

// Name of the relay.
func (o HybridConnectionConfigResponseOutput) Relay() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionConfigResponse) string { return v.Relay }).(pulumi.StringOutput)
}

// Sender access token
func (o HybridConnectionConfigResponseOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionConfigResponse) string { return v.Token }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource modification (UTC).
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// SystemDataResponseInput is an input type that accepts SystemDataResponseArgs and SystemDataResponseOutput values.
// You can construct a concrete instance of `SystemDataResponseInput` via:
//
//          SystemDataResponseArgs{...}
type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseArgs struct {
	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrInput `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrInput `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrInput `pulumi:"createdByType"`
	// The timestamp of resource modification (UTC).
	LastModifiedAt pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}

// SystemDataResponsePtrInput is an input type that accepts SystemDataResponseArgs, SystemDataResponsePtr and SystemDataResponsePtrOutput values.
// You can construct a concrete instance of `SystemDataResponsePtrInput` via:
//
//          SystemDataResponseArgs{...}
//
//  or:
//
//          nil
type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource modification (UTC).
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse { return *v }).(SystemDataResponseOutput)
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

// The timestamp of resource modification (UTC).
func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedClusterIdentityOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityPtrOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityResponseOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseArrayOutput{})
	pulumi.RegisterOutputType(HybridConnectionConfigResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
