package esxiSession

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vmware/govmomi/object"
)

func TurnVm(vm *object.VirtualMachine,ctx *gin.Context, status bool) error {
	var task *object.Task
	var err error

	if status {
		task, err = vm.PowerOn(ctx)
		if err != nil {
			return fmt.Errorf("failed to power on VM: %w", err)
		}
	} else {
		task, err = vm.PowerOff(ctx)
		if err != nil {
			return fmt.Errorf("failed to power off VM: %w", err)
		}
	}

	err = task.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for task: %w", err)
	}

	time.Sleep(5 * time.Second) // Wait for operations to complete
	return nil
}

func SuspendVm(vm *object.VirtualMachine, ctx *gin.Context) error {
	suspendTask, err := vm.Suspend(ctx)
	if err != nil {
		return fmt.Errorf("failed to suspend VM: %w", err)
	}

	err = suspendTask.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to wait for suspend task: %w", err)
	}

	return nil
}