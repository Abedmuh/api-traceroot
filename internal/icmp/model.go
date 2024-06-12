package icmp

type IcmpSsh struct {
	Address string   `json:"address" validate:"required"`
	Command string   `json:"command" validate:"required"`
	Router  []string `json:"routers" validate:"required"`
}
