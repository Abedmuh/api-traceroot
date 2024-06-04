package users

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//interfaces

type UserCtrlInter interface {
	PostUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type UserCtrlImpl struct {
	service  UserSvcInter
	DB       *sql.DB
	validate *validator.Validate
}

func NewUserController(service UserSvcInter, DB *sql.DB, validate *validator.Validate) UserCtrlInter {
	return &UserCtrlImpl{
		service:  service,
		DB:       DB,
		validate: validate,
	}
}

func (c *UserCtrlImpl) PostUser(ctx *gin.Context) {
	var req ReqUserReg
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.validate.Struct(req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	output, err := c.service.CreateUser(req, c.DB, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": output})
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

	user, err := c.service.CheckUserLog(req.Email, c.DB, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.LoginUser(user, req, c.DB, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "User logged in successfully",
		"data":    res,
	})
}
