package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", Test)
	r.Run(":8080")
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"result": "cool"})
}
