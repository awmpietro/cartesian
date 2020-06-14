package controllers

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/awmpietro/cartesian/models"
	"github.com/awmpietro/cartesian/utils"
	"github.com/gin-gonic/gin"
)

type Query struct {
	X        int64 `form:"x"`
	Y        int64 `form:"y"`
	Distance int64 `form:"distance"`
}

func FindPoints(c *gin.Context) {
	points, ok := c.MustGet("points").([]models.Point)
	if !ok {
		fmt.Println(ok)
	}
	var query Query
	var res []models.Point
	var mht int64
	if c.ShouldBind(&query) == nil {
		max := len(points)
		for i := 0; i < max; i++ {
			mht = (utils.Abs(query.X-points[i].X) + utils.Abs(query.Y-points[i].Y))
			if mht <= query.Distance {
				points[i].Manhattan = mht
				res = append(res, points[i])
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Manhattan < res[j].Manhattan
	})
	c.JSON(http.StatusOK, res)
}
