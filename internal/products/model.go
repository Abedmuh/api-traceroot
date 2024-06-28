package products

type Products struct {
	Name     string `json:"name" validate:"required"`
	Os       string `json:"os" validate:"required"`
	Cpu      int32  `json:"cpu" validate:"required"`
	Ram      int64  `json:"ram" validate:"required"`
	Storage  int64  `json:"storage" validate:"required"`
	Firewall bool   `json:"firewall" validate:"required"`
	Selinux  string `json:"selinux" validate:"required"`
	Location string `json:"location" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//request
