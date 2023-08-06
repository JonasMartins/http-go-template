package routes

import (
	httpHandler "project/src/services/main/internal/handler/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, h *httpHandler.Handler) {
	rg := LoadPrefix(r)
	rg.GET("/ping", h.GetPingHttp)
}

func LoadPrefix(r *gin.Engine) *gin.RouterGroup {
	return r.Group("/api/v1")
}
