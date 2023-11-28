package handler

import (
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/gin-gonic/gin"
)

func getQuery(c *gin.Context) (*valueobject.Query, error) {
	pageQuery := c.DefaultQuery("page", "1")
	perPageQuery := c.DefaultQuery("per_page", "0")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		return nil, err
	}
	perPage, err := strconv.Atoi(perPageQuery)
	if err != nil {
		return nil, err
	}
	query := valueobject.NewQuery().Paginate(page, perPage).Order(c.Query("order"))

	return query, nil
}

func filterCondition(conditions []*valueobject.Condition) []*valueobject.Condition {
	filtered := make([]*valueobject.Condition, 0)
	for _, condition := range conditions {
		if condition != nil {
			filtered = append(filtered, condition)
		}
	}
	return filtered
}
