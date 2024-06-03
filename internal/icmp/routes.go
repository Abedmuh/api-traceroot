package icmp

import (
	"github.com/gin-gonic/gin"
)

func IcmpRoutes(route *gin.RouterGroup) {

	controller := NewIcmpController()

	endpoint := route.Group("/icmp")
	{
		endpoint.POST("/ping", controller.PostPing)
		endpoint.POST("/traceroute", controller.PostTraceroute)
		endpoint.POST("/testsse", controller.PostCountSSE)
	}
}
