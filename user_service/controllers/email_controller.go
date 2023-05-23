package controllers

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"net/http"
	"user_service/models"
	"user_service/utils"
)

type EmailController struct {
	producer sarama.SyncProducer
}

func NewEmailController(producer sarama.SyncProducer) *EmailController {
	return &EmailController{producer: producer}
}

func (e *EmailController) SendEmail(c *gin.Context) {
	var notification models.EmailNotification

	if err := c.BindJSON(&notification); err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// Encode notification as JSON
	payload, err := json.Marshal(notification)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// Send email notification message
	message := &sarama.ProducerMessage{
		Topic: utils.ReadEnv("KAFKA_EMAIL_TOPIC"),
		Value: sarama.ByteEncoder(payload),
	}
	_, _, err = e.producer.SendMessage(message)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	} else {
		c.JSON(http.StatusOK, "Message sent")
	}
}
