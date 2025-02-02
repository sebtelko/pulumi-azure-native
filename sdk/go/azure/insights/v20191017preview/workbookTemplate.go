// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191017preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Application Insights workbook template definition.
type WorkbookTemplate struct {
	pulumi.CustomResourceState

	// Information about the author of the workbook template.
	Author pulumi.StringPtrOutput `pulumi:"author"`
	// Workbook galleries supported by the template.
	Galleries WorkbookTemplateGalleryResponseArrayOutput `pulumi:"galleries"`
	// Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
	Localized WorkbookTemplateLocalizedGalleryResponseArrayMapOutput `pulumi:"localized"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Valid JSON object containing workbook template payload.
	TemplateData pulumi.AnyOutput `pulumi:"templateData"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkbookTemplate registers a new resource with the given unique name, arguments, and options.
func NewWorkbookTemplate(ctx *pulumi.Context,
	name string, args *WorkbookTemplateArgs, opts ...pulumi.ResourceOption) (*WorkbookTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Galleries == nil {
		return nil, errors.New("invalid value for required argument 'Galleries'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TemplateData == nil {
		return nil, errors.New("invalid value for required argument 'TemplateData'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights/v20191017preview:WorkbookTemplate"),
		},
		{
			Type: pulumi.String("azure-native:insights:WorkbookTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights:WorkbookTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkbookTemplate
	err := ctx.RegisterResource("azure-native:insights/v20191017preview:WorkbookTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkbookTemplate gets an existing WorkbookTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkbookTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkbookTemplateState, opts ...pulumi.ResourceOption) (*WorkbookTemplate, error) {
	var resource WorkbookTemplate
	err := ctx.ReadResource("azure-native:insights/v20191017preview:WorkbookTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkbookTemplate resources.
type workbookTemplateState struct {
	// Information about the author of the workbook template.
	Author *string `pulumi:"author"`
	// Workbook galleries supported by the template.
	Galleries []WorkbookTemplateGalleryResponse `pulumi:"galleries"`
	// Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
	Localized map[string][]WorkbookTemplateLocalizedGalleryResponse `pulumi:"localized"`
	// Resource location
	Location *string `pulumi:"location"`
	// Azure resource name.
	Name *string `pulumi:"name"`
	// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
	Priority *int `pulumi:"priority"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Valid JSON object containing workbook template payload.
	TemplateData interface{} `pulumi:"templateData"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type WorkbookTemplateState struct {
	// Information about the author of the workbook template.
	Author pulumi.StringPtrInput
	// Workbook galleries supported by the template.
	Galleries WorkbookTemplateGalleryResponseArrayInput
	// Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
	Localized WorkbookTemplateLocalizedGalleryResponseArrayMapInput
	// Resource location
	Location pulumi.StringPtrInput
	// Azure resource name.
	Name pulumi.StringPtrInput
	// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
	Priority pulumi.IntPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Valid JSON object containing workbook template payload.
	TemplateData pulumi.Input
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (WorkbookTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookTemplateState)(nil)).Elem()
}

type workbookTemplateArgs struct {
	// Information about the author of the workbook template.
	Author *string `pulumi:"author"`
	// Workbook galleries supported by the template.
	Galleries []WorkbookTemplateGallery `pulumi:"galleries"`
	// Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
	Localized map[string][]WorkbookTemplateLocalizedGallery `pulumi:"localized"`
	// Resource location
	Location *string `pulumi:"location"`
	// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
	Priority *int `pulumi:"priority"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName *string `pulumi:"resourceName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Valid JSON object containing workbook template payload.
	TemplateData interface{} `pulumi:"templateData"`
}

// The set of arguments for constructing a WorkbookTemplate resource.
type WorkbookTemplateArgs struct {
	// Information about the author of the workbook template.
	Author pulumi.StringPtrInput
	// Workbook galleries supported by the template.
	Galleries WorkbookTemplateGalleryArrayInput
	// Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
	Localized WorkbookTemplateLocalizedGalleryArrayMapInput
	// Resource location
	Location pulumi.StringPtrInput
	// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
	Priority pulumi.IntPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Valid JSON object containing workbook template payload.
	TemplateData pulumi.Input
}

func (WorkbookTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookTemplateArgs)(nil)).Elem()
}

type WorkbookTemplateInput interface {
	pulumi.Input

	ToWorkbookTemplateOutput() WorkbookTemplateOutput
	ToWorkbookTemplateOutputWithContext(ctx context.Context) WorkbookTemplateOutput
}

func (*WorkbookTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplate)(nil))
}

func (i *WorkbookTemplate) ToWorkbookTemplateOutput() WorkbookTemplateOutput {
	return i.ToWorkbookTemplateOutputWithContext(context.Background())
}

func (i *WorkbookTemplate) ToWorkbookTemplateOutputWithContext(ctx context.Context) WorkbookTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateOutput)
}

type WorkbookTemplateOutput struct {
	*pulumi.OutputState
}

func (WorkbookTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplate)(nil))
}

func (o WorkbookTemplateOutput) ToWorkbookTemplateOutput() WorkbookTemplateOutput {
	return o
}

func (o WorkbookTemplateOutput) ToWorkbookTemplateOutputWithContext(ctx context.Context) WorkbookTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkbookTemplateOutput{})
}
