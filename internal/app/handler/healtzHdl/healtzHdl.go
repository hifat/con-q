package healtzHdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealtzHandler struct{}

func New() HealtzHandler {
	return HealtzHandler{}
}

func (h *HealtzHandler) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
