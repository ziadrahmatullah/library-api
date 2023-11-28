package handler

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/gin-gonic/gin"
)

func getQuery(c *gin.Context) (*valueobject.Query, error) {
	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per_page", "0")
	query, err := valueobject.NewQuery(page, perPage, c.Query("order"))
	if err != nil {
		return nil, err
	}
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
