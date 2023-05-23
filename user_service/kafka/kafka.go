package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"user_service/utils"
)

func CreateKafkaProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create Kafka producer
	producer, err := sarama.NewSyncProducer([]string{utils.ReadEnv("KAFKA_URL")}, config)
	if err != nil {
		log.Fatal(err)
	}

	return producer
}
