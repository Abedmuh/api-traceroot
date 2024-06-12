package icmp

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func IcmpRoutes(route *gin.RouterGroup, validate *validator.Validate) {

	controller := NewIcmpController(validate)

	endpoint := route.Group("/icmp")
	{
		endpoint.POST("/ping", controller.PostPing)
		endpoint.POST("/testsse", controller.PostCountSSE)
	}
}
