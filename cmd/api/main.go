package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", hello)
	r.Run()
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
