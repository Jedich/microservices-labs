package main

import (
	"fmt"
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"user_service/controllers"
	"user_service/db"
	"user_service/kafka"
	"user_service/utils"
)

var delay bool = false

func main() {
	conn := db.OpenConnection()

	producer := kafka.CreateKafkaProducer()
	_ = utils.ReadEnv("KAFKA_EMAIL_TOPIC")

	r := gin.Default()
	r.Use(utils.ErrorHandler())
	r.Use(fluentdMiddleware(NewFluentLogger()))

	r.GET("/api/golang-service/", func(c *gin.Context) {
		if delay {
			time.Sleep(10 * time.Second)
		}
		c.JSON(http.StatusOK, gin.H{"response": "I am a golang microservice!"})
	})

	r.GET("/api/golang-service/kill", func(c *gin.Context) {
		delay = true
		c.JSON(http.StatusTeapot, gin.H{"response": "im unalived"})
	})

	userController := controllers.NewUserController(conn.Connection)
	r.POST("/api/user", userController.CreateUser)

	emailController := controllers.NewEmailController(producer)
	r.POST("/api/send-email", emailController.SendEmail)

	defer producer.Close()
	defer conn.CloseConnection()

	err := r.Run(":8083")

	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}

func fluentdMiddleware(logger *fluent.Fluent) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log request details
		logger.Post("http_request", map[string]string{
			"url":    c.Request.URL.Path,
			"method": c.Request.Method,
			"ip":     c.ClientIP(),
		})

		// Continue processing the request
		c.Next()
	}
}

func NewFluentLogger() *fluent.Fluent {
	logger, err := fluent.New(fluent.Config{
		FluentHost: utils.ReadEnv("FLUENT_HOST"), // Fluentd server host
		FluentPort: 24224,                        // Fluentd server port
		TagPrefix:  "user-mserv",                 // Fluentd tag prefix
		Async:      false,                        // Whether to send logs asynchronously (true) or synchronously (false)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to fluentd")
	return logger
}
