// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource representation of a service topology.
type ServiceTopology struct {
	pulumi.CustomResourceState

	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId pulumi.StringPtrOutput `pulumi:"artifactSourceId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceTopology registers a new resource with the given unique name, arguments, and options.
func NewServiceTopology(ctx *pulumi.Context,
	name string, args *ServiceTopologyArgs, opts ...pulumi.ResourceOption) (*ServiceTopology, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager/v20191101preview:ServiceTopology"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager:ServiceTopology"),
		},
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager:ServiceTopology"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager/v20180901preview:ServiceTopology"),
		},
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager/v20180901preview:ServiceTopology"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceTopology
	err := ctx.RegisterResource("azure-native:deploymentmanager/v20191101preview:ServiceTopology", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceTopology gets an existing ServiceTopology resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceTopology(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceTopologyState, opts ...pulumi.ResourceOption) (*ServiceTopology, error) {
	var resource ServiceTopology
	err := ctx.ReadResource("azure-native:deploymentmanager/v20191101preview:ServiceTopology", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceTopology resources.
type serviceTopologyState struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId *string `pulumi:"artifactSourceId"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type ServiceTopologyState struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (ServiceTopologyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTopologyState)(nil)).Elem()
}

type serviceTopologyArgs struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId *string `pulumi:"artifactSourceId"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service topology .
	ServiceTopologyName *string `pulumi:"serviceTopologyName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ServiceTopology resource.
type ServiceTopologyArgs struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the service topology .
	ServiceTopologyName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ServiceTopologyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTopologyArgs)(nil)).Elem()
}

type ServiceTopologyInput interface {
	pulumi.Input

	ToServiceTopologyOutput() ServiceTopologyOutput
	ToServiceTopologyOutputWithContext(ctx context.Context) ServiceTopologyOutput
}

func (*ServiceTopology) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTopology)(nil))
}

func (i *ServiceTopology) ToServiceTopologyOutput() ServiceTopologyOutput {
	return i.ToServiceTopologyOutputWithContext(context.Background())
}

func (i *ServiceTopology) ToServiceTopologyOutputWithContext(ctx context.Context) ServiceTopologyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTopologyOutput)
}

type ServiceTopologyOutput struct {
	*pulumi.OutputState
}

func (ServiceTopologyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTopology)(nil))
}

func (o ServiceTopologyOutput) ToServiceTopologyOutput() ServiceTopologyOutput {
	return o
}

func (o ServiceTopologyOutput) ToServiceTopologyOutputWithContext(ctx context.Context) ServiceTopologyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceTopologyOutput{})
}
