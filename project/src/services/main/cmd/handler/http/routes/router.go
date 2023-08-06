package routes

import (
	swgDocs "project/src/services/main/cmd/docs"
	httpHandler "project/src/services/main/cmd/handler/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gsw "github.com/swaggo/gin-swagger"
)

func Router(r *gin.Engine, h *httpHandler.Handler) {
	rg := LoadPrefix(r)
	swgDocs.SwaggerInfo.BasePath = "/api/v1"
	rg.GET("/docs/*any", gsw.WrapHandler(swaggerFiles.Handler))
	rg.GET("/ping", h.GetPingHttp)
}

func LoadPrefix(r *gin.Engine) *gin.RouterGroup {
	return r.Group("/api/v1")
}
