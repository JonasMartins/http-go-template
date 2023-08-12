package usecases

import "github.com/gin-gonic/gin"

type AddUser interface {
	AddUser(ctx *gin.Context, data *AddUserParams) (*AddUserResult, error)
}

type AddUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AddUserResult struct {
	Id int `json:"id"`
}
