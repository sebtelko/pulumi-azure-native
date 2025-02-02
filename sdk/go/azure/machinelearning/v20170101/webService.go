// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Instance of an Azure ML web service resource.
type WebService struct {
	pulumi.CustomResourceState

	// Specifies the location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Contains the property payload that describes the web service.
	Properties WebServicePropertiesForGraphResponseOutput `pulumi:"properties"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebService registers a new resource with the given unique name, arguments, and options.
func NewWebService(ctx *pulumi.Context,
	name string, args *WebServiceArgs, opts ...pulumi.ResourceOption) (*WebService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearning/v20170101:WebService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearning:WebService"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearning:WebService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearning/v20160501preview:WebService"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearning/v20160501preview:WebService"),
		},
	})
	opts = append(opts, aliases)
	var resource WebService
	err := ctx.RegisterResource("azure-native:machinelearning/v20170101:WebService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebService gets an existing WebService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebServiceState, opts ...pulumi.ResourceOption) (*WebService, error) {
	var resource WebService
	err := ctx.ReadResource("azure-native:machinelearning/v20170101:WebService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebService resources.
type webServiceState struct {
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name *string `pulumi:"name"`
	// Contains the property payload that describes the web service.
	Properties *WebServicePropertiesForGraphResponse `pulumi:"properties"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type *string `pulumi:"type"`
}

type WebServiceState struct {
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Specifies the name of the resource.
	Name pulumi.StringPtrInput
	// Contains the property payload that describes the web service.
	Properties WebServicePropertiesForGraphResponsePtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Specifies the type of the resource.
	Type pulumi.StringPtrInput
}

func (WebServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*webServiceState)(nil)).Elem()
}

type webServiceArgs struct {
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Contains the property payload that describes the web service.
	Properties WebServicePropertiesForGraph `pulumi:"properties"`
	// Name of the resource group in which the web service is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// The name of the web service.
	WebServiceName *string `pulumi:"webServiceName"`
}

// The set of arguments for constructing a WebService resource.
type WebServiceArgs struct {
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Contains the property payload that describes the web service.
	Properties WebServicePropertiesForGraphInput
	// Name of the resource group in which the web service is located.
	ResourceGroupName pulumi.StringInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// The name of the web service.
	WebServiceName pulumi.StringPtrInput
}

func (WebServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webServiceArgs)(nil)).Elem()
}

type WebServiceInput interface {
	pulumi.Input

	ToWebServiceOutput() WebServiceOutput
	ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput
}

func (*WebService) ElementType() reflect.Type {
	return reflect.TypeOf((*WebService)(nil))
}

func (i *WebService) ToWebServiceOutput() WebServiceOutput {
	return i.ToWebServiceOutputWithContext(context.Background())
}

func (i *WebService) ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceOutput)
}

type WebServiceOutput struct {
	*pulumi.OutputState
}

func (WebServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebService)(nil))
}

func (o WebServiceOutput) ToWebServiceOutput() WebServiceOutput {
	return o
}

func (o WebServiceOutput) ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebServiceOutput{})
}
