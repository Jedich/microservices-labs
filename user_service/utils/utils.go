package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			log.Println(ginErr.Err)
		}

		// status -1 doesn't overwrite existing status code
		c.Status(-1)
	}
}

func ReadEnv(key string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal("APP_DSN not set")
	}
	return env
}
