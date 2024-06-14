package icmp

import (
	"errors"

	"github.com/spf13/viper"
)

type IcmpSsh struct {
	Address string `json:"address" validate:"required"`
	Command string `json:"command" validate:"required,command"`
	Router  string `json:"router" validate:"required"`
}

type SshTargeting struct {
	Address   string
	Command   string
	Host 	  string
	Username  string
	Password  string
}

type IcmpSSHs struct {
	Target  string    `json:"target" validate:"required"`
	Command string    `json:"command" validate:"required,command"`
	Routers []Routers `json:"routers" validate:"required"`
}

type Routers struct {
	Name     string  `json:"name" validate:"required"`
	Address  *string `json:"address"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type Server struct {
	Name     string
	Host     string
	Username string
	Password string
}

var errSSHconnection = errors.New("ssh connection failed")

var commandList = map[string]string{
	"ping":       "ping -c 3",
	"traceroute": "/usr/sbin/mtr -rnc 10",
}

var servers = []Server{
	{
		Name:     "jakarta",
		Host:  	  "103.130.198.130:22",
		Username: viper.GetString("SSH_JAKARTA_USERNAME"),
		Password: viper.GetString("SSH_JAKARTA_PASSWORD"),
	},
	{
		Name:     "bandung",
		Host:     "190.xxx.xx.xx",
		Username: viper.GetString("SSH_BANDUNG_USERNAME"),
		Password: viper.GetString("SSH_BANDUNG_PASSWORD"),
	},
}