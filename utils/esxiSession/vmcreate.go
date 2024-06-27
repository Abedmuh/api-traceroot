package esxiSession

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vmware/govmomi"
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
						FileName: "[" + resources.Datastore.Name() + "] Operating System/CentOS-7-x86_64-Minimal-2009.iso",
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
		GuestId:  "centos7_64Guest", // Use appopriate GuestId
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

	// Power on the VM
	powerOnTask, err := vm.PowerOn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to power on VM: %w", err)
	}

	// Wait for the power-on task to complete
	err = powerOnTask.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for power-on task: %w", err)
	}

	// Wait for a few seconds to ensure the VM is fully powered on
	time.Sleep(10 * time.Second)
	return vm, nil
}