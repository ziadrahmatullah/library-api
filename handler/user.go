package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	cl, err := getClause(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	name := c.Query("name")
	email := c.Query("email")
	phone := c.Query("phone")

	conditions := []*valueobject.Condition{
		valueobject.NewCondition("name", valueobject.Ilike, name),
		valueobject.NewCondition("email", valueobject.Equal, email),
		valueobject.NewCondition("phone", valueobject.Equal, phone),
	}
	var books []*entity.User
	books = h.userUsecase.GetUsers(*cl, filterCondition(conditions))
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}