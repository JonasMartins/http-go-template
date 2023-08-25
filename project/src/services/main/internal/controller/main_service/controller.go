package mainservice

import (
	usecases "project/src/services/main/domain/usecases"

	"github.com/gin-gonic/gin"
)

type userRepository interface {
	GetPing(ctx *gin.Context) (*usecases.GetPingResult, error)
	AddUser(ctx *gin.Context, data *usecases.AddUserParams) (*usecases.AddUserResult, error)
	UpdateUser(ctx *gin.Context, data *usecases.UpdateUserParams) (*usecases.UpdateUserResult, error)
	Login(ctx *gin.Context, data *usecases.LoginParams) (*usecases.LoginResult, error)
	GetUsers(ctx *gin.Context, params *usecases.GetUsersParams) (*usecases.GetUsersResult, error)
}

type Controller struct {
	ur userRepository
}

func New(
	ur userRepository,
) *Controller {
	return &Controller{ur}
}

func (c *Controller) GetUsers(ctx *gin.Context, params *usecases.GetUsersParams) (*usecases.GetUsersResult, error) {
	res, err := c.ur.GetUsers(ctx, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Controller) GetPing(ctx *gin.Context) (*usecases.GetPingResult, error) {
	res, err := c.ur.GetPing(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Controller) Login(ctx *gin.Context, data *usecases.LoginParams) (*usecases.LoginResult, error) {
	res, err := c.ur.Login(ctx, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Controller) UpdateUser(ctx *gin.Context, data *usecases.UpdateUserParams) (*usecases.UpdateUserResult, error) {
	res, err := c.ur.UpdateUser(ctx, data)
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
