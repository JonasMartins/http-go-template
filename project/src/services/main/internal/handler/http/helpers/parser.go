package helpers

import (
	"project/src/services/main/domain/usecases"

	"github.com/gin-gonic/gin"
)

// * a method to get and parse params from
// * gin context to a valid GetUsersParams struct
func ParseGinContextToGetUsersParams(ctx *gin.Context) (*usecases.GetUsersParams, error) {
	// * bind a param to a variable, if not passed, set a default
	// * or use ctx.Query to return empty string if not found
	/*
		page := ctx.DefaultQuery("page", "1")
		limit := ctx.DefaultQuery("limit", "10")
		status := ctx.DefaultQuery("status", "active")
	*/

	var params usecases.GetUsersParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		return nil, err
	}
	return &params, nil
}

func ParseGinContextToGetUserParams(ctx *gin.Context) (*usecases.GetUserParams, error) {
	var params usecases.GetUserParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		return nil, err
	}
	return &params, nil
}
