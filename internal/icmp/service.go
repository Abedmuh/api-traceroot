package icmp

import (
	"fmt"

	"github.com/spf13/viper"
)

type icmpSvcInter interface {
	LookingGlass(req IcmpSsh) (string, error)
	ListedLG(req IcmpSSHs) (map[string]string, error)
}

type icmpSvcImpl struct {
}

func NewIcmpSvc() icmpSvcInter {
	return &icmpSvcImpl{}
}

func (i *icmpSvcImpl) LookingGlass(req IcmpSsh) (string, error) {
	targetSSH := viper.GetString("SSH_TARGET")
	username := viper.GetString("SSH_JAKARTA_USERNAME")
	password := viper.GetString("SSH_JAKARTA_PASSWORD")

	commandList := map[string]string{
		"ping":       "ping -c 10",
		"traceroute": "/usr/sbin/mtr -rnc 10",
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

func (i *icmpSvcImpl) ListedLG(req IcmpSSHs) (map[string]string, error) {

	commandList := map[string]string{
		"ping":       "ping -c 3",
		"traceroute": "/usr/sbin/mtr -rnc 10",
	}

	command, exists := commandList[req.Command]
	if !exists {
		return nil, fmt.Errorf("unsupported command: %s", req.Command)
	}

	servers := []Server{
		{
			Name:     "jakarta",
			Address:  "103.130.198.130:22",
			Username: viper.GetString("SSH_JAKARTA_USERNAME"),
			Password: viper.GetString("SSH_JAKARTA_PASSWORD"),
		},
		{
			Name:     "bandung",
			Address:  "190.xxx.xx.xx",
			Username: viper.GetString("SSH_BANDUNG_USERNAME"),
			Password: viper.GetString("SSH_BANDUNG_PASSWORD"),
		},
	}

	if err := fillTheRouter(req.Routers, servers); err != nil {
		return nil, err
	}

	results := make(map[string]string)
	for _, router := range req.Routers {
		target := SshTargeting{
			Address:   req.Target,
			Command:   command,
			TargetSSH: *router.Address,
			Username:  *router.Username,
			Password:  *router.Password,
		}

		result, err := SshTarget(target)
		if err != nil {
			return nil, err
		}

		results[router.Name] = result
	}


	return results, nil
}
