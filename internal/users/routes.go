package users

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RoutesUser(route *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	service := NewUserService()
	controller := NewUserController(service, db, validate)

	endpoint := route.Group("/users")
	{
		endpoint.POST("/signup", controller.PostUser)
		endpoint.POST("/signin", controller.LoginUser)
	}
}
