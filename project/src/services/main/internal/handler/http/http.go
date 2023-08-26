package http

import (
	"net/http"
	"project/src/pkg/utils"
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
		switch err.Error() {
		case utils.Server_Error:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
	ctx.JSON(http.StatusOK, res)
}
