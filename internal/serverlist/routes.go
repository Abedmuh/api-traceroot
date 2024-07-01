package serverlist

import (
	"github.com/Abedmuh/api-traceroot/utils/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ServerListRoutes(route *gin.RouterGroup, tx *gorm.DB, validate *validator.Validate) {
	service := NewServerListService()
	controller := NewServerListCtrl(service, tx, validate)
	endpoint := route.Group("/serverlist")
	{
		endpoint.Use(middleware.Authentication())
		endpoint.POST("/", controller.PostServerList)
		endpoint.GET("/", controller.GetServerList)
		endpoint.GET("/:id", controller.GetServerListById)
		endpoint.PUT("/:id", controller.PutServerList)
		endpoint.DELETE("/:id", controller.DeleteServerList)
	}
}