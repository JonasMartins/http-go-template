package routes

import (
	httpHandler "project/src/services/main/internal/handler/http"

	"github.com/gin-gonic/gin"

	cfg "project/src/services/main/configs"
	mdw "project/src/services/main/internal/handler/http/middleware"
)

func Router(r *gin.Engine, h *httpHandler.Handler, config *cfg.Config) {
	rg := LoadPrefix(r, config)
	rg.GET("/ping", h.GetPingHttp)
	rg.POST("/users", h.AddUserHttp)
	rg.PUT("/users/:id", h.UpdateUserHttp)
}

func LoadPrefix(r *gin.Engine, config *cfg.Config) *gin.RouterGroup {
	return r.Group("/api/v1", mdw.AuthApiKeyMiddleware(&config.API.AvaliableApiKeys))
}
