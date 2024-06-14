package icmp

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type icmpCtrlInter interface {
	PostLookingGlass(c *gin.Context)
	PostCountSSE(c *gin.Context)
	PostLGlist(c *gin.Context)
}

type icmpCtrlImpl struct {
	service icmpSvcInter
	validate *validator.Validate
}

func NewIcmpController(service icmpSvcInter,validate *validator.Validate) icmpCtrlInter {
	return &icmpCtrlImpl{
		service: service,
		validate: validate,
	}
}

func (c *icmpCtrlImpl) PostLookingGlass(ctx *gin.Context) {
	var req IcmpSsh
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.validate.Struct(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	output, err := c.service.LookingGlass(req)
	if errors.Is(err, errSSHconnection){
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"Message": "ssh successfully completed",
		"data": output,
	})
}

func (c *icmpCtrlImpl) PostCountSSE(ctx *gin.Context) {
	// Set header untuk SSE
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")

	flusher, ok := ctx.Writer.(http.Flusher)
	if !ok {
		ctx.String(http.StatusInternalServerError, "Streaming unsupported!")
		return
	}

	// Kirim perhitungan dari 1 sampai 10
	for i := 1; i <= 10; i++ { // Loop terbatas untuk pengiriman hingga 10
		fmt.Fprintf(ctx.Writer, "data: %d\n\n", i)
		flusher.Flush() // Flush data ke klien

		// Tunggu 1 detik sebelum mengirim angka berikutnya
		time.Sleep(1 * time.Second)
	}

	// Akhiri SSE dengan mengirim pesan selesai
	fmt.Fprint(ctx.Writer, "data: selesai\n\n")
	flusher.Flush()
}

func (c *icmpCtrlImpl) PostLGlist(ctx *gin.Context) {
	var req IcmpSSHs
    if err := ctx.ShouldBindJSON(&req); err!= nil {
        ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
        return
    }

    if err := c.validate.Struct(&req); err!= nil {
        ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
        return
    }
    output, err := c.service.ListedLG(req)
	if errors.Is(err, errSSHconnection){
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
    if err!= nil {
        ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(200, gin.H{
        "Message": "ssh successfully completed",
        "data":    output,
    })
}
