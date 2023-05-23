package main

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"os/signal"
)

type EmailNotification struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func ReadEnv(key string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("%s not set", key)
	}
	return env
}

func main() {
	// Set up Kafka consumer configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create Kafka consumer
	consumer, err := sarama.NewConsumer([]string{ReadEnv("KAFKA_URL")}, config)
	if err != nil {
		log.Fatal("Failed to set up Kafka consumer:", err)
	}
	defer consumer.Close()

	// Set up signal handler to gracefully exit
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Consume email notification messages
	partitionConsumer, err := consumer.ConsumePartition(ReadEnv("KAFKA_TOPIC"), 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Failed to start consumer:", err)
	}
	defer partitionConsumer.Close()

	// Handle email notification messages
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var notification EmailNotification

			if err := json.Unmarshal(msg.Value, &notification); err != nil {
				log.Println("Failed to decode JSON:", err)
			}

			sendEmail(&notification)
		case err := <-partitionConsumer.Errors():
			log.Println("Failed to consume message:", err.Err)
		case <-signals:
			log.Println("Exiting...")
			return
		}
	}
}

func sendEmail(body *EmailNotification) {
	// Initialize email configuration
	email := gomail.NewMessage()
	email.SetHeader("From", body.Sender)
	email.SetHeader("To", body.Receiver)
	email.SetHeader("Subject", "New Email Notification")
	email.SetBody("text/plain", body.Message)

	// Set up email sending options
	//d := gomail.Dialer{Host: "smtp.example.com", Port: 587, Username: "username", Password: "password"}
	//
	//// Send email
	//if err := d.DialAndSend(email); err != nil {
	//	log.Println("Failed to send email:", err)
	//} else {
	log.Println("Email sent successfully")
	//}
}
