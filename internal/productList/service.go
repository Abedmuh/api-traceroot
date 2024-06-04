package productlist

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// interface
type ProdListSvcInter interface {
	CreateProductList(req ReqProdList, tx *sql.DB, ctx *gin.Context) error
	GetProductsLists(req ReqProdList, tx *sql.DB, ctx *gin.Context) error
	GetProductsList(req ReqProdList, tx *sql.DB, ctx *gin.Context) error
	UpdateProductList(req ReqProdList, tx *sql.DB, ctx *gin.Context) error
}

type ProdListSvcImpl struct {
}

func NewProdListSvc() ProdListSvcInter {
	return &ProdListSvcImpl{}
}

func (p *ProdListSvcImpl) CreateProductList(req ReqProdList, tx *sql.DB, ctx *gin.Context) error {
	return nil
}

func (p *ProdListSvcImpl) GetProductsList(req ReqProdList, tx *sql.DB, ctx *gin.Context) error {
	return nil
}

func (p *ProdListSvcImpl) GetProductsLists(req ReqProdList, tx *sql.DB, ctx *gin.Context) error {
	return nil
}

func (p *ProdListSvcImpl) UpdateProductList(req ReqProdList, tx *sql.DB, ctx *gin.Context) error {
	return nil
}
