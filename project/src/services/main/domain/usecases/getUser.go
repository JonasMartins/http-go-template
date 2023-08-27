package usecases

import (
	"project/src/services/main/domain/model"

	"github.com/gin-gonic/gin"
)

type GetUser interface {
	GetUser(ctx *gin.Context, data *GetUserParams) (*GetUserResult, error)
}

type GetUserParams struct {
	Id    *uint32 `form:"id"`
	Email *string `form:"email"`
	Uuid  *string `form:"uuid"`
}

type GetUserResult struct {
	Data *model.User `json:"data"`
}
