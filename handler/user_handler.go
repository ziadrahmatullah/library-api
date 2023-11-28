package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: uu,
	}
}

func (h *UserHandler) HandleGetUsers(ctx *gin.Context) {
	resp := dto.Response{}
	name := ctx.Query("name")
	var users []models.User
	var err error
	if name != "" {
		users, err = h.userUsecase.GetUserByName(ctx, name)
	} else {
		users, err = h.userUsecase.GetAllUsers(ctx)
	}
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = users
	ctx.JSON(http.StatusOK, resp)
}

func (h *UserHandler) HandleUserRegister(ctx *gin.Context) {
	resp := dto.Response{}
	var data dto.RegisterReq
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	user, err := h.userUsecase.CreateUser(ctx, data)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = user
	ctx.JSON(http.StatusOK, resp)

}

func (h *UserHandler) HandleUserLogin(ctx *gin.Context) {
	var data dto.LoginReq
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	resp, err := h.userUsecase.UserLogin(ctx, data)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
