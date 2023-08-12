package mainservice

import (
	usecases "project/src/services/main/domain/usecases"

	"github.com/gin-gonic/gin"
)

type userRepository interface {
	GetPing(ctx *gin.Context) (*usecases.GetPingResult, error)
	AddUser(ctx *gin.Context, data *usecases.AddUserParams) (*usecases.AddUserResult, error)
}

type Controller struct {
	ur userRepository
}

func New(
	ur userRepository,
) *Controller {
	return &Controller{ur}
}

func (c *Controller) GetPing(ctx *gin.Context) (*usecases.GetPingResult, error) {
	res, err := c.ur.GetPing(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Controller) AddUser(ctx *gin.Context, data *usecases.AddUserParams) (*usecases.AddUserResult, error) {
	res, err := c.ur.AddUser(ctx, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
