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
	client, err := ssh.Dial("tcp", target.TargetSSH, config)
	if err != nil {
		return "", fmt.Errorf("failed to dial: %v", err)
	}
	defer client.Close()

	// Create a session
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	// Execute the command with the IP address
	fullCommand := fmt.Sprintf("%s %s", target.Command, target.Address)
	output, err := session.CombinedOutput(fullCommand)
	if err != nil {
		return "", fmt.Errorf("failed to run command: %v", err)
	}

	return string(output), nil
}

func ConvertRouterToLocations(routers []string) ([]string, error) {
	var result []string
	var locationList = map[string]string{
		"jakarta": "103.130.198.130:22",
		"bandung": "103.130.198.190:22",
		// Add more locations as needed
	}
	for _, router := range routers {
		location, exists := locationList[router]
		if !exists {
			return nil, fmt.Errorf("unsupported router location: %s", router)
		}
		result = append(result, location)
	}
	return result, nil
}
