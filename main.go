package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type News struct {
	Date int    `json:"date" binding: “required”`
	Name string `json:"name" binding: “required”`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"response": "ping",
		})
	})

	r.GET("/pong/:id", func(c *gin.Context) {

		paramsId := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"response": paramsId,
		})
	})

	r.POST("/add", func(c *gin.Context) {

		var data News

		if err := c.ShouldBind(&data); err != nil {

			fmt.Println(err)

			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{

			"data": data,
		})

	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
