// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents the serial port of the parent resource.
type SerialPort struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether the port is enabled for a serial console connection.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSerialPort registers a new resource with the given unique name, arguments, and options.
func NewSerialPort(ctx *pulumi.Context,
	name string, args *SerialPortArgs, opts ...pulumi.ResourceOption) (*SerialPort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentResource == nil {
		return nil, errors.New("invalid value for required argument 'ParentResource'")
	}
	if args.ParentResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourceType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ResourceProviderNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:serialconsole/v20180501:SerialPort"),
		},
		{
			Type: pulumi.String("azure-native:serialconsole:SerialPort"),
		},
		{
			Type: pulumi.String("azure-nextgen:serialconsole:SerialPort"),
		},
	})
	opts = append(opts, aliases)
	var resource SerialPort
	err := ctx.RegisterResource("azure-native:serialconsole/v20180501:SerialPort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSerialPort gets an existing SerialPort resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSerialPort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SerialPortState, opts ...pulumi.ResourceOption) (*SerialPort, error) {
	var resource SerialPort
	err := ctx.ReadResource("azure-native:serialconsole/v20180501:SerialPort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SerialPort resources.
type serialPortState struct {
	// Resource name
	Name *string `pulumi:"name"`
	// Specifies whether the port is enabled for a serial console connection.
	State *string `pulumi:"state"`
	// Resource type
	Type *string `pulumi:"type"`
}

type SerialPortState struct {
	// Resource name
	Name pulumi.StringPtrInput
	// Specifies whether the port is enabled for a serial console connection.
	State pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SerialPortState) ElementType() reflect.Type {
	return reflect.TypeOf((*serialPortState)(nil)).Elem()
}

type serialPortArgs struct {
	// The resource name, or subordinate path, for the parent of the serial port. For example: the name of the virtual machine.
	ParentResource string `pulumi:"parentResource"`
	// The resource type of the parent resource.  For example: 'virtualMachines' or 'virtualMachineScaleSets'
	ParentResourceType string `pulumi:"parentResourceType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The namespace of the resource provider.
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	// The name of the serial port to create.
	SerialPort *string `pulumi:"serialPort"`
	// Specifies whether the port is enabled for a serial console connection.
	State *string `pulumi:"state"`
}

// The set of arguments for constructing a SerialPort resource.
type SerialPortArgs struct {
	// The resource name, or subordinate path, for the parent of the serial port. For example: the name of the virtual machine.
	ParentResource pulumi.StringInput
	// The resource type of the parent resource.  For example: 'virtualMachines' or 'virtualMachineScaleSets'
	ParentResourceType pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The namespace of the resource provider.
	ResourceProviderNamespace pulumi.StringInput
	// The name of the serial port to create.
	SerialPort pulumi.StringPtrInput
	// Specifies whether the port is enabled for a serial console connection.
	State *SerialPortStateEnum
}

func (SerialPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serialPortArgs)(nil)).Elem()
}

type SerialPortInput interface {
	pulumi.Input

	ToSerialPortOutput() SerialPortOutput
	ToSerialPortOutputWithContext(ctx context.Context) SerialPortOutput
}

func (*SerialPort) ElementType() reflect.Type {
	return reflect.TypeOf((*SerialPort)(nil))
}

func (i *SerialPort) ToSerialPortOutput() SerialPortOutput {
	return i.ToSerialPortOutputWithContext(context.Background())
}

func (i *SerialPort) ToSerialPortOutputWithContext(ctx context.Context) SerialPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SerialPortOutput)
}

type SerialPortOutput struct {
	*pulumi.OutputState
}

func (SerialPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SerialPort)(nil))
}

func (o SerialPortOutput) ToSerialPortOutput() SerialPortOutput {
	return o
}

func (o SerialPortOutput) ToSerialPortOutputWithContext(ctx context.Context) SerialPortOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SerialPortOutput{})
}
