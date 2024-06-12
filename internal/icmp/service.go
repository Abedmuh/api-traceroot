package icmp

import (
	"fmt"

	"github.com/spf13/viper"
)

type icmpSvcInter interface {
	LookingGlass(req IcmpSsh) (string, error)
	TraceRoute(addr string) (int, error)
}

type icmpSvcImpl struct {
}

func NewIcmpSvc() icmpSvcInter {
	return &icmpSvcImpl{}
}

func (i *icmpSvcImpl) LookingGlass(req IcmpSsh) (string, error) {
	targetSSH := viper.GetString("SSH_TARGET")
	username := viper.GetString("SSH_USERNAME")
	password := viper.GetString("SSH_PASSWORD")

	commandList := map[string]string{
		"ping":       "ping -c 8",
		"traceroute": "/usr/sbin/mtr -rnc 3",
	}

	replacement, exists := commandList[req.Command]
	if !exists {
		return "", fmt.Errorf("unsupported command: %s", req.Command)
	}

	// locations, err := convertRouterToLocations(req.Router)
	// if err != nil {
	// 	fmt.Println("Conversion failed:", err)
	// 	return "", fmt.Errorf("Router not available: %s", req.Command)
	// }


	target := SshTargeting{
		Address:   req.Address,
		Command:   replacement,
		TargetSSH: targetSSH,
		Username:  username,
		Password:  password,
	}

	result, err := SshTarget(target)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (i *icmpSvcImpl) TraceRoute(addr string) (int, error) {
	return 0, nil
}
