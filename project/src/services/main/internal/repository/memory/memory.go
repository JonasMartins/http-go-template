package memory

import (
	"project/src/services/main/domain/model"
	"project/src/services/main/domain/usecases"

	"github.com/gin-gonic/gin"

	base "project/src/pkg/model"
)

type Memory struct{}

func New() *Memory {
	return &Memory{}
}

func (m *Memory) GetPing(ctx *gin.Context) (*usecases.GetPingResult, error) {
	return &usecases.GetPingResult{
		Data: model.Ping{
			Base:    base.Base{},
			Message: "Pong Message",
		},
	}, nil
}
