package icmp

type IcmpSsh struct {
	Address string `json:"address" validate:"required"`
	Command string `json:"command" validate:"required,command"`
	// Router  []string `json:"routers" validate:"required"`
}

type SshTargeting struct {
	Address   string
	Command   string
	TargetSSH string
	Username  string
	Password  string
}