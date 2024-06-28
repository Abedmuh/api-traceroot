package productlist

import (
	"github.com/Abedmuh/api-traceroot/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// interface
type ProdListCtrlInter interface {
	GetProductLists(c *gin.Context)
	GetProductListById(c *gin.Context)
	PostProductList(c *gin.Context)
	PutProductList(c *gin.Context)
	DeleteProductList(c *gin.Context)
}

type ProdListCtrlImpl struct {
	service  ProdListSvcInter
	Db       *gorm.DB
	validate *validator.Validate
}

func NewProdListCtrl(service ProdListSvcInter, Db *gorm.DB, validate *validator.Validate) ProdListCtrlInter {
	return &ProdListCtrlImpl{
		service:  service,
		Db:       Db,
		validate: validate,
	}
}

func (c *ProdListCtrlImpl) GetProductLists(ctx *gin.Context) {

}

func (c *ProdListCtrlImpl) GetProductListById(ctx *gin.Context) {
	var req products.Products
    if err := ctx.ShouldBindJSON(&req); err!= nil {
        ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := c.validate.Struct(req); err!= nil {
        ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
        return
    }
	
}

func (c *ProdListCtrlImpl) PostProductList(ctx *gin.Context) {
	var req products.Products
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.validate.Struct(req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	
	res, err := c.service.CreateProductList(req, c.Db, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	err = CreateVmWithESXI(ctx, res) 
	if err!= nil {
		ctx.AbortWithStatusJSON(503, gin.H{"error": err.Error()})
        return 
    }

	ctx.JSON(200, gin.H{
		"data":    res,
		"message": "Successfully created product",
	})
}

func (c *ProdListCtrlImpl) PutProductList(ctx *gin.Context) {}

func (c *ProdListCtrlImpl) DeleteProductList(ctx *gin.Context) {}
