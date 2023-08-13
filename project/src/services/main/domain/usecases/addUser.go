package usecases

import "github.com/gin-gonic/gin"

type AddUser interface {
	AddUser(ctx *gin.Context, data *AddUserParams) (*AddUserResult, error)
}

type AddUserParams struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type AddUserResult struct {
	Id int `json:"id"`
}
