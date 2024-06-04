package users

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserRoutes(route *gin.RouterGroup, db *sql.DB, validate *validator.Validate) {
	service := NewUserService()
	controller := NewUserController(service, db, validate)

	endpoint := route.Group("/users")
	{
		endpoint.POST("/signup", controller.PostUser)
		endpoint.POST("/signin", controller.LoginUser)
	}
}
