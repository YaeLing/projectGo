package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := setupRoute()
	server.Run(":8888")
}

func setupRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		fmt.Printf("ctx.Request.Header: %v\n", c.Request.Header)
	}, func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	return router
}
