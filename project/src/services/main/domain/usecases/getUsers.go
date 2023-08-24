package usecases

import (
	"project/src/services/main/domain/model"

	"github.com/gin-gonic/gin"
)

type GetUsers interface {
	GetUsers(ctx *gin.Context, data *GetUsersParams) (*GetUsersResult, error)
}

type GetUsersParams struct {
	Page   uint32           `json:"page" binding:"required"`
	Limit  uint32           `json:"limit" binding:"required"`
	Status model.UserStatus `json:"status"`
}

type GetUsersResult struct {
	CurrentPage uint32        `json:"current_page"`
	TotalIPages uint32        `json:"total_pages"`
	TotalItems  uint32        `json:"total_items"`
	Items       *[]model.User `json:"items"`
}
