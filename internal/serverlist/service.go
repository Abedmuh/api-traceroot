package serverlist

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/Abedmuh/api-traceroot/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServerListSvcInter interface {
	CreateServerList(req ReqServerList, tx *gorm.DB, ctx *gin.Context) (ServerList,error)
	GetServerLists(tx *gorm.DB, ctx *gin.Context) ([]ServerList, error)
	GetServerListById(id string, tx *gorm.DB, ctx *gin.Context) (ServerList, error)
	UpdateServerList(id string, req ServerList, tx *gorm.DB, ctx *gin.Context) error
	DeleteServerList(tx *gorm.DB, ctx *gin.Context) error
	TestAnsibleServer(ctx *gin.Context) (string, error)
}

type ServerListSvcImpl struct {
}

func NewServerListService() ServerListSvcInter {
    return &ServerListSvcImpl{}
}

func (s *ServerListSvcImpl) CreateServerList(req ReqServerList, tx *gorm.DB, ctx *gin.Context) (ServerList,error) {
	user, err := utils.GetTokenEmail(ctx)
	if err!= nil {
        return ServerList{}, err
    }

	serverList := ServerList{
		Owner: user,
        Username: req.Username,
		Password: req.Rootpass,
		Timelimit: time.Now().Add(time.Duration(24) * time.Hour), 
		Name: req.Name,
        Os:    req.Os,
        Cpu:   req.Cpu,
        Ram:   req.Ram,
        Storage: req.Storage,
		Firewall: req.Firewall,
        Selinux: req.Selinux,
        Location: req.Location,       
	}

	if err := tx.Create(&serverList).Error; err!= nil {
        return ServerList{},err
    }
	return serverList, nil
}

func (s *ServerListSvcImpl) GetServerLists(tx *gorm.DB, ctx *gin.Context) ([]ServerList, error) {
	var serverLists []ServerList
    user, err := utils.GetTokenEmail(ctx)
    if err!= nil {
        return nil, err
    }

    if err := tx.Where("owner =?", user).Find(&serverLists).Error; err!= nil {
        return nil, err
    }
    return serverLists, nil
}

func (s *ServerListSvcImpl) GetServerListById(id string, tx *gorm.DB, ctx *gin.Context) (ServerList, error) {
	var serverList ServerList
    user, err := utils.GetTokenEmail(ctx)
    if err!= nil {
        return serverList, err
	}

	if err := tx.Where("id =? AND owner =?", id, user).First(&serverList).Error; err!= nil {
        return serverList, err
    }
	return serverList, nil
}

func (s *ServerListSvcImpl) UpdateServerList(id string, req ServerList, tx *gorm.DB, ctx *gin.Context) error {
    // Your code here
	return nil
}

func (s *ServerListSvcImpl) DeleteServerList(tx *gorm.DB, ctx *gin.Context) error {
	return nil
}

func (s *ServerListSvcImpl) TestAnsibleServer(ctx *gin.Context) (string,error) {

	// Example arguments
    esxiValidateCerts := "false"
    esxiName := "Test-Athena-Cloner"
    esxiOvf := "./res/Test-Athena-centos709.ovf"
    esxiNetworks := "vl99-POC-VPS"
    guestHardwareEsxiNumCpu := "2"
    guestHardwareEsxiMemoryMb := "4096"
    guestHardwareEsxiStorage := "24"

    scriptPath := "./commandscript/deployovf.sh"

    cmd := exec.Command("sh", scriptPath, 
	esxiValidateCerts, 
	esxiName, 
	esxiOvf, 
	esxiNetworks, 
	guestHardwareEsxiNumCpu, 
	guestHardwareEsxiMemoryMb, 
	guestHardwareEsxiStorage)
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Log the detailed error message
		log.Printf("Ansible playbook error: %s\nOutput: %s", err.Error(), string(output))
		return "", fmt.Errorf("ansible playbook execution failed: %s", string(output))
	}
    
	return scriptPath, nil
}



