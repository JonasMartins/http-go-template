package routes

import (
	"fmt"
	httpHandler "project/src/services/main/internal/handler/http"

	auth "project/src/services/main/internal/handler/auth"

	"github.com/gin-gonic/gin"

	cfg "project/src/services/main/configs"
	mdw "project/src/services/main/internal/handler/http/middleware"
)

const BasicSubDomain = "/api/v1"

func Router(r *gin.Engine, h *httpHandler.Handler, config *cfg.Config, auth auth.TokenFactory) {
	rg := ApieyGroup(r, config)
	authGroup := AuthGroup(r, auth, config)

	// * dont need anything
	r.GET(fmt.Sprintf("%s/ping", BasicSubDomain), h.GetPingHttp)

	// * need api key only
	rg.POST("/login", h.LoginHttp)

	// * need token and api key
	authGroup.POST("/users", h.AddUserHttp)
	authGroup.PUT("/users/:id", h.UpdateUserHttp)
	authGroup.GET("/users", h.GetUsersHttp)
}

// * group that needs ap key to validate
func ApieyGroup(r *gin.Engine, config *cfg.Config) *gin.RouterGroup {
	return r.Group(BasicSubDomain, mdw.AuthApiKeyMiddleware(&config.API.AvaliableApiKeys))
}

// * group that require api_key and token
func AuthGroup(r *gin.Engine, auth auth.TokenFactory, config *cfg.Config) gin.IRoutes {
	return r.Group(BasicSubDomain).Use(mdw.AuthApiKeyMiddleware(&config.API.AvaliableApiKeys)).Use(mdw.AuthTokenMiddleware(auth))
}
