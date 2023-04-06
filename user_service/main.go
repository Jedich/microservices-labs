package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"user_service/controllers"
	"user_service/db"
	"user_service/utils"
)

func main() {
	conn := db.OpenConnection()

	r := gin.Default()
	r.Use(utils.ErrorHandler())

	r.GET("/api/golang-service/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response": "I am a golang microservice!"})
	})

	userController := controllers.NewUserController(conn.Connection)
	r.POST("/user", userController.CreateUser)

	defer conn.CloseConnection()

	err := r.Run(":8083")

	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
