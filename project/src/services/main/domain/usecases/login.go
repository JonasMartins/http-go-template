package usecases

import "github.com/gin-gonic/gin"

type Login interface {
	Login(ctx *gin.Context, data *LoginParams) (*LoginResult, error)
}

type LoginParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResult struct {
	Token string `json:"token"`
}
