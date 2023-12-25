package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	r.GET("/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "items",
		})
	})
	r.GET("/products", func(c *gin.Context) {
		type product struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		var list []product

		for i := 0; i < 10; i++ {
			list = append(list, product{
				ID:   i,
				Name: "clothes",
			})
		}
		c.JSON(http.StatusOK, list)
	})
	r.Run()
}
