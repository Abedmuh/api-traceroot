package productlist

import (
	"github.com/Abedmuh/api-traceroot/utils/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ProductlistRoutes(route *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {

	service := NewProdListSvc()
	controller := NewProdListCtrl(service, db, validate)

	endpoint := route.Group("/productlist")
	{
		endpoint.Use(middleware.Authentication())
		endpoint.POST("/", controller.PostProductList)
		endpoint.GET("/", controller.GetProductLists)
		endpoint.GET("/:id", controller.GetProductList)
		endpoint.PUT("/:id", controller.PutProductList)
		endpoint.DELETE("/:id", controller.DeleteProductList)
	}

}
