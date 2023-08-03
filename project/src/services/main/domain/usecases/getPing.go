package usecases

import (
	"project/src/services/main/domain/model"

	"github.com/gin-gonic/gin"
)

type GetPing interface {
	GetPing(ctx *gin.Context) (*GetPingResult, error)
}

type GetPingResult struct {
	Data model.Ping `json:"data"`
}
