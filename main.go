package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/awmpietro/cartesian/controllers"
	"github.com/awmpietro/cartesian/models"

	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(loadData())
	r.GET("/api/points", controllers.FindPoints)
	return r
}

func loadData() gin.HandlerFunc {
	file, err := ioutil.ReadFile("./data/points.json")
	if err != nil {
		fmt.Println(err)
	}
	var points []models.Point
	json.Unmarshal([]byte(file), &points)

	return func(c *gin.Context) {
		c.Set("points", points)
		c.Next()
	}
}
