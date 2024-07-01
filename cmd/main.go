package main

import (
	"github.com/Abedmuh/api-traceroot/internal/icmp"
	productlist "github.com/Abedmuh/api-traceroot/internal/productList"
	serverlist "github.com/Abedmuh/api-traceroot/internal/serverList"
	"github.com/Abedmuh/api-traceroot/internal/users"
	"github.com/Abedmuh/api-traceroot/utils"
	"github.com/Abedmuh/api-traceroot/utils/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	db, err := utils.GetDBConnection()
	if err != nil {
		panic(err)
	}

	api := gin.Default()
	api.Use(cors.Default())
	api.Use(middleware.RecoveryMiddleware())

	validate := validator.New()
	validate.RegisterValidation("command", utils.CommandValidator)

	v1 := api.Group("/v1")
	{
		icmp.IcmpRoutes(v1, validate)
		productlist.ProductlistRoutes(v1, db, validate)
		users.RoutesUser(v1, db, validate)
		serverlist.ServerListRoutes(v1, db, validate)
	}
	api.Run(":8080")
}
