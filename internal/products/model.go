package products

//main
type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Os       string `json:"os"`
	Cpu      string `json:"cpu"`
	Storage  string `json:"storage"`
	Firewall bool   `json:"firewall"`
	Selinux  string `json:"selinux"`
	Location string `json:"location"`
}

type Products struct {
	Name     string `json:"name" validate:"required"`
	Os       string `json:"os" validate:"required"`
	Cpu      string `json:"cpu" validate:"required"`
	Ram      string `json:"ram" validate:"required"`
	Storage  string `json:"storage" validate:"required"`
	Firewall bool   `json:"firewall" validate:"required"`
	Selinux  string `json:"selinux" validate:"required"`
	Location string `json:"location" validate:"required"`
}

//request
