package productlist

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// interface
type ProdListCtrlInter interface {
	GetProductLists(c *gin.Context)
	GetProductList(c *gin.Context)
	PostProductList(c *gin.Context)
	PutProductList(c *gin.Context)
}

type ProdListCtrlImpl struct {
	service   ProdListSvcInter
	Db        *sql.DB
	validator *validator.Validate
}

func NewProdListCtrl(service ProdListSvcInter, Db *sql.DB, validator *validator.Validate) ProdListCtrlInter {
	return &ProdListCtrlImpl{
		service:   service,
		Db:        Db,
		validator: validator,
	}
}

func (c *ProdListCtrlImpl) GetProductLists(ctx *gin.Context) {}

func (c *ProdListCtrlImpl) GetProductList(ctx *gin.Context) {}

func (c *ProdListCtrlImpl) PostProductList(ctx *gin.Context) {}

func (c *ProdListCtrlImpl) PutProductList(ctx *gin.Context) {}
