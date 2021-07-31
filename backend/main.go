package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "hello world",
			"message1": "hello world",
			"message2": "hello world",
			"message3": "hello world",
			"message4": "hello world",
			"message5": "hello world",
			"message6": "hello world",
			"message7": "hello world",
		})
	})
	engine.Run(":3001")
}
