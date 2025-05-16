package services

import (
	"archive-server/config"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection

func ConnectRabbitMQ(config *config.Config) (*amqp.Connection, error) {
	conn, err := amqp.Dial(config.RabbitMQ.URL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	for _, queueName := range config.RabbitMQ.Queues {
		_, err = ch.QueueDeclare(
			queueName,
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			log.Fatalf("Failed to declare a queue: %v", err)
		}
	}

	defer ch.Close()

	log.Printf("Connected to RabbitMQ and declared queues: %v", config.RabbitMQ.Queues)
	RabbitMQConn = conn
	return conn, nil
}

func Publisher(rabbitMQConn *amqp.Connection, queueName string, data interface{}) error {
	ch, err := rabbitMQConn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclarePassive(queueName, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("queue '%s' does not exist: %w", queueName, err)
	}

	message, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	err = ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message to '%s': %w", queueName, err)
	}
	log.Printf("Published a message to queue: %s", queueName)
	return nil
}

// func ConsumeMessages(rabbitMQConn *amqp.Connection, mongoClient *mongo.Client, queueNames []string) {
// 	for _, queueName := range queueNames {
// 		go func(queueName string) {
// 			ch, err := rabbitMQConn.Channel()
// 			if err != nil {
// 				log.Fatalf("Failed to open a channel for queue %s: %v", queueName, err)
// 			}
// 			defer ch.Close()

// 			msgs, err := ch.Consume(
// 				queueName, // queue
// 				"",        // consumer
// 				false,     // auto-ack
// 				false,     // exclusive
// 				false,     // no-local
// 				false,     // no-wait
// 				nil,       // args
// 			)
// 			if err != nil {
// 				log.Fatalf("Failed to register a consumer for queue %s: %v", queueName, err)
// 			}
// 			log.Printf("Waiting for messages in queue: %s", queueName)

// 			for msg := range msgs {
// 				log.Printf("Received a message from queue %s: %s", queueName, msg.Body)
// 				var journal models.Journal
// 				err := json.Unmarshal(msg.Body, &journal)
// 				if err != nil {
// 					log.Printf("Failed to unmarshal message from queue %s: %v", queueName, err)
// 					msg.Nack(false, false) // Reject the message without requeueing
// 					return
// 				}

// 				database.ArchiveJournals(msg.Body)

// 				// collection := mongoClient.Database("your_database_name").Collection("your_collection_name")
// 				// _, err := collection.InsertOne(nil, map[string]interface{}{
// 				// 	"queue":   queueName,
// 				// 	"message": string(msg.Body),
// 				// 	"timestamp": msg.Timestamp,
// 				// })
// 				// if err != nil {
// 				// 	log.Printf("Failed to insert message into MongoDB for queue %s: %v", queueName, err)
// 				// } else {
// 				// 	log.Printf("Message inserted into MongoDB for queue %s", queueName)
// 				// }

// 				msg.Ack(false) // Acknowledge the message
// 			}
// 			log.Printf("Stopped consuming messages from queue: %s", queueName)
// 		}(queueName)
// 	}
// }
