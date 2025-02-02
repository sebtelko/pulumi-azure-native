// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Cloud shell console
func LookupConsole(ctx *pulumi.Context, args *LookupConsoleArgs, opts ...pulumi.InvokeOption) (*LookupConsoleResult, error) {
	var rv LookupConsoleResult
	err := ctx.Invoke("azure-native:portal/v20181001:getConsole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConsoleArgs struct {
	// The name of the console
	ConsoleName string `pulumi:"consoleName"`
}

// Cloud shell console
type LookupConsoleResult struct {
	// Cloud shell console properties.
	Properties ConsolePropertiesResponse `pulumi:"properties"`
}
