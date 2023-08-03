package mainservice

import (
	usecases "project/src/services/main/domain/usecases"

	"github.com/gin-gonic/gin"
)

type generalRepository interface {
	GetPing(ctx *gin.Context) (*usecases.GetPingResult, error)
}

type Controller struct {
	gr generalRepository
}

func New(
	gr generalRepository,
) *Controller {
	return &Controller{gr}
}

func (c *Controller) GetPing(ctx *gin.Context) (*usecases.GetPingResult, error) {
	res, err := c.gr.GetPing(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
