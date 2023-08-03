package mainservice

import (
	usecases "project/src/services/main/domain/usecases"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	getPing usecases.GetPing
}

func New(
	getPing usecases.GetPing,
) *Controller {
	return &Controller{getPing}
}

func (c *Controller) GetPing(ctx *gin.Context) (*usecases.GetPingResult, error) {
	res, err := c.getPing.GetPing(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
