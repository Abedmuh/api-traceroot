package users

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

//interfaces

type UserCtrlInter interface {
	PostUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type UserCtrlImpl struct {
	service  UserSvcInter
	DB       *gorm.DB
	validate *validator.Validate
}

func NewUserController(service UserSvcInter, DB *gorm.DB, validate *validator.Validate) UserCtrlInter {
	return &UserCtrlImpl{
		service:  service,
		DB:       DB,
		validate: validate,
	}
}

func (c *UserCtrlImpl) PostUser(ctx *gin.Context) {
	var req Users
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.validate.Struct(req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := c.service.CreateUser(req, c.DB, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Successfully created"})
}

func (c *UserCtrlImpl) LoginUser(ctx *gin.Context) {
	var req ReqUserLog
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.validate.Struct(req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.LoginUser(req, c.DB, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"data":    res,
		"message": "User logged in successfully",
	})
}
