package handler

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request dto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(apperror.ErrBinding{ErrBinding: err})
		return
	}
	user := request.ToUser()
	createdUser, err := h.authUsecase.Register(c.Request.Context(), user)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(201, gin.H{
		"data": dto.NewFromUser(createdUser),
	})
}
