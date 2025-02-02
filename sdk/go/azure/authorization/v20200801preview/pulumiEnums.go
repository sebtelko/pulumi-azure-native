// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The principal type of the assigned principal ID.
type PrincipalType pulumi.String

const (
	PrincipalTypeUser                   = PrincipalType("User")
	PrincipalTypeGroup                  = PrincipalType("Group")
	PrincipalTypeServicePrincipal       = PrincipalType("ServicePrincipal")
	PrincipalTypeUnknown                = PrincipalType("Unknown")
	PrincipalTypeDirectoryRoleTemplate  = PrincipalType("DirectoryRoleTemplate")
	PrincipalTypeForeignGroup           = PrincipalType("ForeignGroup")
	PrincipalTypeApplication            = PrincipalType("Application")
	PrincipalTypeMSI                    = PrincipalType("MSI")
	PrincipalTypeDirectoryObjectOrGroup = PrincipalType("DirectoryObjectOrGroup")
	PrincipalTypeEveryone               = PrincipalType("Everyone")
)

func (PrincipalType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e PrincipalType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrincipalType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
