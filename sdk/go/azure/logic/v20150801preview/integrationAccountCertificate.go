// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountCertificate struct {
	pulumi.CustomResourceState

	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The key details in the key vault.
	Key KeyVaultKeyReferenceResponsePtrOutput `pulumi:"key"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The metadata.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The resource name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The public certificate.
	PublicCertificate pulumi.StringPtrOutput `pulumi:"publicCertificate"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewIntegrationAccountCertificate registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountCertificate(ctx *pulumi.Context,
	name string, args *IntegrationAccountCertificateArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:logic/v20150801preview:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20160601:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:IntegrationAccountCertificate"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountCertificate
	err := ctx.RegisterResource("azure-native:logic/v20150801preview:IntegrationAccountCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountCertificate gets an existing IntegrationAccountCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountCertificateState, opts ...pulumi.ResourceOption) (*IntegrationAccountCertificate, error) {
	var resource IntegrationAccountCertificate
	err := ctx.ReadResource("azure-native:logic/v20150801preview:IntegrationAccountCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountCertificate resources.
type integrationAccountCertificateState struct {
	// The changed time.
	ChangedTime *string `pulumi:"changedTime"`
	// The created time.
	CreatedTime *string `pulumi:"createdTime"`
	// The key details in the key vault.
	Key *KeyVaultKeyReferenceResponse `pulumi:"key"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The public certificate.
	PublicCertificate *string `pulumi:"publicCertificate"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type IntegrationAccountCertificateState struct {
	// The changed time.
	ChangedTime pulumi.StringPtrInput
	// The created time.
	CreatedTime pulumi.StringPtrInput
	// The key details in the key vault.
	Key KeyVaultKeyReferenceResponsePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.Input
	// The resource name.
	Name pulumi.StringPtrInput
	// The public certificate.
	PublicCertificate pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (IntegrationAccountCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountCertificateState)(nil)).Elem()
}

type integrationAccountCertificateArgs struct {
	// The integration account certificate name.
	CertificateName *string `pulumi:"certificateName"`
	// The resource id.
	Id *string `pulumi:"id"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The key details in the key vault.
	Key *KeyVaultKeyReference `pulumi:"key"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The public certificate.
	PublicCertificate *string `pulumi:"publicCertificate"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a IntegrationAccountCertificate resource.
type IntegrationAccountCertificateArgs struct {
	// The integration account certificate name.
	CertificateName pulumi.StringPtrInput
	// The resource id.
	Id pulumi.StringPtrInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The key details in the key vault.
	Key KeyVaultKeyReferencePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.Input
	// The resource name.
	Name pulumi.StringPtrInput
	// The public certificate.
	PublicCertificate pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (IntegrationAccountCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountCertificateArgs)(nil)).Elem()
}

type IntegrationAccountCertificateInput interface {
	pulumi.Input

	ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput
	ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput
}

func (*IntegrationAccountCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountCertificate)(nil))
}

func (i *IntegrationAccountCertificate) ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput {
	return i.ToIntegrationAccountCertificateOutputWithContext(context.Background())
}

func (i *IntegrationAccountCertificate) ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountCertificateOutput)
}

type IntegrationAccountCertificateOutput struct {
	*pulumi.OutputState
}

func (IntegrationAccountCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountCertificate)(nil))
}

func (o IntegrationAccountCertificateOutput) ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput {
	return o
}

func (o IntegrationAccountCertificateOutput) ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountCertificateOutput{})
}
