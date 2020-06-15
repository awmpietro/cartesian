package controllers

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/awmpietro/cartesian/models"
	"github.com/awmpietro/cartesian/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Query struct {
	X        int64 `form:"x" validate:"required"`
	Y        int64 `form:"y" validate:"required"`
	Distance int64 `form:"distance" validate:"required"`
}

func FindPoints(c *gin.Context) {
	points, ok := c.MustGet("points").([]models.Point)
	if !ok {
		fmt.Println(ok)
	}
	var query Query
	if c.ShouldBind(&query) == nil {
		err := utils.Validation(query)
		if err != nil {
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.Field()+" is "+err.Tag())
			}
			c.JSON(http.StatusBadRequest, errors)
			return
		}
		var res []models.Point
		var mht int64
		max := len(points)
		for i := 0; i < max; i++ {
			mht = (utils.Abs(query.X-points[i].X) + utils.Abs(query.Y-points[i].Y))
			if mht <= query.Distance {
				points[i].Manhattan = mht
				res = append(res, points[i])
			}
		}
		if len(res) > 0 {
			sort.Slice(res, func(i, j int) bool {
				return res[i].Manhattan < res[j].Manhattan
			})
			c.JSON(http.StatusOK, res)
			return
		}
		c.JSON(http.StatusOK, gin.H{"empty": "no results were found with the provided data"})
	}
}
