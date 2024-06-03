package main

import (
	"github.com/Abedmuh/api-traceroot/internal/icmp"
	"github.com/gin-gonic/gin"
)

func main() {

	api := gin.Default()

	v1 := api.Group("/v1")
	{
		icmp.IcmpRoutes(v1)
	}
	api.Run() // listen and serve on 0.0.0.0:8080
}
