package http

import (
	"net/http"
	controller "project/src/services/main/internal/controller/main_service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ctrl *controller.Controller
}

func New(ctrl *controller.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetPingHttp(ctx *gin.Context) {
	res, err := h.ctrl.GetPing(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, res)
}
