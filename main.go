package main

import (
	"log"
	"persistence-service/cmd"
	"persistence-service/kafka"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cmd.InitDB()
	log.Println("Persistence Service running...")
	kafka.StartKafkaConsumer()
}
