package esxiSession

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
)

func LoginESXI(ctx *gin.Context) (*govmomi.Client, error) {
	var (
		vsphereUser 	 = viper.GetString("VSPHERE_USER")
		vspherePassword  = viper.GetString("VSPHERE_PASSWORD")
		vsphereURL 		 = viper.GetString("VSPHERE_URL")
	)
	// resourcePoolName = "vm"

	fmt.Println("Creating a VIM/SOAP session.")
	vcURL := "https://" + vsphereUser + ":" + vspherePassword + "@" + vsphereURL + "/sdk"
	u, err := url.Parse(vcURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing url %s: %w", vcURL, err)
	}
	client, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		return nil, fmt.Errorf("error logging in: %w", err)
	}
	fmt.Println("Login successful")
	return client, nil
}

func Logout(ctx *gin.Context, client *govmomi.Client) {
	if err := client.Logout(ctx); err != nil {
		fmt.Printf("Error logging out: %s\n", err)
	}
}
func FindResources(ctx *gin.Context, client *govmomi.Client) (*Resources, error) {

	var (
		datacenterName   = viper.GetString("VSPHERE_DATACENTER")
		datastoreName    = viper.GetString("VSPHERE_DATASTORE")
		networkName 	 = viper.GetString("VSPHERE_NETWORK")
		folderName       = "vm"
	)

	finder := find.NewFinder(client.Client, true)

	// Find the datacenter
	dc, err := finder.DatacenterOrDefault(ctx, datacenterName)
	if err != nil {
		return nil, fmt.Errorf("failed to find datacenter: %w", err)
	}
	finder.SetDatacenter(dc)

	// Find the network
	network, err := finder.Network(ctx, networkName)
	if err != nil {
		return nil, fmt.Errorf("failed to find network: %w", err)
	}

	// Find the datastore
	ds, err := finder.Datastore(ctx, datastoreName)
	if err != nil {
		return nil, fmt.Errorf("failed to find datastore: %w", err)
	}

	// Find the resource pool
	rp, err := finder.DefaultResourcePool(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find resource pool: %w", err)
	}

	// Find the VM folder
	folder, err := finder.Folder(ctx, folderName)
	if err != nil {
		return nil, fmt.Errorf("failed to find VM folder: %w", err)
	}

	return &Resources{
		Datacenter:   dc,
		Datastore:    ds,
		ResourcePool: rp,
		Folder:       folder,
		Network:      network,
	}, nil
}