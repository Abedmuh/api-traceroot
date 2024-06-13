package icmp

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func IcmpRoutes(route *gin.RouterGroup, validate *validator.Validate) {

	service := NewIcmpSvc()
	controller := NewIcmpController(service, validate)

	endpoint := route.Group("/icmp")
	{
		endpoint.POST("/", controller.PostLookingGlass)
		endpoint.POST("/list", controller.PostLGlist)
		endpoint.POST("/testsse", controller.PostCountSSE)
	}
}
