package productlist

import (
	"fmt"

	"github.com/Abedmuh/api-traceroot/utils/esxiSession"
	"github.com/gin-gonic/gin"
)

func CreateVmWithESXI(ctx *gin.Context, product ProductList) error {

	os, exists := OsMap[product.Os]
	if !exists {
		return fmt.Errorf("unsupported os location: %s", product.Os)
	}

	//data
	sessionData := esxiSession.SessionData{
		Username: product.Username,
		Password: product.Password,
		VmName: product.Name,
		Cpu: product.Cpu,
		Ram: product.Ram,
		Storage: product.Storage,
		OsGuestId: product.Os,
		Location: os.Location,
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