package icmp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// type icmpmodel struct {
// 	Message string `json:"message"`
// }

type icmpCtrlInter interface {
	PostPing(c *gin.Context)
	PostTraceroute(c *gin.Context)
	PostCountSSE(c *gin.Context)
}

type icmpCtrlImpl struct {
}

func NewIcmpController() icmpCtrlInter {
	return &icmpCtrlImpl{}
}

func (c *icmpCtrlImpl) PostPing(ctx *gin.Context) {
	var req string
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"req": req})
}

func (c *icmpCtrlImpl) PostTraceroute(ctx *gin.Context) {
	var req string
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
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
