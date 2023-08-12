package http

import (
	"net/http"
	"project/src/services/main/domain/usecases"
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

func (h *Handler) AddUserHttp(ctx *gin.Context) {
	var data usecases.AddUserParams
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := h.ctrl.AddUser(ctx, &data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, res)
}
