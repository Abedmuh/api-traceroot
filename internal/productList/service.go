package productlist

import (
	"errors"
	"time"

	"github.com/Abedmuh/api-traceroot/internal/products"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// interface
type ProdListSvcInter interface {
	CreateProductList(req products.Products, tx *gorm.DB, ctx *gin.Context) (ProductList, error)
	GetProductsLists(req ProductList, tx *gorm.DB, ctx *gin.Context) error
	GetProductsList(req ProductList, tx *gorm.DB, ctx *gin.Context) error
	UpdateProductList(req ProductList, tx *gorm.DB, ctx *gin.Context) error
	DeleteProductList(tx *gorm.DB, ctx *gin.Context) error
}

type ProdListSvcImpl struct {
}

func NewProdListSvc() ProdListSvcInter {
	return &ProdListSvcImpl{}
}

func (p *ProdListSvcImpl) CreateProductList(req products.Products, tx *gorm.DB, ctx *gin.Context) (ProductList, error) {
	// decode token auth
	user, err := ctx.Get("email")
	if !err {
		return ProductList{}, errors.New("unathorized cant find user")
	}
	reqUser, err := user.(string)
	if !err {
		return ProductList{}, errors.New("Unathorized")
	}

	//model synchronized
	productList := ProductList{
		Owner:     reqUser,
		Timelimit: time.Now().Add(time.Duration(24) * time.Hour),
		Username:  "myacc",
		Password:  "pass",
		Name:      req.Name,
		Os:        req.Os,
		Cpu:       req.Cpu,
		Storage:   req.Storage,
		Firewall:  req.Firewall,
		Selinux:   req.Storage,
		Location:  req.Location,
	}

	// create product list
	if err := tx.Create(productList).Error; err != nil {
		return ProductList{}, err
	}
	return productList, nil
}

func (p *ProdListSvcImpl) GetProductsList(req ProductList, tx *gorm.DB, ctx *gin.Context) error {
	return nil
}

func (p *ProdListSvcImpl) GetProductsLists(req ProductList, tx *gorm.DB, ctx *gin.Context) error {
	return nil
}

func (p *ProdListSvcImpl) UpdateProductList(req ProductList, tx *gorm.DB, ctx *gin.Context) error {
	return nil
}

func (p *ProdListSvcImpl) DeleteProductList(tx *gorm.DB, ctx *gin.Context) error {
	return nil
}
