package utils

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func SshTarget(addr string, command string) (string, error) {
	// SSH connection configuration
	config := &ssh.ClientConfig{
		User: "abdillah",
		Auth: []ssh.AuthMethod{
			ssh.Password("abdillah"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server
	client, err := ssh.Dial("tcp", "103.130.198.130:22", config)
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
	fullCommand := fmt.Sprintf("%s %s", command, addr)
	output, err := session.CombinedOutput(fullCommand)
	if err != nil {
		return "", fmt.Errorf("failed to run command: %v", err)
	}

	fmt.Println(string(output))

	return string(output), nil
}
