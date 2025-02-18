package main

import (
	// "fmt"
	//"net/http"
	
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello World\n"})
    })
    r.Run(":8080")
}