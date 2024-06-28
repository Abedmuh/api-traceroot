package esxiSession

import (
	"github.com/vmware/govmomi/object"
)

type Resources struct {
	Datacenter   *object.Datacenter
	Datastore    *object.Datastore
	ResourcePool *object.ResourcePool
	Folder       *object.Folder
	Network      object.NetworkReference
}

type SessionData struct {
	Username string
	Password string
	VmName string
	Cpu int32
	Ram int64
	Storage int64
	OsGuestId string
	Location string
}