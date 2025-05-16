package services

import (
	"archive-server/config"
	"archive-server/database"
	"archive-server/models"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConsumeMessages(rabbitMQConn *amqp.Connection, mongoClient *mongo.Client) {
	go JournalConsumer(rabbitMQConn, mongoClient)
	go TransactionConsumer(rabbitMQConn, mongoClient)
}

func JournalConsumer(rabbitMQConn *amqp.Connection, mongoClient *mongo.Client) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %v", err)
	}
	queue := cfg.RabbitMQ.JournalQueue

	go func(queue string) {
		ch, err := rabbitMQConn.Channel()
		if err != nil {
			panic(err)
		}
		defer ch.Close()

		msgs, err := ch.Consume(
			queue,
			"",    // consumer
			false, // auto-ack
			false, // exclusive
			false, // no-local
			false, // no-wait
			nil,   // args
		)
		if err != nil {
			log.Fatalf("Failed to register a consumer for queue %s: %v", queue, err)
		}

		for msg := range msgs {
			var message models.Journal
			err := json.Unmarshal(msg.Body, &message)
			if err != nil {
				log.Printf("Failed to unmarshal message from queue %s: %v", queue, err)
				continue
			}
			if err := database.ArchiveJournal(message); err != nil {
				log.Printf("Failed to archive journal: %v", err)
				msg.Nack(false, true)
				continue
			}
			msg.Ack(false)
			log.Printf("Journal consumed")
		}
	}(queue)

}

func TransactionConsumer(rabbitMQConn *amqp.Connection, mongoClient *mongo.Client) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %v", err)
	}
	queue := cfg.RabbitMQ.TransactionQueue

	go func(queue string) {
		ch, err := rabbitMQConn.Channel()
		if err != nil {
			panic(err)
		}
		defer ch.Close()

		msgs, err := ch.Consume(
			queue,
			"",    // consumer
			false, // auto-ack
			false, // exclusive
			false, // no-local
			false, // no-wait
			nil,   // args
		)
		if err != nil {
			log.Fatalf("Failed to register a consumer for queue %s: %v", queue, err)
		}

		for msg := range msgs {
			var message models.Transaction
			err := json.Unmarshal(msg.Body, &message)
			if err != nil {
				log.Printf("Failed to unmarshal message from queue %s: %v", queue, err)
				continue
			}
			if err := database.ArchiveTransaction(message); err != nil {
				log.Printf("Failed to archive transactions: %v", err)
				msg.Nack(false, true)
				continue
			}
			msg.Ack(false)
		}
	}(queue)

}
