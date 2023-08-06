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

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /ping [get]
func (h *Handler) GetPingHttp(ctx *gin.Context) {
	res, err := h.ctrl.GetPing(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, res)
}
