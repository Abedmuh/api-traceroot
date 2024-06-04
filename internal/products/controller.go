package products

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// interfaces
type ProdCtrlInter interface {
	GetProducts(c *gin.Context)
}

type ProdCtrlImpl struct {
	service ProdSvcInter
	Db      *sql.DB
}

func NewProdCtrl(service ProdSvcInter, db *sql.DB) ProdCtrlInter {
	return &ProdCtrlImpl{
		service: service,
		Db:      db,
	}
}

func (c *ProdCtrlImpl) GetProducts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}
