package esxiSession

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
)

func PrintVMDetails(ctx *gin.Context, vm *object.VirtualMachine) error {
	var vmMo mo.VirtualMachine
	err := vm.Properties(ctx, vm.Reference(), []string{"name", "config", "summary"}, &vmMo)
	if err != nil {
		return fmt.Errorf("error retrieving VM properties: %w", err)
	}

	fmt.Printf("VM Name: %s\n", vmMo.Name)
	fmt.Printf("VM Guest ID: %s\n", vmMo.Config.GuestId)
	fmt.Printf("VM CPU Count: %d\n", vmMo.Config.Hardware.NumCPU)
	fmt.Printf("VM Memory Size: %d MB\n", vmMo.Config.Hardware.MemoryMB)
	fmt.Printf("VM Power State: %s\n", vmMo.Summary.Runtime.PowerState)
	return nil
}