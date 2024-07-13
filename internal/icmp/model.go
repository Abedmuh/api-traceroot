package icmp

import (
	"errors"
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
var servers []Server
var commandList = map[string]string{
	"ping":       "ping -c 10",
	"traceroute": "/usr/sbin/mtr -rnc 10",
}