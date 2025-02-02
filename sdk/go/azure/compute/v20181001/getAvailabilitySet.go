// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
func LookupAvailabilitySet(ctx *pulumi.Context, args *LookupAvailabilitySetArgs, opts ...pulumi.InvokeOption) (*LookupAvailabilitySetResult, error) {
	var rv LookupAvailabilitySetResult
	err := ctx.Invoke("azure-native:compute/v20181001:getAvailabilitySet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAvailabilitySetArgs struct {
	// The name of the availability set.
	AvailabilitySetName string `pulumi:"availabilitySetName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
type LookupAvailabilitySetResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Fault Domain count.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Update Domain count.
	PlatformUpdateDomainCount *int `pulumi:"platformUpdateDomainCount"`
	// Specifies information about the proximity placement group that the availability set should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup *SubResourceResponse `pulumi:"proximityPlacementGroup"`
	// Sku of the availability set, only name is required to be set. See AvailabilitySetSkuTypes for possible set of values. Use 'Aligned' for virtual machines with managed disks and 'Classic' for virtual machines with unmanaged disks. Default value is 'Classic'.
	Sku *SkuResponse `pulumi:"sku"`
	// The resource status information.
	Statuses []InstanceViewStatusResponse `pulumi:"statuses"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// A list of references to all virtual machines in the availability set.
	VirtualMachines []SubResourceResponse `pulumi:"virtualMachines"`
}
