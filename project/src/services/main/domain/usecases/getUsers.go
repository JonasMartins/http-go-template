package usecases

import (
	"project/src/services/main/domain/model"

	"github.com/gin-gonic/gin"
)

type GetUsers interface {
	GetUsers(ctx *gin.Context, data *GetUsersParams) (*GetUsersResult, error)
}

type GetUsersParams struct {
	Page               uint32            `form:"page" binding:"required"`
	Limit              uint32            `form:"limit" binding:"required"`
	Status             *model.UserStatus `form:"status"`
	PreviousTotalItems *uint32           `form:"previous_total_items"`
}

type GetUsersResult struct {
	CurrentPage       uint32        `json:"current_page"`
	TotalPages        uint32        `json:"total_pages"`
	TotalItems        uint32        `json:"total_items"`
	TotalItemsPerPage uint32        `json:"total_items_per_page"`
	Items             []*model.User `json:"items"`
}
