package esxiSession

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

func CreateVM(ctx *gin.Context, client *govmomi.Client, resources *Resources, reqData SessionData) (*object.VirtualMachine, error) {
	ideController := &types.VirtualDeviceConfigSpec{
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
        Device: &types.VirtualIDEController{
			VirtualController: types.VirtualController{
				VirtualDevice: types.VirtualDevice{
					Key: 200,
				},
				BusNumber: 0,
				Device: []int32{},
			},
		},
	}

	dvdDrive := &types.VirtualDeviceConfigSpec{
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
		Device: &types.VirtualCdrom{
			VirtualDevice: types.VirtualDevice{
				Key: 2000,
				Backing: &types.VirtualCdromIsoBackingInfo{
					VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
						FileName: "[" + resources.Datastore.Name() + "] " + reqData.Location,
					},
				},
				ControllerKey: 200, // Key of the controller to which the CD/DVD drive is attached
				UnitNumber:    types.NewInt32(0),
				Connectable: &types.VirtualDeviceConnectInfo{
					StartConnected: true,
					AllowGuestControl: true,
                    Connected: true,
				},
			},
		},
	}

	scsiCOntroller := &types.VirtualDeviceConfigSpec{
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
		Device: &types.VirtualLsiLogicController{
			VirtualSCSIController: types.VirtualSCSIController{
				SharedBus: types.VirtualSCSISharingNoSharing,
				VirtualController: types.VirtualController{
					VirtualDevice: types.VirtualDevice{
						Key: 1000,
					},
					BusNumber: 0,
				},
			},
		},
	}

	storageDisk := &types.VirtualDeviceConfigSpec{
		Operation:     types.VirtualDeviceConfigSpecOperationAdd,
		FileOperation: types.VirtualDeviceConfigSpecFileOperationCreate,
		Device: &types.VirtualDisk{
			CapacityInKB:  reqData.Storage * 1024 * 1024,
			VirtualDevice: types.VirtualDevice{
				Backing: &types.VirtualDiskFlatVer2BackingInfo{
					DiskMode:        string(types.VirtualDiskModePersistent),
					ThinProvisioned: types.NewBool(true),
					VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
						FileName: "[" + resources.Datastore.Name() + "] " + reqData.VmName + "/" + reqData.VmName + ".vmdk",
					},
				},
				ControllerKey: 1000,
				UnitNumber:    types.NewInt32(0),
			},
		},
	}

	networkSrc := &types.VirtualDeviceConfigSpec{
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
		Device: &types.VirtualVmxnet3{
			VirtualVmxnet: types.VirtualVmxnet{
				VirtualEthernetCard: types.VirtualEthernetCard{
					VirtualDevice: types.VirtualDevice{
                        Backing: &types.VirtualEthernetCardNetworkBackingInfo{
                            VirtualDeviceDeviceBackingInfo: types.VirtualDeviceDeviceBackingInfo{
								DeviceName: viper.GetString("VSPHERE_NETWORK"),
							},
                        },
                    },
					// MacAddress: "14:02:ec:34:7a:d0",
				},
			},
		},
	}

	// Define VM configuration
	vmConfigSpec := types.VirtualMachineConfigSpec{
		Name:     reqData.VmName,
		GuestId:  reqData.OsGuestId,
		Files:    &types.VirtualMachineFileInfo{VmPathName: "[" + resources.Datastore.Name() + "]"},
		NumCPUs:  reqData.Cpu,
		MemoryMB: reqData.Ram,
		DeviceChange: []types.BaseVirtualDeviceConfigSpec{
			ideController,
			scsiCOntroller,
			storageDisk,
			dvdDrive,
			networkSrc,
		},
	}

	// Create the VM
	task, err := resources.Folder.CreateVM(ctx, vmConfigSpec, resources.ResourcePool, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create VM: %w", err)
	}

	// Wait for the task to complete
	taskInfo, err := task.WaitForResult(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for result: %w", err)
	}

	vm := object.NewVirtualMachine(client.Client, taskInfo.Result.(types.ManagedObjectReference))

	err = TurnVm(vm, ctx, true)
	if err!= nil {
        return nil, fmt.Errorf("failed to turn VM on: %w", err)
    }
	return vm, nil
}

func DestroyVMByName(ctx context.Context, client *govmomi.Client, vmName string) error {
	finder := find.NewFinder(client.Client, false)

	datacenter, err := finder.DefaultDatacenter(ctx)
	if err != nil {
		return fmt.Errorf("error finding datacenter: %w", err)
	}

	finder.SetDatacenter(datacenter)

	vm, err := finder.VirtualMachine(ctx, vmName)
	if err != nil {
		return fmt.Errorf("error finding VM '%s': %w", vmName, err)
	}

	powerState, err := vm.PowerState(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving VM power state: %w", err)
	}
	if powerState != "poweredOff" {
		task, err := vm.PowerOff(ctx)
		if err != nil {
			return fmt.Errorf("error powering off VM '%s': %w", vmName, err)
		}
		if err := task.Wait(ctx); err != nil {
			return fmt.Errorf("error waiting for VM power off task: %w", err)
		}
		fmt.Printf("VM '%s' powered off successfully.\n", vmName)
	}

	task, err := vm.Destroy(ctx)
	if err != nil {
		return fmt.Errorf("error destroying VM '%s': %w", vmName, err)
	}
	if err := task.Wait(ctx); err != nil {
		return fmt.Errorf("error waiting for VM destroy task: %w", err)
	}
	fmt.Printf("VM '%s' destroyed successfully.\n", vmName)

	return nil
}