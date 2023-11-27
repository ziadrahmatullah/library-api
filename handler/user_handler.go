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
		users, err = h.userUsecase.GetUserByName(name)
	} else {
		users, err = h.userUsecase.GetAllUsers()
	}
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
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
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	user, err := h.userUsecase.CreateUser(data)
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Data = user
	ctx.JSON(http.StatusOK, resp)

}

func (h *UserHandler) HandleUserLogin(ctx *gin.Context) {
	resp := dto.Response{}
	var data dto.LoginReq
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		resp.Message = apperror.ErrCannotBindJSON.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	res, err := h.userUsecase.UserLogin(data)
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
