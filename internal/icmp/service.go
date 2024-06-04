package icmp

type icmpSvc interface {
	Ping(addr string) (int, error)
	TraceRoute(addr string) (int, error)
}

type icmpSvcImpl struct {
}

func NewIcmpSvc() icmpSvc {
	return &icmpSvcImpl{}
}

func (i *icmpSvcImpl) Ping(addr string) (int, error) {

	return 0, nil
}

func (i *icmpSvcImpl) TraceRoute(addr string) (int, error) {
	return 0, nil
}
