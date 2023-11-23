package handler

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/gin-gonic/gin"
)

func getClause(c *gin.Context) (*valueobject.Clause, error) {
	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per_page", "0")
	cl, err := valueobject.NewClause(page, perPage, c.Query("order"))
	if err != nil {
		return nil, err
	}
	return cl, nil
}
