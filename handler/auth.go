package handler

import (
	"log"
	"net/http"

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
func (h *AuthHandler) Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(apperror.ErrBinding{ErrBinding: err})
	}
	log.Println(request)
	user := request.ToUser()
	token, err := h.authUsecase.Login(c.Request.Context(), user)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})
}
