package main

import (
	"archive-server/config"
	"archive-server/database"
	"archive-server/routes"
	"archive-server/services"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	rabbitMQConn, err := services.ConnectRabbitMQ(cfg)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	defer rabbitMQConn.Close()

	mongoClient, err := database.ConnectMongoDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	go services.ConsumeMessages(rabbitMQConn, mongoClient)

	app := fiber.New()
	app.Get("/matrix", monitor.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.SetupRoutes(app)
	log.Println("Starting server on :3000")
	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}

	select {}
}
