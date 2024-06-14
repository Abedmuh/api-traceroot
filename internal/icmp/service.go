package icmp

import (
	"fmt"
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
	replacement, exists := commandList[req.Command]
	if !exists {
		return "", fmt.Errorf("unsupported command: %s", req.Command)
	}

	serverMap := make(map[string]Server)
	for _, server := range servers {
		serverMap[server.Name] = server
	}
	server, exists := serverMap[req.Router]
	if !exists {
		return "",fmt.Errorf("unsupported router location: %s", req.Router)
	}

	target := SshTargeting{
		Address:   req.Address,
		Command:   replacement,
		Host: 	   server.Host,
		Username:  server.Username,
		Password:  server.Password,
	}

	result, err := SshTarget(target)
	if err != nil {
		return "", err
	}

	return result, nil
}



func (i *icmpSvcImpl) ListedLG(req IcmpSSHs) (map[string]string, error) {

	command, exists := commandList[req.Command]
	if !exists {
		return nil, fmt.Errorf("unsupported command: %s", req.Command)
	}

	if err := fillTheRouter(req.Routers, servers); err != nil {
		return nil, err
	}

	results := make(map[string]string)
	for _, router := range req.Routers {
		target := SshTargeting{
			Address:   req.Target,
			Command:   command,
			Host: 	   *router.Address,
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
