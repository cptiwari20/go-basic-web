package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main() {
	router = gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
