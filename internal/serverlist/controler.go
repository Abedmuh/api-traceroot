package serverlist

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ServerListCtrlInter interface {
	GetServerList(c *gin.Context)
	GetServerListById(c *gin.Context)
	PostServerList(c *gin.Context)
	PutServerList(c *gin.Context)
	DeleteServerList(c *gin.Context)
	TestAnsibleServerList(c *gin.Context)
}

type ServerListCtrlImpl struct {
	service  ServerListSvcInter
	Db       *gorm.DB
	validate *validator.Validate
}

func (c *ServerListCtrlImpl) GetServerList(ctx *gin.Context) {
	res, err := c.service.GetServerLists(c.Db, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, res)
}

func (c *ServerListCtrlImpl) GetServerListById(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.service.GetServerListById(id, c.Db, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

func (c *ServerListCtrlImpl) PostServerList(ctx *gin.Context) {
	var req ReqServerList
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.validate.Struct(req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.CreateServerList(req, c.Db, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	err = CreateVmWithESXI(ctx, res)
	if err!= nil {
		ctx.AbortWithStatusJSON(503, gin.H{"error": err.Error()})
	    return
	}

	ctx.JSON(201, gin.H{
		"data":    res,
		"message": "Successfully created server list",
	})
}

func (c *ServerListCtrlImpl) PutServerList(ctx *gin.Context) {}

func (c *ServerListCtrlImpl) DeleteServerList(ctx *gin.Context) {}

func (c *ServerListCtrlImpl) TestAnsibleServerList(ctx *gin.Context) {
	var req ReqServerList
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.validate.Struct(req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	res,err:= c.service.TestAnsibleServer(ctx)
	if err!= nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
        return 
    }
    ctx.JSON(201, gin.H{
		"data":    res,
		"message": "Successfully created server list",
	})
}