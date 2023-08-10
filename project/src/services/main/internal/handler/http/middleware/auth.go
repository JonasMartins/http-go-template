package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"project/src/pkg/utils"
)

const (
	api_key = "api_key"
)

func AuthMiddleware(avaliableApiKey *[]string) gin.HandlerFunc {
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
