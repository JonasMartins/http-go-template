package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"project/src/pkg/utils"
	"project/src/services/main/internal/handler/auth"

	"github.com/gin-gonic/gin"
)

const (
	api_key   = "api_key"
	authToken = "authorization"
)

func AuthTokenMiddleware(tf auth.TokenFactory) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authToken)
		if len(authHeader) == 0 {
			err := errors.New("authorization is required")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		if strings.ToLower(fields[0]) != "bearer" {
			err := fmt.Errorf(`invalid authorization type, required %s`, `"bearer"`)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		payload, err := tf.VerifyToken(fields[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		ctx.Set("authorization_token", payload)
		ctx.Next()
	}
}

func AuthApiKeyMiddleware(avaliableApiKey *[]string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apikey := ctx.GetHeader(api_key)
		if len(apikey) == 0 {
			err := errors.New("api_key is required")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		if !utils.StringContains(&apikey, avaliableApiKey) {
			err := errors.New("invalid api_key")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		ctx.Next()
	}
}
