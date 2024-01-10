package healtzHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/domain/httpDomain"
)

type HealtzHandler struct{}

func New() HealtzHandler {
	return HealtzHandler{}
}

// @Summary		Healt Check
// @Accept		json
// @Produce		json
// @Success		200 {string} {} "ok"
// @Router		/healtz [get]
// @Param		Body body authDomain.ReqRegister true "Register request"
func (h *HealtzHandler) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, httpDomain.ResSucces[any]{
		Message: "ok",
	})
}
