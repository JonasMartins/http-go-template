package memory

import (
	"project/src/services/main/domain/model"
	"project/src/services/main/domain/usecases"
	"time"

	"github.com/gin-gonic/gin"

	base "project/src/pkg/model"
	"project/src/pkg/utils"
)

type Memory struct{}

func New() *Memory {
	return &Memory{}
}

func (m *Memory) GetPing(ctx *gin.Context) (*usecases.GetPingResult, error) {

	newUUID, err := utils.GenerateNewUUid()
	if err != nil {
		return nil, err
	}
	return &usecases.GetPingResult{
		Data: model.Ping{
			Base: base.Base{
				ID:        1,
				Uuid:      string(newUUID),
				Version:   1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Message: "Pong Message",
		},
	}, nil
}
