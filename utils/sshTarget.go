package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

func SshTarget(addr string, command string) (string, error) {

	targetSSH := viper.GetString("SSH_TARGET")
	username := viper.GetString("SSH_USERNAME")
	password := viper.GetString("SSH_PASSWORD")

	// SSH connection configuration
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server
	client, err := ssh.Dial("tcp", targetSSH, config)
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

	return string(output), nil
}
