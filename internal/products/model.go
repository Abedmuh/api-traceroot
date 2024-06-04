package products

//main
type Products struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Os       string `json:"os"`
	Cpu      string `json:"cpu"`
	Storage  string `json:"storage"`
	Firewall bool   `json:"firewall"`
	Selinux  string `json:"selinux"`
	Location string `json:"location"`
}

//request
