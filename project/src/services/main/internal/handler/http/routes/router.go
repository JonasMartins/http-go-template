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
	rg := ApiKeyGroup(r, config)
	authGroup := AuthGroup(r, auth, config)

	/*
		  // custom logger message
			r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
				// your custom format
				return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
					param.ClientIP,
					param.TimeStamp.Format(time.RFC1123),
					param.Method,
					param.Path,
					param.Request.Proto,
					param.StatusCode,
					param.Latency,
					param.Request.UserAgent(),
					param.ErrorMessage,
				)
			}))
			r.Use(gin.Recovery())
	*/

	// * dont need anything
	r.GET(fmt.Sprintf("%s/ping", BasicSubDomain), h.GetPingHttp)

	// * need api key only
	rg.POST("/login", h.LoginHttp)

	// * need token and api key
	authGroup.POST("/users", h.AddUserHttp)
	authGroup.PUT("/users/:id", h.UpdateUserHttp)
	authGroup.GET("/users", h.GetUsersHttp)
	authGroup.GET("/user", h.GetUserHttp)
}

// * group that needs ap key to validate
func ApiKeyGroup(r *gin.Engine, config *cfg.Config) *gin.RouterGroup {
	return r.Group(BasicSubDomain, mdw.AuthApiKeyMiddleware(&config.API.AvaliableApiKeys))
}

// * group that require api_key and token
func AuthGroup(r *gin.Engine, auth auth.TokenFactory, config *cfg.Config) gin.IRoutes {
	return r.Group(BasicSubDomain).Use(mdw.AuthApiKeyMiddleware(&config.API.AvaliableApiKeys)).Use(mdw.AuthTokenMiddleware(auth))
}
