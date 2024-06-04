package products

//interface

type ProdSvcInter interface {
	GetProducts() ([]Products, error)
}

type ProdSvcImpl struct {
}

func NewProdSvc() ProdSvcInter {
	return &ProdSvcImpl{}
}

func (p *ProdSvcImpl) GetProducts() ([]Products, error) {
	return nil, nil
}
