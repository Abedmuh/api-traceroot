package products

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RoutesProducts(route *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	service := NewProductService()
	controller := NewProductController(service, db, validate)

	endpoint := route.Group("/products")
	{
		endpoint.GET("/", controller.GetProducts)
		endpoint.GET("/:id", controller.GetProductByID)
		endpoint.POST("/", controller.PostProduct)
		endpoint.PUT("/:id", controller.PutProduct)
		endpoint.DELETE("/:id", controller.DeleteProduct)
	}
}
