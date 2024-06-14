package icmp

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func SshTarget(target SshTargeting) (string, error) {

	// SSH connection configuration
	config := &ssh.ClientConfig{
		User: target.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(target.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server
	client, err := ssh.Dial("tcp", target.Host, config)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to dial: %v", err))
		return "", errSSHconnection
	}
	defer client.Close()

	// Create a session
	session, err := client.NewSession()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to create session: %v", err))
		return "", errSSHconnection
	}
	defer session.Close()

	// Execute the command with the IP address
	fullCommand := fmt.Sprintf("%s %s", target.Command, target.Address)
	output, err := session.CombinedOutput(fullCommand)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to run command: %v", err))
		return "", errSSHconnection
	}

	return string(output), nil
}

func fillTheRouter(routers []Routers, servers []Server) error {
	serverMap := make(map[string]Server)
	for _, server := range servers {
		serverMap[server.Name] = server
	}

	for i, router := range routers {
		server, exists := serverMap[router.Name]
		if !exists {
			return fmt.Errorf("unsupported router location: %s", router.Name)
		}

		if router.Address == nil {
			routers[i].Address = &server.Host
		}

		if router.Username == nil {
			routers[i].Username = &server.Username
		}

		if router.Password == nil {
			routers[i].Password = &server.Password
		}
	}
	return nil
}
