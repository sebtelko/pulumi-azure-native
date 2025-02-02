// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a DNS record set (a collection of DNS records with the same name and type).
type RecordSet struct {
	pulumi.CustomResourceState

	// The list of A records in the record set.
	ARecords ARecordResponseArrayOutput `pulumi:"aRecords"`
	// The list of AAAA records in the record set.
	AaaaRecords AaaaRecordResponseArrayOutput `pulumi:"aaaaRecords"`
	// The list of CAA records in the record set.
	CaaRecords CaaRecordResponseArrayOutput `pulumi:"caaRecords"`
	// The CNAME record in the  record set.
	CnameRecord CnameRecordResponsePtrOutput `pulumi:"cnameRecord"`
	// The etag of the record set.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Fully qualified domain name of the record set.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The metadata attached to the record set.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The list of MX records in the record set.
	MxRecords MxRecordResponseArrayOutput `pulumi:"mxRecords"`
	// The name of the record set.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of NS records in the record set.
	NsRecords NsRecordResponseArrayOutput `pulumi:"nsRecords"`
	// provisioning State of the record set.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The list of PTR records in the record set.
	PtrRecords PtrRecordResponseArrayOutput `pulumi:"ptrRecords"`
	// The SOA record in the record set.
	SoaRecord SoaRecordResponsePtrOutput `pulumi:"soaRecord"`
	// The list of SRV records in the record set.
	SrvRecords SrvRecordResponseArrayOutput `pulumi:"srvRecords"`
	// A reference to an azure resource from where the dns resource value is taken.
	TargetResource SubResourceResponsePtrOutput `pulumi:"targetResource"`
	// The TTL (time-to-live) of the records in the record set.
	Ttl pulumi.Float64PtrOutput `pulumi:"ttl"`
	// The list of TXT records in the record set.
	TxtRecords TxtRecordResponseArrayOutput `pulumi:"txtRecords"`
	// The type of the record set.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRecordSet registers a new resource with the given unique name, arguments, and options.
func NewRecordSet(ctx *pulumi.Context,
	name string, args *RecordSetArgs, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RecordType == nil {
		return nil, errors.New("invalid value for required argument 'RecordType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ZoneName == nil {
		return nil, errors.New("invalid value for required argument 'ZoneName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20180501:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150504preview:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20150504preview:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160401:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160401:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171001:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301preview:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180301preview:RecordSet"),
		},
	})
	opts = append(opts, aliases)
	var resource RecordSet
	err := ctx.RegisterResource("azure-native:network/v20180501:RecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecordSet gets an existing RecordSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordSetState, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	var resource RecordSet
	err := ctx.ReadResource("azure-native:network/v20180501:RecordSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RecordSet resources.
type recordSetState struct {
	// The list of A records in the record set.
	ARecords []ARecordResponse `pulumi:"aRecords"`
	// The list of AAAA records in the record set.
	AaaaRecords []AaaaRecordResponse `pulumi:"aaaaRecords"`
	// The list of CAA records in the record set.
	CaaRecords []CaaRecordResponse `pulumi:"caaRecords"`
	// The CNAME record in the  record set.
	CnameRecord *CnameRecordResponse `pulumi:"cnameRecord"`
	// The etag of the record set.
	Etag *string `pulumi:"etag"`
	// Fully qualified domain name of the record set.
	Fqdn *string `pulumi:"fqdn"`
	// The metadata attached to the record set.
	Metadata map[string]string `pulumi:"metadata"`
	// The list of MX records in the record set.
	MxRecords []MxRecordResponse `pulumi:"mxRecords"`
	// The name of the record set.
	Name *string `pulumi:"name"`
	// The list of NS records in the record set.
	NsRecords []NsRecordResponse `pulumi:"nsRecords"`
	// provisioning State of the record set.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The list of PTR records in the record set.
	PtrRecords []PtrRecordResponse `pulumi:"ptrRecords"`
	// The SOA record in the record set.
	SoaRecord *SoaRecordResponse `pulumi:"soaRecord"`
	// The list of SRV records in the record set.
	SrvRecords []SrvRecordResponse `pulumi:"srvRecords"`
	// A reference to an azure resource from where the dns resource value is taken.
	TargetResource *SubResourceResponse `pulumi:"targetResource"`
	// The TTL (time-to-live) of the records in the record set.
	Ttl *float64 `pulumi:"ttl"`
	// The list of TXT records in the record set.
	TxtRecords []TxtRecordResponse `pulumi:"txtRecords"`
	// The type of the record set.
	Type *string `pulumi:"type"`
}

type RecordSetState struct {
	// The list of A records in the record set.
	ARecords ARecordResponseArrayInput
	// The list of AAAA records in the record set.
	AaaaRecords AaaaRecordResponseArrayInput
	// The list of CAA records in the record set.
	CaaRecords CaaRecordResponseArrayInput
	// The CNAME record in the  record set.
	CnameRecord CnameRecordResponsePtrInput
	// The etag of the record set.
	Etag pulumi.StringPtrInput
	// Fully qualified domain name of the record set.
	Fqdn pulumi.StringPtrInput
	// The metadata attached to the record set.
	Metadata pulumi.StringMapInput
	// The list of MX records in the record set.
	MxRecords MxRecordResponseArrayInput
	// The name of the record set.
	Name pulumi.StringPtrInput
	// The list of NS records in the record set.
	NsRecords NsRecordResponseArrayInput
	// provisioning State of the record set.
	ProvisioningState pulumi.StringPtrInput
	// The list of PTR records in the record set.
	PtrRecords PtrRecordResponseArrayInput
	// The SOA record in the record set.
	SoaRecord SoaRecordResponsePtrInput
	// The list of SRV records in the record set.
	SrvRecords SrvRecordResponseArrayInput
	// A reference to an azure resource from where the dns resource value is taken.
	TargetResource SubResourceResponsePtrInput
	// The TTL (time-to-live) of the records in the record set.
	Ttl pulumi.Float64PtrInput
	// The list of TXT records in the record set.
	TxtRecords TxtRecordResponseArrayInput
	// The type of the record set.
	Type pulumi.StringPtrInput
}

func (RecordSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetState)(nil)).Elem()
}

type recordSetArgs struct {
	// The list of A records in the record set.
	ARecords []ARecord `pulumi:"aRecords"`
	// The list of AAAA records in the record set.
	AaaaRecords []AaaaRecord `pulumi:"aaaaRecords"`
	// The list of CAA records in the record set.
	CaaRecords []CaaRecord `pulumi:"caaRecords"`
	// The CNAME record in the  record set.
	CnameRecord *CnameRecord `pulumi:"cnameRecord"`
	// The etag of the record set.
	Etag *string `pulumi:"etag"`
	// The metadata attached to the record set.
	Metadata map[string]string `pulumi:"metadata"`
	// The list of MX records in the record set.
	MxRecords []MxRecord `pulumi:"mxRecords"`
	// The list of NS records in the record set.
	NsRecords []NsRecord `pulumi:"nsRecords"`
	// The list of PTR records in the record set.
	PtrRecords []PtrRecord `pulumi:"ptrRecords"`
	// The type of DNS record in this record set. Record sets of type SOA can be updated but not created (they are created when the DNS zone is created).
	RecordType string `pulumi:"recordType"`
	// The name of the record set, relative to the name of the zone.
	RelativeRecordSetName *string `pulumi:"relativeRecordSetName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SOA record in the record set.
	SoaRecord *SoaRecord `pulumi:"soaRecord"`
	// The list of SRV records in the record set.
	SrvRecords []SrvRecord `pulumi:"srvRecords"`
	// A reference to an azure resource from where the dns resource value is taken.
	TargetResource *SubResource `pulumi:"targetResource"`
	// The TTL (time-to-live) of the records in the record set.
	Ttl *float64 `pulumi:"ttl"`
	// The list of TXT records in the record set.
	TxtRecords []TxtRecord `pulumi:"txtRecords"`
	// The name of the DNS zone (without a terminating dot).
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a RecordSet resource.
type RecordSetArgs struct {
	// The list of A records in the record set.
	ARecords ARecordArrayInput
	// The list of AAAA records in the record set.
	AaaaRecords AaaaRecordArrayInput
	// The list of CAA records in the record set.
	CaaRecords CaaRecordArrayInput
	// The CNAME record in the  record set.
	CnameRecord CnameRecordPtrInput
	// The etag of the record set.
	Etag pulumi.StringPtrInput
	// The metadata attached to the record set.
	Metadata pulumi.StringMapInput
	// The list of MX records in the record set.
	MxRecords MxRecordArrayInput
	// The list of NS records in the record set.
	NsRecords NsRecordArrayInput
	// The list of PTR records in the record set.
	PtrRecords PtrRecordArrayInput
	// The type of DNS record in this record set. Record sets of type SOA can be updated but not created (they are created when the DNS zone is created).
	RecordType pulumi.StringInput
	// The name of the record set, relative to the name of the zone.
	RelativeRecordSetName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The SOA record in the record set.
	SoaRecord SoaRecordPtrInput
	// The list of SRV records in the record set.
	SrvRecords SrvRecordArrayInput
	// A reference to an azure resource from where the dns resource value is taken.
	TargetResource SubResourcePtrInput
	// The TTL (time-to-live) of the records in the record set.
	Ttl pulumi.Float64PtrInput
	// The list of TXT records in the record set.
	TxtRecords TxtRecordArrayInput
	// The name of the DNS zone (without a terminating dot).
	ZoneName pulumi.StringInput
}

func (RecordSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetArgs)(nil)).Elem()
}

type RecordSetInput interface {
	pulumi.Input

	ToRecordSetOutput() RecordSetOutput
	ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput
}

func (*RecordSet) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSet)(nil))
}

func (i *RecordSet) ToRecordSetOutput() RecordSetOutput {
	return i.ToRecordSetOutputWithContext(context.Background())
}

func (i *RecordSet) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetOutput)
}

type RecordSetOutput struct {
	*pulumi.OutputState
}

func (RecordSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSet)(nil))
}

func (o RecordSetOutput) ToRecordSetOutput() RecordSetOutput {
	return o
}

func (o RecordSetOutput) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RecordSetOutput{})
}
