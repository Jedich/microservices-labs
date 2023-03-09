package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/hello-go", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response": "I am a golang microservice!"})
	})

	err := r.Run(":8083")

	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
