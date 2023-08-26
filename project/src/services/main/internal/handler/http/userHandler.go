package http

import (
	"net/http"
	"project/src/pkg/utils"
	"project/src/services/main/domain/usecases"
	ginParser "project/src/services/main/internal/handler/http/helpers"

	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsersHttp(ctx *gin.Context) {
	params, err := ginParser.ParseGinContextToGetUsersParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ctrl.GetUsers(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
func (h *Handler) LoginHttp(ctx *gin.Context) {
	var data usecases.LoginParams
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := h.ctrl.Login(ctx, &data)
	if err != nil {
		switch err.Error() {
		case utils.Wrong_password:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case utils.Not_found:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}

func (h *Handler) AddUserHttp(ctx *gin.Context) {
	var data usecases.AddUserParams
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := h.ctrl.AddUser(ctx, &data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *Handler) UpdateUserHttp(ctx *gin.Context) {
	var data usecases.UpdateUserParams
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error at getting user id": err.Error()})
	}
	data.Id = intId
	res, err := h.ctrl.UpdateUser(ctx, &data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, res)
}
