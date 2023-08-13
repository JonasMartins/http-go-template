package usecases

import "github.com/gin-gonic/gin"

type UpdateUser interface {
	UpdateUser(ctx *gin.Context, data *UpdateUserParams) (*UpdateUserResult, error)
}

type UpdateUserParams struct {
	Id       int     `json:"id"`
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}
type UpdateUserResult struct {
	Id int `json:"id"`
}
