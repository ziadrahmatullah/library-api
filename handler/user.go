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
	title := c.Query("title")
	quantity := c.Query("qty")
	description := c.Query("desc")

	conditions := []*valueobject.Condition{
		valueobject.NewCondition("title", valueobject.Ilike, title),
		valueobject.NewCondition("description", valueobject.Ilike, description),
		valueobject.NewCondition("quantity", valueobject.Equal, quantity),
	}
	var books []*entity.User
	books = h.userUsecase.GetUsers(*cl, filterCondition(conditions))
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
