package kafka

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"persistence-service/cmd"
	"persistence-service/model"

	"github.com/segmentio/kafka-go"
)

func StartKafkaConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("BROKERS")},
		Topic:   "message-topic",
		GroupID: "persistence-group",
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Kafka read error:", err)
			continue
		}

		log.Println("value", string(msg.Value))
		log.Println("topic", string(msg.Topic))

		var data model.Message
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			log.Println("Invalid message JSON:", err)
			continue
		}

		log.Println(data)

		if err := cmd.DB.Create(&data).Error; err != nil {
			log.Println("DB insert error:", err)
			continue
		}

		log.Printf("Stored message: %+v\n", data)
	}
}
