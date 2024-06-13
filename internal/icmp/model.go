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
	Address  string
	Username string
	Password string
}