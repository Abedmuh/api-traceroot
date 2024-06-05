package products

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// interfaces
type ProdCtrlInter interface {
	GetProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	PostProduct(c *gin.Context)
	PutProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type ProdCtrlImpl struct {
	service  ProdSvcInter
	Db       *gorm.DB
	validate *validator.Validate
}

func NewProductController(service ProdSvcInter, db *gorm.DB, validate *validator.Validate) ProdCtrlInter {
	return &ProdCtrlImpl{
		service:  service,
		Db:       db,
		validate: validate,
	}
}

func (c *ProdCtrlImpl) GetProducts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (c *ProdCtrlImpl) GetProductByID(ctx *gin.Context) {}

func (c *ProdCtrlImpl) PostProduct(ctx *gin.Context) {}

func (c *ProdCtrlImpl) PutProduct(ctx *gin.Context) {}

func (c *ProdCtrlImpl) DeleteProduct(ctx *gin.Context) {}
