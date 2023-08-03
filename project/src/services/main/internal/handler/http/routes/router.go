package routes

import (
	httpHandler "project/src/services/main/internal/handler/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, h *httpHandler.Handler) {
	r.GET("/ping", h.GetPingHttp)
}
