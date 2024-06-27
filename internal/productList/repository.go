package productlist

import (
	"fmt"
	"strconv"

	"github.com/Abedmuh/api-traceroot/utils/esxiSession"
	"github.com/gin-gonic/gin"
)

func CreateVmWithESXI(ctx *gin.Context, product ProductList) error {

	vcpu, err := strconv.ParseInt(product.Cpu, 10, 32)
    if err != nil {
        fmt.Println("Error:", err)
        return err
    }
	vram, err := strconv.ParseInt(product.Ram, 10, 32)
	if err!= nil {
        fmt.Println("Error:", err)
        return err
    }
	hdd, err := strconv.ParseInt(product.Storage, 10, 32)
	if err!= nil {
        fmt.Println("Error:", err)
        return err
    }

	//data
	sessionData := esxiSession.SessionData{
		Username: product.Username,
		Password: product.Password,
		VmName: product.Name,
		Cpu: int32(vcpu),
		Ram: vram,
		Storage: hdd,
		OsGuestId: product.Os,
	}

	//stage 1: login to ESXI server
	client, err := esxiSession.LoginESXI(ctx)
	if err != nil {
		return err
	}
	defer esxiSession.Logout(ctx, client)

	// stage 2: find resources
	resources, err := esxiSession.FindResources(ctx, client)
	if err != nil {
		fmt.Printf("Error finding resources: %s\n", err)
		return err
	}

	// Stage 3: Create VM
	vm, err := esxiSession.CreateVM(ctx, client, resources, sessionData)
	if err != nil {
		fmt.Printf("Error creating VM: %s\n", err)
		return err
	}
	fmt.Printf("VM %s created successfully\n", vm.Name())

	// Print VM details
	err = esxiSession.PrintVMDetails(ctx, vm)
	if err != nil {
		fmt.Printf("Error retrieving VM details: %s\n", err)
		return err
	}

	return nil
}