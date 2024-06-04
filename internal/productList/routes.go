package productlist

import (
	"database/sql"

	"github.com/Abedmuh/api-traceroot/utils/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ProductlistRoutes(route *gin.RouterGroup, db *sql.DB, validate *validator.Validate) {

	service := NewProdListSvc()
	controller := NewProdListCtrl(service, db, validate)

	route.Use(middleware.Authentication())
	endpoint := route.Group("/productlist")
	{
		endpoint.POST("/", controller.PostProductList)
		endpoint.GET("/", controller.GetProductLists)
		endpoint.GET("/:id", controller.GetProductList)
		endpoint.PUT("/:id", controller.PutProductList)
	}

}
