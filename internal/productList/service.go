package productlist

import (
	"time"

	"github.com/Abedmuh/api-traceroot/internal/products"
	"github.com/Abedmuh/api-traceroot/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// interface
type ProdListSvcInter interface {
	CreateProductList(req products.Products, tx *gorm.DB, ctx *gin.Context) (ProductList, error)
	GetProductsLists(tx *gorm.DB, ctx *gin.Context) ([]ProductList, error)
	GetProductListById(id string, tx *gorm.DB, ctx *gin.Context) (products.Products, error)
	UpdateProductList(id string,req ProductList, tx *gorm.DB, ctx *gin.Context) error
	DeleteProductList(tx *gorm.DB, ctx *gin.Context) error
}

type ProdListSvcImpl struct {
}

func NewProdListSvc() ProdListSvcInter {
	return &ProdListSvcImpl{}
}

func (p *ProdListSvcImpl) CreateProductList(req products.Products, tx *gorm.DB, ctx *gin.Context) (ProductList, error) {
	// decode token auth
	user, err := utils.GetTokenEmail(ctx)
	if err!= nil {
        return ProductList{}, err
    }

	//model synchronized
	productList := ProductList{
		Owner:     user,
		Timelimit: time.Now().Add(time.Duration(24) * time.Hour),
		Username:  req.Username,
		Password:  req.Password,
		Name:      req.Name,
		Os:        req.Os,
		Cpu:       req.Cpu,
		Ram:       req.Ram,
		Storage:   req.Storage,
		Firewall:  req.Firewall,
		Selinux:   req.Selinux,
		Location:  req.Location,
	}

	// create product list
	if err := tx.Create(productList).Error; err != nil {
		return ProductList{}, err
	}

	return productList, nil
}

func (p *ProdListSvcImpl) GetProductListById(req string, tx *gorm.DB, ctx *gin.Context) (products.Products, error) {
	user, err := utils.GetTokenEmail(ctx)
	if err!= nil {
        return products.Products{}, err
    }

	var productList ProductList
	if err := tx.Where("id =? AND owner =?", req, user).First(&productList).Error; err!= nil {
        return products.Products{}, err
    }
	return products.Products{}, nil
}

func (p *ProdListSvcImpl) GetProductsLists(tx *gorm.DB, ctx *gin.Context) ([]ProductList, error) {
	user, err := utils.GetTokenEmail(ctx)
	if err != nil {
		return nil, err
	}

	var productLists []ProductList
	if err := tx.Where("owner =?", user).Find(&productLists).Error; err!= nil {
        return nil, err
    }
	return productLists, nil
}

func (p *ProdListSvcImpl) UpdateProductList(id string,req ProductList, tx *gorm.DB, ctx *gin.Context) error {
	return nil
}

func (p *ProdListSvcImpl) DeleteProductList(tx *gorm.DB, ctx *gin.Context) error {
	return nil
}
