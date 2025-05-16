package controllers

import (
	"archive-server/config"
	"archive-server/models"
	"archive-server/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateJournalArchive(c *fiber.Ctx) error {

	var journal models.Journal
	if err := c.BodyParser(&journal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load configuration",
		})
	}

	err = services.Publisher(services.RabbitMQConn, cfg.RabbitMQ.JournalQueue, journal)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Journal archive request submitted successfully",
	})
}

func CreateTransactionArchive(c *fiber.Ctx) error {
	var transaction models.Transaction
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load configuration",
		})
	}

	err = services.Publisher(services.RabbitMQConn, cfg.RabbitMQ.TransactionQueue, transaction)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "transaction archive request submitted successfully",
	})
}
