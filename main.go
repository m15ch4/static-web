package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func main() {
	r := gin.Default()

	r.GET("/", rootHandler)

	r.Run(":8080")
}
