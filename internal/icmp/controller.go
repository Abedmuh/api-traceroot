package icmp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Abedmuh/api-traceroot/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// type icmpmodel struct {
// 	Message string `json:"message"`
// }

type icmpCtrlInter interface {
	PostPing(c *gin.Context)
	PostCountSSE(c *gin.Context)
}

type icmpCtrlImpl struct {
	validate *validator.Validate
}

func NewIcmpController(validate *validator.Validate) icmpCtrlInter {
	return &icmpCtrlImpl{
		validate: validate,
	}
}

func (c *icmpCtrlImpl) PostPing(ctx *gin.Context) {
	var req IcmpSsh
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.validate.Struct(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	output, err := utils.SshTarget(req.Address, req.Command)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"Message": "ssh successfully completed",
		"data":    output,
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
